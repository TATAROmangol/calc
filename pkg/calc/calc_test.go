package calc

import (
	"testing"

	"example.com/m/errors"
)

var calc = NewCalculator()

// func TestCalculation(t *testing.T){
// 	test_cases := []struct{
// 		inp string
// 		res float64
// 	}{
// 		{"2/2", 1},
// 		{""}
// 	}
// }

func TestParse(t *testing.T){
	test_cases := []struct{
		inp string
		operands []float64
		numbers []byte
		err error 
	}{
		{"2/2", []float64{2,2}, []byte{bDiv}, nil},
		{"2*2+3*(123-231)", []float64{2,2,3,123,231}, []byte{bMult, bPlus, bMult,bOQuot,bMinus,bCQuot}, nil},
		{"2\\", []float64{}, []byte{}, errors.ErrUnknownFormat},
	}

	for _, tCase := range test_cases{
		exO, exN, exE := calc.parse(tCase.inp)
		if exE != tCase.err{
			t.Errorf("For input %q, expected operands %v, got %v", tCase.inp, tCase.operands, exO)
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