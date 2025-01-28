package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

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

func (h *handler) contactPageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, templates.GetHomeWithAddOns(templates.ContactPageTempl))
}
func (h *handler) eventsHandler(w http.ResponseWriter, r *http.Request) {
	events, err := database.GetEvents(h.db)
	if err != nil {
		fmt.Fprintf(w, templates.GetHomeWithAddOns(templates.ErrTemplate))
		return
	}
	res := templates.ParseEvents(events)

	fmt.Fprintf(w, res)
}

func (h *handler) eventsOptions(w http.ResponseWriter, r *http.Request) {
	events, err := database.GetEvents(h.db)
	if err != nil {
		fmt.Fprintf(w, templates.GetHomeWithAddOns(templates.ErrTemplate))
		return
	}
	res := templates.GetOptionsEvent(events)
	fmt.Fprintf(w, res)
}

func (h *handler) registrationHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	defer func(err error) {
		if err != nil {
			fmt.Fprintf(w, templates.GetHomeWithAddOns(templates.ErrTemplate))
		}
	}(err)
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
	eventID := r.FormValue("event")
	firstName := r.FormValue("first_name")
	familyName := r.FormValue("last_name")
	email := r.FormValue("email")
	sex := r.FormValue("sex")
	phone := r.FormValue("phone")
	birthdate := r.FormValue("birthdate")
	parsedTime, err := time.Parse(time.DateOnly, birthdate)
	if err != nil {
		log.Println("Error parsing date:", err)
		return
	}
	subscriber := database.Subsriber{
		FirstName:   firstName,
		FamilyName:  familyName,
		Email:       email,
		PhoneNumber: phone,
		Birthday:    database.DateOnly(parsedTime),
		Sex:         sex,
	}
	err, subID := database.SaveSubscriber(h.db, subscriber)
	if err != nil {
		fmt.Fprintf(w, templates.GetHomeWithAddOns(templates.ErrTemplate))
		return
	}
	evID, err := strconv.Atoi(eventID)
	if err != nil {
		fmt.Fprintf(w, templates.GetHomeWithAddOns(templates.ErrTemplate))
		return
	}

	ok, _, _, err := database.GetEligibleBySex(h.db, uint(evID), sex)
	if err != nil {
		fmt.Fprintf(w, templates.GetHomeWithAddOns(templates.ErrTemplate))
		return
	}
	if !ok {
		fmt.Fprintf(w, templates.GetHomeWithAddOns(templates.ErrTemplate))
		return
	}

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
		return
	}

	database.SaveTempSubscription(h.db, uint(evID), subID, s.ID)
	w.Header().Set("HX-Redirect", s.URL)
	w.WriteHeader(http.StatusSeeOther)

}

func (h *handler) success(w http.ResponseWriter, r *http.Request) {
	sID := r.URL.Query().Get("session_id")
	sub, event, err := database.SuccessPaidSubscription(h.db, sID)
	if err != nil {
		fmt.Fprintf(w, "error happened contact admin")
		return
	}
	sendEmail(event, sub)
	fmt.Fprint(w, templates.GetHomeWithAddOns(templates.SuccessPopUp))
}

func (h *handler) cancel(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, templates.GetHomeWithAddOns())

}
