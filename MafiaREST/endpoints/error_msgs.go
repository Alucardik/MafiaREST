package endpoints

const (
	_SUCCESS           = "operation succeeded"
	_INTERNAL_ERR      = "Something went wrong"
	_DUP_MAIL          = "User with the given email already exists"
	_GET_USERS_ERR     = "Couldn't get the list of users"
	_GET_USER_ERR      = "Couldn't find user with the specified ID"
	_INVALID_USER_INFO = "Wrong scheme format"
	_INVALID_UID       = "Invalid user id"
	_INVALID_REPORT    = "Invalid session report scheme"
	_FAILED_REPORT     = "There was an error in report generation, try again later"
	_MISSING_REPORT    = "The requested report is not ready yet or have not been requested"
	_ADD_USR_ERR       = "Couldn't add user to the database"
	_MISSING_STATS     = "Couldn't find associated stats"
)

type respMsg struct {
	Msg string `json:"msg"`
}

func fillInMsg(msg string) respMsg {
	em := respMsg{Msg: msg}
	return em
}
