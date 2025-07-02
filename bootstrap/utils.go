package bootstrap

import "github.com/lyonnee/go-template/pkg/idgen"

func initUtils(hostId int64) error {
	if err := idgen.Initialize(hostId); err != nil {
		return err
	}
	return nil
}
