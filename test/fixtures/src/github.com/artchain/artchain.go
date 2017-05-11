package main

import (
    "bytes"
    "encoding/json"
    "fmt"
    "strconv"
    "time"
    "errors"

    "github.com/hyperledger/fabric/core/chaincode/shim"
    pb "github.com/hyperledger/fabric/protos/peer"
)

var (
    layout = "2006-01-02 15:04:05"
    data   = "20060102150405"
    loc    *time.Location
)

func init() {
    loc, _ = time.LoadLocation("Asia/Shanghai")
}

type User struct {
    DocType    string `json:"docType"`
    UserId     string `json:"userId"`
    UserName   string `json:"userName"`
    Coin       int64  `json:"coin"`
    Version    string `json:"version"`
    CreateTime string `json:"createTime"`
    UpdateTime string `json:"updateTime"`
}

func (user *User) PutUser(stub shim.ChaincodeStubInterface) error {
    userbytes, err := json.Marshal(user)
    if err != nil {
        fmt.Println("PutUser Marshal fail:", err.Error())
        return errors.New("PutUser Marshal fail:" + err.Error())
    }

    err = stub.PutState("User:"+user.UserId, userbytes)
    if err != nil {
        fmt.Println("PutUser PutState fail:", err.Error())
        return errors.New("PutUser PutState Error" + err.Error())
    }

    return nil
}

type Org struct {
    DocType    string `json:"docType"`
    OrgId      string `json:"orgId"`
    OrgName    string `json:"orgName"`
    Coin       int64  `json:"coin"`
    Version    string `json:"version"`
    CreateTime string `json:"createTime"`
    UpdateTime string `json:"updateTime"`
}

func (org *Org) PutOrg(stub shim.ChaincodeStubInterface) error {
    orgbytes, err := json.Marshal(org)
    if err != nil {
        fmt.Println("PutOrg Marshal fail:", err.Error())
        return errors.New("PutOrg Marshal fail:" + err.Error())
    }

    err = stub.PutState("Org:"+org.OrgId, orgbytes)
    if err != nil {
        fmt.Println("PutOrg PutState fail:", err.Error())
        return errors.New("PutOrg PutState Error" + err.Error())
    }

    return nil
}

type IP struct {
    DocType       string `json:"docType"`
    IPId          string `json:"ipId"`
    IPName        string `json:"ipName"`
    Author        string `json:"author"`
    Description   string `json:"description"`
    Authorization string `json:"description"`
    ProposalUrl   string `json:"proposalUrl"`
    PictureUrl    string `json:"pictureUrl"`
    SubId         string `json:"subId"`
    Owner         string `json:"owner"`
    Price         int64  `json:"Price"`
    State         string `json:"state"` // 1-在售 2-持有 3-消耗
    Version       string `json:"version"`
    CreateTime    string `json:"createTime"`
    UpdateTime    string `json:"updateTime"`
}

func (ip *IP) PutIP(stub shim.ChaincodeStubInterface) error {
    ipbytes, err := json.Marshal(ip)
    if err != nil {
        fmt.Println("PutIP Marshal fail:", err.Error())
        return errors.New("PutIP Marshal fail:" + err.Error())
    }

    err = stub.PutState("IP:"+ip.IPId, ipbytes)
    if err != nil {
        fmt.Println("PutIP PutState fail:", err.Error())
        return errors.New("PutIP PutState Error" + err.Error())
    }

    return nil
}

type Transaction struct {
    DocType    string `json:"docType"`
    TxId       string `json:"txId"`
    IPName     string `json:"ipName"`
    IPId       string `json:"ipId"`
    SubId      string `json:"subId"`
    From       string `json:"from"`
    To         string `json:"to"`
    Price      int64  `json:"price"`
    Version    string `json:"version"`
    CreateTime string `json:"createTime"`
    UpdateTime string `json:"updateTime"`
}

func (tx *Transaction) PutTransaction(stub shim.ChaincodeStubInterface) error {
    txbytes, err := json.Marshal(tx)
    if err != nil {
        fmt.Println("PutTransaction Marshal fail:", err.Error())
        return errors.New("PutTransaction Marshal fail:" + err.Error())
    }

    err = stub.PutState("Transaction:" + tx.TxId, txbytes)
    if err != nil {
        fmt.Println("PutTransaction PutState fail:", err.Error())
        return errors.New("PutTransaction PutState Error" + err.Error())
    }

    return nil
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

    org1 := &Org{
        DocType:    "Org",
        OrgId:      "org001",
        OrgName:    "Art",
        Coin:       1000000,
        Version:    "v1.0.0",
        CreateTime: time.Now().In(loc).Format(layout),
        UpdateTime: time.Now().In(loc).Format(layout),
    }
    err := org1.PutOrg(stub)
    if err != nil {
        return shim.Error("Unknown function call:" + err.Error())
    }

    org2 := &Org{
        DocType:    "Org",
        OrgId:      "org002",
        OrgName:    "Artorg2",
        Coin:       1000000,
        Version:    "v1.0.0",
        CreateTime: time.Now().In(loc).Format(layout),
        UpdateTime: time.Now().In(loc).Format(layout),
    }

    err = org2.PutOrg(stub)
    if err != nil {
        return shim.Error("Unknown function call:" + err.Error())
    }

    user := &User{
        DocType:    "User",
        UserId:     "test001",
        UserName:   "测试1",
        Coin:       10000,
        Version:    "v1.0.0",
        CreateTime: time.Now().In(loc).Format(layout),
        UpdateTime: time.Now().In(loc).Format(layout),
    }
    err = user.PutUser(stub)
    if err != nil {
        return shim.Error("Init PutUser fail:" + err.Error())
    }

    user1 := &User{
        DocType:    "User",
        UserId:     "atest001",
        UserName:   "测试1",
        Coin:       10000,
        Version:    "v1.0.0",
        CreateTime: time.Now().In(loc).Format(layout),
        UpdateTime: time.Now().In(loc).Format(layout),
    }
    err = user1.PutUser(stub)
    if err != nil {
        return shim.Error("Init PutUser fail:" + err.Error())
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

    /*
       if len(args) != 2 {
           return shim.Error("Incorrect number of arguments. Expecting at least 2")
       }
    */

    if args[0] == "apply" {
        return t.apply(stub, args)
    } else if args[0] == "buy" {
        return t.buy(stub, args)
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
    } else if args[0] == "queryUserIPList" {
        return t.queryUserIPList(stub, args[1])
    } else if args[0] == "queryTransaction" {
        return t.queryTransaction(stub, args[1])
    } else if args[0] == "queryUserTransaction" {
        return t.queryUserTransaction(stub, args[1])
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

func (t *SimpleChaincode) queryUserIPList(stub shim.ChaincodeStubInterface, args string) pb.Response {
    fmt.Println(args)

    queryString := fmt.Sprintf("{\"selector\":{\"docType\":\"IP\", \"owner\":\"%s\"}}", args)
    queryResults, err := getQueryResultForQueryString(stub, queryString)
    if err != nil {
        fmt.Println("queryUser getQueryResultForQueryString fail:", err.Error())
        return shim.Error(err.Error())
    }

    return shim.Success(queryResults)
}

func (t *SimpleChaincode) queryIPList(stub shim.ChaincodeStubInterface, args string) pb.Response {
    fmt.Println(args)
    queryString := fmt.Sprintf("{\"selector\":{\"docType\":\"IP\", \"state\":\"1\"}}")
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

func (t *SimpleChaincode) queryTransaction(stub shim.ChaincodeStubInterface, args string) pb.Response {
    fmt.Println(args)

    queryString := fmt.Sprintf("{\"selector\":{\"docType\":\"Transaction\"}}")
    queryResults, err := getQueryResultForQueryString(stub, queryString)
    if err != nil {
        fmt.Println("queryUser getQueryResultForQueryString fail:", err.Error())
        return shim.Error(err.Error())
    }

    return shim.Success(queryResults)
}

func (t *SimpleChaincode) queryUserTransaction(stub shim.ChaincodeStubInterface, args string) pb.Response {
    fmt.Println(args)

    queryString := fmt.Sprintf("{\"selector\":{\"docType\":\"Transaction\", \"$or\": [{\"from\":\"%s\"},{\"to\":\"%s\"}]}}", args, args)
    queryResults, err := getQueryResultForQueryString(stub, queryString)
    if err != nil {
        fmt.Println("queryUser getQueryResultForQueryString fail:", err.Error())
        return shim.Error(err.Error())
    }

    return shim.Success(queryResults)
}

func (t *SimpleChaincode) apply(stub shim.ChaincodeStubInterface, args[] string) pb.Response {
    fmt.Println(args)

    var reqIP IP
    reqIP.Owner = args[1]
    reqIP.IPName = args[2]
    reqIP.Author = args[3]
    reqIP.Description = args[4]
    reqIP.ProposalUrl = args[5]
    reqIP.PictureUrl  = args[6]

    var err error
    reqIP.Price, err = strconv.ParseInt(args[7], 10, 64)
    if err != nil {
        fmt.Println("apply ParseInt total fail:", err.Error())
        return shim.Error(err.Error())
    }

    total, err := strconv.Atoi(args[8])
    if err != nil {
        fmt.Println("apply atoi total fail:", err.Error())
        return shim.Error(err.Error())
    }


    /*

fmt.Println("1")
    queryString := fmt.Sprintf("{\"selector\":{\"docType\":\"Org\"}}")
    queryResults, err := getQueryResultList(stub, queryString)
    if err != nil {
        fmt.Println("apply getQueryResultList fail:", err.Error())
        return shim.Error(err.Error())
    }

    orgnum := len(queryResults)
    if orgnum <= 0 {
        fmt.Println("orgnum is :", strconv.Itoa(orgnum))
        return shim.Error("orgnum is :" + strconv.Itoa(orgnum))
    }

    userbytes, err := t.GetState(stub, "User:" + req.UserId)
    if err != nil {
        fmt.Println("GetState User:", err.Error())
        return shim.Error("GetState User:" + err.Error())
    }

fmt.Println("2")
    var user User
    err = json.Unmarshal(userbytes, &user)
    if err != nil {
        fmt.Println("Apply Unmarshal fail:", err.Error())
        return shim.Error("Apply Unmarshal fail:" + err.Error())
    }

fmt.Println("3")
    if user.Coin < int64(orgnum) {
        fmt.Println("user coin is to low:", strconv.FormatInt(user.Coin, 10))
        return shim.Error("user coin is to low:" + strconv.FormatInt(user.Coin, 10))
    }

    user.Coin -= int64(orgnum)
    user.PutUser(stub)

    var org Org
    for _, orgbytes := range queryResults {
        fmt.Println(string(orgbytes))
        err = json.Unmarshal(orgbytes, &org)
        if err != nil {
            fmt.Println("Apply Unmarshal fail:", err.Error())
            return shim.Error("Apply Unmarshal fail:" + err.Error())
        }

fmt.Println("4")
        org.Coin += 1
        org.PutOrg(stub)
    }
    */

    for i := 1; i <= total; i++ {
        ip := &IP{
            DocType:     "IP",
            IPId:        time.Now().In(loc).Format(data) + fmt.Sprintf("%06d", i),
            IPName:      reqIP.IPName,
            Author:      reqIP.Author,
            Description: reqIP.Description,
            ProposalUrl: reqIP.ProposalUrl,
            PictureUrl:  reqIP.PictureUrl,
            SubId:       strconv.Itoa(i) + "/" + strconv.Itoa(total),
            Owner:       reqIP.Owner,
            Price:       reqIP.Price,
            State:       "1",
            Version:     "v1.0.0",
            CreateTime:  time.Now().In(loc).Format(layout),
            UpdateTime:  time.Now().In(loc).Format(layout),
        }

        err := ip.PutIP(stub)
        if err != nil {
            fmt.Println("apply PutIP fail:" + err.Error())
            return shim.Error("apply PutIP fail:" + err.Error())
        }
    }

    return shim.Success(nil)
}

func (t *SimpleChaincode) buy(stub shim.ChaincodeStubInterface, args []string) pb.Response {
    fmt.Println(args)
    /*
    queryString := fmt.Sprintf("{\"selector\":{\"docType\":\"Org\"}}")
    queryResults, err := getQueryResultList(stub, queryString)
    if err != nil {
        fmt.Println("apply getQueryResultList fail:", err.Error())
        return shim.Error(err.Error())
    }

    orgnum := len(queryResults)
    if orgnum <= 0 {
        fmt.Println("orgnum is :", strconv.Itoa(orgnum))
        return shim.Error("orgnum is :" + strconv.Itoa(orgnum))
    }
    */

    touser, err := t.GetUser(stub, "User:" + args[1])
    if err != nil {
        fmt.Println("GetState User:", err.Error())
        return shim.Error("GetState User:" + err.Error())
    }

    /*
    if touser.Coin < int64(orgnum) {
        fmt.Println("user coin is to low:", strconv.FormatInt(touser.Coin, 10))
        return shim.Error("user coin is to low:" + strconv.FormatInt(touser.Coin, 10))
    }
    */

    ip, err := t.GetIP(stub, "IP:" + args[2])
    if err != nil {
        fmt.Println("GetState IP:", err.Error())
        return shim.Error("GetState IP:" + err.Error())
    }

    fromuser, err := t.GetUser(stub, "User:" + ip.Owner)
    if err != nil {
        fmt.Println("GetState User:", err.Error())
        return shim.Error("GetState User:" + err.Error())
    }

    //touser.Coin -= (int64(orgnum) + ip.Price)
    touser.Coin -= ip.Price
    touser.PutUser(stub)
    fromuser.Coin += ip.Price
    fromuser.PutUser(stub)
    ip.Owner = touser.UserId
    ip.State = "2"
    ip.PutIP(stub)

    /*
    var org Org
    for _, orgbytes := range queryResults {
        err = json.Unmarshal(orgbytes, &org)
        if err != nil {
            fmt.Println("Apply Unmarshal fail:", err.Error())
            return shim.Error("Apply Unmarshal fail:" + err.Error())
        }

        org.Coin += 1
        err = org.PutOrg(stub)
        if err != nil {
            fmt.Println("Apply PutOrg fail:", err.Error())
            return shim.Error("Apply PutOrg fail:" + err.Error())
        }
    }
    */

    tx := &Transaction{
        DocType: "Transaction",
        TxId:    time.Now().In(loc).Format(data),
        IPName:  ip.IPName,
        IPId:    ip.IPId,
        SubId:   ip.SubId,
        From: fromuser.UserId,
        To: touser.UserId,
        Price:   ip.Price,
        Version: "v1.0.0",
        CreateTime: time.Now().In(loc).Format(layout),
        UpdateTime: time.Now().In(loc).Format(layout),
    }
    err = tx.PutTransaction(stub)
    if err != nil {
        fmt.Println("Apply PutOrg fail:", err.Error())
        return shim.Error("Apply PutOrg fail:" + err.Error())
    }

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
    err = ip.PutIP(stub)
    if err != nil {
        fmt.Println("use PutIP fail:" + err.Error())
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
    err = ip.PutIP(stub)
    if err != nil {
        fmt.Println("sell PutIP fail:", err.Error())
        return shim.Error("sell PutIP fail:" + err.Error())
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

func getQueryResultList(stub shim.ChaincodeStubInterface, queryString string) ([][]byte, error) {

    fmt.Printf("- getQueryResultList queryString:\n%s\n", queryString)

    resultsIterator, err := stub.GetQueryResult(queryString)
    if err != nil {
        return nil, err
    }
    defer resultsIterator.Close()

    var list [][]byte

    for resultsIterator.HasNext() {
        _, value, err := resultsIterator.Next()
        if err != nil {
            return nil, err
        }
        list = append(list, value)
    }

    fmt.Printf("- getQueryResultList queryResult:\n%v\n", list)

    return list, nil
}

func (t *SimpleChaincode) GetState(stub shim.ChaincodeStubInterface, key string) ([]byte, error) {
    fmt.Println(key)
    respBytes, err := stub.GetState(key)
    if err != nil {
        fmt.Println("GetState fail:", err.Error())
        return nil, err
    }
    return respBytes, nil
}

func (t *SimpleChaincode) GetUser(stub shim.ChaincodeStubInterface, key string) (*User, error) {
    fmt.Println(key)
    var user User
    userBytes, err := t.GetState(stub, key)
    if err != nil {
        fmt.Println("queryUser GetState fail:", err.Error())
        return nil, err
    }

    if string(userBytes) == "" {
        fmt.Println("no key:", key)
        return nil, errors.New("no key:" + key)
    }

    err = json.Unmarshal(userBytes, &user)
    if err != nil {
        fmt.Println("queryUser Unmarshal fail:", err.Error())
        return nil, err
    }

    fmt.Println(user)
    return &user, nil
}

func (t *SimpleChaincode) GetIP(stub shim.ChaincodeStubInterface, key string) (*IP, error) {
    fmt.Println(key)
    var ip IP
    ipBytes, err := t.GetState(stub, key)
    if err != nil {
        fmt.Println("GetIP GetState fail:", err.Error())
        return nil, err
    }
    err = json.Unmarshal(ipBytes, &ip)
    if err != nil {
        fmt.Println("GetIP Unmarshal fail:", err.Error())
        return nil, err
    }

    fmt.Println(ip)
    return &ip, nil
}

func main() {
    err := shim.Start(new(SimpleChaincode))
    if err != nil {
        fmt.Printf("Error starting Simple chaincode: %s", err)
    }
}
