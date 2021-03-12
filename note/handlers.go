package note

import (
	"net/http"
	"strconv"

	"local/gonotes/common"
	"local/gonotes/db"

	"github.com/gin-gonic/gin"
)

//getNotesView lista as notas
func getNotesView(c *gin.Context) {
	db := db.GetDbConnection()
	defer db.Close()

	var notes []Note
	db.Model(&notes).Select()
	c.JSON(http.StatusOK, notes)
}

//getNoteView retorna uma nota especifica
func getNoteView(c *gin.Context) {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		c.JSON(http.StatusBadRequest, common.ErrorHolder{Error: err.Error()})
		return
	}

	db := db.GetDbConnection()
	defer db.Close()

	note := &Note{ID: id}
	err = db.Model(note).WherePK().Select()
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorHolder{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, note)
}

//postNoteView cria uma nova nota
func postNoteView(c *gin.Context) {
	var note Note
	if err := c.BindJSON(&note); err != nil {
		c.JSON(http.StatusBadRequest, common.ErrorHolder{Error: err.Error()})
		return
	}

	db := db.GetDbConnection()
	defer db.Close()

	db.Insert(&note)

	c.JSON(http.StatusOK, note)
}

//deleteNoteView deleta uma nota
func deleteNoteView(c *gin.Context) {
	idString := c.Param("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	db := db.GetDbConnection()
	defer db.Close()

	note := &Note{ID: id}
	err = db.Delete(note)
	if err != nil {
		c.JSON(http.StatusNotFound, common.ErrorHolder{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, note)
}
