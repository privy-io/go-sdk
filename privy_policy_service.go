// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.
package privyclient

type PrivyPolicyService struct {
	// Directly embed the generated PolicyService to expose all its methods through PrivyPolicyService
	PolicyService
}

// newPrivyPolicyService creates a new wrapped policy service.
// This is unexported so only PrivyClient can construct it.
func newPrivyPolicyService(service PolicyService) *PrivyPolicyService {
	return &PrivyPolicyService{
		PolicyService: service,
	}
}
