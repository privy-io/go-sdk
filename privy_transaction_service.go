// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.
package privyclient

type PrivyTransactionService struct {
	// Directly embed the generated TransactionService to expose all its methods through PrivyTransactionService
	TransactionService
}

// newPrivyTransactionService creates a new wrapped transaction service.
// This is unexported so only PrivyClient can construct it.
func newPrivyTransactionService(service TransactionService) *PrivyTransactionService {
	return &PrivyTransactionService{
		TransactionService: service,
	}
}
