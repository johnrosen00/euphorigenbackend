package main

import (
	"context"
	"database/sql"
	"errors"
	"euphorigenbackend/servers/gateway/handlers"
	"euphorigenbackend/servers/gateway/models/users"
	"euphorigenbackend/servers/gateway/sessions"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
)

//main is the main entry point for the server
func main() {

	//summary api stuff
	addr := os.Getenv("ADDR")
	if len(addr) == 0 {
		addr = ":443"
	}

	tlsKeyPath := os.Getenv("TLSKEY")
	tlsCertPath := os.Getenv("TLSCERT")

	//auth api stuff
	sessionKey := os.Getenv("SESSIONKEY")
	redisAddr := os.Getenv("REDISADDR")
	dsn := os.Getenv("DSN")

	sessionStore := &sessions.RedisStore{}
	userStore := &users.MySQLStore{}

	sessionStore.SessionDuration, _ = time.ParseDuration("7m")
	nctx := context.Background()
	sessionStore.Client = redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	sessionStore.Context = nctx
	err := errors.New("")

	//TODO: Handle errors better here. Unsure how to do so.
	if userStore.DB, err = sql.Open("mysql", dsn); err != nil {
		fmt.Printf("error opening database: %v\n", err)
	}

	defer userStore.DB.Close()

	cx := &handlers.HandlerContext{
		Key:          sessionKey,
		SessionStore: sessionStore,
		UserStore:    userStore,
	}

	mux2 := http.NewServeMux()
	mux2.HandleFunc("/v1/users", cx.UserHandler)
	mux2.HandleFunc("/v1/users/", cx.SpecificUserHandler)
	mux2.HandleFunc("/v1/sessions", cx.SessionsHandler)
	mux2.HandleFunc("/v1/sessions/", cx.SpecificSessionsHandler)

	wrappedMux := handlers.NewWrappedCORSHandler(mux2)

	log.Fatal(http.ListenAndServeTLS(addr, tlsCertPath, tlsKeyPath, wrappedMux))

}
