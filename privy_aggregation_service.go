package privyclient

type PrivyAggregationService struct {
	// Directly embed the generated AggregationService to expose all its methods through PrivyAggregationService
	AggregationService
}

// newPrivyAggregationService creates a new wrapped aggregation service.
// This is unexported so only PrivyClient can construct it.
func newPrivyAggregationService(service AggregationService) *PrivyAggregationService {
	return &PrivyAggregationService{
		AggregationService: service,
	}
}
