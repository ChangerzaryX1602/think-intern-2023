package services

type Request struct {
	Input int32 `json:"input"`
}
type Response struct {
	Result int32 `json:"result"`
}
type ThinkService interface {
	Think(Request) (Response, error)
}

// ไม่จำเป็นต้องใช้ service
// ไม่จำเป็นต้องใช้ service
// ไม่จำเป็นต้องใช้ service
