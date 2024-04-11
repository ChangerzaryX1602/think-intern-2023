package services

import repository "think-intern-2023/repositories"

type thinkService struct {
	thinkRepo repository.ThinkRepository
}

func NewThinkService(thinkRepo repository.ThinkRepository) thinkService {
	return thinkService{
		thinkRepo: thinkRepo,
	}
}
func (s thinkService) Think(request Request) (Response, error) {
	return Response{Result: request.Input}, nil
}
// ไม่จำเป็นต้องใช้ service
// ไม่จำเป็นต้องใช้ service
// ไม่จำเป็นต้องใช้ service