package enterprise

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strings"

	"github.com/i9si-sistemas/cep/api"
	"github.com/i9si-sistemas/nine"
	"github.com/i9si-sistemas/nine/pkg/client"
)

func WithCNPJ(ctx context.Context, cnpj string) (Enterprise, error) {
	url := fmt.Sprintf("https://www.receitaws.com.br/v1/cnpj/%s", cleanCNPJ(cnpj))
	res, err := nine.New(ctx).Get(url, &client.Options{})
	if err != nil {
		return nil, fmt.Errorf("erro ao fazer a requisição: %v", err)
	}
	enterprise := enterpriseWithCNPJ{
		ctx:    ctx,
		reader: res.Body,
		cnpj:   cnpj,
	}
	var payload nine.JSON
	if err := json.NewDecoder(res.Body).Decode(&payload); err != nil {
		return nil, fmt.Errorf("erro ao decodificar a resposta: %v", err)
	}
	enterprise.json = payload
	return &enterprise, nil
}

func cleanCEP(cep string) string {
	return strings.NewReplacer(".", "", "-", "", "/", "").Replace(cep)
}

func cleanCNPJ(cnpj string) string {
	return strings.NewReplacer(".", "", "-", "", "/", "").Replace(cnpj)
}

type enterpriseWithCNPJ struct {
	ctx    context.Context
	cnpj   string
	json   nine.JSON
	reader io.Reader
}

func (e *enterpriseWithCNPJ) Data() nine.JSON {
	return e.json
}

func (e *enterpriseWithCNPJ) CEP() *api.CEP {
	data := e.Data()
	cep, err := api.Find(e.ctx, cleanCEP(data["cep"].(string)))
	if err != nil {
		return nil
	}
	return cep
}

func (e *enterpriseWithCNPJ) Reader() io.Reader {
	return e.reader
}

func (e *enterpriseWithCNPJ) String() string {
	return fmt.Sprintf("CNPJ: %s", e.cnpj)
}
