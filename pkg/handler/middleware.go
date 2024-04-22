package handler

import (
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/mbredikhin/snippets"
)

const (
	authorizationHeader = "Authorization"
	userContext         = "userID"
	paginationContext   = "pagination"
)

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader(authorizationHeader)
	if header == "" {
		newErrorResponse(c, http.StatusUnauthorized, "empty auth header")
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		newErrorResponse(c, http.StatusUnauthorized, "invalid auth header")
		return
	}

	if len(headerParts[1]) == 0 {
		newErrorResponse(c, http.StatusUnauthorized, "token is empty")
		return
	}

	userID, _, err := h.services.Authorization.ParseToken(headerParts[1])
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	isTokenBlacklisted := h.services.Authorization.CheckIfTokenBlacklisted(headerParts[1])
	if isTokenBlacklisted {
		newErrorResponse(c, http.StatusUnauthorized, "token is blacklisted")
		return
	}
	c.Set(userContext, userID)
}

func getUserID(c *gin.Context) (int, error) {
	id, ok := c.Get(userContext)
	if !ok {
		return 0, errors.New("user id not found")
	}
	idInt, ok := id.(int)
	if !ok {
		return 0, errors.New("user id is of invalid type")
	}
	return idInt, nil
}

func getPaginationParams(c *gin.Context) *snippets.PaginationParams {
	params, ok := c.Get(paginationContext)
	if ok {
		pagination := params.(snippets.PaginationParams)
		return &snippets.PaginationParams{
			Page:  pagination.Page,
			Limit: pagination.Limit,
		}
	}
	return nil
}

func (h *Handler) collectPaginationParams(c *gin.Context) {
	paginationParams := snippets.PaginationParams{}
	var err error
	paginationParams.Page, err = strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "page param is invalid")
		return
	}
	paginationParams.Limit, err = strconv.Atoi(c.DefaultQuery("limit", "25"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "limit param is invalid")
		return
	}
	c.Set(paginationContext, paginationParams)
}
