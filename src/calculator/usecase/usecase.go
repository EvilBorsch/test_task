package usecase

import (
	"context"
	"test_task/src/calculator/models"
	"time"
)

type CalculatorUsecase struct {
	contextTimeout time.Duration
}

func NewCalculatorUsecase(timeout time.Duration) *CalculatorUsecase {
	return &CalculatorUsecase{contextTimeout: timeout}
}

func (cu *CalculatorUsecase) Calculate(c context.Context, params models.InputStruct) models.ResultStruct {
	_, cancel := context.WithTimeout(c, cu.contextTimeout)
	defer cancel()
	res := models.ResultStruct{
		Success: true,
		ErrCode: "",
		Value:   0,
	}
	switch params.Operation {
	case "+":
		res.Value = params.A + params.B
	case "-":
		res.Value = params.A - params.B
	case "*":
		res.Value = params.A * params.B
	case "/":
		if params.B == 0 {
			return models.ResultStruct{
				Success: false,
				ErrCode: "can`t divide by 0",
				Value:   0,
			}
		}
		res.Value = params.A / params.B

	}

	return res
}
