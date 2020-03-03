package main

import (
	"os"
	"log"
	"github.com/gin-gonic/gin"
	"github.com/justym/todo/controller"
	//"github.com/justym/todo/model"
)

func main(){
	r := gin.Default()
	r.LoadHTMLGlob("views/*.html")
	handler := controller.Handler{}
	if err := handler.DB.Init(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
	r.GET("/",handler.All)
	r.POST("/create",handler.Insert)
	r.GET("/item/:id",handler.Get)
	r.POST("/update/:id",handler.Update)
	r.GET("/confirm/:id",handler.Confirm)
	r.POST("/delete/:id",handler.Delete)
	

	if err := r.Run(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}