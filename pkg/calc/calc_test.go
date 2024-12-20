package calc

import (
	"testing"

	"example.com/m/errors"
)

func TestCalculation(t *testing.T) {
	test_cases := []struct {
		numbers  []float64
		operands []byte
		res      float64
		err      error
	}{
		{[]float64{2, 2}, []byte{bPlus}, 4, nil},
		{[]float64{2, 2}, []byte{bMult}, 4, nil},
		{[]float64{2, 2, 3, 4}, []byte{bDiv, bMult, bPlus}, 7, nil},
		{[]float64{1, 2, 3, 5, 5}, []byte{bPlus, bMult, bMinus, bDiv}, 6, nil},
	}

	for _, tCase := range test_cases {
		eRes, eErr := calculation(tCase.numbers, tCase.operands)
		if eErr != tCase.err || eRes != tCase.res {
			t.Errorf("Expected res:%v and err:%v, got res:%v and err:%v", tCase.res, tCase.err, eRes, eErr)
		}
	}
}

func TestParse(t *testing.T) {
	test_cases := []struct {
		inp      string
		operands []float64
		numbers  []byte
		err      error
	}{
		{"2/2", []float64{2, 2}, []byte{bDiv}, nil},
		{"2+3*2", []float64{2, 3, 2}, []byte{bPlus, bMult}, nil},
		{"2++2", []float64{}, []byte{}, errors.ErrUnknownFormat},
		{"+2+2", []float64{}, []byte{}, errors.ErrUnknownFormat},
		{"2+2+", []float64{}, []byte{}, errors.ErrUnknownFormat},
	}

	for _, tCase := range test_cases {
		exO, exN, exE := parse(tCase.inp)
		if exE != tCase.err {
			t.Errorf("For input %q, expected err %v, got %v", tCase.inp, tCase.err, exE)
		}
		if !equalFloat64Slices(exO, tCase.operands) {
			t.Errorf("For input %q, expected operands %v, got %v", tCase.inp, tCase.operands, exO)
		}

		// Проверка чисел
		if !equalByteSlices(exN, tCase.numbers) {
			t.Errorf("For input %q, expected numbers %v, got %v", tCase.inp, tCase.numbers, exN)
		}
	}
}

func equalFloat64Slices(a, b []float64) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func equalByteSlices(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestCalc(t *testing.T) {
	test_cases := []struct {
		inp string
		res float64
		err error
	}{
		{"2/2", 1.0, nil},
		{"2+3*2", 8.0, nil},
	}

	for _, tCase := range test_cases {
		eRes, eErr := Calc(tCase.inp)
		if eErr != tCase.err || eRes != tCase.res {
			t.Errorf("Expected res:%v and err:%v, got res:%v and err:%v", tCase.res, tCase.err, eRes, eErr)
		}
	}
}

func TestCalcQuots(t *testing.T) {
	test_cases := []struct {
		inp string
		res string
		err error
	}{
		{"1+(1+1)+1", "1+2+1", nil},
		{"2*((1+1)/(3-2*1))+1", "2*2+1", nil},
		{"2*((1+1)/(3-2*1))+1+((1)", "", errors.ErrNotCloseQuot},
		{"2+(1))", "", errors.ErrNotOpenQuot},
	}

	for _, tCase := range test_cases {
		res := tCase.inp
		exErr := сalcQuots(&res)
		if tCase.err != exErr {
			t.Errorf("For input %v, expected err %v, got err:%v", tCase.inp, tCase.err, exErr)
		} else if tCase.err == nil && res != tCase.res {
			t.Errorf("For input %v, expected res: %v, got res: %v", tCase.inp, tCase.res, res)
		}
	}
}
