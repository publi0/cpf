package pkg

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
)

type CPFData struct {
	CPF        string `json:"cpf,omitempty"`
	Nome       string `json:"nome,omitempty"`
	Sexo       string `json:"sexo,omitempty"`
	Nascimento string `json:"nascimento,omitempty"`
	Mae        string `json:"mae,omitempty"`
	Pai        string `json:"pai,omitempty"`
	Idade      string `json:"idade,omitempty"`
}

const (
	CPFURL = "https://api.consultanacional.org:3000/consulta/"
)

func fetchCPFData(ctx context.Context, cpf string) (CPFData, error) {
	client := &http.Client{}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, CPFURL+cpf, nil)
	if err != nil {
		return CPFData{}, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return CPFData{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return CPFData{}, err
	}

	var cpfData CPFData

	json.Unmarshal(body, &cpfData)

	return cpfData, nil
}
