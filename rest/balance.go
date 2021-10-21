package rest

import (
	"encoding/json"
	"net/http"

	"github.com/dalyoon-practice/nomadcoin/blockchain"
	"github.com/dalyoon-practice/nomadcoin/utils"
	"github.com/gorilla/mux"
)

type balanceResponse struct {
	Address string `json:"address"`
	Balance int    `json:"balance"`
}

func balance(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	address := vars["address"]
	total := r.URL.Query().Get("total")
	switch total {
	case "true":
		amount := blockchain.BalanceByAddress(address, blockchain.Blockchain())
		json.NewEncoder(rw).Encode(balanceResponse{address, amount})
	default:
		utils.HandleErr(json.NewEncoder(rw).Encode(blockchain.UTxOutsByAddress(address, blockchain.Blockchain())))
	}
}
