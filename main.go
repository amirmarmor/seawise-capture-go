package main

import "seawise.com/capture/server"

func main() {
	entryPoint := server.Entrypoint{}
	entryPoint.Execute()
}

