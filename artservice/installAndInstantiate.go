package main

import (
    "net/http"
    "io/ioutil"
    "encoding/json"
)

type ReqInstallAndInstantiate struct {
    ChainCodePath string `json:"chainCodePath"`
    ChainCodeVersion string `json:"chainCodeVersion"`
}

func InstallAndInstantiate(w http.ResponseWriter, r *http.Request) {
    if r.Method != "POST" {
        OutputJson(w, -1, "requset method is not post", nil)
        return
    }

    body, _ := ioutil.ReadAll(r.Body)

    var req ReqInstallAndInstantiate
    err := json.Unmarshal(body, &req)
    if err != nil {
        OutputJson(w, -1, err.Error(), nil)
        return
    }

    if req.ChainCodePath == "" || req.ChainCodeVersion == "" {
        OutputJson(w, -1, "chainCodePath or chainCodeVersion is null", nil)
        return
    }

    if err = base.InstallAndInstantiateArtChainCC(req.ChainCodePath, req.ChainCodeVersion); err != nil {
	    OutputJson(w, -1, err.Error(), nil)
        return
    }

	OutputJson(w, 0, "InstallAndInstantiate ok", nil)
}

