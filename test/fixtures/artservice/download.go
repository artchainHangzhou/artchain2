package main

import (
    "net/http"
	"fmt"
    "os"
    "io"
)

func DownLoad(w http.ResponseWriter, r *http.Request) {
    if origin := r.Header.Get("Origin"); origin != "" {
        w.Header().Set("Access-Control-Allow-Origin", "*")
        w.Header().Set("Access-Control-Allow-Methods", "POST,GET,OPTIONS,PUT,DELETE")
        w.Header().Set("Access-Control-Allow-Headers", "Action, Module")
    }

    if r.Method != "GET" {
        OutputJson(w, -1, "requset method is not get", nil)
        return
    }

	r.ParseMultipartForm(32 << 20)
	file, handler, err := r.FormFile("uploadfile")
	if err != nil {
		fmt.Println(err)
		OutputJson(w, 0, err.Error(), nil)
		return
	}

	defer file.Close()
	fmt.Fprintf(w, "%v", handler.Header)

    var f *os.File
    if _, err = os.Stat("./file/" + handler.Filename); os.IsNotExist(err) {
        f, err = os.Create("./file/" + handler.Filename)
        if err != nil {
            fmt.Println(err)
            OutputJson(w, 0, err.Error(), nil)
            return
        }
    } else {
	    f, err = os.OpenFile("./file/" + handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
        if err != nil {
            fmt.Println(err)
            OutputJson(w, 0, err.Error(), nil)
            return
        }
    }

	defer f.Close()
	io.Copy(f, file)

	OutputJson(w, 0, "DownLoad ok", handler.Filename)
}

