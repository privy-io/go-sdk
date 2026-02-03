// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.
package privyclient

type PrivyUserService struct {
	// Directly embed the generated UserService to expose all its methods through PrivyUserService
	UserService
}

// newPrivyUserService creates a new wrapped user service.
// This is unexported so only PrivyClient can construct it.
func newPrivyUserService(service UserService) *PrivyUserService {
	return &PrivyUserService{
		UserService: service,
	}
}
