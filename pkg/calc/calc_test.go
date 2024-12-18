package calc_test

import (
	"fmt"
	"testing"
	"github.com/vizurth/calc_go/pkg/calc"
)

func TestCalc(t *testing.T){
	testCasesSuccess := []struct {
		name           string
		expression     string
		expectedResult float64
	}{
		{
			name:           "simple",
			expression:     "1+1",
			expectedResult: 2,
		},
		{
			name:           "priority",
			expression:     "(2+2)*2",
			expectedResult: 8,
		},
		{
			name:           "priority",
			expression:     "2+2*2",
			expectedResult: 6,
		},
		{
			name:           "/",
			expression:     "1/2",
			expectedResult: 0.5,
		},
	}

	for _, testCase := range testCasesSuccess {
		t.Run(testCase.name, func(t *testing.T) {
			val, err := calc.Calc(testCase.expression)
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
			name:       "simple",
			expression: "1+1*",
			expectedErr: fmt.Errorf("Expression is not valid"),
		},
		{
			name:       "priority",
			expression: "1+1*)",
			expectedErr: fmt.Errorf("Expression is not valid"),
		},
		{
			name:       "priority",
			expression: "1+1(*)",
			expectedErr: fmt.Errorf("Expression is not valid"),
		},
	}
	
	for _, testCase := range testCasesFail{
		t.Run(testCase.name, func(t *testing.T){
			_, err := calc.Calc(testCase.expression)
			if err != fmt.Errorf("Expression is not valid"){
				t.Fatalf("dont needed error", testCase.expression)
			}
		})
	}
}