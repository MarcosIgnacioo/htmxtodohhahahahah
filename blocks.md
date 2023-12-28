package main

import (
	"html/template"
	"io"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

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

type Block struct {
    Id int
}

type Blocks struct {
    Start int
    Next int
    More bool
    Blocks []Block
}

func main() {
	e := echo.New()
    e.Renderer = NewTemplates()
    e.Use(middleware.Logger())

    e.GET("/blocks", func(c echo.Context) error {
        // Agarramos de nuestra url el parametro start si no esta en la url lo creamos y lo inicializamos en 0
        startStr := c.QueryParam("start")
        start, err := strconv.Atoi(startStr)
        if err != nil {
            start = 0
        }

        blocks := []Block{}
        // Creamos nuestro array de blocks que comienza en nuestro start que si no existe en los query params es decir la primera vez que entramos a la pagina sera start = 0 pero cuando vayamos bajando el start sera diferente
        for i := start; i < start + 10; i++ {
            blocks = append(blocks, Block{Id: i})
        }


        template := "blocks"
        if start == 0 {
            template = "blocks-index"
        }
        return c.Render(http.StatusOK, template, Blocks{
            Start: start,
            Next: start + 10,
            More: start + 10 < 500,
            Blocks: blocks,
        });
        // En nuestro html estamos actualizando el start en los queryparams cuando sea cumpla la condiciion de que no se excedan los 500 elementos y el start tomara el valor de next que seran los siguientes 10 elementos que se generaran. Este evento se trigerea cuando se scrollea al final de la pagina por lo que, por ejemplo si en nuestra pantallla ya eran visibles los primeros 20, los siguientes 10 se generaran de la siguiente manera

        // En el html se va a activar el trigger de revealed cuando nuestro div que contiene el trigger revealed
            // Este trigger lo que hace es un get a nuestro endpoint de blocks 
            // El request que se hace a este endpoint se hace ya con un queryparam de start  = .Next que es basicamente los siguientes bloques por lo que por ejemplo en el caso de los segundos 10 bloques ya el start seria 10 (el start de los primeros 10 bloques es 0)
            // Y pues se hace de nuevo todo, creamos nuestro arreglo de bloques a los que les damos el id a partir de nuestro valor de start (en este caso 10), y luego verificamos que no sea la priimera vez que se hace este get checando si start vale 0 o no, en caso de no valerlo la template se queda como blocks
            // LE pasamos ya nuestro BLocks que contendra toda la informacion de estos bloques, como el iinicio, donde sera el siguiente inicio (next) y si aun se pueden generar mas bloques si hay menos de 500 (more) y en si el arreglo de blocks, tambien le pasamos la template de blocks que lo que hace es iterar sobre el arreglo de Blocks y los pone, despues pone el div que va a detectar el scroll para asi volver a hacer todo esto hasta que hayamos generado los 500 divs
        //  Esto esta chido porque podemos guardar en cada div el id de nuestros registros y asi poderp or ejemplo borrarlo de nuestra tabla o actualizarlo y toda esta ya es informacion que cargo el usuario asi que no hay pedo por lo que nos evita tener que cargar toda la tabla de nuevo cada vez que hay un cambio
    });

    e.Logger.Fatal(e.Start(":3023"))
}

