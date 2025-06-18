package id_generator

import (
	"fmt"

	"github.com/bwmarrin/snowflake"
)

var node *snowflake.Node

func Initialize(nodeId int64) error {
	var err error
	node, err = snowflake.NewNode(nodeId)
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
