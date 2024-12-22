package calculation_test

import (
	"testing"

	"github.com/LeonidSelivanov/Yandex-Calculator-Service/pkg/calculation"
)

func TestCalc(t *testing.T) {
	testCasesSuccess := []struct {
		name           string
		expression     string
		expectedResult float64
	}{
		{
			name:           "addition",
			expression:     "5+5",
			expectedResult: 10,
		},
		{
			name:           "subtraction",
			expression:     "5-5",
			expectedResult: 0,
		},
		{
			name:           "multiplication",
			expression:     "5*5",
			expectedResult: 25,
		},
		{
			name:           "division",
			expression:     "5/5",
			expectedResult: 1,
		},
		{
			name:           "priority",
			expression:     "(5+5)+(5+5)",
			expectedResult: 20,
		},
		{
			name:           "priority",
			expression:     "5*5+5*5",
			expectedResult: 50,
		},
	}

	for _, testCase := range testCasesSuccess {
		t.Run(testCase.name, func(t *testing.T) {
			val, err := calculation.Calc(testCase.expression)
			if err != nil {
				t.Fatalf("successful case %s returns error", testCase.expression)
			}
			if val != testCase.expectedResult {
				t.Fatalf("%f should be equal %f", val, testCase.expectedResult)
			}
		})
	}

	testCasesFail := []struct {
		name        string
		expression  string
		expectedErr error
	}{
		{
			name:       "division by zero",
			expression: "5/0",
		},
		{
			name:       "invalid expression",
			expression: "5*5++5*5",
		},
	}

	for _, testCase := range testCasesFail {
		t.Run(testCase.name, func(t *testing.T) {
			val, err := calculation.Calc(testCase.expression)
			if err == nil {
				t.Fatalf("expression %s is invalid but result  %f was obtained", testCase.expression, val)
			}
		})
	}
}
