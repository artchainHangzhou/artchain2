package main

import (
	"net/http"
    "strconv"
)

func Apply(w http.ResponseWriter, r *http.Request) {
    if origin := r.Header.Get("Origin"); origin != "" {
        w.Header().Set("Access-Control-Allow-Origin", "*")
        w.Header().Set("Access-Control-Allow-Methods", "POST,GET,OPTIONS,PUT,DELETE")
        w.Header().Set("Access-Control-Allow-Headers", "Action, Module")
    }

	if r.Method != "POST" {
		OutputJson(w, -1, "requset method is not post", nil)
		return
	}

	if r.PostFormValue("userId") == "" || r.PostFormValue("ipName") == "" || r.PostFormValue("author") == "" || 
        r.PostFormValue("description") == "" || r.PostFormValue("proposalUrl") == "" || r.PostFormValue("pictureUrl") == "" || 
        r.PostFormValue("price") == "" || r.PostFormValue("total") == "" {
		OutputJson(w, -1, "request is null", nil)
		return
	}

    price, err := strconv.Atoi(r.PostFormValue("price"))
    if err != nil {
		OutputJson(w, -1, err.Error(), nil)
		return
    }

    if price <= 0 {
		OutputJson(w, -1, "price is to low", nil)
		return
    }

    total, err := strconv.Atoi(r.PostFormValue("price"))
    if err != nil {
		OutputJson(w, -1, err.Error(), nil)
		return
    }

    if total <= 0 {
		OutputJson(w, -1, "total is to low", nil)
		return
    }


	var args []string
	args = append(args, "invoke")
	args = append(args, "apply")
	args = append(args, r.PostFormValue("userId"))
	args = append(args, r.PostFormValue("ipName"))
	args = append(args, r.PostFormValue("author"))
	args = append(args, r.PostFormValue("description"))
	args = append(args, r.PostFormValue("proposalUrl"))
	args = append(args, r.PostFormValue("pictureUrl"))
	args = append(args, r.PostFormValue("price"))
	args = append(args, r.PostFormValue("total"))

	value, err := base.Invoke(args)
	if err != nil {
		OutputJson(w, 0, err.Error(), nil)
		return
	}

	OutputJson(w, 0, "Apply ok", value)
}
