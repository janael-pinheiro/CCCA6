package refectoring

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	name  string
	input string
	want  bool
}

func TestGivenInvalidCPFReturnFalse(t *testing.T) {
	var testCases = []testCase{
		{
			name:  "Invalid CPF with special characters",
			input: "111.444.777-05",
			want:  false,
		},
		{
			name:  "Invalid CPF without special characters",
			input: "11144477705",
			want:  false,
		},
	}
	for _, test := range testCases {
		isCPFValid := IsValid(test.input)
		assert.Equal(t, isCPFValid, test.want, test.name)
	}

}

func TestGivenEmptyStringReturnFalse(t *testing.T) {
	fake_cpf := ""
	isCPFValid := IsValid(fake_cpf)
	var expected bool = false
	assert.Equal(t, isCPFValid, expected)
}

func TestGivenCPFWithInvalidNumberCharactersReturnFalse(t *testing.T) {
	var testCases = []testCase{
		{
			name:  "CPF with less than 11 digits",
			input: "111.444.777-3",
			want:  false,
		},
		{
			name:  "CPF with less than 11 digits",
			input: "111",
			want:  false,
		},
		{
			name:  "CPF with more than 15 digits",
			input: "111.444.777358945",
			want:  false,
		},
	}
	for _, test := range testCases {
		isCPFValid := IsValid(test.input)
		assert.Equal(t, isCPFValid, test.want, test.name)
	}
}
func TestGivenValidCPFReturnTrue(t *testing.T) {
	var testCases = []testCase{
		{
			name:  "Valid CPF with special characters",
			input: "111.444.777-35",
			want:  true,
		},
		{
			name:  "Valid CPF without special characters",
			input: "11144477735",
			want:  true,
		},
		{
			name:  "Valid CPF with points",
			input: "111.444.77735",
			want:  true,
		},
		{
			name:  "Valid CPF with hyphen",
			input: "111444777-35",
			want:  true,
		},
	}
	for _, test := range testCases {
		isCPFValid := IsValid(test.input)
		assert.Equal(t, isCPFValid, test.want, test.name)
	}

}

func TestGivenCPFWithAllSameDigitsReturnFalse(t *testing.T) {
	var testCases = []testCase{
		{
			name:  "Invalid CPF with all digits equal to 1",
			input: "111.111.111-11",
			want:  false,
		},
		{
			name:  "Invalid CPF with all digits equal to 9",
			input: "999.999.999-99",
			want:  false,
		},
	}
	for _, test := range testCases {
		isCPFValid := IsValid(test.input)
		assert.Equal(t, isCPFValid, test.want, test.name)
	}

}
