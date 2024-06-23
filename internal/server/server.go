package server

import (
	"context"
	"log/slog"
	"os"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	protoApi "github.com/ikaliuzh/card-validator/api/proto"
	"github.com/ikaliuzh/card-validator/pkg/card"
	"github.com/ikaliuzh/card-validator/pkg/errorcodes"
	"github.com/ikaliuzh/card-validator/pkg/validator"
	"github.com/ikaliuzh/card-validator/pkg/validator/combined"
	"github.com/ikaliuzh/card-validator/pkg/validator/expirydate"
	"github.com/ikaliuzh/card-validator/pkg/validator/luhn"
)

type Server struct {
	protoApi.UnimplementedCardValidatorServer

	log         *slog.Logger
	validator   validator.Validator
	respTimeout time.Duration
}

func New(options ...func(*Server)) *Server {
	svr := &Server{
		respTimeout: 1 * time.Second,
		validator:   combined.New(expirydate.New(), luhn.New()),
		log:         slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo})),
	}
	for _, o := range options {
		o(svr)
	}
	return svr
}

func WithTimeout(timeout time.Duration) func(*Server) {
	return func(s *Server) {
		s.respTimeout = timeout
	}
}

func WithValidator(v validator.Validator) func(*Server) {
	return func(s *Server) {
		s.validator = v
	}
}

func WithLog(log *slog.Logger) func(*Server) {
	return func(s *Server) {
		s.log = log
	}
}

func (s *Server) ValidateCard(ctx context.Context, req *protoApi.Card) (*protoApi.CardValidationResponse, error) {
	log := s.log.With(
		slog.Group("card",
			slog.String("number", req.CardNumber),
			slog.String("expiration_month", req.ExpirationMonth),
			slog.String("expiration_year", req.ExpirationYear),
		),
	)

	log.Debug("validating card")

	validationResult := make(chan error, 1)
	go func() {
		defer close(validationResult)
		creditCard, err := card.NewCardFromProto(req)
		if err != nil {
			validationResult <- err
			return
		}
		validationResult <- s.validator.Validate(ctx, creditCard)
	}()

	select {
	case validationError := <-validationResult:
		if validationError == nil {
			log.Info("card is valid")
			return &protoApi.CardValidationResponse{IsValid: true}, nil
		}

		code, ok := errorcodes.Extract(validationError)
		if !ok {
			log.Error("unexpected error", slog.Any("error", validationError))
			return nil, status.Errorf(codes.Internal, "unexpected error: %v", validationError)
		}

		log.Error("card is invalid",
			slog.String("code", code), slog.Any("reason", validationError))
		return &protoApi.CardValidationResponse{
			IsValid: false,
			Error:   &protoApi.Error{Code: code, Message: validationError.Error()},
		}, nil

	case <-time.After(s.respTimeout):
		log.Warn("timeout validating card")
		return nil, status.Errorf(codes.DeadlineExceeded, "timeout validating card after %v seconds", s.respTimeout)

	case <-ctx.Done():
		log.Warn("request canceled", slog.Any("reason", ctx.Err()))
		return nil, status.Errorf(codes.Canceled, "request canceled: %v", ctx.Err())
	}
}
