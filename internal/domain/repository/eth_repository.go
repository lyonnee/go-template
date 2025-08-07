package repository

import "context"

type EthRepository interface {
	ETHBalance(ctx context.Context, address string) (string, error)
}
