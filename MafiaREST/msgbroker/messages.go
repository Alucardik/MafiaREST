package msgbroker

import "MafiaREST/schemes"

type Task struct {
	User  schemes.User      `json:"user"`
	Stats schemes.UserStats `json:"stats"`
}
