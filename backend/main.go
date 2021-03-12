package main

import (
	"github.com/etienne-bourganel/skymill_blog_project/backend/handlers"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.Static("/", "../web/public")

	e.GET("/", handlers.Index)
	e.GET("/post/:postID", handlers.Post)

	e.Logger.Fatal(e.Start(":8000"))

}
