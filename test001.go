package main

import (
	"fmt"

	api "helloworld/api"

	"github.com/labstack/echo"
)

func main() {
	fmt.Printf("hello chain world test001 \n")

	e := echo.New()
	e.GET("/", api.HelloWorld)
	e.Logger.Fatal(e.Start(":1323"))
}
