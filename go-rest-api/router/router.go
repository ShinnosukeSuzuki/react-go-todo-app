package router

import (
	"go-rest-api/controller"
	"go-rest-api/repository"
	"go-rest-api/usecase"
	"os"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func NewRouter(db *gorm.DB) *echo.Echo {
	e := echo.New()

	ur := repository.NewUserRepository(db)
	uu := usecase.NewUserUsecase(ur)
	uc := controller.NewUserController(uu)

	tr := repository.NewTaskRepository(db)
	tu := usecase.NewTaskUsecase(tr)
	tc := controller.NewTaskController(tu)

	e.POST("/signup", uc.SignUp)
	e.POST("/login", uc.LogIn)
	e.POST("/logout", uc.LogOut)

	t := e.Group("/tasks")
	t.Use(echojwt.WithConfig(echojwt.Config{
		SigningKey:  []byte(os.Getenv("SECRET")),
		TokenLookup: "cookie:token",
	}))
	t.GET("", tc.GetAll)
	t.GET("/:id", tc.GetByID)
	t.POST("", tc.Create)
	t.PUT("/:id", tc.Update)
	t.DELETE("/:id", tc.Delete)

	return e
}
