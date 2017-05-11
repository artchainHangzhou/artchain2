package main

import (
    "fmt"
    "net/http"
    "log"
	"os"
    "encoding/json"
)

var base BaseSetupImpl

type Result struct {
	Code int `json:"code"`
	Message string`json:"message"`
	//Data []byte`json:"data"`
	Data interface{}`json:"data"`
}

func init() {
	base = BaseSetupImpl{
		ConfigFile:      "../config/config_test.yaml",
		ChainID:         "testchannel",
		ChannelConfig:   "../channel/testchannel.tx",
		ConnectEventHub: true,
	}

	if err := base.Initialize(); err != nil {
		fmt.Printf("Initialize: %v", err)
		os.Exit(-1)
	}
}

func OutputJson(w http.ResponseWriter, code int, reason string, data interface{}) {
    out := &Result{code, reason, data}
    b, err := json.Marshal(out)
    if err != nil {
		fmt.Println("OutputJson fail:" + err.Error())
        return
    }

    w.Write(b)
}

func main() {
    http.HandleFunc("/InstallAndInstantiate", InstallAndInstantiate)
    http.HandleFunc("/queryOrg", QueryOrg)
    http.HandleFunc("/queryUser", QueryUser)
    http.HandleFunc("/queryUserList", QueryUserList)
    http.HandleFunc("/queryUserIPList", QueryUserIPList)
    http.HandleFunc("/queryIPList", QueryIPList)
    http.HandleFunc("/queryTransaction", QueryTransaction)
    http.HandleFunc("/queryUserTransaction", QueryUserTransaction)

    http.HandleFunc("/apply", Apply)
    http.HandleFunc("/buy", Buy)
    http.HandleFunc("/use", Use)
    http.HandleFunc("/sell", Sell)

    http.HandleFunc("/upload", UpLoad)
    http.HandleFunc("/download", DownLoad)

    err := http.ListenAndServe(":12345", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
