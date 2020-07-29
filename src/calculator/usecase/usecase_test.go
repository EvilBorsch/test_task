package usecase

import (
	"context"
	"github.com/stretchr/testify/assert"
	"test_task/src/calculator/models"
	"testing"
)

func TestCalculation(t *testing.T) {
	type readInputTestCase struct {
		input  models.InputStruct
		output models.ResultStruct
	}
	usecase := NewCalculatorUsecase(10)
	ctx := context.Background()
	testCases := []readInputTestCase{
		{
			input: models.InputStruct{
				A:         2,
				B:         3,
				Operation: "+",
			},
			output: models.ResultStruct{
				Success: true,
				ErrCode: "",
				Value:   5,
			},
		},
		{
			input: models.InputStruct{
				A:         2.0,
				B:         3.1,
				Operation: "*",
			},
			output: models.ResultStruct{
				Success: true,
				ErrCode: "",
				Value:   6.2,
			},
		},
		{
			input: models.InputStruct{
				A:         -20,
				B:         -5,
				Operation: "-",
			},
			output: models.ResultStruct{
				Success: true,
				ErrCode: "",
				Value:   -15,
			},
		},
		{
			input: models.InputStruct{
				A:         30,
				B:         0,
				Operation: "/",
			},
			output: models.ResultStruct{
				Success: false,
				ErrCode: "can`t divide by 0",
				Value:   0,
			},
		},
	}

	for _, testCase := range testCases {
		output := usecase.Calculate(ctx, testCase.input)
		assert.Equal(t, output, testCase.output)
	}
}
