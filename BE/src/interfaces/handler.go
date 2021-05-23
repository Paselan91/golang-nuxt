package interfaces

import (
	"app/src/config"
	// "app/src/usecase"
	"fmt"
	"github.com/go-playground/validator"
	"github.com/labstack/echo"
	// "github.com/labstack/echo/middleware"
	"log"
	"net/http"
)

// Run start server
func Run(e *echo.Echo, port string) {
	log.Printf("Server running at http://localhost:%s/", port)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", port)))
}

type Validator struct {
	validator *validator.Validate
}

func (v *Validator) Validate(i interface{}) error {
	return v.validator.Struct(i)
}

func BindValidate(c echo.Context, i interface{}) error {
	if err := c.Bind(i); err != nil {
		return c.String(http.StatusBadRequest, "Request is failed: "+err.Error())
	}
	if err := c.Validate(i); err != nil {
		return c.String(http.StatusBadRequest, "Validate is failed: "+err.Error())
	}
	return nil
}

// Routes returns the initialized router
func Routes(e *echo.Echo) {
	e.Validator = &Validator{validator: validator.New()}

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Good morning, Golang + Nuxt.js !")
	})
	// Migration Route
	e.GET("/api/v1/migrate", migrate)
	e.GET("/api/v1/seed", Seeds)
}

// =============================
//    MIGRATE
// =============================
func migrate(c echo.Context) error {
	_, err := config.DBMigrate()
	if err != nil {
		return c.String(http.StatusNotFound, "failed migrate")
	} else {
		return c.String(http.StatusOK, "success migrate")
	}
}

func Seeds(c echo.Context) error {
	_, err := config.Seeds()
	if err != nil {
		return c.String(http.StatusNotFound, "failed seed")
	} else {
		return c.String(http.StatusOK, "success seed")
	}
}
