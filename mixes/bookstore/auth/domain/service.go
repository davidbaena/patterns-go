package domain

// PasswordHasher defines the methods for hashing
// and comparing passwords.
type PasswordHasher interface {
	Hash(password string) (string, error)
	Compare(hash, password string) error
}

// AuthService provides methods for user registration
// and authentication.
type AuthService struct {
	userRepository UserRepository
	passwordHasher PasswordHasher
}

// NewAuthService creates a new AuthService.
func NewAuthService(userRepo UserRepository, hasher PasswordHasher) *AuthService {
	return &AuthService{
		userRepository: userRepo,
		passwordHasher: hasher,
	}
}
