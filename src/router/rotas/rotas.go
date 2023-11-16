package rotas

import (
	"net/http"

	"github.com/gorilla/mux"
)

// representa as rotas da API
type Rota struct {
	URI       string
	Metodo    string
	Funcao    func(http.ResponseWriter, *http.Request)
	RequerAut bool
}

func Configurar(r *mux.Router) *mux.Router {
	rotas := rotasUsuarios
	rotas = append(rotas, rotaLogin)

	for _, rota := range rotas {
		r.HandleFunc(rota.URI, rota.Funcao).Methods(rota.Metodo)
	}

	return r
}
