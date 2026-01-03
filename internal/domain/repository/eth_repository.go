package repository

import "context"

type EthRepository interface {
	BaseRepository
	ETHBalance(ctx context.Context, address string) (string, error)
}
