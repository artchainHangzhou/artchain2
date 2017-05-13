package main

import (
    "net/http"
)

func Search(w http.ResponseWriter, r *http.Request) {
    if origin := r.Header.Get("Origin"); origin != "" {
        w.Header().Set("Access-Control-Allow-Origin", "*")
        w.Header().Set("Access-Control-Allow-Methods", "POST,GET,OPTIONS,PUT,DELETE")
        w.Header().Set("Access-Control-Allow-Headers", "Action, Module")
    }

    if r.Method != "POST" {
        OutputJson(w, -1, "requset method is not post", nil)
        return
    }


    if r.PostFormValue("fileStyle") == "" {
        OutputJson(w, -1, "fileStyle is null", nil)
        return
    }
	

	var args []string
	args = append(args, "invoke")
	args = append(args, "search")
	args = append(args, r.PostFormValue("fileStyle"))

	value, err := base.Query(base.ChainID, base.ChainCodeID, args)
	if err != nil {
		OutputJson(w, -1, err.Error(), nil)
		return
	}

	OutputJson(w, 0, "Search ok", value)
}

