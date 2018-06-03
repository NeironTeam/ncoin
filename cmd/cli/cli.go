package main

import (
	"fmt"
	ncw "github.com/NeironTeam/ncoin-wallet"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

const DEFAULT_WALLET_PORT = "11811"
const DEFAULT_LOCAL_HOST = "http://localhost:"

const HELP_MESSAGE = `Usage:
	new
	start <address>
	balance <address>
	mine
	transaction	<address>	<amount>
	chain	<nBlocks>
`
const ERROR_NO_ARGUMENTS = "No arguments. \n" + HELP_MESSAGE
const ERROR_INVALID_COMMAND = "Invalid command usage.\n" + HELP_MESSAGE
const ERROR_COMMAND_NOT_FOUND = "Command not found.\n" + HELP_MESSAGE
const ERROR_GET_FAILED = "Something went wrong.\n"

type transaction_params []string

func main() {
	var size int = len(os.Args)
	if size < 2 {
		log.Fatal(ERROR_NO_ARGUMENTS)
	}

	switch os.Args[1] {
	case "new":
		if size != 2 {
			log.Fatal(ERROR_INVALID_COMMAND)
		} else {
			newWallet()
		}

	case "start":
		if size != 3 {
			log.Fatal(ERROR_INVALID_COMMAND)
		} else {
			start(os.Args[2])
		}

	case "balance":
		if size != 3 {
			log.Fatal(ERROR_INVALID_COMMAND)
		} else {
			balance(os.Args[2])
		}

	case "mine":
		if size != 2 {
			log.Fatal(ERROR_INVALID_COMMAND)
		} else {
			mine()
		}

	case "transaction":
		if size != 4 {
			log.Fatal(ERROR_INVALID_COMMAND)
		} else {
			amount, err := strconv.ParseFloat(os.Args[3], 64)
			if err != nil {
				log.Fatal(ERROR_INVALID_COMMAND)
			}
			transaction(os.Args[2], amount)
		}

	case "check-transaction":
		if size != 3 {
			log.Fatal(ERROR_INVALID_COMMAND)
		} else {
			checkTransaction(os.Args[2])
		}

	case "chain":
		if size != 3 {
			log.Fatal(ERROR_INVALID_COMMAND)
		} else {
			nBlocks, err := strconv.Atoi(os.Args[2])
			if err != nil {
				log.Fatal(ERROR_INVALID_COMMAND)
			}

			chain(nBlocks)
		}

	default:
		log.Fatal(ERROR_COMMAND_NOT_FOUND)
	}

	return
}

// Returns the enviroment var
func getPort() (walletPort string) {
	if walletPort = os.Getenv("WALLET_PORT"); walletPort == "" {
		walletPort = DEFAULT_WALLET_PORT
	}
	return
}

func newWallet() {
	log.Println("Creating new wallet ...")

	if wallet, err := ncw.NewWallet(); err != nil {
		log.Fatal(ERROR_GET_FAILED + err.Error())
	} else {
		log.Printf("Wallet created successfully with address %s", wallet.Address())
	}
}

func start(address string) {
	//TODO: Run server with docker
	log.Printf("Starting wallet on address %s", address)
}

func balance(address string) {
	log.Printf("Checking balance on wallet %s", address)
	var (
		e        error
		client   *http.Client = &http.Client{}
		endpoint *url.URL
		uri      string = fmt.Sprintf("%s%s/balance", DEFAULT_LOCAL_HOST, getPort())
	)

	if endpoint, e = url.Parse(uri); e != nil {
		log.Fatal(ERROR_GET_FAILED + e.Error())
	}

	var q url.Values = endpoint.Query()
	q.Add("address", address)
	endpoint.RawQuery = q.Encode()

	var req *http.Request
	if req, e = http.NewRequest(http.MethodGet, endpoint.String(), nil); e != nil {
		log.Fatal(ERROR_GET_FAILED + e.Error())
	}

	var resp *http.Response
	if resp, e = client.Do(req); e != nil {
		log.Fatal(ERROR_GET_FAILED + e.Error())
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		bodyString := string(bodyBytes)
		log.Println(bodyString)
	}
}

func mine() {
	log.Println("Starting mining ...")
	//TODO: add miners
}

func transaction(address string, amount float64) {
	log.Printf("Stating %f ncoins transaction to %s", amount, address)
	var (
		e        error
		client   *http.Client = &http.Client{}
		endpoint *url.URL
		uri      string = fmt.Sprintf("%s%s/transaction", DEFAULT_LOCAL_HOST, getPort())
	)
	if endpoint, e = url.Parse(uri); e != nil {
		log.Fatal(ERROR_GET_FAILED + e.Error())
	}

	var q url.Values = endpoint.Query()
	q.Add("address", address)
	q.Add("amount", fmt.Sprintf("%f", amount))
	endpoint.RawQuery = q.Encode()

	var req *http.Request
	if req, e = http.NewRequest(http.MethodGet, endpoint.String(), nil); e != nil {
		log.Fatal(ERROR_GET_FAILED + e.Error())
	}

	var res *http.Response
	if res, e = client.Do(req); e != nil {
		log.Fatal(ERROR_GET_FAILED + e.Error())
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		bodyString := string(bodyBytes)
		log.Println(bodyString)
	}
}

func checkTransaction(hash string) {
	log.Printf("Check the state of transaction %s", hash)
	var (
		e        error
		client   *http.Client = &http.Client{}
		endpoint *url.URL
		uri      string = fmt.Sprintf("%s%s/check-transaction", DEFAULT_LOCAL_HOST, getPort())
	)
	if endpoint, e = url.Parse(uri); e != nil {
		log.Fatal(ERROR_GET_FAILED + e.Error())
	}

	var q url.Values = endpoint.Query()
	q.Add("hash", hash)
	endpoint.RawQuery = q.Encode()

	var req *http.Request
	if req, e = http.NewRequest(http.MethodGet, endpoint.String(), nil); e != nil {
		log.Fatal(ERROR_GET_FAILED + e.Error())
	}

	var res *http.Response
	if res, e = client.Do(req); e != nil {
		log.Fatal(ERROR_GET_FAILED + e.Error())
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		bodyString := string(bodyBytes)
		log.Println(bodyString)
	}
}

func chain(nBlocks int) {
	log.Printf("Displaying last %d blocks", nBlocks)
	var (
		e        error
		client   *http.Client = &http.Client{}
		endpoint *url.URL
		uri      string = fmt.Sprintf("%s%s/chain", DEFAULT_LOCAL_HOST, getPort())
	)
	if endpoint, e = url.Parse(uri); e != nil {
		log.Fatal(ERROR_GET_FAILED + e.Error())
	}

	var q url.Values = endpoint.Query()
	q.Add("nblocks", fmt.Sprintf("%d", nBlocks))
	endpoint.RawQuery = q.Encode()

	var req *http.Request
	if req, e = http.NewRequest(http.MethodGet, endpoint.String(), nil); e != nil {
		log.Fatal(ERROR_GET_FAILED + e.Error())
	}

	var res *http.Response
	if res, e = client.Do(req); e != nil {
		log.Fatal(ERROR_GET_FAILED + e.Error())
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		bodyString := string(bodyBytes)
		log.Println(bodyString)
	}
}
