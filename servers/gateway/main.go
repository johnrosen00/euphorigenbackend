package main

import (
	"context"
	"database/sql"
	"euphorigenbackend/servers/gateway/handlers"
	"euphorigenbackend/servers/gateway/models/gamepass"
	"euphorigenbackend/servers/gateway/models/metrics"
	"euphorigenbackend/servers/gateway/models/players"
	"euphorigenbackend/servers/gateway/sessions"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
	"golang.org/x/crypto/bcrypt"
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
	manpass, err := bcrypt.GenerateFromPassword([]byte(os.Getenv("MANPASS")), 13)

	if err != nil {
		return
	}

	sessionStore := &sessions.RedisStore{}

	sessionStore.SessionDuration, _ = time.ParseDuration("7m")
	nctx := context.Background()
	sessionStore.Client = redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	sessionStore.Context = nctx
	//DB Connections
	dsn := os.Getenv("DSN")
	gamepassstore := &gamepass.PassStore{}
	if gamepassstore.DB, err = sql.Open("mysql", dsn); err != nil {
		fmt.Printf("error opening database: %v\n", err)
	}

	defer gamepassstore.DB.Close()

	playerstore := &players.PlayerStore{}
	if playerstore.DB, err = sql.Open("mysql", dsn); err != nil {
		fmt.Printf("error opening database: %v\n", err)
	}

	defer playerstore.DB.Close()

	metricstore := &metrics.MetricStore{}
	if metricstore.DB, err = sql.Open("mysql", dsn); err != nil {
		fmt.Printf("error opening database: %v\n", err)
	}

	defer metricstore.DB.Close()

	cx := &handlers.HandlerContext{
		Key:           sessionKey,
		SessionStore:  sessionStore,
		ManPass:       manpass,
		GamePassStore: gamepassstore,
		PlayerStore:   playerstore,
		MetricStore:   metricstore,
	}

	mux2 := http.NewServeMux()

	mux2.HandleFunc("/v1/sessions", cx.SessionsHandler)
	mux2.HandleFunc("v1/game/manage/password", cx.GamePassHandler)
	mux2.HandleFunc("v1/game/metrics", cx.MetricHandler)
	wrappedMux := handlers.NewWrappedCORSHandler(mux2)

	log.Fatal(http.ListenAndServeTLS(addr, tlsCertPath, tlsKeyPath, wrappedMux))

}
