package main

import (
	"github.com/dalyoon-practice/nomadcoin/cli"
	"github.com/dalyoon-practice/nomadcoin/db"
)

func main() {
	// wallet.Wallet()
	defer db.Close()
	cli.Start()
}
