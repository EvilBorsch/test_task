package models

type InputStruct struct {
	A         float64
	B         float64
	Operation string
}

type ResultStruct struct {
	Success bool
	ErrCode string
	Value   float64
}
