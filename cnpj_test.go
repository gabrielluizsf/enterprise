package enterprise

import (
	"context"
	"testing"

	"github.com/i9si-sistemas/assert"
)

func TestWithCNPJ(t *testing.T) {
	cnpj := "00.000.000/0001-91"
	ctx := context.Background()
	enterprise, err := WithCNPJ(ctx, cnpj)
	if err != nil {
		t.Skipf("Erro ao consultar a empresa através do CNPJ: %s", err.Error())
	}
	assert.NotNil(t, enterprise)
	data := enterprise.Data()
	assert.NotNil(t, data)
	assert.Equal(t, data["cnpj"].(string), cnpj)
	assert.Equal(t, data["nome"].(string), "BANCO DO BRASIL SA")
	assert.Equal(t, data["cep"].(string), "70.040-912")
	cep := enterprise.CEP()
	if cep == nil {
		t.Skip("Erro ao consultar o CEP")
	}
	assert.Equal(t, cep.ZipCode, "70040-912")
}
