package utils

import (
	"fmt"
	"testing"
)

func TestSticker(t *testing.T) {

	cases := []struct {
		expected Sticker
		input    []string
	}{
		{
			expected: Sticker{
				Name:       "Gold Web",
				Tournament: "none",
				Condition:  "Foil",
			},
			input: []string{"Sticker", "Gold Web (Foil) ", "(Foil)"},
		},
		{
			expected: Sticker{
				Name:       "Gold Web",
				Tournament: "none",
				Condition:  "Foil",
			},
			input: []string{"Sticker", "Gold Web (Foil) ", "(Foil)"},
		},
		{
			expected: Sticker{
				Name:       "Gold Web",
				Tournament: "none",
				Condition:  "Foil",
			},
			input: []string{"Sticker", "Gold Web (Foil)", "(Foil)"},
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			ParseSticker(c.input)

		})

	}
}
