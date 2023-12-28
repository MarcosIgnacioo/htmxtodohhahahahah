package main

import (
	"html/template"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Page struct {
    Data Data
    FormData FormData
}

func NewPage() Page {
    return Page{
        Data: NewData(),
        FormData: NewFormData(),
    }
}

type FormData struct {
    Values map[string]string
    Errors map[string]string
}

func NewFormData() FormData {
    return FormData{
        Values: make(map[string]string),
        Errors: make(map[string]string),
    }
}
var id int = 0

type Contact struct {
    Id int
    Name string
    Email string
}

type Contacts []Contact

func NewContact(name string, email string) Contact  {
    id++
    return Contact{ 
        Id:id,
        Name:name,
        Email:email,
    }
}

func (c Contacts) IndexOfContact(id int) int  {
    for i, v := range c{
        if v.Id == id {
            return i
        }
    }
    return -1
}

type Data struct{
    Contacts Contacts
}

func NewData() Data{
    return Data{
        // Esto puede ser un pooco confuso pero aqui te digo que esta pasando
        // Tenemos que nuestro Data tiene una propiedad de tipo Contacts la cual es un arreglo de Contact por lo que tenemos que darle un arreglo de Contact pues ese es el valor de Contacts, porque asi funcionan los structs si fuera un string le estariamos dando un string pero esto es un arreglo de Contact porque eso es lo que en realidad vale el struct Conctacts, entonces a la propiedad Contacts que tiene Data que es de tipo Contacts es decir un arreglo de Contact le tenemos que lpasar pues un arreglo de Contact

        Contacts : [] Contact {
            NewContact("Alexelcapo", "EvilAFM@gmail.com"),
            NewContact("CHinchetoj", "chin@gmail.com"),
            NewContact("Tonacho", "to@gmail.com"),
            NewContact("Felipez360", "F360@gmail.com"),
        },
    }
}

func (d Data) hasEmail(email string) bool {
    for _, contact := range d.Contacts{
        if contact.Email == email {
            return true
        }
    }
    return false
}

type Templates struct {
    templates *template.Template
    // Tenemos un struct con una template del tipo template que extraemos del package template
}

func (t *Templates) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
    return t.templates.ExecuteTemplate(w, name, data)
    // Renderizamos a nuestra template de la cual pasamos su nombre y el objeto que queramos que se renderice
}

func NewTemplate() *Templates {
   return &Templates{
        templates: template.Must(template.ParseGlob("views/*.html")),
    } 
}

func main()  {
    page := NewPage()
    e := echo.New()
    e.Use(middleware.Logger())
    e.Renderer = NewTemplate()
    e.Static("/css", "css")
    e.Static("/images", "images")

    e.GET("/", func (c echo.Context) error {
        return c.Render(http.StatusOK, "index", page)
    })

    e.POST("/contacts", func (c echo.Context) error {
        name := c.FormValue("name") 
        email := c.FormValue("email") 

        if (page.Data.hasEmail(email)) {
            // Usamos el formdata para que en caso de que haya errores tengamos un hashmap con toda la informacion que quiso hacer post para que se le ponga de nuevo y simplemente le marcamos su error
            formData := NewFormData()
            formData.Values["name"] = name
            formData.Values["email"] = email
            formData.Errors["email"] = "Email already exists biiitch"
            return c.Render(http.StatusUnprocessableEntity, "form", formData)
        }


        contact := NewContact(name, email)
        page.Data.Contacts = append(page.Data.Contacts, contact)

        c.Render(http.StatusOK, "form", NewFormData())

        return c.Render(http.StatusOK, "oob-contact", contact)
    })

    e.DELETE("/contacts/:id", func(c echo.Context) error {
        idParam := c.Param("id")
        time.Sleep(3 * time.Second)

        id, err := strconv.Atoi(idParam)
        if err != nil {
            return c.String(http.StatusBadRequest, "Invalid ID")
        }

        index := page.Data.Contacts.IndexOfContact(id)
        if index == -1 {
            return c.String(http.StatusNotFound, "Contact not found")
        }

        page.Data.Contacts = append(page.Data.Contacts[:index], page.Data.Contacts[index+1:]...)

        return c.NoContent(http.StatusOK)
    })

    e.Logger.Fatal(e.Start(":3022"))
}
