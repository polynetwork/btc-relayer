package main

import (
	"flag"
	"fmt"
	"github.com/polynetwork/btc-relayer"
	"github.com/polynetwork/btc-relayer/log"
	"github.com/polynetwork/btc-relayer/utils"
	"github.com/polynetwork/poly/common/password"
	"os"
	"time"
)

var (
	confFile  string
	walletPwd string
)

func init() {
	flag.StringVar(&confFile, "conf-file", "./conf.json", "configuration file for btc relayer")
	flag.StringVar(&walletPwd, "wallet-pwd", "", "your poly chain wallet password")
}

func main() {
	flag.Parse()

	conf, err := btc_relayer.NewRelayerConfig(confFile)
	if err != nil {
		log.Fatalf("failed to new a config: %v", err)
		return
	}
	var pwd []byte
	if walletPwd != "" {
		pwd = []byte(walletPwd)
	} else if conf.PolyObConf.WalletPwd == "" {
		fmt.Println("enter your poly wallet password:")
		if pwd, err = password.GetPassword(); err != nil {
			log.Fatalf("password is not found in config file and enter password failed: %v", err)
			os.Exit(1)
		}
		fmt.Println("done")
	} else {
		pwd = []byte(conf.PolyObConf.WalletPwd)
	}
	log.InitLog(conf.LogLevel, os.Stdout)
	r, err := btc_relayer.NewBtcRelayer(conf, pwd)
	if err != nil {
		log.Fatalf("Failed to new a relayer: %v", err)
		return
	}
	if conf.SleepTime > 0 {
		utils.SleepTime = time.Duration(conf.SleepTime)
	}
	go r.BtcListen()
	go r.RelayHeaders()
	go r.RelayTx()
	go r.SendCCIFromDB()

	go r.PolyListen()
	go r.Broadcast()
	go r.ReBroadcast()

	select {}
}
