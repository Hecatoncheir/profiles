package main

import "profiles/httpengine"

func main() {
	httpServer := httpengine.NewHTTPEngine("v1.0")
	httpServer.PowerUp("localhost", 8081)
}
