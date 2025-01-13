package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/perso-proj/database"
	"github.com/perso-proj/templates"
	"github.com/stripe/stripe-go/v81"
	"github.com/stripe/stripe-go/v81/checkout/session"

	"github.com/gorilla/mux"
)

type handler struct {
	db *sqlx.DB
}

func (h *handler) homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, templates.HomeTempl)
}

func (h *handler) eventsHandler(w http.ResponseWriter, r *http.Request) {
	events := database.GetEvents(h.db)
	res := templates.ParseEvents(events)

	fmt.Fprintf(w, res)
}

func (h *handler) eventsOptions(w http.ResponseWriter, r *http.Request) {
	events := database.GetEvents(h.db)
	res := templates.GetOptionsEvent(events)
	fmt.Fprintf(w, res)
}

func (h *handler) registrationHandler(w http.ResponseWriter, r *http.Request) {
	// Ensure the method is POST
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse the form data
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Failed to parse form", http.StatusBadRequest)
		return
	}

	// Retrieve form values
	event := r.FormValue("event")
	firstName := r.FormValue("first_name")
	lastName := r.FormValue("last_name")
	email := r.FormValue("email")
	phone := r.FormValue("phone")

	fmt.Printf("Received: Event=%s, FirstName=%s, LastName=%s, Email=%s, Phone=%s",
		event, firstName, lastName, email, phone)
	domain := "http://localhost:8080"
	params := &stripe.CheckoutSessionParams{
		LineItems: []*stripe.CheckoutSessionLineItemParams{
			&stripe.CheckoutSessionLineItemParams{
				// Provide the exact Price ID (for example, pr_1234) of the product you want to sell
				Price:    stripe.String("price_1QgpBMBgG9LkRx4rRHkAmLVf"),
				Quantity: stripe.Int64(1),
			},
		},
		Mode:       stripe.String(string(stripe.CheckoutSessionModePayment)),
		SuccessURL: stripe.String(domain + "/api/success"),
		CancelURL:  stripe.String(domain + "/api/cancel"),
	}

	s, err := session.New(params)

	if err != nil {
		log.Printf("session.New: %v", err)
	}

	// http.Redirect(w, r, s.URL, http.StatusSeeOther)
	// Return the URL as JSON
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("HX-Redirect", s.URL)

}

func main() {
	stripe.Key = "sk_test_51Qg0zKBgG9LkRx4rH0FZgS6Dk0eSv4m4Qh1i7YrAp69JXgJGVHPbMrOFCUY1qJdpIE5aKKHehMBmw51UEICY5VeW00l7RQs0po"
	GlobalDB, err := database.DBinit()
	if err != nil {
		log.Fatal("cannot init database, reason -> %s", err.Error())
		panic("db failure")
	}

	hdler := handler{db: GlobalDB}
	initRouter(hdler)

}

func initRouter(h handler) {
	r := mux.NewRouter()

	r.Use(loggingMiddleware)
	// Define routes
	r.HandleFunc("/", h.homeHandler).Methods("GET")
	r.HandleFunc("/api/events", h.eventsHandler).Methods("GET")
	r.HandleFunc("/api/event-opts", h.eventsOptions).Methods("GET")
	r.HandleFunc("/api/register", h.registrationHandler).Methods("POST")
	//http.Handle("/", http.FileServer(http.Dir("stripe-public")))
	r.HandleFunc("/api/success", h.Success).Methods("GET")

	// Start the server
	fmt.Println("Starting server on :8080...")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal("cannot init server, reason -> %s", err.Error())
		panic("db failure")
	}
}

func (h *handler) Success(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, templates.Success_page)

}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		log.Println(r.RequestURI)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}
