package litestorage

import (
	"context"

	"github.com/tonkeeper/opentonapi/pkg/core"
	"github.com/tonkeeper/tongo"
)

func (s *LiteStorage) GetSubscriptions(ctx context.Context, address tongo.AccountID) ([]core.Subscription, error) {
	return []core.Subscription{}, nil
}

func (s *LiteStorage) GetSeqno(ctx context.Context, account tongo.AccountID) (uint32, error) {
	return s.client.GetSeqno(ctx, account)
}
