package card_test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ikaliuzh/card-validator/gen/proto"
	"github.com/ikaliuzh/card-validator/pkg/card"
	"github.com/ikaliuzh/card-validator/pkg/errorcodes"
)

func TestNewCardFromProto(t *testing.T) {
	testCases := []struct {
		description   string
		protoCard     *proto.Card
		expectedCard  *card.Card
		expectedError string
	}{
		{
			description: "valid card",
			protoCard: &proto.Card{
				CardNumber:      "4111111111111111",
				ExpirationMonth: "12",
				ExpirationYear:  "2024",
			},
			expectedCard: &card.Card{
				Number:          card.CardNumber{4, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
				ExpirationMonth: 12,
				ExpirationYear:  2024,
			},
		},
		{
			description: "valid card, expired",
			protoCard: &proto.Card{
				CardNumber:      "4111111111111111",
				ExpirationMonth: "12",
				ExpirationYear:  "1999",
			},
			expectedCard: &card.Card{
				Number:          card.CardNumber{4, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
				ExpirationMonth: 12,
				ExpirationYear:  1999,
			},
		},
		{
			description: "invalid card number containing letters",
			protoCard: &proto.Card{
				CardNumber:      "411111111111111a",
				ExpirationMonth: "12",
				ExpirationYear:  "2024",
			},
			expectedError: "invalid card number format: card number must contain only digits, got 'a'",
		},
		{
			description: "invalid card number length: too short",
			protoCard: &proto.Card{CardNumber: "4111",
				ExpirationMonth: "12",
				ExpirationYear:  "2024",
			},
			expectedError: "invalid card number format: card number length should be at least 8 digits: got 4",
		},
		{
			description: "invalid card number length: too long",
			protoCard: &proto.Card{
				CardNumber:      "4111222222222222222222222",
				ExpirationMonth: "12",
				ExpirationYear:  "2024",
			},
			expectedError: "invalid card number format: card number length should be at most 19 digits: got 25",
		},
		{
			description: "invalid expiration month: not a number",
			protoCard: &proto.Card{
				CardNumber:      "4111111111111111",
				ExpirationMonth: "12a",
				ExpirationYear:  "2024",
			},
			expectedError: "invalid expiration date format: expiration month is invalid: strconv.Atoi: parsing \"12a\": invalid syntax",
		},
		{
			description: "invalid expiration month: out of range",
			protoCard: &proto.Card{
				CardNumber:      "4111111111111111",
				ExpirationMonth: "13",
				ExpirationYear:  "2024",
			},
			expectedError: "invalid expiration date format: expiration month is invalid: got 13",
		},
		{
			description: "invalid expiration year: not a number",
			protoCard: &proto.Card{
				CardNumber:      "4111111111111111",
				ExpirationMonth: "12",
				ExpirationYear:  "2024a",
			},
			expectedError: "invalid expiration date format: expiration year is invalid: strconv.Atoi: parsing \"2024a\": invalid syntax",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			c, err := card.NewCardFromProto(tc.protoCard)

			if tc.expectedError != "" {
				require.EqualError(t, err, tc.expectedError)
				if strings.Contains(tc.expectedError, "invalid c number format") {
					require.ErrorIs(t, err, errorcodes.ErrInvalidCardNumberFormat)
				}
				if strings.Contains(tc.expectedError, "invalid expiration date format") {
					require.ErrorIs(t, err, errorcodes.ErrExpirationDateInvalidFormat)
				}
			}

			if tc.expectedCard != nil {
				require.Equal(t, tc.expectedCard.Number, c.Number)
				require.Equal(t, tc.expectedCard.ExpirationMonth, c.ExpirationMonth)
				require.Equal(t, tc.expectedCard.ExpirationYear, c.ExpirationYear)
			}
		})
	}
}
