package main

import (
    "net/http"
	"fmt"
    "os"
    "io"
)

func UpLoad(w http.ResponseWriter, r *http.Request) {
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
    fullfilename := "/file/" + handler.Filename

    var f *os.File
    if _, err = os.Stat("." + fullfilename); os.IsNotExist(err) {
        f, err = os.Create("." + fullfilename)
        if err != nil {
            fmt.Println(err)
            OutputJson(w, -1, err.Error(), nil)
            return
        }
    } else {
	    f, err = os.OpenFile("." + fullfilename, os.O_WRONLY|os.O_CREATE, 0666)
        if err != nil {
            fmt.Println(err)
            OutputJson(w, -1, err.Error(), nil)
            return
        }
    }

	defer f.Close()
	io.Copy(f, file)

	//OutputJson(w, 0, "UpLoad ok", "/download" + fullfilename)
	OutputJson(w, 0, "UpLoad ok", handler.Filename)
}

