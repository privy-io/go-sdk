package privyclient

type PrivyKeyQuorumService struct {
	// Directly embed the generated KeyQuorumService to expose all its methods through PrivyKeyQuorumService
	KeyQuorumService
	logger logger
}

// newPrivyKeyQuorumService creates a new wrapped key quorum service.
// This is unexported so only PrivyClient can construct it.
func newPrivyKeyQuorumService(service KeyQuorumService, logger logger) *PrivyKeyQuorumService {
	return &PrivyKeyQuorumService{
		KeyQuorumService: service,
		logger:           logger,
	}
}
