package genErr

import (
	"errors"
	"fmt"
	"reflect"
)

type Error struct {
	prevError error
	message   string
	data      map[string]any
}

func (se *Error) Error() string {
	return fmt.Sprintf("%s", se.message)
}

func NewError(prevError error, message error, data ...any) *Error {
	d := make(map[string]any)
	var k string

	if len(data)%2 != 0 {
		data = data[:len(data)-1]
	} else {
		for i, v := range data {
			if i%2 == 0 {
				k = fmt.Sprintf("%v", v)
			} else {
				d[k] = v
			}
		}
	}

	se := &Error{
		prevError: prevError,
		message:   message.Error(),
		data:      d,
	}

	if prevError != nil && message != nil {
		pe, pOk := prevError.(*Error)

		if pOk {
			if reflect.DeepEqual(pe.data, se.data) && pe.message == se.message {
				return pe
			}
		}
	}

	return se
}

func (se *Error) Unwrap() error {
	return se.prevError
}

func New(text string) *Error {
	return NewError(nil, errors.New(text))
}
