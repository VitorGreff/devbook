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
	}, {
		URI:       "/usuarios/{usuarioId}/parar-de-seguir",
		Metodo:    http.MethodPost,
		Funcao:    controllers.PararDeSeguirUsuario,
		RequerAut: true,
	}, {
		URI:       "/usuarios/{usuarioId}/seguidores",
		Metodo:    http.MethodGet,
		Funcao:    controllers.BuscarSeguidores,
		RequerAut: true,
	}, {
		URI:       "/usuarios/{usuarioId}/seguindo",
		Metodo:    http.MethodGet,
		Funcao:    controllers.BuscarSeguindo,
		RequerAut: true,
	},
}
