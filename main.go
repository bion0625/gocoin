package main

import (
	"github.com/gocoin/explorer"
	"github.com/gocoin/rest"
)

func main() {
	go explorer.Start(3000)
	rest.Start(4000)
}
