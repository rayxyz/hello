package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"path"
	"strconv"

	_ "k8s.io/api/core/v1"
)

var logPath = path.Join(os.Getenv("HOME"), "file/palm-log.txt")

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}

type PalmLog struct {
	SchoolId int64  `json:"school_id,string"`
	CampusId int64  `json:"campus_id,string"`
	UserId   int64  `json:"user_id,string"`
	RoleId   int32  `json:"role_id, string"`
	Data     string `json:"data"`
}

func savePalmLog(w http.ResponseWriter, r *http.Request) {
	var palmLog PalmLog
	json.NewDecoder(r.Body).Decode(&palmLog)
	file, err := os.OpenFile(logPath, os.O_APPEND|os.O_WRONLY, 0600)
	defer file.Close()
	if err != nil {
		return
	}
	n, err := file.WriteString("\n\n" + fmt.Sprintf("%#v", palmLog))
	if err != nil {
		return
	}
	log.Println("Have written lenth of string => ", strconv.Itoa(n))
}

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/", defaultHandler)
	// http.HandleFunc("/", defaultHandler)
	router.HandleFunc("/savePalmLog", savePalmLog)
	log.Println("Hello world!")
	http.ListenAndServe(":8085", router)
}
