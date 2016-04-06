package main

import (
    "fmt"

    "github.com/docker/engine-api/client"
    "github.com/docker/engine-api/types"
    "golang.org/x/net/context"
    "sync"
)

func main() {
    defaultHeaders := map[string]string{"User-Agent": "engine-api-cli-1.0"}

    cli, err := client.NewClient(
        "unix:///var/run/docker.sock",
        "v1.22",
        nil,
        defaultHeaders,
    )

    if err != nil {
        panic(err)
    }


    closeChan := make(chan error)

    // waitFirst is a WaitGroup to wait first stat data's reach for each container
    waitFirst := &sync.WaitGroup{}

    cStats := stats{}

    // getContainerList simulates creation event for all previously existing
    // containers (only used when calling `docker stats` without arguments).
    getContainerList := func() {
        options := types.ContainerListOptions{All: true}

        cs, err := cli.ContainerList(context.Background(), options)
        if err != nil {
            closeChan <- err
        }
        for _, container := range cs {
            s := &containerStats{Name: container.ID[:12]}
            if cStats.add(s) {
                waitFirst.Add(1)
                go s.Collect(cli, true, waitFirst)
            }
        }
    }

    getContainerList()

    options := types.ContainerListOptions{All: true}
    containers, err := cli.ContainerList(context.Background(), options)
    if err != nil {
        panic(err)
    }

    for _, c := range containers {
        fmt.Println(c.ID)

        cli.ContainerStats(context.Background(), c.ID, false)
    }
}