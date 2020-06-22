package models

// Book model
type Book struct {
	ID     uint `json:"id" gorm:"primary_key"`
	Title  uint `json:"title"`
	Author uint `json:"author"`
}
