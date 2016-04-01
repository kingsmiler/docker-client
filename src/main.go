package main

import (
    "fmt"

    "github.com/docker/engine-api/client"
    "github.com/docker/engine-api/types"
    "golang.org/x/net/context"
    "net/http"
)

func main() {
    // 明确指定不使用代理
    transport := &http.Transport{Proxy: nil}

    defaultHeaders := map[string]string{"User-Agent": "engine-api-cli-1.0"}
    cli, err := client.NewClient(
        "unix:///var/run/docker.sock",
        "v1.22",
        &http.Client{Transport: transport},
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
    }
}