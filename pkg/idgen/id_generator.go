package idgen

import (
	"fmt"

	"github.com/bwmarrin/snowflake"
	"github.com/lyonnee/go-template/internal/infrastructure/config"
	"github.com/lyonnee/go-template/internal/infrastructure/di"
)

var node *snowflake.Node

func init() {
	config := di.Get[config.Config]()

	newNode, err := snowflake.NewNode(config.App.HostId)
	if err != nil {
		panic(fmt.Sprintf("failed to initialize ID generator: %v", err))
	}

	node = newNode
}

func Initialize(hostId int64) error {
	var err error
	node, err = snowflake.NewNode(hostId)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func GenerateID() int64 {
	if node == nil {
		panic("ID generator not initialized")
	}

	id := node.Generate()
	return id.Int64()
}

func GenerateStringId() string {
	if node == nil {
		panic("ID generator not initialized")
	}

	id := node.Generate()
	return id.String()
}
