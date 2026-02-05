package privyclient

type PrivyTransactionService struct {
	// Directly embed the generated TransactionService to expose all its methods through PrivyTransactionService
	TransactionService
	logger logger
}

// newPrivyTransactionService creates a new wrapped transaction service.
// This is unexported so only PrivyClient can construct it.
func newPrivyTransactionService(service TransactionService, logger logger) *PrivyTransactionService {
	return &PrivyTransactionService{
		TransactionService: service,
		logger:             logger,
	}
}
