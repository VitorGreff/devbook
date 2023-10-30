package rotas

import (
	"api/src/controllers"
	"net/http"
)

var rotasUsuarios = []Rota{
	{
		URI:       "/usuarios",
		Metodo:    http.MethodPost,
		Funcao:    controllers.CriarUsuario,
		RequerAut: false,
	}, {
		URI:       "/usuarios",
		Metodo:    http.MethodGet,
		Funcao:    controllers.BuscarUsuarios,
		RequerAut: false,
	}, {
		URI:       "/usuarios/{usuarioId}",
		Metodo:    http.MethodGet,
		Funcao:    controllers.BuscarUsuario,
		RequerAut: false,
	}, {
		URI:       "/usuarios/{usuarioId}",
		Metodo:    http.MethodPut,
		Funcao:    controllers.AtualizarUsuario,
		RequerAut: false,
	}, {
		URI:       "/usuarios/{usuarioId}",
		Metodo:    http.MethodDelete,
		Funcao:    controllers.DeletarUsuario,
		RequerAut: false,
	},
}
