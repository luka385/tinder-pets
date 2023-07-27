package application

import "github.com/luka385/tinder-pets/domain"

type UserService struct {
	userRepository domain.UserRepository
}

func NewUserService(UserRepository *domain.UserRepository) *UserService {
	return &UserService{
		userRepository: *UserRepository,
	}
}

func (s *UserService) CreateUser(user *domain.User) error {
	return s.userRepository.Create(user)
}

func (s *UserService) UpdateUser(user *domain.User) error {
	return s.userRepository.Update(user)
}

func (s *UserService) DeleteUser(id string) error {
	return s.userRepository.Delete(id)
}

func (s *UserService) GetUserByID(id string) (*domain.User, error) {
	return s.userRepository.GetByID(id)
}

func (s *UserService) GetAllUser() ([]*domain.User, error) {
	return s.userRepository.GetAll()
}
