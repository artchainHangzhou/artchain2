package main

import (
    "net/http"
    "os"
    "fmt"
    "io"
    "path"
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

    fmt.Println(r.URL.Path)
    fileName := path.Base(r.URL.Path)
    fmt.Println(fileName)

    file, err := os.Open("./file/" + fileName)  
    if err != nil {  
        OutputJson(w, -1, err.Error(), nil)
        return  
    } 

    defer file.Close()  
    //fileName = url.QueryEscape(fileName) // 防止中文乱码  
    //w.Header().Set("Content-Type", "application/octet-stream")  
    //w.Header().Set("content-disposition", "attachment; filename=\""+fileName+"\"")  

    _, err = io.Copy(w, file)  
    if err != nil {
        OutputJson(w, -1, err.Error(), nil)
        return  
    }

	OutputJson(w, 0, "DownLoad ok", fileName)
}

