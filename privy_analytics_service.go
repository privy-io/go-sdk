package privyclient

type PrivyAnalyticsService struct {
	// Directly embed the generated AnalyticsService to expose all its methods through PrivyAnalyticsService
	AnalyticsService
	logger Logger
}

// newPrivyAnalyticsService creates a new wrapped analytics service.
// This is unexported so only PrivyClient can construct it.
func newPrivyAnalyticsService(service AnalyticsService, logger Logger) *PrivyAnalyticsService {
	return &PrivyAnalyticsService{
		AnalyticsService: service,
		logger:           logger,
	}
}
