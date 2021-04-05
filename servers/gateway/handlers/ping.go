package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
)

func (cx *HandlerContext) Ping(w http.ResponseWriter, r *http.Request) {

	msg := "no error"
	x, err := sql.Open("mysql", os.Getenv("DSN"))

	w.WriteHeader(http.StatusOK)
	if err != nil {
		msg = fmt.Sprintf("error opening database: %v\n", err)
		w.Write([]byte(msg))
	}

	defer x.Close()

	if err = x.Ping(); err != nil {
		msg = fmt.Sprintf("error pinging database: %v\n", err)
		w.Write([]byte(msg))
	}

	w.Write([]byte("UserStore all good"))

}
