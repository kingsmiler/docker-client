package main

import (
    "fmt"

    "github.com/docker/engine-api/client"
    "github.com/docker/engine-api/types"
    "golang.org/x/net/context"
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