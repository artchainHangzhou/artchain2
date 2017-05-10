package main

import (
    "net/http"
    "io/ioutil"
    "encoding/json"
)

type ReqBuy struct {
    UserId string `json:"userId"`
    IPId string `json:"ipId"`
}

func Buy(w http.ResponseWriter, r *http.Request) {
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

    var req ReqBuy
    err := json.Unmarshal(body, &req)
    if err != nil {
        OutputJson(w, -1, err.Error(), nil)
        return
    }

    if req.IPId == "" {
        OutputJson(w, -1, "IPId is null", nil)
        return
    }
	
	var args []string
	args = append(args, "invoke")
	args = append(args, "buy")
	args = append(args, req.UserId)
	args = append(args, req.IPId)

	value, err := base.Invoke(args)
	if err != nil {
		OutputJson(w, 0, err.Error(), nil)
		return
	}

	OutputJson(w, 0, "Buy ok", value)
}

