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

func (s *WalletServer) HttpHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "COMEME LOS HUEVOS")
}

// Inicializa el servidor, lee la lista de nodos e inicializa las carteras.
func (s *WalletServer) Run() {
    // var command string

    // TODO: Cargar Wallets guardadas

    fmt.Println("Initializing walletServer...")
    s.Sync()
    fmt.Println("Wallet sync completed succesfully." )

    fmt.Println("Starting HTTP server")
    http.HandleFunc("/", s.HttpHandler)
    http.ListenAndServe(":11811", nil)
    fmt.Println("Server started; and terminated?")

    // START VERY DEPRECATED CODE
    // server := &http.Server{
    //     Addr: ":11811",
    //     Handler: s.HttpHandler,
    //     ReadTimeout: 10 * time.Seconds,
    //     WriteTimeout: 10 * time.Seconds,
    //     MaxHeaderBytes: 1 << 20
    // }
    // END VERY DEPRECATED CODE
    fmt.Println("Server started")

    // START DEPRECATED CODE
    // for {
    //     fmt.Scanln(&command)
    //     if command == "stop" {
    //         s.Stop()
    //         break
    //     } else {
    //     fmt.Println("Uknown command")
    //     fmt.Println(command)
    //     }
    // }
    // END DEPRECADTED CODE

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
