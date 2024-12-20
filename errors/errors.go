package errors

import "fmt"

var(
	ErrUnknownFormat = fmt.Errorf("Неправильный формат")
	ErrDivisionByZero = fmt.Errorf("Нельзя делить на 0")
	ErrNotOpenQuot = fmt.Errorf("Нет открывающей скобки")
	ErrNotCloseQuot = fmt.Errorf("Нет закрывающей скобки")
	ErrUnknownError = fmt.Errorf("Что-то пошло не так")
)