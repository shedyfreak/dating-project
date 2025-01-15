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
)

type handler struct {
	db *sqlx.DB
}

func (h *handler) homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, templates.GetHomeWithAddOns())
}

func (h *handler) eventsHandler(w http.ResponseWriter, r *http.Request) {
	events := database.GetEvents(h.db)
	res := templates.ParseEvents(events)
	fmt.Println(res)

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

	fmt.Println("Received: Event=%s, FirstName=%s, LastName=%s, Email=%s, Phone=%s",
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
		SuccessURL: stripe.String(domain + "/api/success?session_id={CHECKOUT_SESSION_ID}"),
		CancelURL:  stripe.String(domain + "/api/cancel"),
	}

	s, err := session.New(params)
	if err != nil {
		log.Printf("session.New: %v", err)
	}

	fmt.Println(s.ID)
	w.Header().Set("HX-Redirect", s.URL)
	w.WriteHeader(http.StatusSeeOther)

}

func (h *handler) success(w http.ResponseWriter, r *http.Request) {
	sID := r.URL.Query().Get("session_id")
	fmt.Println(sID)
	fmt.Fprint(w, templates.GetHomeWithAddOns(templates.SuccessPopUp))

}

func (h *handler) cancel(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, templates.GetHomeWithAddOns())

}
