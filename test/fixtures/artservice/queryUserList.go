package main

import (
    "net/http"
)

func queryUserList(w http.ResponseWriter, r *http.Request) {
    if r.Method != "POST" {
        OutputJson(w, -1, "requset method is not post", nil)
        return
    }

	var args []string
	args = append(args, "invoke")
	args = append(args, "queryUserList")
	args = append(args, "list")

	value, err := base.Query(base.ChainID, base.ChainCodeID, args)
	if err != nil {
		OutputJson(w, 0, err.Error(), nil)
		return
	}

	OutputJson(w, 0, "QueryUserList ok", value)
}

