package main

import (
	"fmt"
	"time"
	"github.com/btcsuite/btcutil"
	"github.com/btcsuite/btcd/chaincfg"
	libbitcoin "github.com/OpenBazaar/go-libbitcoinclient"
)

func main() {
	servers := []libbitcoin.Server{
		libbitcoin.Server{
			Url:"tcp://libbitcoin3.openbazaar.org:9091",
			PublicKey:"",
		},
	}
	client := libbitcoin.NewLibbitcoinClient(servers, &chaincfg.MainNetParams)

	tx := "2d3024e7d75d4f12c4b879916fa0ffeca7e3d3d2885a789841542888304463a2"
	client.FetchUnconfirmedTransaction(tx, func(i interface{}, err error){
		if err != nil {
			fmt.Println(err.Error())

		} else {
			fmt.Println(i.(btcutil.Tx))
		}
	})
	time.Sleep(10 *time.Second)
}