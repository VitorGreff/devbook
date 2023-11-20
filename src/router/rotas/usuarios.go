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
		RequerAut: true,
	}, {
		URI:       "/usuarios/{usuarioId}",
		Metodo:    http.MethodGet,
		Funcao:    controllers.BuscarUsuario,
		RequerAut: true,
	}, {
		URI:       "/usuarios/{usuarioId}",
		Metodo:    http.MethodPut,
		Funcao:    controllers.AtualizarUsuario,
		RequerAut: true,
	}, {
		URI:       "/usuarios/{usuarioId}",
		Metodo:    http.MethodDelete,
		Funcao:    controllers.DeletarUsuario,
		RequerAut: true,
	}, {
		URI:       "/usuarios/{usuarioId}/seguir",
		Metodo:    http.MethodPost,
		Funcao:    controllers.SeguirUsuario,
		RequerAut: true,
	},
}
