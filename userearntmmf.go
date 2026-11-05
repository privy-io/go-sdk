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
// UserEarnTmmfService contains methods and other services that help with
// interacting with the Privy API API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUserEarnTmmfService] method instead.
type UserEarnTmmfService struct {
	Options []option.RequestOption
}

// NewUserEarnTmmfService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewUserEarnTmmfService(opts ...option.RequestOption) (r UserEarnTmmfService) {
	r = UserEarnTmmfService{}
	r.Options = opts
	return
}

// Returns a user's registration with each tMMF provider and a page of the wallets
// registered for it, with each status read live from the provider.
func (r *UserEarnTmmfService) Get(ctx context.Context, userID string, query UserEarnTmmfGetParams, opts ...option.RequestOption) (res *TmmfUserResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if userID == "" {
		err = errors.New("missing required user_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/users/%s/earn/tmmf", url.PathEscape(userID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type UserEarnTmmfGetParams struct {
	// Only return the registration of this wallet.
	WalletID param.Opt[string] `query:"wallet_id,omitzero" json:"-"`
	// Cursor returned by the previous page.
	WalletsCursor param.Opt[string] `query:"wallets_cursor,omitzero" json:"-"`
	// Maximum number of wallet registrations to return.
	WalletsLimit param.Opt[int64] `query:"wallets_limit,omitzero" json:"-"`
	// The fund provider.
	Provider TmmfProvider `query:"provider,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [UserEarnTmmfGetParams]'s query parameters as `url.Values`.
func (r UserEarnTmmfGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
