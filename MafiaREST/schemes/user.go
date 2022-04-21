package schemes

import (
	"MafiaREST/utils"
	"log"
	"net/mail"
	"net/url"
)

type sex uint8

func (s sex) toString() string {
	switch s {
	case SEX_MALE:
		return "male"
	case SEX_FEMALE:
		return "female"
	default:
		return ""
	}
}

const (
	SEX_MALE   sex = 0
	SEX_FEMALE sex = 1
)

type User struct {
	Name   string `json:"name" bson:"name" binding:"required"`
	Avatar string `json:"avatar" bson:"avatar" binding:"required"`
	Sex    sex    `json:"sex" bson:"sex" binding:"required"`
	Email  string `json:"email" bson:"email" binding:"required"`
}

// TODO: maybe replace bool with err and add custom errors
func (u *User) Validate() bool {
	parsedMail, err := mail.ParseAddress(u.Email)
	utils.NotifyOnError("", err)
	if err != nil {
		return false
	}

	if res, err := url.Parse(u.Avatar); err != nil || res.Scheme == "" {
		log.Println("Invalid Avatar URL")
		return false
	}

	if u.Sex != SEX_MALE && u.Sex != SEX_FEMALE {
		log.Println("Invalid Sex")
		return false
	}

	if len(u.Name) == 0 {
		log.Println("Empty Name")
		return false
	}

	u.Email = parsedMail.Address
	return true
}
