package privyclient

import (
	"context"
	"net/url"

	"github.com/privy-io/go-sdk/option"
)

// PrivyWalletActionService wraps the generated WalletActionService to
// path-escape action IDs.
type PrivyWalletActionService struct {
	WalletActionService
}

// Get fetches a wallet action, path-escaping the action ID the generated method interpolates verbatim.
func (s *PrivyWalletActionService) Get(ctx context.Context, actionID string, params WalletActionGetParams, opts ...option.RequestOption) (*WalletActionResponseUnion, error) {
	return s.WalletActionService.Get(ctx, url.PathEscape(actionID), params, opts...)
}
