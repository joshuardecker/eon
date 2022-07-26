package client

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/Sucks-To-Suck/LuncheonNetwork/mempool"
	"github.com/Sucks-To-Suck/LuncheonNetwork/transactions"
)

type Server struct {
	mem *mempool.Mempool
}

// Inits the ServerClient.
// Input is the mempool that will be used by the server.
// Returns the new ServerClient
func Init(mempool *mempool.Mempool) Server {

	sc := new(Server)
	sc.mem = mempool

	return *sc
}

// Initiates the mux for the server.
// Returns the ServerMux of all of the Handled functions of the client.
func (s Server) InitMux() http.ServeMux {

	mux := http.NewServeMux()

	// Handled Funcs
	mux.HandleFunc("/tx", s.AddTx)

	return *mux
}

// Adds a tx to the mempool.
// Inputs are the requests and writer from the http request.
// Returns nothing.
func (s Server) AddTx(w http.ResponseWriter, r *http.Request) {

	// Get the body of the http message.
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {

		w.WriteHeader(http.StatusBadRequest)
		return
	}

	tx := new(transactions.LuTx)

	// Get the tx into the tx struct
	err = json.Unmarshal(body, tx)

	if err != nil {

		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Add the tx to the mempool
	s.mem.AddTx(tx)
}
