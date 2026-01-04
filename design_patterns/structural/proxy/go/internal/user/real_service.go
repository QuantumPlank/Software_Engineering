package user

import (
	"fmt"
	"time"
)

type RealService struct{}

func NewRealService() *RealService {
	return &RealService{}
}

func (r *RealService) GetUser(id int) (string, error) {
	time.Sleep(2 * time.Second)

	return fmt.Sprintf("User-%d", id), nil
}
