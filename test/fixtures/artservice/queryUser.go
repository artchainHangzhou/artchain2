package main

import (
    "net/http"
    "io/ioutil"
    "encoding/json"
)

type ReqQueryUser struct {
    UserId string `json:"userId"`
}

func queryUser(w http.ResponseWriter, r *http.Request) {
    if r.Method != "POST" {
        OutputJson(w, -1, "requset method is not post", nil)
        return
    }

    body, _ := ioutil.ReadAll(r.Body)

    var req ReqQueryUser
    err := json.Unmarshal(body, &req)
    if err != nil {
        OutputJson(w, -1, err.Error(), nil)
        return
    }

    if req.UserId == "" {
        OutputJson(w, -1, "UserId is null", nil)
        return
    }
	
	var args []string
	args = append(args, "invoke")
	args = append(args, "queryUser")
	args = append(args, req.UserId)

	value, err := base.Query(base.ChainID, base.ChainCodeID, args)
	if err != nil {
		OutputJson(w, 0, err.Error(), nil)
		return
	}

	OutputJson(w, 0, "QueryUser ok", value)
}

