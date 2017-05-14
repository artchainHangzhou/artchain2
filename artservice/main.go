package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "os"
)

var base BaseSetupImpl

type Result struct {
    Code    int    `json:"code"`
    Message string `json:"message"`
    //Data []byte`json:"data"`
    Data interface{} `json:"data"`
}

func init() {
    base = BaseSetupImpl{
        ConfigFile:      "../test/fixtures/config/config_test.yaml",
        ChainID:         "testchannel",
        ChannelConfig:   "../test/fixtures/channel/testchannel.tx",
        ConnectEventHub: true,
        ChainCodeID: "artchain",
        AppraisalChainCodeID: "appraisal",
        SearchChainCodeID: "search",
    }

    if err := base.Initialize(); err != nil {
        fmt.Printf("Initialize: %v", err)
        os.Exit(-1)
    }

    fmt.Println("InstallAndInstantiateArtChainCC ing ...")
    if err := base.InstallAndInstantiateArtChainCC("github.com/artchain", "v1.0"); err != nil {
        fmt.Printf("InstallAndInstantiateArtChainCC: %v", err)
        os.Exit(-1)
    }
    fmt.Println("InstallAndInstantiateArtChainCC succ!")

    fmt.Println("InstallAndInstantiateAppraisalCC ing ...")
    if err := base.InstallAndInstantiateAppraisalCC("github.com/appraisal", "v1.0"); err != nil {
        fmt.Printf("InstallAndInstantiateAppraisalCC: %v", err)
        os.Exit(-1)
    }
    fmt.Println("InstallAndInstantiateAppraisalCC succ!")

    fmt.Println("InstallAndInstantiateSearchCC ing ...")
    if err := base.InstallAndInstantiateSearchCC("github.com/search", "v1.0"); err != nil {
        fmt.Printf("InstallAndInstantiateSearchCC: %v", err)
        os.Exit(-1)
    }
    fmt.Println("InstallAndInstantiateSearchCC succ!")
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

    http.HandleFunc("/signIn", SignIn)
    http.HandleFunc("/apply", Apply)
    http.HandleFunc("/buy", Buy)
    http.HandleFunc("/use", Use)
    http.HandleFunc("/sell", Sell)

    http.HandleFunc("/upload", UpLoad)
    http.HandleFunc("/download/", DownLoad)

    http.HandleFunc("/search", Search)
    http.HandleFunc("/block", Block)

    err := http.ListenAndServe(":12345", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
