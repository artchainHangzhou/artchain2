package main

import (
    "net/http"
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

    if r.PostFormValue("userId") == "" || r.PostFormValue("ipId") == "" {
        OutputJson(w, -1, "userId or ipId is null", nil)
        return
    }
	
	var args []string
	args = append(args, "invoke")
	args = append(args, "buy")
	args = append(args, r.PostFormValue("userId"))
	args = append(args, r.PostFormValue("ipId"))

	value, err := base.Invoke(args)
	if err != nil {
		OutputJson(w, -1, err.Error(), nil)
		return
	}

	OutputJson(w, 0, "Buy ok", value)
}

