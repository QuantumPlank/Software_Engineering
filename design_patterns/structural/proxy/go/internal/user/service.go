package user

type Service interface {
	GetUser(int) (string, error)
}
