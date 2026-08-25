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
// UserExternalFiatAccountService contains methods and other services that help
// with interacting with the Privy API API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUserExternalFiatAccountService] method instead.
type UserExternalFiatAccountService struct {
	Options []option.RequestOption
}

// NewUserExternalFiatAccountService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewUserExternalFiatAccountService(opts ...option.RequestOption) (r UserExternalFiatAccountService) {
	r = UserExternalFiatAccountService{}
	r.Options = opts
	return
}

// Creates an external fiat account linked to a user for use in offramp transfers.
func (r *UserExternalFiatAccountService) New(ctx context.Context, userID string, body UserExternalFiatAccountNewParams, opts ...option.RequestOption) (res *ExternalFiatAccountResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if userID == "" {
		err = errors.New("missing required user_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/users/%s/external_fiat_accounts", url.PathEscape(userID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Returns a list of external fiat accounts linked to a user.
func (r *UserExternalFiatAccountService) List(ctx context.Context, userID string, query UserExternalFiatAccountListParams, opts ...option.RequestOption) (res *ListExternalFiatAccountsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if userID == "" {
		err = errors.New("missing required user_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/users/%s/external_fiat_accounts", url.PathEscape(userID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Deletes an external fiat account linked to a user.
func (r *UserExternalFiatAccountService) Delete(ctx context.Context, accountID string, body UserExternalFiatAccountDeleteParams, opts ...option.RequestOption) (res *SuccessResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if body.UserID == "" {
		err = errors.New("missing required user_id parameter")
		return nil, err
	}
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/users/%s/external_fiat_accounts/%s", url.PathEscape(body.UserID), url.PathEscape(accountID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Returns a single external fiat account linked to a user.
func (r *UserExternalFiatAccountService) Get(ctx context.Context, accountID string, query UserExternalFiatAccountGetParams, opts ...option.RequestOption) (res *ExternalFiatAccountResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.UserID == "" {
		err = errors.New("missing required user_id parameter")
		return nil, err
	}
	if accountID == "" {
		err = errors.New("missing required account_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/users/%s/external_fiat_accounts/%s", url.PathEscape(query.UserID), url.PathEscape(accountID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type UserExternalFiatAccountNewParams struct {
	// Request body for creating a Bridge external fiat account.
	CreateExternalFiatAccountRequestBody CreateExternalFiatAccountRequestBody
	paramObj
}

func (r UserExternalFiatAccountNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.CreateExternalFiatAccountRequestBody)
}
func (r *UserExternalFiatAccountNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserExternalFiatAccountListParams struct {
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

// URLQuery serializes [UserExternalFiatAccountListParams]'s query parameters as
// `url.Values`.
func (r UserExternalFiatAccountListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type UserExternalFiatAccountDeleteParams struct {
	// The DID of the user.
	UserID string `path:"user_id" api:"required" json:"-"`
	paramObj
}

type UserExternalFiatAccountGetParams struct {
	// The DID of the user.
	UserID string `path:"user_id" api:"required" json:"-"`
	paramObj
}
