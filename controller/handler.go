package controller

import (
	"log"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/justym/todo/model"
)

//Handler bundle Request Handler
type Handler struct{
	model.DB
}

//All shows all todos
func (h *Handler) All(ctx *gin.Context){
	list,err := h.DB.All()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	
	ctx.HTML(200,"index.html",gin.H{"List":list})
}

//Insert insert new todo into DB
func (h *Handler) Insert(ctx *gin.Context){
	title := ctx.PostForm("title")
	status := ctx.PostForm("status")

	err := h.DB.Insert(title,status)
	if err != nil{
		log.Println(err)
		os.Exit(1)
	}
	ctx.Redirect(302,"/")
} 

//Get get a todo using id from DB
func (h *Handler) Get(ctx *gin.Context){
	param := ctx.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	item,err := h.DB.Get(id)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	ctx.HTML(200,"item.html",gin.H{"Item":item})
}

//Update update item in DB
func (h *Handler) Update(ctx *gin.Context){
	param := ctx.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	
	title := ctx.PostForm("title")
	status := ctx.PostForm("status")
	err = h.DB.Update(id,title,status)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	ctx.Redirect(302,"/")
}

//Delete delete todo from DB
func (h *Handler) Delete(ctx *gin.Context){
	param := ctx.Param("id")
	id,err := strconv.Atoi(param)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	err = h.DB.Delete(id)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	ctx.Redirect(302,"/")
}

func (h *Handler) Confirm(ctx *gin.Context){
	param := ctx.Param("id")
	id,err := strconv.Atoi(param)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	item, err := h.DB.Get(id)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	ctx.HTML(200,"confirm.html",gin.H{"Item":item})
}