/* walletServer.go
+  Autor: NeironTeam
+  Licencia: MIT License, Copyright (c) 2018 Neiron
*/
package main

import ("fmt"
        "time"
        "wallet"
        "transaction")

/* Manager de wallets, capaz de conectarse a la red. */
type WalletServer struct {
    online_wallets       []Wallet       // Wallet "instances"
    node_list            []string       // Server-IPs
    pending_transactions []Transaction  // Transaciones pendientes de enviar
  }

/* Inicializa el servidor, lee la lista de nodos e inicializa las carteras. */
func (s *WalletServer) Run() {
    var command string

    fmt.Println("Initializing walletServer...")
    s.Sync()
    fmt.Println("Sync for 0 wallets completed succesfully." )
    fmt.Println("Waiting for orders...")

    for {
        fmt.Scanln(&command)
        if command == "stop" {
            s.Stop()
            break
        } else {
        fmt.Println("Uknown command")
        fmt.Println(command)
        }
    }
}

/* Sincroniza con la blockchain para actualizar el saldo de todas las
+  carteras que maneja y enviar las transacciones pendientes,
+  además, actualiza su lista de nodos. */
func (s *WalletServer) Sync() {
  fmt.Println("Syncing...")
}

/* Actualiza el estado de una sola cartera. Toma la dirección de la cartera
+  como argumento */
func (s *WalletServer) WalletSync(address uint64) {

}

/* Crea una nueva cartera y la añade a sus carteras */
func NewWallet() Wallet {

}

/* Detiene el servidor, asegura los cambios, cierra las conexiones, etc... */
func (s *WalletServer) Stop() {
  fmt.Println("Stopping server...")
}

func main() {
    var a []string
    server := WalletServer{a}

    server.Run()
    time.Sleep(time.Seconds * 1)
}
