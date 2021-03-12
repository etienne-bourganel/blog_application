package handlers

import (
	"text/template"

	"github.com/labstack/echo/v4"
)

type IndexData struct {
	Posts []BlogPost
}

type BlogPost struct {
	ID          string
	Title       string
	Summary     string
	PublishedAt string
}

func Index(c echo.Context) error {
	tmpl, err := template.ParseFiles("../web/index.html", "../web/header.html")
	if err != nil {
		return err
	}
	w := c.Response().Writer

	posts := make([]BlogPost, 0)
	posts = append(posts, BlogPost{
		ID:    "1",
		Title: "The very first blog post",
		Summary: `Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nullam pretium aliquet elit quis facilisis.
		Orci varius natoque penatibus et magnis dis parturient montes, nascetur ridiculus mus. Phasellus scelerisque
		inenim sed pellentesque. Morbi ac semper tortor, ut volutpat diam. Suspendisse potenti. Mauris vel dui vel
		massa rhoncus ultricies in vitae massa. Morbi vitae neque in diam posuere bibendum a eget nulla. Ut non
		nibh at enim accumsan faucibus. Suspendisse potenti. Sed sollicitudin, sapien sed ultrices consectetur, dui
		risus pellentesque justo, quis lacinia ligula sapien ut sapien. Proin ac finibus eros. Curabitur auctor massa
		ac leo scelerisque ornare. Nulla facilisi. Integer efficitur mauris eget nibh volutpat, nec accumsan augue
		feugiat. Pellentesque luctus fringilla lorem.`[:100],
		PublishedAt: "March 12, 2021",
	})
	posts = append(posts, BlogPost{
		ID:          "2",
		Title:       "The second blog post",
		Summary:     "Jajemen",
		PublishedAt: "March 11, 2021",
	})
	data := IndexData{
		Posts: posts,
	}

	return tmpl.Execute(w, data)
}
