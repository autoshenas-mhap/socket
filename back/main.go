package main

import (
    "log"
    "net/http"
    "github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
	// Add this line to bypass origin check (for development only!)
    CheckOrigin: func(r *http.Request) bool { return true },
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
    // Upgrade initial GET request to a websocket
    ws, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Fatal(err)
    }
    defer ws.Close()

    for {
        // Read message from browser
        messageType, p, err := ws.ReadMessage()
        if err != nil {
            log.Println(err)
            return
        }
        // Print the message to the console
        log.Printf("%s sent: %s\n", ws.RemoteAddr(), p)

        // Write message back to browser
        if err := ws.WriteMessage(messageType, p); err != nil {
            log.Println(err)
            return
        }
    }
}

func main() {
    http.HandleFunc("/websocket", handleConnections)
    log.Println("http server started on :8000")
    err := http.ListenAndServe(":8000", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
