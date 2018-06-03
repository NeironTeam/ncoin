// server.go
// Autor: NeironTeam
// Licencia: MIT License, Copyright (c) 2018 Neiron

package main

import (
	"fmt"
	ncw "github.com/NeironTeam/ncoin-wallet"
	"net/http"
	"time"
)

// Manager de wallets, capaz de conectarse a la red.
type WalletServer struct {
	onlineWallets       []ncw.Wallet      // Wallet "instances"
	nodeList            []string          // Server-IPs
	pendingTransactions []ncw.Transaction // Transaciones pendientes de enviar
	server *http.Server
}

type WalletHandler func(w http.ResponseWriter, r *http.Request)

func (s *WalletServer) BalanceHandler() WalletHandler {
	return function(w http.ResponseWriter, r *htpp.Request) {
		var address string = r.URL.Query().Get("address")
		// TODO: Read address wallet ballance
		fmt.Println(address)
		w.Write([]byte(address))
	}
}

func (s *WalletServer) ChainHandler(w http.ResponseWriter, r *http.Request) {}

// Inicializa el servidor, lee la lista de nodos e inicializa las carteras.
func (ws *WalletServer) Run() {
	var s *http.ServeMux = http.NewServeMux()
	s.HandleFunc("/balance", ws.BalanceHandler())
	s.HandleFunc("/chain", ws.ChainHandler())

	ws.server = &http.Server{Addr: host, Handler: s}
	
	// TODO: Cargar Wallets guardadas
	fmt.Println("Initializing walletServer...")
	s.Sync()
	fmt.Println("Wallet sync completed succesfully.")

	// TODO: Load address and host from enviroment or set by default
	fmt.Println("Starting HTTP server")
	http.ListenAndServe(":11811", nil)
	fmt.Println("Server terminated?")

	// TODO: Guardar Wallets cargadas?
}

// Detiene el servidor, asegura los cambios, cierra las conexiones, etc...
func (s *WalletServer) Stop() {
	fmt.Println("Stopping server...")
}

func main() {
	var a []ncw.Wallet
	var s []string
	var t []ncw.Transaction
	server := WalletServer{a, s, t}

	server.Run()
	time.Sleep(time.Second * 1)
}
