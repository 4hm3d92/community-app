package models

import (
	"net/http"
	"time"
)

type Payment struct {
	ID          int       `json:"id"`
	PaymentType int       `json:"payment_type"`
	ReceiptNo   string    `json:"receipt_no"`
	MemberID    string    `json:"member_id"`
	Amount      int       `json:"amount"`
	AddedOn     time.Time `json:"added_on"`
	AddedBy     string    `json:"added_by"`
	PaidOn      time.Time `json:"paid_on"`
	UpdatedBy   string    `json:"updated_by"`
	UpdatedOn   time.Time `json:"updated_on"`
}

type PaymentList struct {
	Payments []Payment `json:"payments"`
}

func (i *Payment) Bind(r *http.Request) error {
	return nil
}

func (*PaymentList) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (*Payment) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
