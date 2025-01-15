package main

import (
	"log"
	"net/http"

	"github.com/perso-proj/database"
	"github.com/stripe/stripe-go/v81"
)

func main() {
	stripe.Key = "sk_test_51Qg0zKBgG9LkRx4rH0FZgS6Dk0eSv4m4Qh1i7YrAp69JXgJGVHPbMrOFCUY1qJdpIE5aKKHehMBmw51UEICY5VeW00l7RQs0po"
	db, err := database.DBinit()
	if err != nil {
		log.Fatal("cannot init database, reason -> %s", err.Error())
		panic("db failure")
	}

	handler := handler{db: db}
	initRouter(handler)

}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		log.Println(r.RequestURI)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}
