package note

import "github.com/gin-gonic/gin"

//MapRouteGroup mapeia rotas de notas
func MapRouteGroup(r *gin.Engine) {
	notesGroup := r.Group("/note")
	{
		notesGroup.GET("", getNotesView)
		notesGroup.GET("/:id", getNoteView)
		notesGroup.POST("", postNoteView)
		notesGroup.DELETE("/:id", deleteNoteView)
	}
}
