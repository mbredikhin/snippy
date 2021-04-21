package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mbredikhin/snippets"
)

func (h *Handler) getSyntaxList(c *gin.Context) {
	syntaxList, err := h.services.Snippet.GetSyntaxList()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"data": syntaxList,
	})
}

func (h *Handler) addSyntax(c *gin.Context) {
	var input snippets.Syntax
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.services.Snippet.CreateSyntax(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}
