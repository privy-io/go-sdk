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
// WalletEarnEthereumIncentiveService contains methods and other services that help
// with interacting with the Privy API API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWalletEarnEthereumIncentiveService] method instead.
type WalletEarnEthereumIncentiveService struct {
	Options []option.RequestOption
}

// NewWalletEarnEthereumIncentiveService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewWalletEarnEthereumIncentiveService(opts ...option.RequestOption) (r WalletEarnEthereumIncentiveService) {
	r = WalletEarnEthereumIncentiveService{}
	r.Options = opts
	return
}

// Claim incentive rewards for a wallet.
func (r *WalletEarnEthereumIncentiveService) Claim(ctx context.Context, walletID string, params WalletEarnEthereumIncentiveClaimParams, opts ...option.RequestOption) (res *EarnIncentiveClaimActionResponse, err error) {
	if !param.IsOmitted(params.PrivyAuthorizationSignature) {
		opts = append(opts, option.WithHeader("privy-authorization-signature", fmt.Sprintf("%v", params.PrivyAuthorizationSignature.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if walletID == "" {
		err = errors.New("missing required wallet_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/wallets/%s/earn/ethereum/incentive/claim", url.PathEscape(walletID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

type WalletEarnEthereumIncentiveClaimParams struct {
	// Input for claiming incentive rewards.
	EarnIncentiveClaimRequestBody EarnIncentiveClaimRequestBody
	// Request authorization signature. If multiple signatures are required, they
	// should be comma separated.
	PrivyAuthorizationSignature param.Opt[string] `header:"privy-authorization-signature,omitzero" json:"-"`
	paramObj
}

func (r WalletEarnEthereumIncentiveClaimParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.EarnIncentiveClaimRequestBody)
}
func (r *WalletEarnEthereumIncentiveClaimParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
