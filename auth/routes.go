package auth

import "github.com/gin-gonic/gin"

//MapRouteGroup mapeia rotas de autenticação
func MapRouteGroup(r *gin.Engine) {
	authGroup := r.Group("/auth")
	{
		authGroup.POST("/register", postRegisterView)
		authGroup.POST("/login", postLoginView)
	}
}
