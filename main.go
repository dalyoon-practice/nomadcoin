package main

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/dalyoon-practice/nomadcoin/blockchain"
)

const port string = ":4000"

type homeData struct {
	PageTitle string
	Blocks    []*blockchain.Block
}

func home(rw http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/home.html"))
	data := homeData{"This is Home", blockchain.GetBlockchain().AllBlocks()}

	tmpl.Execute(rw, data)
}

func main() {
	http.HandleFunc("/", home)
	fmt.Printf("Server listening on port http://localhost%s\n", port)
	http.ListenAndServe(port, nil)
}
