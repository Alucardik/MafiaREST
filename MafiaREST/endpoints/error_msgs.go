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
	_ADD_USR_ERR       = "Couldn't add user to the database"
	_MISSING_STATS     = "Couldn't find associated stats"
)

type errorMsg struct {
	Msg string `json:"msg"`
}

func fillInError(msg string) errorMsg {
	em := errorMsg{Msg: msg}
	return em
}
