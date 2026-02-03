// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.
package privyclient

type PrivyClientAuthService struct {
	// Directly embed the generated ClientAuthService to expose all its methods through PrivyClientAuthService
	ClientAuthService
}

// newPrivyClientAuthService creates a new wrapped client auth service.
// This is unexported so only PrivyClient can construct it.
func newPrivyClientAuthService(service ClientAuthService) *PrivyClientAuthService {
	return &PrivyClientAuthService{
		ClientAuthService: service,
	}
}
