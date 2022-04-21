package endpoints

import (
	"MafiaREST/schemes"
	"MafiaREST/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
)

func AddUser(ctx *gin.Context) {
	if !verifyHandle(ctx) {
		return
	}

	var newUser schemes.User
	if err := ctx.BindJSON(&newUser); err != nil || !newUser.Validate() {
		ctx.JSON(http.StatusBadRequest, fillInError(_INVALID_USER_INFO))
		return
	}

	res, err := handle.AddUser(&newUser)
	utils.NotifyOnError("Couldn't add user", err)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			ctx.JSON(http.StatusForbidden, fillInError(_DUP_MAIL))
		} else {
			ctx.JSON(http.StatusInternalServerError, fillInError(_ADD_USR_ERR))
		}

		return
	}
	log.Println("Added user", res.InsertedID)

	ctx.JSON(http.StatusCreated, newUser)
}

func GetUser(ctx *gin.Context) {
	if !verifyHandle(ctx) {
		return
	}

	uid, err := primitive.ObjectIDFromHex(ctx.Param("id"))
	utils.NotifyOnError("", err)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, fillInError(_INVALID_UID))
		return
	}

	res, err := handle.GetUserById(uid)
	utils.NotifyOnError("", err)
	if err != nil {
		ctx.JSON(http.StatusNotFound, _GET_USER_ERR)
		return
	}

	ctx.JSON(http.StatusOK, *res)
}

func GetUsers(ctx *gin.Context) {
	if !verifyHandle(ctx) {
		return
	}

	res, err := handle.GetAllUsers()
	utils.NotifyOnError(_GET_USERS_ERR, err)
	if err != nil {
		ctx.JSON(http.StatusNotFound, fillInError(_GET_USERS_ERR))
	} else {
		ctx.JSON(http.StatusOK, *res)
	}
}

func UpdateUser(ctx *gin.Context) {
	if !verifyHandle(ctx) {
		return
	}

	uid, err := primitive.ObjectIDFromHex(ctx.Param("id"))
	utils.NotifyOnError("", err)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, fillInError(_INVALID_UID))
		return
	}

	var newUser schemes.User
	if err := ctx.BindJSON(&newUser); err != nil || !newUser.Validate() {
		ctx.JSON(http.StatusBadRequest, fillInError(_INVALID_USER_INFO))
		return
	}

	err = handle.UpdateUserById(uid, &newUser)
	utils.NotifyOnError("", err)
	if err != nil {
		ctx.JSON(http.StatusNotFound, _GET_USER_ERR)
		return
	}

	ctx.JSON(http.StatusOK, fillInError(_SUCCESS))
}

func DeleteUser(ctx *gin.Context) {
	if !verifyHandle(ctx) {
		return
	}

	uid, err := primitive.ObjectIDFromHex(ctx.Param("id"))
	utils.NotifyOnError("", err)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, fillInError(_INVALID_UID))
		return
	}

	err = handle.DeleteUserById(uid)
	utils.NotifyOnError("", err)
	if err != nil {
		ctx.JSON(http.StatusNotFound, _GET_USER_ERR)
		return
	}

	// deleting user deletes associated stats
	err = handle.DeleteUserStatsByUID(uid)
	utils.NotifyOnError("", err)
	if err != nil {
		ctx.JSON(http.StatusNotFound, _MISSING_STATS)
		return
	}

	ctx.JSON(http.StatusOK, fillInError(_SUCCESS))
}
