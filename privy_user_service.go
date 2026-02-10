package privyclient

type PrivyUserService struct {
	// Directly embed the generated UserService to expose all its methods through PrivyUserService
	UserService
	logger logger
}

// newPrivyUserService creates a new wrapped user service.
// This is unexported so only PrivyClient can construct it.
func newPrivyUserService(service UserService, logger logger) *PrivyUserService {
	return &PrivyUserService{
		UserService: service,
		logger:      logger,
	}
}
