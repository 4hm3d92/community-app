package db

import (
	"context"

	"github.com/jackc/pgx/v5"

	"github.com/4hm3d92/community-app/backend/models"
)

func (db Database) GetAllPayments() (*models.PaymentList, error) {
	list := &models.PaymentList{}

	rows, err := db.Pool.Query(context.Background(),
		`SELECT id, member_id, amount, receipt_no, payment_type, paid_on, u1.name as added_by, u2.name as updated_by
			FROM payments p
			INNER JOIN users u1
				ON p.added_by = u1.id
			INNER JOIN users u2
				ON p.updated_by = u2.id
			ORDER BY id DESC;`)
	if err != nil {
		return list, err
	}

	for rows.Next() {
		var payment models.Payment
		err := rows.Scan(&payment.ID, &payment.MemberID, &payment.Amount, &payment.ReceiptNo, &payment.PaymentType, &payment.PaidOn, &payment.AddedBy)
		if err != nil {
			return list, err
		}
		list.Payments = append(list.Payments, payment)
	}
	return list, nil
}

func (db Database) AddPayment(payment *models.Payment, actor int) error {

	query := `INSERT INTO payments (member_id, amount, receipt_no, payment_type, paid_on, added_by, updated_on) VALUES ($1, $2, $3, $4, $5, $6, NULL)`
	_, err := db.Pool.Exec(context.Background(), query, &payment.MemberID, &payment.Amount, &payment.ReceiptNo, &payment.PaymentType, &payment.PaidOn, actor)
	if err != nil {
		return err
	}

	return nil
}

/*
func (db Database) GetPaymentById(paymentId int) (models.Payment, error) {
	payment := models.Payment{}

	query := `SELECT * FROM payments WHERE id = $1;`
	row := db.Pool.QueryRow(context.Background(), query, paymentId)
	switch err := row.Scan(&payment.ID, &payment.FirstName, &payment.MiddleName, &payment.LastName); err {
	case sql.ErrNoRows:
		return payment, ErrNoMatch
	default:
		return payment, err
	}
}
*/

func (db Database) DeletePayment(paymentId int) error {
	query := `DELETE FROM payments WHERE id = $1;`
	_, err := db.Pool.Exec(context.Background(), query, paymentId)
	switch err {
	case pgx.ErrNoRows:
		return ErrNoMatch
	default:
		return err
	}
}

func (db Database) UpdatePayment(paymentId int, paymentData *models.Payment, actor int) error {
	payment := models.Payment{}
	query := `UPDATE payments SET member_id=$1, amount=$2, receipt_no=$3, payment_type=$4, paid_on=$5, updated_by=$6 WHERE id=$7`
	_, err := db.Pool.Exec(context.Background(), query, &payment.MemberID, &payment.Amount, &payment.ReceiptNo, &payment.PaymentType, &payment.PaidOn, actor, paymentId)
	if err != nil {
		if err == pgx.ErrNoRows {
			return ErrNoMatch
		}
		return err
	}

	return nil
}
