package main

import (
	"fmt"
	"unsafe"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/api/types/container"
)

func main(){
	// 360
	imageVar := types.ImageInspect{
		ID: "asjdfaskldfjsklafjd",
	}
	// 48
	containerVar := types.ContainerJSON{}
	jsonVar := types.ContainerJSONBase{}
	netVar := types.NetworkSettings{}
	ep := network.EndpointSettings{}
	// 304
	confVar := container.Config{}
	fmt.Println("image:", unsafe.Sizeof(imageVar))
	fmt.Println("container:", unsafe.Sizeof(containerVar), " json:", unsafe.Sizeof(jsonVar), " netVar:", unsafe.Sizeof(netVar), " ep:", unsafe.Sizeof(ep))
	fmt.Println("image-conf:", unsafe.Sizeof(confVar))
}
