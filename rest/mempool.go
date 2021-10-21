package rest

import (
	"encoding/json"
	"net/http"

	"github.com/dalyoon-practice/nomadcoin/blockchain"
)

func mempool(rw http.ResponseWriter, r *http.Request) {
	json.NewEncoder(rw).Encode(blockchain.Mempool.Txs)
}
