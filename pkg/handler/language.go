package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mbredikhin/snippets"
)

func (h *Handler) getLanguages(c *gin.Context) {
	languages, err := h.services.Snippet.GetLanguages()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, response[[]snippets.Language]{
		languages,
	})
}

func (h *Handler) addLanguage(c *gin.Context) {
	var input snippets.Language
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	id, err := h.services.Snippet.CreateLanguage(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, response[snippets.AddLanguageResponse]{
		snippets.AddLanguageResponse{ID: &id},
	})
}
