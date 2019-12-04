package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	fmt.Printf("hello chain world \n")

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
