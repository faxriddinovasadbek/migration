package main

import (
	"encoding/json"
	"fmt"
	"handlar_tes/model"
	"handlar_tes/storge"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {

	router := gin.Default()

	router.POST("/user/create", CreateUser)
	router.GET("/user/get", GetUser)
	router.DELETE("/user/delete", DeleteUser)
	router.PUT("/user/update", UpdateUser)
	router.GET("/user/getall", GetAllUsers)

	fmt.Println("Server is running...")

	err := router.Run("localhost:7777")

	if err != nil {
		fmt.Println("Error while running server", err)
	}
}

func CreateUser(ctx *gin.Context) {
	bodyByte, err := io.ReadAll(ctx.Request.Body)

	if err != nil {
		log.Println("error while getting body", err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var user *model.User

	if err = json.Unmarshal(bodyByte, &user); err != nil {
		log.Println("error while unmarshalling body", err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	respUser, err := storge.CreateUser(user)

	if err != nil {
		log.Println("error while creating body", err)
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	ctx.JSONP(http.StatusCreated, respUser)
}

func GetAllUsers(ctx *gin.Context) {
	page := ctx.Query("page")

	intPage, err := strconv.Atoi(page)
	if err != nil {
		log.Println("Error while converting page")
		ctx.AbortWithError(http.StatusBadRequest, err)
	}

	limit := ctx.Query("limit")

	intLimit, err := strconv.Atoi(limit)
	if err != nil {
		log.Println("Error while converting page")
		ctx.AbortWithError(http.StatusBadRequest, err)
	}

	users, err := storge.GetAll(intPage, intLimit)
	if err != nil {
		log.Println("Error while getting all users", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
	}

	ctx.JSONP(http.StatusOK, users)
}

func GetUser(ctx *gin.Context) {

	id := ctx.Query("id")
	// id := ctx.Param("id")

	user, err := storge.Get(id)

	if err != nil {
		log.Println("Error while getting user")
		ctx.AbortWithError(http.StatusInternalServerError, err)
	}

	ctx.JSONP(http.StatusOK, user)
}

func DeleteUser(ctx *gin.Context) {

	id := ctx.Query("id")

	err := storge.DeleteUser(id)

	if err != nil {
		log.Println("Error while deleting user")
		ctx.AbortWithError(http.StatusInternalServerError, err)
	}

	ctx.JSONP(http.StatusOK, "delete")
}

func UpdateUser(ctx *gin.Context) {

	bodyByte, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println("error while getting body", err)
		ctx.AbortWithError(http.StatusBadRequest, err)
	}

	var user *model.User
	err = json.Unmarshal(bodyByte, &user)
	if err != nil {
		panic(err)
	}

	id := ctx.Query("id")

	respouser, err := storge.UptadeUser(id, user)
	if err != nil {
		log.Println("Error not update user", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
	}

	ctx.JSONP(http.StatusOK, respouser)
}

func GetAllRole(ctx *gin.Context) {
	page := ctx.Query("page")

	intPage, err := strconv.Atoi(page)

	if err != nil {
		log.Println("Error while converting page")
		ctx.AbortWithError(http.StatusBadRequest, err)
	}

	limit := ctx.Query("limit")

	intLimit, err := strconv.Atoi(limit)

	if err != nil{
		log.Println("Error while converting page")
		ctx.AbortWithError(http.StatusBadRequest, err)
	}

	role := ctx.Query("role")	
}