package calc

import (
	"strconv"
	"strings"

	"example.com/m/errors"
)

var (
	bSpace = byte(' ')
	bPlus  = byte('+')
	bMinus = byte('-')
	bDiv   = byte('/')
	bMult  = byte('*')
	bOQuot = byte('(')
	bCQuot = byte(')')
)

func Calc(src string) (float64, error) {
	if err := сalcQuots(&src); err != nil {
		return 0, err
	}

	numbers, operators, err := parse(src)
	if err != nil {
		return 0, err
	}

	if len(numbers) <= len(operators) {
		return 0, errors.ErrUnknownFormat
	}

	return calculation(numbers, operators)
}

func сalcQuots(src *string) error {
	i := 0
	for i < len(*src) {
		let := (*src)[i]
		switch let {
		case bOQuot:
			cI, err := findCloseQuot(i, *src)
			if err != nil {
				return err
			}
			fRes, err := Calc((*src)[i+1 : cI])
			if err != nil {
				return err
			}
			sRes := strconv.FormatFloat(fRes, 'g', -1, 64)
			*src = replaceSubstring(src, sRes, i, cI)
		case bCQuot:
			return errors.ErrNotOpenQuot
		}
		i++
	}

	return nil
}

func replaceSubstring(src *string, res string, i, cI int) string {
	if cI+1 < len(*src) {
		return (*src)[:i] + res + (*src)[cI+1:]
	}

	return (*src)[:i] + res
}

func findCloseQuot(i int, src string) (int, error) {
	open := 0

	for i < len(src) {
		exp := src[i]
		if exp == bOQuot {
			open++
		} else if exp == bCQuot {
			open--
			if open == 0 {
				return i, nil
			}
		}
		i++
	}

	return 0, errors.ErrNotCloseQuot
}

func parse(src string) ([]float64, []byte, error) {
	if len(src) < 1 || strings.Contains(string(src[0]), "-+*/") || (len(src) > 1 && strings.Contains(string(src[1]), "-+*/")) {
		return nil, nil, errors.ErrUnknownFormat
	}

	numbers := make([]float64, 0)
	operands := make([]byte, 0)

	builder := new(strings.Builder)
	addNumber := func() error {
		defer builder.Reset()
		if builder.Len() < 1 {
			return errors.ErrUnknownFormat
		}
		number, err := strconv.ParseFloat(builder.String(), 64)
		if err != nil {
			return errors.ErrUnknownFormat
		}
		numbers = append(numbers, number)
		return nil
	}

	for i := 0; i < len(src); i++ {
		bl := src[i]
		switch bl {
		case bSpace:
			continue
		case bPlus, bMinus, bDiv, bMult:
			if err := addNumber(); err != nil {
				return nil, nil, err
			}
			operands = append(operands, bl)
		default:
			builder.WriteByte(bl)
		}
	}
	if err := addNumber(); err != nil {
		return nil, nil, err
	}

	return numbers, operands, nil
}

func calculation(numbers []float64, operands []byte) (float64, error) {
	for i := 0; i < len(operands); {
		op := operands[i]
		if op == bDiv || op == bMult {
			res, err := makeOperation(numbers[i], numbers[i+1], op)
			if err != nil {
				return 0, err
			}
			numbers[i+1] = res
			deleteSliceElement(&numbers, i)
			deleteSliceElement(&operands, i)
			continue
		}
		i++
	}

	for i := 0; i < len(operands); i++ {
		op := operands[i]
		if op == bMinus || op == bPlus {
			res, err := makeOperation(numbers[0], numbers[1], op)
			if err != nil {
				return 0, err
			}
			numbers[0] = res
			deleteSliceElement(&numbers, 1)
		}
	}

	return numbers[0], nil
}

func deleteSliceElement[T any](sl *[]T, i int) {
	if i < len(*sl)-1 {
		*sl = append((*sl)[:i], (*sl)[i+1:]...)
		return
	}
	*sl = (*sl)[:i]
}

func makeOperation(a, b float64, o byte) (float64, error) {
	switch o {
	case bPlus:
		return a + b, nil
	case bMinus:
		return a - b, nil
	case bDiv:
		if b == 0 {
			return 0, errors.ErrDivisionByZero
		}
		return a / b, nil
	default:
		return a * b, nil
	}
}
