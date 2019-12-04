package main

import (
	"fmt"

	api "helloworld/api"

	"github.com/labstack/echo"
)

//测试跑该模块，使用 go run test001.go
//测试部署该模块，使用 go install helloworld

func main() {
	fmt.Printf("hello chain world app.go \n")

	e := echo.New()
	e.GET("/", api.HelloWorld)
	e.Logger.Fatal(e.Start(":1323"))
}
