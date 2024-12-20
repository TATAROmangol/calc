package server

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"example.com/m/errors"
)

func TestCalcHandlerStatus200(t *testing.T){
	test_cases := []struct{
		inp Input
		res OkResult
	}{
		{Input{"2+2"},OkResult{4.0}},
		{Input{"2+2*2"},OkResult{6.0}},
		{Input{"2/2+2*4"},OkResult{9.0}},
		{Input{"1+(1+1)*2/2"},OkResult{3.0}},
		{Input{"2*((1+1)/(3-2*1))+1"},OkResult{5.0}},
	}
	for _, tCase := range test_cases{
		body, _ := json.Marshal(tCase.inp)
		r := httptest.NewRequest(http.MethodPost, "/api/v1/calculate", bytes.NewBuffer(body))
		w := httptest.NewRecorder()

		calcHandler(w,r)

		res := w.Result()
		defer res.Body.Close()

		if res.StatusCode != 200{
			t.Errorf("For input %v, wrong status code", tCase.inp)
		}

		data, err := io.ReadAll(res.Body)
		if err != nil {
			t.Errorf("For input %v, error: %v", tCase.inp, err)
		}

		var gotRes OkResult
        if err := json.Unmarshal(data, &gotRes); err != nil {
            t.Errorf("For input %v, error unmarshalling response: %v", tCase.inp, err)
        }

        if gotRes.Result != tCase.res.Result {
            t.Errorf("For input %v, expected result %v, got %v", tCase.inp, tCase.res.Result, gotRes.Result)
        }
	}
}

func TestCalcHandlerStatus402(t *testing.T){
	test_cases := []struct{
		inp Input
		res ErrResult
	}{
		{Input{"2/"}, ErrResult{errors.ErrUnknownFormat.Error()}},
		{Input{"2*"}, ErrResult{errors.ErrUnknownFormat.Error()}},
		{Input{"2+"}, ErrResult{errors.ErrUnknownFormat.Error()}},
		{Input{"2-"}, ErrResult{errors.ErrUnknownFormat.Error()}},
		{Input{"/2"}, ErrResult{errors.ErrUnknownFormat.Error()}},
		{Input{"*2"}, ErrResult{errors.ErrUnknownFormat.Error()}},
		{Input{"+2"}, ErrResult{errors.ErrUnknownFormat.Error()}},
		{Input{"-2"}, ErrResult{errors.ErrUnknownFormat.Error()}},
		{Input{"2--2"}, ErrResult{errors.ErrUnknownFormat.Error()}},
		{Input{"(2+2))"}, ErrResult{errors.ErrNotOpenQuot.Error()}},
		{Input{"((2+2)"}, ErrResult{errors.ErrNotCloseQuot.Error()}},
		{Input{"1+(2/0+1)"}, ErrResult{errors.ErrDivisionByZero.Error()}},

	}
	for _, tCase := range test_cases{
		body, _ := json.Marshal(tCase.inp)
		r := httptest.NewRequest(http.MethodPost, "/api/v1/calculate", bytes.NewBuffer(body))
		w := httptest.NewRecorder()

		calcHandler(w,r)

		res := w.Result()
		defer res.Body.Close()

		if res.StatusCode != 402{
			t.Errorf("For input %v, wrong status code", tCase.inp)
		}

		data, err := io.ReadAll(res.Body)
		if err != nil {
			t.Errorf("For input %v, error: %v", tCase.inp, err)
		}

		var gotRes ErrResult
        if err := json.Unmarshal(data, &gotRes); err != nil {
            t.Errorf("For input %v, error unmarshalling response: %v", tCase.inp, err)
        }

        if gotRes.Err != tCase.res.Err {
            t.Errorf("For input %v, expected result %v, got %v", tCase.inp, tCase.res.Err, gotRes.Err)
        }
	}
}