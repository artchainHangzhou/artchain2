package main

import (
    "net/http"
    "io/ioutil"
    "encoding/json"
)

type ReqQueryOrg struct {
    OrgId string `json:"orgId"`
}

func QueryOrg(w http.ResponseWriter, r *http.Request) {
    if r.Method != "POST" {
        OutputJson(w, -1, "requset method is not post", nil)
        return
    }

    body, _ := ioutil.ReadAll(r.Body)

    var req ReqQueryOrg
    err := json.Unmarshal(body, &req)
    if err != nil {
        OutputJson(w, -1, err.Error(), nil)
        return
    }

    if req.OrgId == "" {
        OutputJson(w, -1, "OrgId is null", nil)
        return
    }
	
	var args []string
	args = append(args, "invoke")
	args = append(args, "queryOrg")
	args = append(args, req.OrgId)

	value, err := base.Query(base.ChainID, base.ChainCodeID, args)
	if err != nil {
		OutputJson(w, 0, err.Error(), nil)
		return
	}

	OutputJson(w, 0, "QueryOrg ok", value)
}

