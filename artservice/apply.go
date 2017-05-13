package main

import (
    "encoding/hex"
    "io/ioutil"
    "net/http"
    "os"
    "strconv"
)

var (
    apilaw = "http://47.89.22.214:8080/law/LawServlet"
    apiimages = "http://47.89.22.214:3000/images"
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
        OutputJson(w, -1, "price is too low", nil)
        return
    }

    total, err := strconv.Atoi(r.PostFormValue("price"))
    if err != nil {
        OutputJson(w, -1, err.Error(), nil)
        return
    }

    if total <= 0 {
        OutputJson(w, -1, "total is too low", nil)
        return
    }

    var fileStyle string

    // appraisal
    if true {

        fi, err := os.Open("./file/" + r.PostFormValue("proposalUrl"))
        if err != nil {
            OutputJson(w, -1, err.Error(), nil)
            return
        }
        defer fi.Close()
        fd, err := ioutil.ReadAll(fi)

        var args []string
        args = append(args, "invoke")
        args = append(args, "appraisal")
        args = append(args, r.PostFormValue("proposalUrl"))
        args = append(args, hex.EncodeToString(fd))
        args = append(args, apilaw)

        _, err = base.InvokeAppraisal(args)
        if err != nil {
            OutputJson(w, -1, err.Error(), nil)
            return
        }
    }

    // search
    if true {

        fi, err := os.Open("./file/" + r.PostFormValue("pictureUrl"))
        if err != nil {
            OutputJson(w, -1, err.Error(), nil)
            return
        }
        defer fi.Close()
        fd, err := ioutil.ReadAll(fi)

        var args []string
        args = append(args, "invoke")
        args = append(args, "search")
        args = append(args, r.PostFormValue("pictureUrl"))
        args = append(args, hex.EncodeToString(fd))
        args = append(args, apiimages)

        value, err := base.InvokeSearch(args)
        if err != nil {
            OutputJson(w, -1, err.Error(), nil)
            return
        }

        if value == "" {
            OutputJson(w, -1, "FileStyle is null", nil)
            return
        }

        fileStyle = value
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
    args = append(args, fileStyle)

    value, err := base.Invoke(args)
    if err != nil {
        OutputJson(w, -1, err.Error(), nil)
        return
    }

    OutputJson(w, 0, "Apply ok", value)
}
