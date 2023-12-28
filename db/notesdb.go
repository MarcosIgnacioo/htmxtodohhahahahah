package db

import "fmt"

type User struct {
    Id int        `json:"id"`
    Nombre string `json:"nombre"`
    Edad int      `json:"edad"`
    Email string  `json:"email"`
}

func (u *User) GetAll() ([]User, error) {
    println("whatsapp")
    db := GetConnection()
    q := `SELECT
            *
            FROM usuarios`    // Ejecutamos la query
    rows, err := db.Query(q)
    if err != nil {
        return []User{}, err
    }    // Cerramos el recurso
    defer rows.Close()    // Declaramos un slice de notas para que almacene las
    // notas que retorna la petición.
    usuarios := []User{}    // El método Next retorna un bool, mientras sea true indicará
    // que existe un valor siguiente para leer.
    for rows.Next() {
        // Escaneamos el valor actual de la fila e insertamos el
        // retorno en los correspondientes campos de la nota.
        rows.Scan(
            &u.Id,
            &u.Nombre,
            &u.Edad,
            &u.Email,
        )        // Añadimos cada nueva nota al slice de notas que
        // declaramos antes.
        usuarios = append(usuarios, *u)
    }
    fmt.Printf("usuarios: %v\n", usuarios)
    return usuarios, nil
}
