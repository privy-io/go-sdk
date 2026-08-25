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
	"github.com/privy-io/go-sdk/internal/apiquery"
	shimjson "github.com/privy-io/go-sdk/internal/encoding/json"
	"github.com/privy-io/go-sdk/internal/requestconfig"
	"github.com/privy-io/go-sdk/option"
)

// Operations related to fiat onramping and offramping
//
// WalletDepositAccountFiatService contains methods and other services that help
// with interacting with the Privy API API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWalletDepositAccountFiatService] method instead.
type WalletDepositAccountFiatService struct {
	Options []option.RequestOption
}

// NewWalletDepositAccountFiatService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewWalletDepositAccountFiatService(opts ...option.RequestOption) (r WalletDepositAccountFiatService) {
	r = WalletDepositAccountFiatService{}
	r.Options = opts
	return
}

// Creates a Bridge Virtual Account linked to a wallet. Fiat sent to the returned
// deposit instructions will be converted to the specified crypto asset and
// delivered to the wallet.
func (r *WalletDepositAccountFiatService) New(ctx context.Context, walletID string, body WalletDepositAccountFiatNewParams, opts ...option.RequestOption) (res *FiatDepositAccountResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if walletID == "" {
		err = errors.New("missing required wallet_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/wallets/%s/deposit_accounts/fiat", url.PathEscape(walletID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Returns a list of fiat deposit accounts linked to a wallet.
func (r *WalletDepositAccountFiatService) List(ctx context.Context, walletID string, query WalletDepositAccountFiatListParams, opts ...option.RequestOption) (res *ListFiatDepositAccountsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if walletID == "" {
		err = errors.New("missing required wallet_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/wallets/%s/deposit_accounts/fiat", url.PathEscape(walletID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Returns a single fiat deposit account linked to a wallet.
func (r *WalletDepositAccountFiatService) Get(ctx context.Context, depositAccountID string, query WalletDepositAccountFiatGetParams, opts ...option.RequestOption) (res *FiatDepositAccountResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.WalletID == "" {
		err = errors.New("missing required wallet_id parameter")
		return nil, err
	}
	if depositAccountID == "" {
		err = errors.New("missing required deposit_account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/wallets/%s/deposit_accounts/fiat/%s", url.PathEscape(query.WalletID), url.PathEscape(depositAccountID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type WalletDepositAccountFiatNewParams struct {
	// Request body for creating a Bridge fiat deposit account linked to a wallet.
	CreateFiatDepositAccountRequestBody CreateFiatDepositAccountRequestBody
	paramObj
}

func (r WalletDepositAccountFiatNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.CreateFiatDepositAccountRequestBody)
}
func (r *WalletDepositAccountFiatNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WalletDepositAccountFiatListParams struct {
	// Supported fiat orchestration providers.
	//
	// Any of "bridge".
	Provider OrchestrationProvider `query:"provider,omitzero" api:"required" json:"-"`
	// The Privy API environment.
	//
	// Any of "sandbox", "production".
	Environment Environment `query:"environment,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [WalletDepositAccountFiatListParams]'s query parameters as
// `url.Values`.
func (r WalletDepositAccountFiatListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type WalletDepositAccountFiatGetParams struct {
	// The ID of the wallet.
	WalletID string `path:"wallet_id" api:"required" json:"-"`
	paramObj
}
