package main

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"strings"
	"time"
	"unsafe"
)


func main() {
	ctx := context.Background()
	fmt.Printf("%#v\n", client.FromEnv)
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithTimeout(ctx, 10 * time.Minute)
	defer cancel()

	images, err := cli.ImageList(ctx, types.ImageListOptions{})
	if err != nil {
		panic(err)
	}
	for _, i := range images {
			fmt.Println("tags:", len(i.RepoTags))
		if strings.Contains(i.ID, "3964ce7b8458") {
			for _, v := range (i.RepoTags) {
				fmt.Println("tag=", v)
			}
		}
	}
	fmt.Println("len:", len(images), unsafe.Sizeof(images))
}
