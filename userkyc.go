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
)

// Operations related to fiat onramping and offramping
//
// UserKYCService contains methods and other services that help with interacting
// with the Privy API API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUserKYCService] method instead.
type UserKYCService struct {
	Options []option.RequestOption
}

// NewUserKYCService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewUserKYCService(opts ...option.RequestOption) (r UserKYCService) {
	r = UserKYCService{}
	r.Options = opts
	return
}

// Returns KYC status for all providers the user has initiated KYC with.
func (r *UserKYCService) List(ctx context.Context, userID string, opts ...option.RequestOption) (res *KYCStatusListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if userID == "" {
		err = errors.New("missing required user_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/users/%s/kyc", url.PathEscape(userID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Generates a hosted KYC link for the user and returns the current KYC status
// snapshot.
func (r *UserKYCService) InitiateLinks(ctx context.Context, userID string, body UserKYCInitiateLinksParams, opts ...option.RequestOption) (res *KYCStatusResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if userID == "" {
		err = errors.New("missing required user_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/users/%s/kyc/links", url.PathEscape(userID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Generates a Bridge terms-of-service acceptance link for the user.
func (r *UserKYCService) InitiateTos(ctx context.Context, userID string, body UserKYCInitiateTosParams, opts ...option.RequestOption) (res *KyxTosResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if userID == "" {
		err = errors.New("missing required user_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("v1/users/%s/kyc/tos", url.PathEscape(userID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type UserKYCInitiateLinksParams struct {
	// Request body for initiating a hosted KYC flow.
	KYCLinksRequestBody KYCLinksRequestBody
	paramObj
}

func (r UserKYCInitiateLinksParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.KYCLinksRequestBody)
}
func (r *UserKYCInitiateLinksParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserKYCInitiateTosParams struct {
	// Request body for initiating Terms of Service acceptance.
	KyxTosRequestBody KyxTosRequestBody
	paramObj
}

func (r UserKYCInitiateTosParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.KyxTosRequestBody)
}
func (r *UserKYCInitiateTosParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
