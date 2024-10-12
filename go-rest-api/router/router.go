package router

import (
	"go-rest-api/controller"
	"go-rest-api/repository"
	"go-rest-api/usecase"
	"go-rest-api/validator"
	"net/http"
	"os"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

func NewRouter(db *gorm.DB) *echo.Echo {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000", os.Getenv("FE_URL")},
		AllowHeaders: []string{echo.HeaderOrigin,
			echo.HeaderContentType,
			echo.HeaderAccept,
			echo.HeaderAccessControlAllowHeaders,
			echo.HeaderXCSRFToken},
		AllowMethods:     []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
		AllowCredentials: true,
	}))
	e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		CookiePath:     "/",
		CookieDomain:   os.Getenv("API_DOMAIN"),
		CookieHTTPOnly: true,
		CookieSecure:   true,
		CookieSameSite: http.SameSiteDefaultMode,
	}))

	ur := repository.NewUserRepository(db)
	uv := validator.NewUserValidator()
	uu := usecase.NewUserUsecase(ur, uv)
	uc := controller.NewUserController(uu)

	tr := repository.NewTaskRepository(db)
	tv := validator.NewTaskValidator()
	tu := usecase.NewTaskUsecase(tr, tv)
	tc := controller.NewTaskController(tu)

	e.POST("/signup", uc.SignUp)
	e.POST("/login", uc.LogIn)
	e.POST("/logout", uc.LogOut)
	e.GET("/csrf", uc.CsrfToken)

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
