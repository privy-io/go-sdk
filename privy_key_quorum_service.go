// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.
package privyclient

type PrivyKeyQuorumService struct {
	// Directly embed the generated KeyQuorumService to expose all its methods through PrivyKeyQuorumService
	KeyQuorumService
}

// newPrivyKeyQuorumService creates a new wrapped key quorum service.
// This is unexported so only PrivyClient can construct it.
func newPrivyKeyQuorumService(service KeyQuorumService) *PrivyKeyQuorumService {
	return &PrivyKeyQuorumService{
		KeyQuorumService: service,
	}
}
