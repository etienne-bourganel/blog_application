package handlers

import (
	"text/template"

	"github.com/labstack/echo/v4"
)

type PostData struct {
	ID          string
	Title       string
	Body        string
	PublishedAt string
	Author      string
}

func Post(c echo.Context) error {
	postID := c.Param("postID")

	tmpl, err := template.ParseFiles("../web/blogpost.html")
	if err != nil {
		return err
	}
	w := c.Response().Writer

	data := PostData{
		ID:          postID,
		Title:       "En viktig bloggpost",
		Body:        "kjdsflk ajlfkjafjsdkfj skldjfd s",
		PublishedAt: "March 12 2021",
		Author:      "Etienne",
	}

	return tmpl.Execute(w, data)
}
