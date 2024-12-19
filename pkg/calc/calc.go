package calc

import (
	"strconv"
	"strings"
	"example.com/m/errors"
)

type Calculator struct{}

func NewCalculator() *Calculator{
	return &Calculator{}
}

var (
	bSpace = byte(' ')
	bPlus = byte('+')
	bMinus = byte('-')
	bDiv = byte('/')
	bMult = byte('*')
	bOQuot = byte('(')
	bCQuot = byte(')')
)

func (c *Calculator) Calculation(src string) (float64, error){
	numbers, operands, err := c.parse(src)
	if err != nil{
		return 0, err
	}

	return c.calc(numbers, operands)
}

func (c *Calculator) parse(src string)([]float64, []byte, error){
	operands := make([]byte, 0)
	numbers := make([]float64, 0)

	builder := new(strings.Builder)

	addNumberF := func() error{
		if builder.Len() > 0{
			number, err := strconv.ParseFloat(builder.String(), 64)
			if err != nil{
				return errors.ErrUnknownFormat
			}

			numbers = append(numbers, number)
			builder.Reset()
		}

		return nil
	}

	for i := range src{
		bl := src[i]
		switch bl{
		case bSpace:
			continue
		case bPlus, bMinus, bDiv, bMult, bOQuot, bCQuot:
			if err := addNumberF(); err != nil{
				return nil, nil, err
			}

			operands = append(operands, bl)
			builder.Reset()
		default:
			builder.WriteByte(bl)
		}
	}
	if err := addNumberF(); err != nil{
		return nil, nil, err
	}

	return numbers, operands, nil
}

func (c *Calculator) calc(numbers []float64, operands []byte) (float64, error){
	
	return 0, nil
}