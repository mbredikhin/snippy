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
	}

	api := router.Group("/api", h.userIdentity)
	{
		lists := api.Group("/lists")
		{
			lists.POST("/", h.createList)
			lists.GET("/", h.getAllLists)
			lists.GET("/:id", h.getList)
			lists.PUT("/:id", h.updateList)
			lists.DELETE("/:id", h.deleteList)
			snippets := lists.Group(":id/snippets")
			{
				snippets.POST("/", h.createSnippet)
				snippets.GET("/", h.getAllSnippets)
			}
		}
		snippets := api.Group("snippets")
		{
			snippets.GET("/:id", h.getSnippet)
			snippets.PUT("/:id", h.updateSnippet)
			snippets.DELETE("/:id", h.deleteSnippet)
		}
		favouriteSnippets := api.Group("/favourite-snippets")
		{
			favouriteSnippets.GET("/", h.getFavouriteSnippets)
			favouriteSnippets.POST("/:id", h.addFavouriteSnippet)
			favouriteSnippets.DELETE("/:id", h.removeFavouriteSnippet)
		}
		tags := api.Group("tags")
		{
			tags.POST("/", h.createTag)
			tags.GET("/", h.getAllTags)
			tags.GET("/:id", h.getTag)
			tags.PUT("/:id", h.updateTag)
			tags.DELETE("/:id", h.deleteTag)
		}
		languages := api.Group("languages")
		{
			languages.POST("/", h.addLanguage)
			languages.GET("/", h.getLanguages)
		}
	}

	return router
}
