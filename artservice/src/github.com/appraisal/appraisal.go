package main

import (
	"bytes"
	"encoding/json"
	"encoding/hex"
    "mime/multipart"
	"errors"
	"fmt"
    "io/ioutil"
	"time"
	"net/http"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

var (
	layout = "2006-01-02 15:04:05"
	date   = "20060102150405"
	loc    *time.Location
)

func init() {
	loc, _ = time.LoadLocation("Asia/Shanghai")
}

type File struct {
	DocType    string `json:"docType"`
	AppraisalId   string `json:"appraisalId"`
	FileName   string `json:"fileName"`
	FileData   []byte `json:"fileData"`
	Status     string `json:"status"`
	Version    string `json:"version"`
	CreateTime string `json:"createTime"`
	UpdateTime string `json:"updateTime"`
}

func (file *File) PutFile(stub shim.ChaincodeStubInterface) error {
	filebytes, err := json.Marshal(file)
	if err != nil {
		fmt.Println("PutFile Marshal fail:", err.Error())
		return errors.New("PutFile Marshal fail:" + err.Error())
	}

	err = stub.PutState("File:"+file.AppraisalId, filebytes)
	if err != nil {
		fmt.Println("PutFile PutState fail:", err.Error())
		return errors.New("PutFile PutState Error" + err.Error())
	}

	return nil
}

type SimpleChaincode struct {
}

func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("########### Appraisal Init ###########")
	return shim.Success(nil)
}

func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("########### Appraisal Invoke ###########")
	function, args := stub.GetFunctionAndParameters()
	fmt.Println("function:" + function)
	for i, v := range args {
        if i != 2 {
		    fmt.Println("args:" + v)
        }
	}

	if function != "invoke" {
		return shim.Error("Unknown function call:" + function)
	}

	/*
	   if len(args) != 2 {
	       return shim.Error("Incorrect number of arguments. Expecting at least 2")
	   }
	*/

	switch args[0] {
	case "appraisal":
		return t.appraisal(stub, args)
	default:
		return shim.Error("Unknown action, check the first argument:" + args[0])
	}
}

func (t *SimpleChaincode) appraisal(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	//fmt.Println(args)
	fileName := args[1]
	hexdata := args[2]
	api := args[3]
    fmt.Println(fileName)
    fmt.Println(api)

	fileData, err := hex.DecodeString(hexdata)
	if err != nil {
		fmt.Println("appraisal DecodeString:" + err.Error())
        return shim.Error("appraisal DecodeString fail:" + err.Error())
	}

	var buff bytes.Buffer
	writer := multipart.NewWriter(&buff)
    writer.WriteField("fileName", fileName)
	w, _ := writer.CreateFormFile("file", "temp.pdf")
	w.Write(fileData)
	writer.Close()
	var client http.Client
	resp, err := client.Post(api, writer.FormDataContentType(), &buff)
	if err != nil {
		fmt.Println("appraisal Post fail:", err.Error())
        return shim.Error("appraisal Post fail:" + err.Error())
	}
	defer resp.Body.Close()
	data, _ := ioutil.ReadAll(resp.Body)
	status := string(data)
	fmt.Println(status)

	file := &File{
		DocType:    "appraisalFile",
		AppraisalId:   time.Now().In(loc).Format(date),
		FileName:   fileName,
		FileData:   fileData,
		Status:  status,
		Version:    "v1.0.0",
		CreateTime: time.Now().In(loc).Format(layout),
		UpdateTime: time.Now().In(loc).Format(layout),
	}

	err = file.PutFile(stub)
	if err != nil {
		fmt.Println("appraisal PutFile fail:" + err.Error())
		return shim.Error("appraisal PutFile fail:" + err.Error())
	}

    if status == "YES" {
	    return shim.Success(nil)
    } else {
        return shim.Error("appraisal fail:" + status)
    }

}

func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}
