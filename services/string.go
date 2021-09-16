package services

import (
	"errors"
	"strconv"
)

/* type FormatString interface {
	ParseString(c string) (*Result, error)
}*/

type StringFormat struct {
	Type   string
	Value  string
	Length int
}

func (r *StringFormat) ParseString(c string) (*StringFormat, error) {
	if len(c) > 4 {
		r.Type = c[0:2]
		r.Length, _ = strconv.Atoi(c[2:4])
		r.Value = c[4:]
		if r.Length == len(r.Value) {
			return r, nil
		}
		return nil, errors.New("la longitud indicada es incorrecta")
	}
	return nil, errors.New("formato de cadena invalido")
}
