package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListAgendamentosFromCliente(ctx *gin.Context) {
	nome := ctx.DefaultQuery("username", "")
	if nome == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Nome de usário não encontrado",
		})
		return
	}

}
