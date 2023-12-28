package main

import (
	"html/template"
	"io"
	"github.com/labstack/echo/v4"
	"todo.com/go/db"
)

// HTML / Server stuff

type Templates struct {
	templates *template.Template
}

func (t *Templates) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func NewTemplates() *Templates {
	return &Templates{
		templates: template.Must(template.ParseGlob("views/*.html")),
    }
}

// Utilities (structs)

type Note struct{
    Id int
    Title string
    Content string
    IsDone bool
}

func NewNote(id int, title string, content string, isDone bool) Note  {
    return Note{
        Id: id,
        Title: title,
        Content: content,
        IsDone: isDone,
    }
}



func main()  {
    println("whatsaaper")
    u := db.User {Id: 0, Nombre: "hola", Edad: 2, Email: "asdf"}
    u.GetAll()
}
