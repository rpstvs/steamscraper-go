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
				Condition:  "none",
			},
			input: []string{"Sticker", "Gold Web"},
		},
		{
			expected: Sticker{
				Name:       "Gold Web",
				Tournament: "none",
				Condition:  "Foil",
			},
			input: []string{"Sticker", "Gold Web (Foil) "},
		},
		{
			expected: Sticker{
				Name:       "Gold Web",
				Tournament: "none",
				Condition:  "Foil",
			},
			input: []string{"Sticker", "Gold Web (Foil)"},
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			actual := ParseSticker(c.input)
			fmt.Println(actual)
			if actual != c.expected {
				t.Errorf("parsing failed")
			}

		})

	}
}

func TestSkin(t *testing.T) {

	cases := []struct {
		expected Skin
		input    []string
	}{
		{
			expected: Skin{
				GunName:   "CZ-75",
				SkinName:  "Gold Web",
				Condition: "field-tested",
			},
			input: []string{"CZ-75", "Gold Web (field-tested)"},
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			actual := parseSkin(c.input)
			fmt.Println(actual)
			if actual != c.expected {
				t.Errorf("parsing failed")
			}

		})

	}
}

func TestPrice(t *testing.T) {

	cases := []struct {
		expected float64
		input    string
	}{
		{
			expected: 0.03,
			input:    "3",
		},
		{
			expected: 1,
			input:    "100",
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			actual := priceConverter(c.input)
			fmt.Println(actual)
			if actual != c.expected {
				t.Errorf("parsing failed")
			}

		})

	}
}
