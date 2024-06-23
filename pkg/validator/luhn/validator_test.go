package luhn

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/ikaliuzh/card-validator/pkg/card"
)

func TestLuhnValidator(t *testing.T) {
	testCases := []struct {
		description string
		card        card.Card
		expectedErr string
	}{
		{
			description: "valid card number",
			card:        card.Card{Number: card.MustGetNumberFromString("4111111111111111")},
		},
		{
			description: "valid card number - 2",
			card:        card.Card{Number: card.MustGetNumberFromString("4485157033981857")},
		},
		{
			description: "valid card number - 3",
			card:        card.Card{Number: card.MustGetNumberFromString("5389984418273855")},
		},
		{
			description: "valid card number - 4",
			card:        card.Card{Number: card.MustGetNumberFromString("4016256492036126")},
		},
		{
			description: "valid card number - 5",
			card:        card.Card{Number: card.MustGetNumberFromString("342342553112031")},
		},
		{
			description: "invalid card number",
			card:        card.Card{Number: card.MustGetNumberFromString("342342553112030")},
			expectedErr: "invalid card number: card number failed Luhn algorithm validation",
		},
		{
			description: "invalid card number - 2",
			card:        card.Card{Number: card.MustGetNumberFromString("4976030843039062")},
			expectedErr: "invalid card number: card number failed Luhn algorithm validation",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			v := Validator{}

			err := v.Validate(context.TODO(), tc.card)
			if tc.expectedErr != "" {
				require.Error(t, err)
				require.EqualError(t, err, tc.expectedErr)
				require.ErrorIs(t, err, ErrInvalidCardNumber)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestLuhnAlgorithm(t *testing.T) {
	testCases := []struct {
		number     string
		checkDigit int
	}{
		{number: "1", checkDigit: 8},  // 10 - (1 * 2) = 8
		{number: "2", checkDigit: 6},  // 10 - (2 * 2) = 6
		{number: "12", checkDigit: 5}, // 10 - (2 * 2 + 1) = 5
		{number: "132", checkDigit: 1},
		{number: "455618712206700", checkDigit: 7},
		{number: "402400717695169", checkDigit: 6},
		{number: "448592961601501", checkDigit: 8},
		{number: "518589827971647", checkDigit: 9},
		{number: "523502969404999", checkDigit: 5},
		{number: "34235196127130", checkDigit: 5},
		{number: "37474918503482", checkDigit: 4},
		{number: "601161520314725", checkDigit: 0},
		{number: "601184452024446", checkDigit: 3},
	}
	for idx, tc := range testCases {
		t.Run(fmt.Sprintf("Test case %d", idx+1), func(t *testing.T) {
			cn, _ := card.NumberFromString(tc.number)
			checkDigit := luhnAlgorithm(cn)
			require.Equal(t, tc.checkDigit, checkDigit)
		})
	}
}

func cardNumberFromString(s string) card.CardNumber {
	digits := make([]int, len(s))
	for i, r := range s {
		digits[i] = int(r - '0')
	}
	return digits
}
