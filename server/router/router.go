package router

import (
	"server/handler/v1/user"
	"server/router/middleware"
	_ "server/docs" // docs is generated by Swag CLI, you have to import it.
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// Load loads the middlewares, routes, handlers.
func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	// Middlewares.
	g.Use(gin.Recovery())
	g.Use(middleware.NoCache)
	g.Use(middleware.Options)
	g.Use(middleware.Secure)
	g.Use(mw...)
	// 404 Handler.
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API route.")
	})
	// swagger api 文档
	// g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// // 登录
	// g.POST("/v1/login", user.Login)
	// rUser.POST("/v1/register", user.Register)
	// // 用户相关
	// rUser := g.Group("/v1/user")
	// //rUser.Use(middleware.AuthMiddleware())
	// {
	// 	rUser.DELETE("/:id", user.Delete)
	// 	rUser.PUT("/:id", user.Update)
	// 	rUser.GET("", user.List)
	// 	rUser.GET("/:id", user.GetUserById)
	// }
	return g
}
