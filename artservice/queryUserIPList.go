package main

import (
    "net/http"
)

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


    if r.PostFormValue("userId") == "" {
        OutputJson(w, -1, "UserId is null", nil)
        return
    }
	

	var args []string
	args = append(args, "invoke")
	args = append(args, "queryUserIPList")
	args = append(args, r.PostFormValue("userId"))

	value, err := base.Query(base.ChainID, base.ChainCodeID, args)
	if err != nil {
		OutputJson(w, -1, err.Error(), nil)
		return
	}

	OutputJson(w, 0, "queryUserIPList ok", value)
}

