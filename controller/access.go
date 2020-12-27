package controller

import (
	"encoding/json"
	"fmt"
	"github.com/MBoegers/go-state-store/datastore"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"net/http"
	"sync"
)

var connections []*websocket.Conn
var updatesChan chan []string

func InitReadCtl(port int, updates chan []string, wg *sync.WaitGroup) {
	var router = mux.NewRouter()
	updatesChan = updates

	router.HandleFunc("/updates", handleRegister).Methods(http.MethodGet)
	router.HandleFunc("/get/{key}", handleGet).Methods(http.MethodGet)

	// Where ORIGIN_ALLOWED is like `scheme://dns[:port]`, or `*` (insecure)
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET"})

	var server = http.Server{
		Addr:    fmt.Sprint(":", port),
		Handler: handlers.CORS(headersOk, originsOk, methodsOk)(router),
		//Handler: router,
	}

	go handleUpdateEvent()
	var err = server.ListenAndServe()
	if err != nil {
		fmt.Errorf(err.Error(), "Error in access controller")
	}
	wg.Done()
}

//region<Handler>
func handleGet(resp http.ResponseWriter, req *http.Request) {
	var vars = mux.Vars(req)
	var key = vars["key"]
	if datastore.Contains(key) {
		var value = datastore.Get(key)
		var data, _ = json.Marshal(value)
		resp.Header().Set("Content-Type", "application/json")
		resp.WriteHeader(http.StatusOK)
		resp.Write(data)
	} else {
		resp.WriteHeader(http.StatusBadRequest)
		resp.Write([]byte(fmt.Sprintf("Key: \"%v\" not found", key)))
	}
}

func handleRegister(resp http.ResponseWriter, req *http.Request) {
	var conn, err = upgrader.Upgrade(resp, req, nil)
	if err != nil {
		fmt.Errorf("Failed to upgrade connection due to: %v", err)
	}
	connections = append(connections, conn)
}

func handleUpdateEvent() {
	for {
		var event = <-updatesChan
		for _, con := range connections {
			con.WriteJSON(event)
		}
	}
}

//endregion

//region<Helper>
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

//endregion
