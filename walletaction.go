// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package privyclient

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/privy-io/go-sdk/internal/apiquery"
	"github.com/privy-io/go-sdk/internal/requestconfig"
	"github.com/privy-io/go-sdk/option"
	"github.com/privy-io/go-sdk/packages/param"
)

// Operations related to wallet actions
//
// WalletActionService contains methods and other services that help with
// interacting with the Privy API API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWalletActionService] method instead.
type WalletActionService struct {
	Options []option.RequestOption
}

// NewWalletActionService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewWalletActionService(opts ...option.RequestOption) (r WalletActionService) {
	r = WalletActionService{}
	r.Options = opts
	return
}

// Get the current status of a wallet action by its ID. Use `?include=steps` to
// include step-level details.
func (r *WalletActionService) Get(ctx context.Context, actionID string, params WalletActionGetParams, opts ...option.RequestOption) (res *WalletActionResponseUnion, err error) {
	if !param.IsOmitted(params.PrivyAuthorizationSignature) {
		opts = append(opts, option.WithHeader("privy-authorization-signature", fmt.Sprintf("%v", params.PrivyAuthorizationSignature.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if params.WalletID == "" {
		err = errors.New("missing required wallet_id parameter")
		return nil, err
	}
	if actionID == "" {
		err = errors.New("missing required action_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/wallets/%s/actions/%s", url.PathEscape(params.WalletID), actionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

type WalletActionGetParams struct {
	// ID of the wallet.
	WalletID string `path:"wallet_id" api:"required" json:"-"`
	// Request authorization signature. If multiple signatures are required, they
	// should be comma separated.
	PrivyAuthorizationSignature param.Opt[string] `header:"privy-authorization-signature,omitzero" json:"-"`
	// Expandable relations to include on a wallet action response.
	//
	// Any of "steps".
	Include WalletActionInclude `query:"include,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [WalletActionGetParams]'s query parameters as `url.Values`.
func (r WalletActionGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
