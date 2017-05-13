package main

import (
    "fmt"
    "bytes"
    "net/http"
    "io/ioutil"
    "mime/multipart"
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

	r.ParseMultipartForm(32 << 20)
	file, handler, err := r.FormFile("uploadfile")
	if err != nil {
		fmt.Println(err)
		OutputJson(w, -1, err.Error(), nil)
		return
	}
	defer file.Close()
    fd, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
		OutputJson(w, -1, err.Error(), nil)
		return
	}

	var buff bytes.Buffer
	writer := multipart.NewWriter(&buff)
	req, _ := writer.CreateFormFile("file", handler.Filename)
	req.Write(fd)
	writer.Close()
	var client http.Client
	resp, err := client.Post(apiimages, writer.FormDataContentType(), &buff)
	if err != nil {
		fmt.Println("search Post fail:", err.Error())
		OutputJson(w, -1, err.Error(), nil)
        return 
	}
	defer resp.Body.Close()
	data, _ := ioutil.ReadAll(resp.Body)
	fileStyle := string(data)
	fmt.Println(fileStyle)

    if fileStyle == "" {
		OutputJson(w, -1, "fileStyle is null", nil)
		return
    }

	var args []string
	args = append(args, "invoke")
	args = append(args, "search")
	args = append(args, fileStyle)

	value, err := base.Query(base.ChainID, base.ChainCodeID, args)
	if err != nil {
		OutputJson(w, -1, err.Error(), nil)
		return
	}

	OutputJson(w, 0, "Search ok", value)
}

