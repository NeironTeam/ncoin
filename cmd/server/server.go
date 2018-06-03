// server.go
// Autor: NeironTeam
// Licencia: MIT License, Copyright (c) 2018 Neiron

package main

import (
	"fmt"
	ncw "github.com/NeironTeam/ncoin-wallet"
	internal "github.com/NeironTeam/ncoin-wallet/internal"
	"net/http"
	"time"
)

// Manager de wallets, capaz de conectarse a la red.
type WalletServer struct {
	wallet *ncw.Wallet
	pendingTransactions []ncw.Transaction // Transaciones pendientes de enviar
	server *http.Server
}

type WalletHandler func(w http.ResponseWriter, r *http.Request)

func (s *WalletServer) BalanceHandler() WalletHandler {
	return func(w http.ResponseWriter, r *http.Request) {
		var address string = r.URL.Query().Get("address")
		// TODO: Read address wallet ballance
		fmt.Println(address)
		w.Write([]byte(address))
	}
}

func (s *WalletServer) ChainHandler() WalletHandler {
	return func(w http.ResponseWriter, r *http.Request) {}
}

// Inicializa el servidor, lee la lista de nodos e inicializa las carteras.
func (ws *WalletServer) Run() {
	var s *http.ServeMux = http.NewServeMux()
	s.HandleFunc("/balance", ws.BalanceHandler())
	s.HandleFunc("/chain", ws.ChainHandler())

	var host string = internal.GetHost()
	fmt.Println(host)
	ws.server = &http.Server{Addr: host, Handler: s}
	
	// TODO: Cargar Wallets guardadas
	fmt.Println("Initializing walletServer...")
	ws.Sync()
	fmt.Println("Wallet sync completed succesfully.")

	// TODO: Load address and host from enviroment or set by default
	fmt.Println("Starting HTTP server")
	if e := http.ListenAndServe(host, s); e != nil {
		fmt.Println(e)
	}
	fmt.Println("Server terminated?")

	// TODO: Guardar Wallets cargadas?
}

func (ws *WalletServer) Sync() {}

// Detiene el servidor, asegura los cambios, cierra las conexiones, etc...
func (s *WalletServer) Stop() {
	fmt.Println("Stopping server...")
}

func main() {
	// TODO: Load wallet
	var server *WalletServer = &WalletServer{}

	server.Run()
	time.Sleep(time.Second * 1)
}
