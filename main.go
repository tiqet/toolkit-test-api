package main

import (
    "github.com/tiqet/toolkit/v2"
    "log"
    "net/http"
)

func main() {
    mux := routes()

    log.Println("Starting application on 8080")
    err := http.ListenAndServe(":8080", mux)

    if err != nil {
        log.Fatalln(err)
    }
}

func routes() http.Handler {
    mux := http.NewServeMux()

    mux.Handle("/", http.FileServer(http.Dir(".")))

    mux.HandleFunc("/api/login", login)
    mux.HandleFunc("/api/logout", logout)

    return mux
}

func login(w http.ResponseWriter, r *http.Request) {
    var tools toolkit.Tools

    var paylaod struct {
        Username string `json:"username"`
        Password string `json:"password"`
    }

    err := tools.ReadJSON(w, r, &paylaod)
    if err != nil {
        tools.ErrorJSON(w, err)
        return
    }

    var respPayload toolkit.JSONResponse

    if paylaod.Username == "me@here.com" && paylaod.Password == "very secret" {
        respPayload.Error = false
        respPayload.Message = "Logged in"
        _ = tools.WriteJSON(w, http.StatusAccepted, respPayload)
        return
    }
    respPayload.Error = true
    respPayload.Message = "Invalid credentials"
    _ = tools.WriteJSON(w, http.StatusUnauthorized, respPayload)
}

func logout(w http.ResponseWriter, r *http.Request) {
    var tools toolkit.Tools

    payload := toolkit.JSONResponse{Message: "Logged out"}

    _ = tools.WriteJSON(w, http.StatusAccepted, payload)
}
