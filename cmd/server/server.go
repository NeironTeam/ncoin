// server.go
// Autor: NeironTeam
// Licencia: MIT License, Copyright (c) 2018 Neiron

package main

import (
    "fmt"
    "time"
    "net/http"
    ncw "github.com/NeironTeam/ncoin-wallet"
)

// Manager de wallets, capaz de conectarse a la red.
type WalletServer struct {
    online_wallets       []ncw.Wallet       // Wallet "instances"
    node_list            []string       // Server-IPs
    pending_transactions []ncw.Transaction  // Transaciones pendientes de enviar
  }

func (s *WalletServer) BalanceHandler(w http.ResponseWriter, r *http.Request) {
    var address string = r.URL.Query().Get("address")
    fmt.Println(address)
    w.Write([]byte(address))
}

func (s *WalletServer) ChainHandler(w http.ResponseWriter, r *http.Request) {

}

// Inicializa el servidor, lee la lista de nodos e inicializa las carteras.
func (s *WalletServer) Run() {
    // var command string

    // TODO: Cargar Wallets guardadas

    fmt.Println("Initializing walletServer...")
    s.Sync()
    fmt.Println("Wallet sync completed succesfully." )

    fmt.Println("Starting HTTP server")
    http.HandleFunc("/balance", s.BalanceHandler)
    http.HandleFunc("/chain", s.ChainHandler)
    http.ListenAndServe(":11811", nil)
    fmt.Println("Server terminated?")

    // TODO: Guardar Wallets cargadas?

}

// Sincroniza con la blockchain para actualizar el saldo de todas las
// carteras que maneja y enviar las transacciones pendientes,
// además, actualiza su lista de nodos.
func (s *WalletServer) Sync() {
  fmt.Println("Syncing...")
}

// Actualiza el estado de una sola cartera. Toma la dirección de la cartera
// como argumento
func (s *WalletServer) WalletSync(wallet ncw.Wallet) {

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
