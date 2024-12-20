package server

import (
	"example.com/m/errors"
)

type Input struct{
	Expression string `json:"expression"`
}

type OkResult struct{
	Result float64 `json:"result"`
}

type ErrResult struct{
	Err string `json:"error"`
}
var (
	error500 = ErrResult{errors.ErrUnknownError.Error()}
)