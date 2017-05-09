package main

import (
    "encoding/json"
    "encoding/hex"
    "errors"
    "fmt"
    "time"
    "bytes"

    "github.com/hyperledger/fabric/core/chaincode/shim"
    pb "github.com/hyperledger/fabric/protos/peer"
)

var (
    layout = "2006-01-02 15:04:05"
    data = "20060102150405"
    loc    *time.Location
)

func init() {
    loc, _ = time.LoadLocation("Asia/Shanghai")
}

type User struct {
    DocType string `json:"docType"`
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

    err = stub.PutState("User:" + user.UserId, userbytes)
    if err != nil {
        fmt.Println("AddUser PutState fail:", err.Error())
        return errors.New("AddUser PutState Error" + err.Error())
    }
    
    return nil
}

type Org struct {
    DocType string `json:"docType"`
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

    err = stub.PutState("Org:" + org.OrgId, orgbytes)
    if err != nil {
        fmt.Println("InitOrg PutState fail:", err.Error())
        return errors.New("InitOrg PutState Error" + err.Error())
    }

    return nil
}

type IP struct {
    DocType string `json:"docType"`
    IPId            string `json:"ipId"`
    IPName          string `json:"ipName"`
    Author        string `json:"author"`
    Description   string `json:"description"`
    Authorization string `json:"description"`
    ProposalUrl           string `json:"proposalUrl"`
    PictureUrl           string `json:"pictureUrl"`
    Owner      string `json:"owner"`
    Price      int64  `json:"Price"`
    State      string `json:"state"` // 1-在售 2-已售 3-可用 4-消耗
    Version       string `json:"version"`
    CreateTime    string `json:"createTime"`
    UpdateTime    string `json:"updateTime"`
}

func (ip *IP) AddIP(stub shim.ChaincodeStubInterface) error {
    ipbytes, err := json.Marshal(ip)
    if err != nil {
        fmt.Println("AddIP Marshal fail:", err.Error())
        return errors.New("AddIP Marshal fail:" + err.Error())
    }

    err = stub.PutState("IP:" + ip.IPId, ipbytes)
    if err != nil {
        fmt.Println("AddIP PutState fail:", err.Error())
        return errors.New("AddIP PutState Error" + err.Error())
    }
    
    return nil
}

type ReqApply struct {
    UserId      string `json:"userId"`
    IPName      string `json:"ipName"`
    Author      string `json:"author"`
    Description string `json:"description"`
    ProposalUrl string `json:"proposalUrl"`
    PictureUrl  string `json:"pictureUrl"`
    Price       int64  `json:"Price"`
    Total       int    `json:"total"`
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
        DocType:"Org",
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
        DocType:"User",
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

    user1 := &User{
        DocType:"User",
        UserId:     "atest001",
        UserName:   "测试1",
        Coin:       10000,
        IPList:     nil,
        IPDetails:  nil,
        Version:    "v1.0.0",
        CreateTime: time.Now().In(loc).Format(layout),
        UpdateTime: time.Now().In(loc).Format(layout),
    }
    err = user1.AddUser(stub)
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

    queryString := fmt.Sprintf("{\"selector\":{\"docType\":\"User\"}}")
    queryResults, err := getQueryResultForQueryString(stub, queryString)
    if err != nil {
        fmt.Println("queryUser getQueryResultForQueryString fail:", err.Error())
        return shim.Error(err.Error())
    }

    return shim.Success(queryResults)
}

func (t *SimpleChaincode) queryIPList(stub shim.ChaincodeStubInterface, args string) pb.Response {
    fmt.Println(args)
    queryString := fmt.Sprintf("{\"selector\":{\"docType\":\"IP\"}}")
    queryResults, err := getQueryResultForQueryString(stub, queryString)
    if err != nil {
        fmt.Println("queryUser getQueryResultForQueryString fail:", err.Error())
        return shim.Error(err.Error())
    }

    return shim.Success(queryResults)
}

func (t *SimpleChaincode) queryUser(stub shim.ChaincodeStubInterface, args string) pb.Response {
    fmt.Println(args)
    var user User
    userBytes, err := stub.GetState("User:" + args)
    if err != nil {
        fmt.Println("queryUser GetState fail:", err.Error())
        return shim.Error(err.Error())
    }
    err = json.Unmarshal(userBytes, &user)
    if err != nil {
        fmt.Println("queryUser Unmarshal fail:", err.Error())
        return shim.Error(err.Error())
    }

    fmt.Println(user)
    return shim.Success(userBytes)
}

func (t *SimpleChaincode) queryOrg(stub shim.ChaincodeStubInterface, args string) pb.Response {
    fmt.Println(args)
    var org Org
    orgBytes, err := stub.GetState("Org:" + args)
    if err != nil {
        fmt.Println("queryOrg GetState fail:", err.Error())
    }
    err = json.Unmarshal(orgBytes, &org)
    if err != nil {
        fmt.Println("queryOrg Unmarshal fail:", err.Error())
    }

    fmt.Println(org)
    return shim.Success(orgBytes)
}

func (t *SimpleChaincode) apply(stub shim.ChaincodeStubInterface, args string) pb.Response {
    fmt.Println(args)
    
    body, err := hex.DecodeString(args)
    if err != nil {
        fmt.Println("apply DecodeString fail:" + err.Error())
        return shim.Error("apply DecodeString fail:" + err.Error())
    }
	var req ReqApply
    err = json.Unmarshal(body, &req)
    if err != nil {
        fmt.Println("apply Unmarshal fail:" + err.Error())
        return shim.Error("apply Unmarshal fail:" + err.Error())
    } 

	fmt.Println(req)
    for i := 1; i <= req.Total; i++ {
        ip := &IP {
            DocType:"IP",
            IPId: time.Now().In(loc).Format(data) + fmt.Sprintf("%06d", i),
            IPName: req.IPName,
            Author: req.Author,
            Description: req.Description,
            ProposalUrl: req.ProposalUrl,
            PictureUrl: req.PictureUrl,
            Owner: req.UserId,
            Price: req.Price,
            State: "1",
            Version: "v1.0.0",
            CreateTime: time.Now().In(loc).Format(layout),
            UpdateTime: time.Now().In(loc).Format(layout),
        }

        err := ip.AddIP(stub)
        if err != nil {
            fmt.Println("apply AddIP fail:" + err.Error())
            return shim.Error("apply AddIP fail:" + err.Error())
        }
    }

    return shim.Success(nil)
}

func (t *SimpleChaincode) buy(stub shim.ChaincodeStubInterface, args string) pb.Response {
    fmt.Println(args)
    return shim.Success(nil)
}

func (t *SimpleChaincode) use(stub shim.ChaincodeStubInterface, args string) pb.Response {
    fmt.Println(args)
    var ip IP
    ipBytes, err := stub.GetState("IP:" + args)
    if err != nil {
        fmt.Println("use GetState fail:" + err.Error())
        return shim.Error(err.Error())
    }
    err = json.Unmarshal(ipBytes, &ip)
    if err != nil {
        fmt.Println("use Unmarshal fail:" + err.Error())
        return shim.Error(err.Error())
    }

    if ip.State != "3" {
        fmt.Println("ip.State:" + ip.State)
        return shim.Error("ip.State:" + ip.State)
    }

    ip.State = "4"
    err = ip.AddIP(stub)
    if err != nil {
        fmt.Println("use AddIP fail:" + err.Error())
        return shim.Error(err.Error())
    }

    return shim.Success(nil)
}

func (t *SimpleChaincode) sell(stub shim.ChaincodeStubInterface, args string) pb.Response {
    fmt.Println(args)
    var ip IP
    ipBytes, err := stub.GetState("IP:" + args)
    if err != nil {
        fmt.Println("sell GetState fail:" + err.Error())
        return shim.Error(err.Error())
    }
    err = json.Unmarshal(ipBytes, &ip)
    if err != nil {
        fmt.Println("sell Unmarshal fail:" + err.Error())
        return shim.Error(err.Error())
    }

    if ip.State != "3" {
        fmt.Println("ip.State:" + ip.State)
        return shim.Error("ip.State:" + ip.State)
    }

    ip.State = "1"
    err = ip.AddIP(stub)
    if err != nil {
        fmt.Println("sell AddIP fail:", err.Error())
        return shim.Error("sell AddIP fail:" + err.Error())
    }

    return shim.Success(nil)
}

func (t *SimpleChaincode) recharge(stub shim.ChaincodeStubInterface, args string) pb.Response {
    fmt.Println(args)
    return shim.Success(nil)
}

func getQueryResultForQueryString(stub shim.ChaincodeStubInterface, queryString string) ([]byte, error) {

	fmt.Printf("- getQueryResultForQueryString queryString:\n%s\n", queryString)

	resultsIterator, err := stub.GetQueryResult(queryString)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing QueryRecords
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		_, value, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString(string(value))
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- getQueryResultForQueryString queryResult:\n%s\n", buffer.String())

	return buffer.Bytes(), nil
}

func main() {
    err := shim.Start(new(SimpleChaincode))
    if err != nil {
        fmt.Printf("Error starting Simple chaincode: %s", err)
    }
}

