package delivery

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"test_task/src/calculator/models"
	usecase2 "test_task/src/calculator/usecase"
	"testing"
)

type testCaseInput struct {
	A   string
	B   string
	AdditionalParams string
	url string
}

type TestCase struct {
	Input  testCaseInput
	Result models.ResultStruct
}

func TestCalculationHandlers(t *testing.T) {

	usecase := usecase2.NewCalculatorUsecase(2)
	handler := CalculatorHandler{CUsecase: usecase}

	testCases := []TestCase{
		{
			Input: testCaseInput{
				A:   "100",
				B:   "20",
				url: "div",
			},
			Result: models.ResultStruct{
				Success: true,
				ErrCode: "",
				Value:   5,
			},
		},
		{
			Input: testCaseInput{
				A:   "10.5",
				B:   "-3",
				url: "sub",
			},
			Result: models.ResultStruct{
				Success: true,
				ErrCode: "",
				Value:   13.5,
			},
		},
		{
			Input: testCaseInput{
				A:   "10.5",
				B:   "3",
				url: "add",
			},
			Result: models.ResultStruct{
				Success: true,
				ErrCode: "",
				Value:   13.5,
			},
		},
		{
			Input: testCaseInput{
				A:   "10.5",
				B:   "3",
				url: "mul",
			},
			Result: models.ResultStruct{
				Success: true,
				ErrCode: "",
				Value:   31.5,
			},
		},
		//not ok
		{
			Input: testCaseInput{
				A:   "not valid",
				B:   "-3",
				url: "mul",
			},
			Result: models.ResultStruct{
				Success: false,
				ErrCode: "a is not number or not exist",
				Value:   0,
			},
		},
		//not ok
		{
			Input: testCaseInput{
				A:   "1",
				B:   "not valid",
				url: "sub",
			},
			Result: models.ResultStruct{
				Success: false,
				ErrCode: "b is not number or not exist",
				Value:   0,
			},
		},
		//not ok
		{
			Input: testCaseInput{
				A:   "1",
				B:   "-3",
				AdditionalParams: "&d=8",
				url: "sub",
			},
			Result: models.ResultStruct{
				Success: false,
				ErrCode: "params count must be 2",
				Value:   0,
			},
		},

	}

	for _, testCase := range testCases {
		url := fmt.Sprintf(`/api/%s?a=%s&b=%s%s`, testCase.Input.url, testCase.Input.A, testCase.Input.B,testCase.Input.AdditionalParams)

		req, err := http.NewRequest("GET", url, nil)

		assert.Nil(t, err)
		respWriter := httptest.NewRecorder()
		switch testCase.Input.url {
		case "add":
			handler.AddHandler(respWriter, req)
		case "div":
			handler.DivHandler(respWriter, req)
		case "mul":
			handler.MulHandler(respWriter, req)
		case "sub":
			handler.SubHandler(respWriter, req)

		}

		resp := respWriter.Result()
		body, err := ioutil.ReadAll(resp.Body)
		assert.NoError(t, err)
		var responseStruct models.ResultStruct
		err = json.Unmarshal(body, &responseStruct)
		assert.Nil(t, err)
		assert.Equal(t, responseStruct, testCase.Result)
	}
}
