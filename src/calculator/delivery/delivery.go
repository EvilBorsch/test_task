package delivery

import (
	"errors"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"test_task/src/calculator"
	"test_task/src/calculator/models"
	"test_task/src/utills"
)

type CalculatorHandler struct {
	CUsecase calculator.Usecase
}

func NewCalculatorHandler(r *mux.Router, us calculator.Usecase) {
	handler := CalculatorHandler{CUsecase: us}
	r.HandleFunc("/api/add", handler.AddHandler).Methods("GET")
	r.HandleFunc("/api/div", handler.DivHandler).Methods("GET")
	r.HandleFunc("/api/mul", handler.MulHandler).Methods("GET")
	r.HandleFunc("/api/sub", handler.SubHandler).Methods("GET")
}

func (c *CalculatorHandler) AddHandler(w http.ResponseWriter, r *http.Request) {
	params, err := fetchParams(r)
	if err != nil {
		utills.SendServerError(err.Error(), 400, w)
		return
	}
	params.Operation = "+"
	res := c.CUsecase.Calculate(r.Context(), params)
	utills.SendOKAnswer(res, w)
}

func (c *CalculatorHandler) SubHandler(w http.ResponseWriter, r *http.Request) {
	params, err := fetchParams(r)
	if err != nil {
		utills.SendServerError(err.Error(), 400, w)
		return
	}
	params.Operation = "-"
	res := c.CUsecase.Calculate(r.Context(), params)
	utills.SendOKAnswer(res, w)
}

func (c *CalculatorHandler) MulHandler(w http.ResponseWriter, r *http.Request) {
	params, err := fetchParams(r)
	if err != nil {
		utills.SendServerError(err.Error(), 400, w)
		return
	}
	params.Operation = "*"
	res := c.CUsecase.Calculate(r.Context(), params)
	utills.SendOKAnswer(res, w)
}

func (c *CalculatorHandler) DivHandler(w http.ResponseWriter, r *http.Request) {
	params, err := fetchParams(r)
	if err != nil {
		utills.SendServerError(err.Error(), 400, w)
		return
	}
	params.Operation = "/"
	res := c.CUsecase.Calculate(r.Context(), params)
	utills.SendOKAnswer(res, w)
}

func fetchParams(r *http.Request) (models.InputStruct, error) {
	st := models.InputStruct{}
	err := r.ParseForm()
	if err != nil {
		return models.InputStruct{}, errors.New("err when parsing form")
	}
	if len(r.Form)!=2 {
		return models.InputStruct{}, errors.New("params count must be 2")
	}
	val, err := strconv.ParseFloat(r.Form.Get("a"), 64)
	if err != nil {
		return models.InputStruct{}, errors.New("a is not number or not exist")
	}
	st.A = val
	val, err = strconv.ParseFloat(r.Form.Get("b"), 64)
	if err != nil {
		return models.InputStruct{}, errors.New("b is not number or not exist")
	}
	st.B = val
	return st, nil
}
