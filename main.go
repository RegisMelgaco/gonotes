package main

import (
	"local/gonotes/auth"
	"local/gonotes/note"
	"os"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	note.MapRouteGroup(r)
	auth.MapRouteGroup(r)

	return r
}

func main() {
	r := setupRouter()

	os.Setenv("ACCESS_SECRET", "h.pWdRT+A8}$Ug}WIHuyRiI6jH[S^J9Ui5a!(4Nj?JBO,y+O7fdUV5};U2.8Mcu")

	err := r.Run()
	panic(err)
}
