package main

import (
    "net/http"
    "io/ioutil"
    "encoding/json"
)

type ReqQueryUserIP struct {
    UserId string `json:"userId"`
}

func QueryUserIPList(w http.ResponseWriter, r *http.Request) {
    if origin := r.Header.Get("Origin"); origin != "" {
        w.Header().Set("Access-Control-Allow-Origin", "*")
        w.Header().Set("Access-Control-Allow-Methods", "POST,GET,OPTIONS,PUT,DELETE")
        w.Header().Set("Access-Control-Allow-Headers", "Action, Module")
    }

    if r.Method != "POST" {
        OutputJson(w, -1, "requset method is not post", nil)
        return
    }

    body, _ := ioutil.ReadAll(r.Body)

    var req ReqQueryUserIP
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
	args = append(args, "queryUserIPList")
	args = append(args, req.UserId)

	value, err := base.Query(base.ChainID, base.ChainCodeID, args)
	if err != nil {
		OutputJson(w, 0, err.Error(), nil)
		return
	}

	OutputJson(w, 0, "queryUserIPList ok", value)
}

