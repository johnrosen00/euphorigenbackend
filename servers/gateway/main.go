package main

import (
	"context"
	"euphorigenbackend/servers/gateway/handlers"
	"euphorigenbackend/servers/gateway/sessions"
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

	cx := &handlers.HandlerContext{
		Key:          sessionKey,
		SessionStore: sessionStore,
		ManPass:      manpass,
	}

	mux2 := http.NewServeMux()

	mux2.HandleFunc("/v1/sessions", cx.SessionsHandler)

	wrappedMux := handlers.NewWrappedCORSHandler(mux2)

	log.Fatal(http.ListenAndServeTLS(addr, tlsCertPath, tlsKeyPath, wrappedMux))

}
