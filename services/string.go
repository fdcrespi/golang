package services

import (
	"errors"
	"strconv"
)

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
		if r.Type != "NN" && r.Type != "TX" {
			return &StringFormat{"", "", 0}, errors.New("el formato indicado es invalido")
		}
		if r.Type == "NN" {
			_, err := strconv.Atoi(r.Value)
			if err != nil {
				return &StringFormat{"", "", 0}, errors.New("el valor no coincide con el tipo indicado")
			}
		}
		if r.Length == len(r.Value) {
			return r, nil
		}
		return &StringFormat{"", "", 0}, errors.New("la longitud indicada es incorrecta")
	}
	return &StringFormat{"", "", 0}, errors.New("formato de cadena invalido")
}
