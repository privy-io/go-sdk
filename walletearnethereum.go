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
	shimjson "github.com/privy-io/go-sdk/internal/encoding/json"
	"github.com/privy-io/go-sdk/internal/requestconfig"
	"github.com/privy-io/go-sdk/option"
	"github.com/privy-io/go-sdk/packages/param"
)

// Operations related to wallet actions
//
// WalletEarnEthereumService contains methods and other services that help with
// interacting with the Privy API API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWalletEarnEthereumService] method instead.
type WalletEarnEthereumService struct {
	Options   []option.RequestOption
	Incentive WalletEarnEthereumIncentiveService
}

// NewWalletEarnEthereumService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewWalletEarnEthereumService(opts ...option.RequestOption) (r WalletEarnEthereumService) {
	r = WalletEarnEthereumService{}
	r.Options = opts
	r.Incentive = NewWalletEarnEthereumIncentiveService(opts...)
	return
}

// Deposit assets into an ERC-4626 vault.
func (r *WalletEarnEthereumService) Deposit(ctx context.Context, walletID string, params WalletEarnEthereumDepositParams, opts ...option.RequestOption) (res *EarnDepositActionResponse, err error) {
	if !param.IsOmitted(params.PrivyAuthorizationSignature) {
		opts = append(opts, option.WithHeader("privy-authorization-signature", fmt.Sprintf("%v", params.PrivyAuthorizationSignature.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if walletID == "" {
		err = errors.New("missing required wallet_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/wallets/%s/earn/ethereum/deposit", url.PathEscape(walletID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Withdraw assets from an ERC-4626 vault.
func (r *WalletEarnEthereumService) Withdraw(ctx context.Context, walletID string, params WalletEarnEthereumWithdrawParams, opts ...option.RequestOption) (res *EarnWithdrawActionResponse, err error) {
	if !param.IsOmitted(params.PrivyAuthorizationSignature) {
		opts = append(opts, option.WithHeader("privy-authorization-signature", fmt.Sprintf("%v", params.PrivyAuthorizationSignature.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if walletID == "" {
		err = errors.New("missing required wallet_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/wallets/%s/earn/ethereum/withdraw", url.PathEscape(walletID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

type WalletEarnEthereumDepositParams struct {
	// Input for depositing assets into an ERC-4626 vault. Exactly one of `amount` or
	// `raw_amount` must be provided.
	EarnDepositRequestBody EarnDepositRequestBody
	// Request authorization signature. If multiple signatures are required, they
	// should be comma separated.
	PrivyAuthorizationSignature param.Opt[string] `header:"privy-authorization-signature,omitzero" json:"-"`
	paramObj
}

func (r WalletEarnEthereumDepositParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.EarnDepositRequestBody)
}
func (r *WalletEarnEthereumDepositParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WalletEarnEthereumWithdrawParams struct {
	// Input for withdrawing assets from an ERC-4626 vault. Exactly one of `amount` or
	// `raw_amount` must be provided.
	EarnWithdrawRequestBody EarnWithdrawRequestBody
	// Request authorization signature. If multiple signatures are required, they
	// should be comma separated.
	PrivyAuthorizationSignature param.Opt[string] `header:"privy-authorization-signature,omitzero" json:"-"`
	paramObj
}

func (r WalletEarnEthereumWithdrawParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.EarnWithdrawRequestBody)
}
func (r *WalletEarnEthereumWithdrawParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
