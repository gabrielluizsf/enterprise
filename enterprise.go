package enterprise

import (
	"io"

	"github.com/i9si-sistemas/cep/api"
	"github.com/i9si-sistemas/nine"
)

type Enterprise interface {
	Data() nine.JSON
	CEP() *api.CEP
	Reader() io.Reader
}
