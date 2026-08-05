// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package privyclient

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/privy-io/go-sdk/internal/apijson"
	"github.com/privy-io/go-sdk/internal/requestconfig"
	"github.com/privy-io/go-sdk/option"
	"github.com/privy-io/go-sdk/packages/respjson"
)

// Operations related to transactions
//
// TransactionService contains methods and other services that help with
// interacting with the Privy API API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTransactionService] method instead.
type TransactionService struct {
	Options []option.RequestOption
}

// NewTransactionService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewTransactionService(opts ...option.RequestOption) (r TransactionService) {
	r = TransactionService{}
	r.Options = opts
	return
}

// Get a transaction by transaction ID.
func (r *TransactionService) Get(ctx context.Context, transactionID string, opts ...option.RequestOption) (res *Transaction, err error) {
	opts = slices.Concat(r.Options, opts)
	if transactionID == "" {
		err = errors.New("missing required transaction_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/transactions/%s", url.PathEscape(transactionID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Status of a blockchain transaction submitted by Privy.
type BlockchainTransactionStatus string

const (
	BlockchainTransactionStatusBroadcasted       BlockchainTransactionStatus = "broadcasted"
	BlockchainTransactionStatusConfirmed         BlockchainTransactionStatus = "confirmed"
	BlockchainTransactionStatusExecutionReverted BlockchainTransactionStatus = "execution_reverted"
	BlockchainTransactionStatusFailed            BlockchainTransactionStatus = "failed"
	BlockchainTransactionStatusReplaced          BlockchainTransactionStatus = "replaced"
	BlockchainTransactionStatusFinalized         BlockchainTransactionStatus = "finalized"
	BlockchainTransactionStatusProviderError     BlockchainTransactionStatus = "provider_error"
	BlockchainTransactionStatusPending           BlockchainTransactionStatus = "pending"
)

// A transaction from a Privy wallet.
type Transaction struct {
	ID        string  `json:"id" api:"required"`
	Caip2     string  `json:"caip2" api:"required"`
	CreatedAt float64 `json:"created_at" api:"required"`
	// Status of a blockchain transaction submitted by Privy.
	//
	// Any of "broadcasted", "confirmed", "execution_reverted", "failed", "replaced",
	// "finalized", "provider_error", "pending".
	Status            BlockchainTransactionStatus `json:"status" api:"required"`
	TransactionHash   string                      `json:"transaction_hash" api:"required"`
	WalletID          string                      `json:"wallet_id" api:"required"`
	ReferenceID       string                      `json:"reference_id" api:"nullable"`
	Sponsored         bool                        `json:"sponsored"`
	UserOperationHash string                      `json:"user_operation_hash"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                respjson.Field
		Caip2             respjson.Field
		CreatedAt         respjson.Field
		Status            respjson.Field
		TransactionHash   respjson.Field
		WalletID          respjson.Field
		ReferenceID       respjson.Field
		Sponsored         respjson.Field
		UserOperationHash respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Transaction) RawJSON() string { return r.JSON.raw }
func (r *Transaction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
