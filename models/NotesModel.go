package models

import "gorm.io/gorm"

type Note struct {
    gorm.Model
    Id int `gorm:"id"`
    Title   string `gorm:"title"`   
    Content string `gorm:"content"`
    IsDone  bool   `gorm:"isDone"`
}
