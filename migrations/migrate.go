package main

import (
	"github.com/MarcosIgnacioo/htmx-go/initializers"
	"github.com/MarcosIgnacioo/htmx-go/models"
)

func init()  {
    // Nos aseguramos de estar conectados a la DB no se que pasa si quitamos esto pero pues luego lo averiguamos
    initializers.LoadEnvVariables()
    initializers.ConnectToDB()
}

func main()  {
    initializers.DB.AutoMigrate(&models.Note{})
}
