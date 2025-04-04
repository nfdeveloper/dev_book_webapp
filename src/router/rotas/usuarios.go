package rotas

import (
	"net/http"

	"github.com/nfdeveloper/dev_book_webapp/src/controllers"
)

var rotaUsuarios = []Rota{
	{
		URI:                "/criar-usuario",
		Metodo:             http.MethodGet,
		Funcao:             controllers.CarregarPaginaDeCadastroDeUsuario,
		RequerAutenticacao: false,
	},
}
