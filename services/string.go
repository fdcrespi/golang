package services

import (
	"errors"
	"strconv"
)

type FormatString interface {
	GenerateString(c string) (*Result, error)
}

type Result struct {
	Type   string
	Value  string
	Length int
}

func (r *Result) GenerateString(c string) (*Result, error) {
	if len(c) > 4 {
		r.Type = c[0:2]
		r.Length, _ = strconv.Atoi(c[2:4])
		r.Value = c[4:]
		if r.Length == len(r.Value) {
			return r, nil
		}
		return nil, errors.New("La longitud indicada no coincide con la de la cadena")
	}
	return nil, errors.New("Formato de cadena invalido")
}
