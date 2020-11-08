package userdm

type UserDomainService struct {
	userRepository UserRepository
}

func NewUserDomainService(userRepo UserRepository) *UserDomainService {
	return &UserDomainService{userRepository: userRepo}
}

func (ds *UserDomainService) IsExists(userID UserID) bool {
	user, err := ds.userRepository.FindByID(userID)
	return !(err != nil || user == nil)
}
