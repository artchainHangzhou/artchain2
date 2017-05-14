package main

import (
    "net/http"
)

func Block(w http.ResponseWriter, r *http.Request) {
    if origin := r.Header.Get("Origin"); origin != "" {
        w.Header().Set("Access-Control-Allow-Origin", "*")
        w.Header().Set("Access-Control-Allow-Methods", "POST,GET,OPTIONS,PUT,DELETE")
        w.Header().Set("Access-Control-Allow-Headers", "Action, Module")
    }

    if r.Method != "GET" {
        OutputJson(w, -1, "requset method is not get", nil)
        return
    }

    value, err := base.Chain.QueryInfo()
    if err != nil {
		OutputJson(w, -1, err.Error(), nil)
		return
    }   

	OutputJson(w, 0, "Block Info ok", value)
}

