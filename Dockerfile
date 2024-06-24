FROM golang:1.22 as builder

WORKDIR /workspace

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64  go build -o /workspace/bin/card-validator-server ./cmd/server

FROM gcr.io/distroless/static:nonroot as runtime
WORKDIR /
COPY --from=builder /workspace/bin/card-validator-server /usr/local/bin/card-validator-server

ENTRYPOINT ["card-validator-server"]