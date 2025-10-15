// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package privyapiclient

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/stainless-sdks/privy-api-client-go/internal/apijson"
	"github.com/stainless-sdks/privy-api-client-go/internal/requestconfig"
	"github.com/stainless-sdks/privy-api-client-go/option"
	"github.com/stainless-sdks/privy-api-client-go/packages/respjson"
)

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
func (r *TransactionService) Get(ctx context.Context, transactionID string, opts ...option.RequestOption) (res *TransactionGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if transactionID == "" {
		err = errors.New("missing required transaction_id parameter")
		return
	}
	path := fmt.Sprintf("v1/transactions/%s", transactionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type TransactionGetResponse struct {
	ID        string  `json:"id,required"`
	Caip2     string  `json:"caip2,required"`
	CreatedAt float64 `json:"created_at,required"`
	// Any of "broadcasted", "confirmed", "execution_reverted", "failed", "replaced",
	// "finalized", "provider_error", "pending".
	Status          TransactionGetResponseStatus `json:"status,required"`
	TransactionHash string                       `json:"transaction_hash,required"`
	WalletID        string                       `json:"wallet_id,required"`
	Sponsored       bool                         `json:"sponsored"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		Caip2           respjson.Field
		CreatedAt       respjson.Field
		Status          respjson.Field
		TransactionHash respjson.Field
		WalletID        respjson.Field
		Sponsored       respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TransactionGetResponse) RawJSON() string { return r.JSON.raw }
func (r *TransactionGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TransactionGetResponseStatus string

const (
	TransactionGetResponseStatusBroadcasted       TransactionGetResponseStatus = "broadcasted"
	TransactionGetResponseStatusConfirmed         TransactionGetResponseStatus = "confirmed"
	TransactionGetResponseStatusExecutionReverted TransactionGetResponseStatus = "execution_reverted"
	TransactionGetResponseStatusFailed            TransactionGetResponseStatus = "failed"
	TransactionGetResponseStatusReplaced          TransactionGetResponseStatus = "replaced"
	TransactionGetResponseStatusFinalized         TransactionGetResponseStatus = "finalized"
	TransactionGetResponseStatusProviderError     TransactionGetResponseStatus = "provider_error"
	TransactionGetResponseStatusPending           TransactionGetResponseStatus = "pending"
)
