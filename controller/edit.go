package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go-state-publisher/datastore"
	"net/http"
	"sync"
)

func SpawnEditCtl(host string, port int, wg *sync.WaitGroup) {
	var router = mux.NewRouter()

	router.HandleFunc("/value/{key}", getValueHandler).Methods(http.MethodGet)
	router.HandleFunc("/value/{key}", setValueHandler).Methods(http.MethodPut, http.MethodPost)
	router.HandleFunc("/value/{key}", removeValueHandler).Methods(http.MethodDelete)

	var editServer = http.Server{
		Addr:    fmt.Sprint(host, ":", port),
		Handler: router,
	}
	var err = editServer.ListenAndServe()
	if err != nil {
		fmt.Errorf(err.Error(), "Error in edit controller")
	}
	wg.Done()
}

//region<Handlers>
func removeValueHandler(resp http.ResponseWriter, req *http.Request) {
	var key = getKey(req)
	var err = datastore.Remove(key)
	sendResponse(resp, err)
}

func setValueHandler(resp http.ResponseWriter, req *http.Request) {
	var key = getKey(req)
	var value = req.URL.Query()["value"]
	var err = datastore.Set(key, value)
	sendResponse(resp, err)
}

func getValueHandler(resp http.ResponseWriter, req *http.Request) {
	var key = getKey(req)

	if datastore.Contains(key) {
		var value = datastore.Get(key)
		var data, _ = json.Marshal(value)
		resp.Header().Set("Content-Type", "application/json")
		resp.WriteHeader(http.StatusOK)
		resp.Write(data)
	} else {
		resp.WriteHeader(http.StatusBadRequest)
		resp.Write([]byte("Value was not present"))
	}
}

//ednregion

//region<Helper>
func getKey(req *http.Request) string {
	var params = mux.Vars(req)
	return params["key"]
}

func sendResponse(resp http.ResponseWriter, err error) {
	if err != nil {
		resp.WriteHeader(http.StatusBadRequest)
		resp.Write([]byte(err.Error()))
	} else {
		resp.WriteHeader(http.StatusOK)
	}
}

//endregion
