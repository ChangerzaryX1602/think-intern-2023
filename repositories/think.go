package repositories

type Request struct {
	Result int32 `json:"result"`
}
type Response struct {
	Input int32 `json:"input"`
}
type ThinkRepository interface {
	Think(request Request) (Response, error)
}

// ไม่จำเป็นต้องใช้ repository
// ไม่จำเป็นต้องใช้ repository
// ไม่จำเป็นต้องใช้ repository
