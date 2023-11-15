package handler

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/4hm3d92/community-app/backend/db"
	"github.com/4hm3d92/community-app/backend/models"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

var paymentIDKey = "paymentID"

func payments(router chi.Router) {
	router.Get("/", getAllPayments)
	router.Post("/", createPayment)
	//router.Options("/", corsFix)

	router.Route("/{paymentId}", func(router chi.Router) {
		router.Use(PaymentContext)
		//router.Get("/", getPayment)
		router.Put("/", updatePayment)
		//router.Delete("/", deletePayment)
	})
}

func PaymentContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		paymentId := chi.URLParam(r, "paymentId")
		if paymentId == "" {
			render.Render(w, r, ErrorRenderer(fmt.Errorf("payment ID is required")))
			return
		}
		id, err := strconv.Atoi(paymentId)
		if err != nil {
			render.Render(w, r, ErrorRenderer(fmt.Errorf("invalid payment ID")))
		}
		ctx := context.WithValue(r.Context(), paymentIDKey, id)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getAllPayments(w http.ResponseWriter, r *http.Request) {

	payments, err := dbInstance.GetAllPayments()
	if err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}
	if err := render.Render(w, r, payments); err != nil {
		render.Render(w, r, ErrorRenderer(err))
	}
}

func createPayment(w http.ResponseWriter, r *http.Request) {
	payment := &models.Payment{}
	if err := render.Bind(r, payment); err != nil {
		render.Render(w, r, ErrBadRequest)
		return
	}
	if err := dbInstance.AddPayment(payment, sm.GetInt(r.Context(), "userId")); err != nil {
		render.Render(w, r, ErrorRenderer(err))
		return
	}
	if err := render.Render(w, r, payment); err != nil {
		render.Render(w, r, ServerErrorRenderer(err))
		return
	}
}

func updatePayment(w http.ResponseWriter, r *http.Request) {
	paymentId := r.Context().Value(paymentIDKey).(int)
	paymentData := &models.Payment{}
	if err := render.Bind(r, paymentData); err != nil {
		render.Render(w, r, ErrBadRequest)
		return
	}
	err := dbInstance.UpdatePayment(paymentId, paymentData, sm.GetInt(r.Context(), "userId"))
	if err != nil {
		if err == db.ErrNoMatch {
			render.Render(w, r, ErrNotFound)
		} else {
			render.Render(w, r, ServerErrorRenderer(err))
		}
		return
	}
	w.WriteHeader(http.StatusOK)
}
