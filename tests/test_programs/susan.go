package test_programs

import (
	"context"
	"fmt"
	
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

func SusanTest() {
	
	fmt.Println("Starting Susan Testing function")
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}

	containers, err := cli.ContainerList(context.Background(), container.ListOptions{})
	if err != nil {
		panic(err)
	}

	for _, ctr := range containers {
		fmt.Printf("%s %s\n", ctr.ID, ctr.Image)
	}
}
