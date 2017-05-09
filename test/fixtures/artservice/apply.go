package main

import (
	"encoding/hex"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type ReqApply struct {
	UserId      string `json:"userId"`
	IPName      string `json:"ipName"`
	Author      string `json:"author"`
	Description string `json:"description"`
	ProposalUrl string `json:"proposalUrl"`
	PictureUrl  string `json:"pictureUrl"`
	Price       int64  `json:"Price"`
	Total       int    `json:"total"`
}

func Apply(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		OutputJson(w, -1, "requset method is not post", nil)
		return
	}

	body, _ := ioutil.ReadAll(r.Body)

	var req ReqApply
	err := json.Unmarshal(body, &req)
	if err != nil {
		OutputJson(w, -1, err.Error(), nil)
		return
	}

	if req.UserId == "" || req.IPName == "" || req.Author == "" || req.Description == "" ||
		req.ProposalUrl == "" || req.PictureUrl == "" || req.Price < 0 || req.Total < 0 {
		OutputJson(w, -1, "UserId is null", nil)
		return
	}

	var args []string
	args = append(args, "invoke")
	args = append(args, "apply")
	args = append(args, hex.EncodeToString(body))

	value, err := base.Invoke(args)
	if err != nil {
		OutputJson(w, 0, err.Error(), nil)
		return
	}

	OutputJson(w, 0, "Apply ok", value)
}
