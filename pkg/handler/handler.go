package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/mbredikhin/snippets/pkg/service"
)

// Handler - route handler
type Handler struct {
	services *service.Service
}

// NewHandler - handler constructor
func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

// InitRoutes - initialize routes
func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("auth")
	{
		auth.POST("/sign-in", h.signIn)
		auth.POST("/sign-up", h.signUp)
		auth.POST("/logout", h.logout)
	}

	api := router.Group("/api", h.userIdentity)
	{
		lists := api.Group("/lists")
		{
			lists.POST("", h.createList)
			lists.GET("", h.getAllLists)
			lists.GET("/:id", h.getList)
			lists.PUT("/:id", h.updateList)
			lists.DELETE("/:id", h.deleteList)
			snippets := lists.Group(":id/snippets")
			{
				snippets.POST("", h.createSnippet)
				snippets.GET("", h.getAllSnippets)
			}
		}
		snippets := api.Group("snippets")
		{
			snippets.GET("/:id", h.getSnippet)
			snippets.PUT("/:id", h.updateSnippet)
			snippets.DELETE("/:id", h.deleteSnippet)
			snippetsTags := snippets.Group(":id/tags")
			{
				snippetsTags.GET("", h.getSnippetTags)
				snippetsTags.POST("", h.addTagToSnippet)
				snippetsTags.DELETE("", h.removeTagFromSnippet)
			}
			favouriteSnippets := snippets.Group("/favourites")
			{
				favouriteSnippets.GET("", h.getFavouriteSnippets)
				favouriteSnippets.POST("", h.addSnippetToFavourites)
				favouriteSnippets.DELETE("", h.removeSnippetFromFavourites)
			}
		}
		tags := api.Group("tags")
		{
			tags.POST("", h.createTag)
			tags.GET("", h.getAllTags)
			tags.GET("/:id", h.getTag)
			tags.PUT("/:id", h.updateTag)
			tags.DELETE("/:id", h.deleteTag)
		}
		languages := api.Group("languages")
		{
			languages.POST("", h.addLanguage)
			languages.GET("", h.getLanguages)
		}
	}

	return router
}
