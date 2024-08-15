package main

import "github.com/JairDavid/Probien-Backend/cmd/provider"

func main() {
	p := provider.New()
	p.Build().Run()
}
