package users

type Service interface {
	Find(username string) (*User, error)
	FindID(username string) (*int, error)
	Login(username string, password string) (*User, error)
	CreateUser(register RegisterUser) (*User, error)
}

type userService struct {
	repo Repository
}

func NewUserService(repo Repository) Service {
	return &userService{
		repo: repo,
	}
}

func (s *userService) Find(username string) (*User, error) {
	return s.repo.GetUserByUsername(username)
}

func (s *userService) FindID(username string) (*int, error) {
	return s.repo.GetIDByUsername(username)
}

func (s *userService) Login(username string, password string) (*User, error) {
	user, err := s.repo.GetUserByUsernameAndPassword(username, password)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userService) CreateUser(register RegisterUser) (*User, error) {
	user, err := s.repo.InsertUser(register)
	if err != nil {
		return nil, err
	}

	return user, nil
}
