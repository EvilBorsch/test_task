package calculator

import (
	"context"
	"test_task/src/calculator/models"
)

type Usecase interface {
	Calculate(c context.Context, inputStruct models.InputStruct) models.ResultStruct
}
