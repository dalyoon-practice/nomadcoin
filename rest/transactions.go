package rest

import (
	"encoding/json"
	"net/http"

	"github.com/dalyoon-practice/nomadcoin/blockchain"
	"github.com/dalyoon-practice/nomadcoin/utils"
)

type addTxPayload struct {
	To     string
	Amount int
}

func transactions(rw http.ResponseWriter, r *http.Request) {
	var payload addTxPayload
	utils.HandleErr(json.NewDecoder(r.Body).Decode(&payload))
	err := blockchain.Mempool.AddTx(payload.To, payload.Amount)
	if err != nil {
		json.NewEncoder(rw).Encode(errResponse{"not enough funds"})
	}
	rw.WriteHeader(http.StatusCreated)
}
