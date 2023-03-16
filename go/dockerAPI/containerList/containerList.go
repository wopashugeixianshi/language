package main

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func main() {

	ctx := context.Background()
	fmt.Printf("%#v\n", client.FromEnv)
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}

	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	for _, container := range containers {
		fmt.Println("name:", container.Names, "imageName:", container.Image, "port:", container.Ports, "status:", container.Status,
			"ID:", container.ID)
	}
}
