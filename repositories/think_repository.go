package repositories

import "gorm.io/gorm"

type thinkRepository struct {
	db *gorm.DB
}

func NewThinkRepository(db *gorm.DB) thinkRepository {
	return thinkRepository{
		db: db,
	}
}
func (r thinkRepository) Think(request Request) (Response, error) {
	return Response{}, nil
}

// ไม่จำเป็นต้องใช้ repository
// ไม่จำเป็นต้องใช้ repository
// ไม่จำเป็นต้องใช้ repository
