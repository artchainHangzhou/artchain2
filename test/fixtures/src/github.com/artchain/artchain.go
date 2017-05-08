package main

import (
    "encoding/json"
    "errors"
    "fmt"
    "time"

    "github.com/hyperledger/fabric/core/chaincode/shim"
    pb "github.com/hyperledger/fabric/protos/peer"
)

var (
    layout = "2006-01-02 15:04:05"
    loc    *time.Location
)

func init() {
    loc, _ = time.LoadLocation("Asia/Shanghai")
}

type User struct {
    UserId     string   `json:"userId"`
    UserName   string   `json:"userName"`
    Coin       int64    `json:"coin"`
    IPList     []string `json:"ipList"`
    IPDetails  []IP     `json:"ipDetails"`
    Version    string   `json:"version"`
    CreateTime string   `json:"createTime"`
    UpdateTime string   `json:"updateTime"`
}

func (user *User) AddUser(stub shim.ChaincodeStubInterface) error {
    userbytes, err := json.Marshal(user)
    if err != nil {
        fmt.Println("AddUser Marshal fail:", err.Error())
        return errors.New("AddUser Marshal fail:" + err.Error())
    }
    err = stub.PutState("User:"+user.UserId, userbytes)
    if err != nil {
        fmt.Println("AddUser PutState fail:", err.Error())
        return errors.New("AddUser PutState Error" + err.Error())
    }

    return nil
}

type Org struct {
    OrgId      string `json:"orgId"`
    OrgName    string `json:"orgName"`
    Coin       int64  `json:"coin"`
    Version    string `json:"version"`
    CreateTime string `json:"createTime"`
    UpdateTime string `json:"updateTime"`
}

func (org *Org) InitOrg(stub shim.ChaincodeStubInterface) error {
    orgbytes, err := json.Marshal(org)
    if err != nil {
        fmt.Println("InitOrg Marshal fail:", err.Error())
        return errors.New("InitOrg Marshal fail:" + err.Error())
    }
    err = stub.PutState("Org:"+org.OrgId, orgbytes)
    if err != nil {
        fmt.Println("InitOrg PutState fail:", err.Error())
        return errors.New("InitOrg PutState Error" + err.Error())
    }

    return nil
}

type IP struct {
    Id            string `json:"id"`
    Name          string `json:"name"`
    Author        string `json:"author"`
    Description   string `json:"description"`
    Authorization string `json:"description"`
    PDF           string `json:"pdf"`
    SubIP         SubIP  `json:"subIP"`
    Version       string `json:"version"`
    CreateTime    string `json:"createTime"`
    UpdateTime    string `json:"updateTime"`
}

type SubIP struct {
    Id         string `json:"id"`
    Price      int64  `json:"Price"`
    State      string `json:"state"` // 1-在售 2-已售 3-可用 4-消耗
    UseRange   string `json:"useRange"`
    Version    string `json:"version"`
    CreateTime string `json:"createTime"`
    UpdateTime string `json:"updateTime"`
}

type SimpleChaincode struct {
}

func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
    fmt.Println("########### ArtChain Init ###########")
    function, args := stub.GetFunctionAndParameters()
    fmt.Println("function:" + function)
    for _, a := range args {
        fmt.Println("args:" + a)
    }

    org := &Org{
        OrgId:      "org001",
        OrgName:    "Art",
        Coin:       100000000,
        Version:    "v1.0.0",
        CreateTime: time.Now().In(loc).Format(layout),
        UpdateTime: time.Now().In(loc).Format(layout),
    }
    err := org.InitOrg(stub)
    if err != nil {
        return shim.Error("Unknown function call:" + err.Error())
    }

    user := &User{
        UserId:     "test001",
        UserName:   "测试1",
        Coin:       100000,
        IPList:     nil,
        IPDetails:  nil,
        Version:    "v1.0.0",
        CreateTime: time.Now().In(loc).Format(layout),
        UpdateTime: time.Now().In(loc).Format(layout),
    }
    err = user.AddUser(stub)
    if err != nil {
        return shim.Error("Init AddUser fail:" + err.Error())
    }

    return shim.Success(nil)
}

func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
    fmt.Println("########### ArtChain Invoke ###########")
    function, args := stub.GetFunctionAndParameters()
    fmt.Println("function:" + function)
    for _, a := range args {
        fmt.Println("args:" + a)
    }

    if function != "invoke" {
        return shim.Error("Unknown function call:" + function)
    }

    if len(args) != 2 {
        return shim.Error("Incorrect number of arguments. Expecting at least 2")
    }

    if args[0] == "apply" {
        return t.apply(stub, args[1])
    } else if args[0] == "buy" {
        return t.buy(stub, args[1])
    } else if args[0] == "use" {
        return t.use(stub, args[1])
    } else if args[0] == "sell" {
        return t.sell(stub, args[1])
    } else if args[0] == "recharge" {
        return t.recharge(stub, args[1])
    } else if args[0] == "queryUserList" {
        return t.queryUserList(stub, args[1])
    } else if args[0] == "queryIPList" {
        return t.queryIPList(stub, args[1])
    } else if args[0] == "queryUser" {
        return t.queryUser(stub, args[1])
    } else if args[0] == "queryOrg" {
        return t.queryOrg(stub, args[1])
    }


    return shim.Error("Unknown action, check the first argument:" + args[0])
}

func (t *SimpleChaincode) queryUserList(stub shim.ChaincodeStubInterface, args string) pb.Response {
    fmt.Println(args)
    return shim.Success(nil)
}

func (t *SimpleChaincode) queryIPList(stub shim.ChaincodeStubInterface, args string) pb.Response {
    fmt.Println(args)
    return shim.Success(nil)
}

func (t *SimpleChaincode) queryUser(stub shim.ChaincodeStubInterface, args string) pb.Response {
    fmt.Println(args)
    return shim.Success(nil)
}

func (t *SimpleChaincode) queryOrg(stub shim.ChaincodeStubInterface, args string) pb.Response {
    fmt.Println(args)
    var org Org
    orgBytes, err := stub.GetState("Org:" + args)
    if err != nil {
        fmt.Println("queryOrg: GetState fail:", err.Error())
    }
    err = json.Unmarshal(orgBytes, &org)
    if err != nil {
        fmt.Println("getBillById: Unmarshal fail:", err.Error())
    }

    fmt.Println(org)
    return shim.Success(orgBytes)
}

func (t *SimpleChaincode) apply(stub shim.ChaincodeStubInterface, args string) pb.Response {
    fmt.Println(args)
    return shim.Success(nil)
}

func (t *SimpleChaincode) buy(stub shim.ChaincodeStubInterface, args string) pb.Response {
    fmt.Println(args)
    return shim.Success(nil)
}

func (t *SimpleChaincode) use(stub shim.ChaincodeStubInterface, args string) pb.Response {
    fmt.Println(args)
    return shim.Success(nil)
}

func (t *SimpleChaincode) sell(stub shim.ChaincodeStubInterface, args string) pb.Response {
    fmt.Println(args)
    return shim.Success(nil)
}

func (t *SimpleChaincode) recharge(stub shim.ChaincodeStubInterface, args string) pb.Response {
    fmt.Println(args)
    return shim.Success(nil)
}

func main() {
    err := shim.Start(new(SimpleChaincode))
    if err != nil {
        fmt.Printf("Error starting Simple chaincode: %s", err)
    }
}
