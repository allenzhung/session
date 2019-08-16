package main

import (
	"database/sql"
	"io"
	"log"
	"net/http"

	"github.com/alexedwards/scs/mysqlstore"
	"github.com/alexedwards/scs/v2"

	_ "github.com/go-sql-driver/mysql"
)

var sessionManager *scs.SessionManager

func main() {
	db, err := sql.Open("mysql", "root:hhxxttxs8@/activiti?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Initialize a new session manager and configure it to use MySQL as
	// the session store.
	sessionManager = scs.New()
	sessionManager.Store = mysqlstore.New(db)

	mux := http.NewServeMux()
	mux.HandleFunc("/put", putHandler)
	mux.HandleFunc("/get", getHandler)

	http.ListenAndServe(":4000", sessionManager.LoadAndSave(mux))
}

func putHandler(w http.ResponseWriter, r *http.Request) {
	sessionManager.Put(r.Context(), "message", "Hello from a session!")
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	msg := sessionManager.GetString(r.Context(), "message")
	io.WriteString(w, msg)
}
