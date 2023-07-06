package models

import "errors"

type Supplier struct {
	Id            int `json:"id"`
	SuppliersName string `json:"suppliers_name"`
	Password	  string `json:"password"`
}

const SupplierSchema string = ``

type Suppliers []Supplier

func (s *Supplier) Validate() error {
	if s.SuppliersName == "" {
		return errors.New("El nombre del proveedor no puede estar vacio")
	}
	if s.Password == "" {
		return errors.New("La contraseña no puede estar vacia")
	}
	return nil
}