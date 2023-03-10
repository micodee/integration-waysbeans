package routes

import (
	"waysbeans/controllers"
	"waysbeans/pkg/middleware"
	"waysbeans/pkg/mysql"
	"waysbeans/repositories"

	"github.com/labstack/echo/v4"
)

func CartRoutes(e *echo.Group) {
	cartRepository := repositories.RepositoryCart(mysql.ConnDB)
	h := controllers.ControlCart(cartRepository)

	e.GET("/cart", h.FindCarts)
	e.GET("/cart/:id", h.GetCart)
	e.POST("/cart/:product_id", middleware.Auth(h.CreateCart))
}
