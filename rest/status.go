package rest

import (
	"encoding/json"
	"net/http"

	"github.com/dalyoon-practice/nomadcoin/blockchain"
)

func status(rw http.ResponseWriter, r *http.Request) {
	json.NewEncoder(rw).Encode(blockchain.Blockchain())
}
