package main

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/api/types/events"
	"github.com/docker/docker/client"
)

func main() {
	ctx := context.Background()

	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}

	eventFilter := filters.NewArgs()
	eventFilter.Add("type", events.ContainerEventType)
	evtOption := types.EventsOptions{
		Filters: eventFilter,
	}

	events, errChan := cli.Events(ctx, evtOption)

	for {
		select {
		case event := <-events:
			fmt.Printf("get event %v action %v\n", event.Actor.Attributes["name"], event.Action)
		case er := <-errChan:
			fmt.Println(er)
		}
	}
}
