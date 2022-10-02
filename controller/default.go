package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type IDefaultController interface {
	DefaultRun(ginContext *gin.Context)
	DefaultQuery(ginContext *gin.Context)
	DefaultCreate(ginContext *gin.Context)
	DefaultModify(ginContext *gin.Context)
	DefaultRemove(ginContext *gin.Context)
}

func (i *Controller) DefaultRun(ctx *gin.Context) {
	if err := i.Logic.DefaultRun(ctx); err != nil {
		fmt.Println(err)
		ctx.Status(http.StatusBadRequest)
	} else {
		ctx.Status(http.StatusOK)
	}
}
func (i *Controller) DefaultQuery(ctx *gin.Context) {
	if err := i.Logic.DefaultQuery(ctx); err != nil {
		fmt.Println(err)
		ctx.Status(http.StatusBadRequest)
	} else {
		ctx.Status(http.StatusOK)
	}
}
func (i *Controller) DefaultCreate(ctx *gin.Context) {
	if err := i.Logic.DefaultCreate(ctx); err != nil {
		fmt.Println(err)
		ctx.Status(http.StatusBadRequest)
	} else {
		ctx.Status(http.StatusOK)
	}
}
func (i *Controller) DefaultModify(ctx *gin.Context) {
	if err := i.Logic.DefaultModify(ctx); err != nil {
		fmt.Println(err)
		ctx.Status(http.StatusBadRequest)
	} else {
		ctx.Status(http.StatusOK)
	}
}
func (i *Controller) DefaultRemove(ctx *gin.Context) {
	if err := i.Logic.DefaultRemove(ctx); err != nil {
		fmt.Println(err)
		ctx.Status(http.StatusBadRequest)
	} else {
		ctx.Status(http.StatusOK)
	}
}
func (i *Controller) DefaultTransaction(ctx *gin.Context) {
	if err := i.Logic.DefaultTransaction(ctx); err != nil {
		fmt.Println(err)
		ctx.Status(http.StatusBadRequest)
	} else {
		ctx.Status(http.StatusOK)
	}
}
