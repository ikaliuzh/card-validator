# Card Validator - an extensible gRPC microservice for validating credit cards.

This is a simple gRPC microservice that validates credit card numbers. It is written in Golang and uses the 
[gRPC](https://grpc.io/) framework. The service is extensible and can be easily modified to add new validation rules.

## API

The service exposes a single RPC method `ValidateCard` which takes a `Card` message as input and returns a 
`CardValidationResponse` message as output.

```protobuf
service CardValidator {
    rpc ValidateCard(Card) returns (CardValidationResponse);
}
```

The `Card` message is defined as follows:

```protobuf
message Card {
  string card_number = 1;
  string expiration_month = 2;
  string expiration_year = 3;
}
```

The service return the validation respons with `is_valid` field and a possible error. Th error is returned only if the
validation failed.

```protobuf
message CardValidationResponse {
  bool is_valid = 1;
  optional Error error = 2;
}

message Error {
  string code = 1;
  string message = 2;
}
```

### Error Codes

These are the possible error codes that can be returned in the `Error` message:

| Code | Name                        | Description                                                                                                   |
|------|-----------------------------|---------------------------------------------------------------------------------------------------------------|
| 101  | ExpirationDateInvalidFormat | Expiration month or expiration year does not specify a valid date.                                            |
| 102  | Expired                     | The card is expired.                                                                                          |
| 199  | ExpirationDateOther         | Other expiration date validation error.                                                                       |
| 201  | CardNumberInvalidFormat     | Card number format is invalid. It must be a string of length between 8 and 19 symbols containing only digits. |
| 202  | CardNumberLuhnFailed        | Card number failed Luhn algorithm verification, the checksum digit is not correct.                            |
| 203  | CardNumberInvalidIIN        | The issuer identification number is not valid.                                                                |
| 299  | CardNumberOther             | Other card number validation error.                                                                           |
| 999  | Other                       | Other card validation error.                                                                                  |

## Usage

This project uses taskfiles to manage the build and run tasks. To run the card validator service locally use
```shell
task run
```
To run the project in a docker container use
```shell
task docker-build docker-run
```
To run an interactive gRPC client locally use [evans](https://github.com/ktr0731/evans)
```shell
task evans
```