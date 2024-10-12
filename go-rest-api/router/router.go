package router

import (
	"go-rest-api/controller"
	"go-rest-api/repository"
	"go-rest-api/usecase"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func NewRouter(db *gorm.DB) *echo.Echo {
	e := echo.New()

	ur := repository.NewUserRepository(db)
	uu := usecase.NewUserUsecase(ur)
	uc := controller.NewUserController(uu)

	e.POST("/signup", uc.SignUp)
	e.POST("/login", uc.LogIn)
	e.POST("/logout", uc.LogOut)

	return e
}
