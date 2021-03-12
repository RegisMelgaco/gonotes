package note

//Note Model que contem as informações sobre uma anotação
type Note struct {
	ID          int
	Name        string `binding:"required"`
	Description string `binding:"required"`
	IsChecked   bool   `binding:"required"`
}
