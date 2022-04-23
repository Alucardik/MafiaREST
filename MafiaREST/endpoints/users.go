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

func (m *Manager) AddUser(ctx *gin.Context) {
	if !m.verify(ctx) {
		return
	}

	var newUser schemes.User
	if err := ctx.BindJSON(&newUser); err != nil || !newUser.Validate() {
		ctx.JSON(http.StatusBadRequest, fillInMsg(_INVALID_USER_INFO))
		return
	}

	res, err := m.handle.AddUser(&newUser)
	utils.NotifyOnError("Couldn't add user", err)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			ctx.JSON(http.StatusForbidden, fillInMsg(_DUP_MAIL))
		} else {
			ctx.JSON(http.StatusInternalServerError, fillInMsg(_ADD_USR_ERR))
		}

		return
	}
	log.Println("Added user", res.InsertedID)

	ctx.JSON(http.StatusCreated, newUser)
}

func (m *Manager) GetUser(ctx *gin.Context) {
	if !m.verify(ctx) {
		return
	}

	uid, err := primitive.ObjectIDFromHex(ctx.Param("id"))
	utils.NotifyOnError("", err)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, fillInMsg(_INVALID_UID))
		return
	}

	res, err := m.handle.GetUserById(uid)
	utils.NotifyOnError("", err)
	if err != nil {
		ctx.JSON(http.StatusNotFound, _GET_USER_ERR)
		return
	}

	ctx.JSON(http.StatusOK, *res)
}

func (m *Manager) GetUsers(ctx *gin.Context) {
	if !m.verify(ctx) {
		return
	}

	res, err := m.handle.GetAllUsers()
	utils.NotifyOnError(_GET_USERS_ERR, err)
	if err != nil {
		ctx.JSON(http.StatusNotFound, fillInMsg(_GET_USERS_ERR))
	} else {
		ctx.JSON(http.StatusOK, *res)
	}
}

func (m *Manager) UpdateUser(ctx *gin.Context) {
	if !m.verify(ctx) {
		return
	}

	uid, err := primitive.ObjectIDFromHex(ctx.Param("id"))
	utils.NotifyOnError("", err)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, fillInMsg(_INVALID_UID))
		return
	}

	var newUser schemes.User
	if err := ctx.BindJSON(&newUser); err != nil || !newUser.Validate() {
		ctx.JSON(http.StatusBadRequest, fillInMsg(_INVALID_USER_INFO))
		return
	}

	err = m.handle.UpdateUserById(uid, &newUser)
	utils.NotifyOnError("", err)
	if err != nil {
		ctx.JSON(http.StatusNotFound, _GET_USER_ERR)
		return
	}

	ctx.JSON(http.StatusOK, fillInMsg(_SUCCESS))
}

func (m *Manager) DeleteUser(ctx *gin.Context) {
	if !m.verify(ctx) {
		return
	}

	uid, err := primitive.ObjectIDFromHex(ctx.Param("id"))
	utils.NotifyOnError("", err)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, fillInMsg(_INVALID_UID))
		return
	}

	err = m.handle.DeleteUserById(uid)
	utils.NotifyOnError("", err)
	if err != nil {
		ctx.JSON(http.StatusNotFound, _GET_USER_ERR)
		return
	}

	// deleting user deletes associated stats
	err = m.handle.DeleteUserStatsByUID(uid)
	utils.NotifyOnError("", err)
	if err != nil {
		ctx.JSON(http.StatusNotFound, _MISSING_STATS)
		return
	}

	ctx.JSON(http.StatusOK, fillInMsg(_SUCCESS))
}
