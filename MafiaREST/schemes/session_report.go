package schemes

const (
	SESSION_WIN  uint8 = 1
	SESSION_LOSE uint8 = 0
)

type SessionReport struct {
	Outcome  uint8  `json:"outcome" binding:"required"`
	Duration uint64 `json:"duration" binding:"required"`
}
