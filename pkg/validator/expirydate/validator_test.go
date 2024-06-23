package expirydate

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/ikaliuzh/card-validator/pkg/card"
	"github.com/ikaliuzh/card-validator/pkg/validator/expirydate/mocks"
)

func TestValidator(t *testing.T) {
	testCases := []struct {
		description  string
		card         card.Card
		currentMonth int
		currentYear  int
		expectedErr  string
	}{
		{
			description: "valid card, expires in 6 month this year",
			card: card.Card{
				ExpirationMonth: 12,
				ExpirationYear:  2024,
			},
			currentMonth: 6,
			currentYear:  2024,
		},
		{
			description: "valid card, expires next year",
			card: card.Card{
				ExpirationMonth: 12,
				ExpirationYear:  2025,
			},
			currentMonth: 6,
			currentYear:  2024,
		},
		{
			description: "invalid card, expired this year",
			card: card.Card{
				ExpirationMonth: 3,
				ExpirationYear:  2024,
			},
			currentMonth: 6,
			currentYear:  2024,
			expectedErr:  "card is expired: got 3/2024",
		},
		{
			description: "invalid card, expired last year",
			card: card.Card{
				ExpirationMonth: 12,
				ExpirationYear:  2023,
			},
			currentMonth: 6,
			currentYear:  2024,
			expectedErr:  "card is expired: got 12/2023",
		},
		{
			description: "invalid card, expired last year - 2",
			card: card.Card{
				ExpirationMonth: 3,
				ExpirationYear:  2023,
			},
			currentMonth: 6,
			currentYear:  2024,
			expectedErr:  "card is expired: got 3/2023",
		},
		{
			description: "valid card, expires this month",
			card: card.Card{
				ExpirationMonth: 6,
				ExpirationYear:  2024,
			},
			currentMonth: 6,
			currentYear:  2024,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			mockDateProvider := mocks.NewMockDateProvider(ctrl)
			mockDateProvider.EXPECT().CurrentMonth().Return(tc.currentMonth).Times(1)
			mockDateProvider.EXPECT().CurrentYear().Return(tc.currentYear).Times(1)

			v := Validator{DateProvider: mockDateProvider}

			err := v.Validate(context.TODO(), tc.card)
			if tc.expectedErr != "" {
				require.Error(t, err)
				require.EqualError(t, err, tc.expectedErr)
				require.ErrorIs(t, err, ErrExpired)

			} else {
				require.NoError(t, err)
			}
		})
	}
}
