package handler

import "fmt"

func errParamIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type %s) is required", name, typ)
}

type CriarComidaRequest struct {
	Nome    string `json:"nome"`
	Tipo    string `json:"tipo"`
	Caloria int64  `json:"caloria"`
}

func (r *CriarComidaRequest) Validate() error {
	if r.Nome == "" && r.Tipo == "" && r.Caloria <= 0 {
		return fmt.Errorf("request body is empty")
	}
	if r.Nome == "" {
		return errParamIsRequired("nome", "string")
	}
	if r.Tipo == "" {
		return errParamIsRequired("tipo", "string")
	}
	if r.Caloria < 0 {
		return errParamIsRequired("caloria", "int64")
	}

	return nil

}

type AlterarComidaRequest struct {
	Nome    string `json:"nome"`
	Tipo    string `json:"tipo"`
	Caloria int64  `json:"caloria"`
}

func (r *AlterarComidaRequest) Validate() error {
	if r.Nome != "" || r.Tipo != "" || r.Caloria > 0 {
		return nil
	}
	return fmt.Errorf("at least one valid field must be provided")
}
