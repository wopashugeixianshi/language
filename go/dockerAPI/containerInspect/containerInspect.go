package main

import (
	"context"
	"fmt"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func main() {
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}
	
	//获取docker信息
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(30)*time.Second)
	defer cancel()
	info, err := cli.Info(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Println("driver:", info.Driver)

	//获取容器列表
	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	for _, container := range containers {
		inspCtx, inspCancel := context.WithTimeout(context.Background(), time.Duration(1)*time.Second)
		// 超时
		//inspCtx, inspCancel := context.WithTimeout(context.Background(), time.Duration(1)*time.Microsecond)
		//info, err := cli.ContainerInspect(inspCtx, container.ID)
		//info, err := cli.ContainerInspect(inspCtx, "1d81f37b4b3a")
		info, err := cli.ContainerInspect(inspCtx, "1asjdfklasdf")
		if err != nil {
			fmt.Println("err:", err)
			continue
		}
		fmt.Println("container.Name:", container.Names, "status:", container.State, "\nlogPath:", info.LogPath)
		fmt.Println("container.Name:", info.Name, "status:", info.State.Status, "\nlogPath:", info.LogPath)
		inspCancel()
		break
	}
}
