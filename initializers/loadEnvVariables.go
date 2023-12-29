package initializers

import (
	"log"

	"github.com/joho/godotenv"
)

func LoadEnvVariables()  {
    // Cargamos nuestras variables de entorno que en este caso ponen el PORT a 3000 y ya gin sabe en que puerto alojar el server
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading env")
    }
}
