package privyclient

type PrivyKeyQuorumService struct {
	// Directly embed the generated KeyQuorumService to expose all its methods through PrivyKeyQuorumService
	KeyQuorumService
	logger Logger
}

// newPrivyKeyQuorumService creates a new wrapped key quorum service.
// This is unexported so only PrivyClient can construct it.
func newPrivyKeyQuorumService(service KeyQuorumService, logger Logger) *PrivyKeyQuorumService {
	return &PrivyKeyQuorumService{
		KeyQuorumService: service,
		logger:           logger,
	}
}
