package server_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/ikaliuzh/card-validator/gen/proto"
	"github.com/ikaliuzh/card-validator/internal/server"
	"github.com/ikaliuzh/card-validator/pkg/errorcodes"
	mock_validator "github.com/ikaliuzh/card-validator/pkg/validator/mock"
)

func TestUserService_GetUser(t *testing.T) {
	type validationResponse struct {
		IsValid   bool
		ErrorCode string
	}
	testCases := []struct {
		description      string
		card             *proto.Card
		validatorVerdict error
		expectedResp     *validationResponse
		expectedError    string
	}{
		{
			description: "valid card",
			card: &proto.Card{
				CardNumber:      "4111111111111111",
				ExpirationMonth: "12",
				ExpirationYear:  "2024",
			},
			expectedResp: &validationResponse{
				IsValid: true,
			},
		},
		{
			description: "invalid card number format - number containing letters",
			card: &proto.Card{
				CardNumber:      "411111111111111a",
				ExpirationMonth: "12",
				ExpirationYear:  "2024",
			},
			expectedResp: &validationResponse{
				IsValid:   false,
				ErrorCode: errorcodes.CardNumberInvalidFormat,
			},
		},
		{
			description: "invalid expiry date format - invalid month",
			card: &proto.Card{
				CardNumber:      "411111111111111",
				ExpirationMonth: "27",
				ExpirationYear:  "2024",
			},
			expectedResp: &validationResponse{
				IsValid:   false,
				ErrorCode: errorcodes.ExpirationDateInvalidFormat,
			},
		},
		{
			description: "invalid expiry date format - invalid year",
			card: &proto.Card{
				CardNumber:      "411111111111111",
				ExpirationMonth: "27",
				ExpirationYear:  "2U2A",
			},
			expectedResp: &validationResponse{
				IsValid:   false,
				ErrorCode: errorcodes.ExpirationDateInvalidFormat,
			},
		},
		{
			description: "some unexpected validator error ",
			card: &proto.Card{
				CardNumber:      "4111111111111111",
				ExpirationMonth: "12",
				ExpirationYear:  "2024",
			},
			validatorVerdict: fmt.Errorf("some validation error"),
			expectedError:    "some validation error",
			expectedResp:     nil,
		},
		{
			description: "lugn algorithm validator failed",
			card: &proto.Card{
				CardNumber:      "4111111111111111",
				ExpirationMonth: "12",
				ExpirationYear:  "2024",
			},
			validatorVerdict: errorcodes.ErrCardNumberLuhnFailed,
			expectedResp: &validationResponse{
				IsValid:   false,
				ErrorCode: errorcodes.CardNumberLuhnFailed,
			},
		},
		{
			description: "expired card",
			card: &proto.Card{
				CardNumber:      "4111111111111111",
				ExpirationMonth: "12",
				ExpirationYear:  "2024",
			},
			validatorVerdict: errorcodes.ErrExpired,
			expectedResp: &validationResponse{
				IsValid:   false,
				ErrorCode: errorcodes.Expired,
			},
		},
		{
			description: "some other expiration date validation error",
			card: &proto.Card{
				CardNumber:      "4111111111111111",
				ExpirationMonth: "12",
				ExpirationYear:  "2024",
			},
			validatorVerdict: errorcodes.ErrExpirationDate,
			expectedResp: &validationResponse{
				IsValid:   false,
				ErrorCode: errorcodes.ExpirationDateOther,
			},
		},
		{
			description: "invalid IIN for card number",
			card: &proto.Card{
				CardNumber:      "4111111111111111",
				ExpirationMonth: "12",
				ExpirationYear:  "2024",
			},
			validatorVerdict: errorcodes.ErrInvalidCardNumberIIN,
			expectedResp: &validationResponse{
				IsValid:   false,
				ErrorCode: errorcodes.CardNumberInvalidIIN,
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			valr := mock_validator.NewMockValidator(ctrl)
			valr.EXPECT().Validate(gomock.Any(), gomock.Any()).Return(tc.validatorVerdict).MaxTimes(1)

			svc := server.New(server.WithValidator(valr))

			resp, err := svc.ValidateCard(context.Background(), tc.card)

			if tc.expectedError != "" {
				require.Error(t, err)
				require.Contains(t, err.Error(), tc.expectedError)
			} else {
				require.NoError(t, err)
			}

			if tc.expectedResp != nil {
				require.NotNil(t, resp)
				require.Equal(t, tc.expectedResp.IsValid, resp.IsValid)
				if tc.expectedResp.IsValid {
					require.Nil(t, resp.Error)
				} else {
					require.NotNil(t, resp.Error)
					require.Equal(t, tc.expectedResp.ErrorCode, resp.Error.Code)
				}

			} else {
				require.Nil(t, resp)
			}
		})
	}
}
