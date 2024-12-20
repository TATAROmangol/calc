package main

import "example.com/m/pkg/server"

func main() {
    srv := server.NewServer()
    srv.Run()
}