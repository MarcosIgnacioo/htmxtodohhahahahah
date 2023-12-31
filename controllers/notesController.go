package controllers

import (
	"net/http"
	"github.com/MarcosIgnacioo/htmx-go/initializers"
	"github.com/MarcosIgnacioo/htmx-go/models"
	"github.com/gin-gonic/gin"
)

func init()  {
    initializers.LoadEnvVariables()
    initializers.ConnectToDB()
}

func PostNote (c *gin.Context) {

    title := c.PostForm("title") 
    content := c.PostForm("content") 

    var body struct {
        Title string `json:"title"` 
        Content string `json:"content"`
        IsDone bool `json:"isDone"`
    }

    // c.Bind(&body) esto sirve para cuando hacemos post request reales asi tipo con postman y asi pero pues htmxgod no nos deja usar esto llorar::

    body.Title = title
    body.Content = content
    body.IsDone = false
    
    note := models.Note{ Title: body.Title, Content: body.Content, IsDone: body.IsDone }

    result := initializers.DB.Create(&note) // Pasamos el pointer porque queremos que modifique o en si pueda hacer lo que sea con ese y gaste menos memoria
    // value interface{} Es el any pero chido gamer porque namas accepta structs d seguro

    if result.Error != nil {
        c.Status(400)
        return 
    }

    c.HTML(http.StatusOK,"test.tmpl", note)

}

func SelectAllNotes() (notes [] models.Note){

    initializers.DB.Order("created_at desc").Find(&notes)
    return

}

func GetNotes(c *gin.Context) {

    var notes [] models.Note
    initializers.DB.Find(&notes)

    c.JSON(200, gin.H{
        "message": notes,
    })

}

func GetNote(c *gin.Context) {

    var note  models.Note
    id := c.Param("id")
    initializers.DB.First(&note, id)

    c.JSON(200, gin.H{
        "message": note,
    })

}


func UpdateNote(c *gin.Context)  {
    title := c.Request.PostFormValue("title")
    content := c.Request.PostFormValue("content")
    // isDone, _ := strconv.ParseBool(c.Request.FormValue("isDone"))
    

    var note  models.Note
    id := c.Param("id")
    initializers.DB.First(&note, id)
    initializers.DB.Model(&note).Updates(models.Note {
        Title: title,
        Content: content,
        IsDone: false,
    })

    initializers.DB.Save(note)

    //c.HTML(http.StatusOK, "note-details.html", note)
    c.Redirect(http.StatusSeeOther,"/")
    //c.JSON(200, note)
}

func DeleteNote(c *gin.Context)  {
    
    id := c.Param("id")

    initializers.DB.Delete(&models.Note{}, id)

    c.JSON(200, gin.H{
        "message": 200,
    })

}
