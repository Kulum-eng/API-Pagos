package adapters

import (
	"database/sql"
	"errors"

	"ModaVane/payments/domain"
)

type MySQLPaymentRepository struct {
	DB *sql.DB
}

func NewMySQLPaymentRepository(db *sql.DB) *MySQLPaymentRepository {
	return &MySQLPaymentRepository{DB: db}
}

func (repo *MySQLPaymentRepository) CreatePayment(payment domain.Payment) (int, error) {
	res, err := repo.DB.Exec(
		"INSERT INTO payments (order_id, amount, status, payment_method) VALUES (?, ?, ?, ?)",
		payment.OrderID, payment.Amount, payment.Status, payment.PaymentMethod,
	)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (repo *MySQLPaymentRepository) GetPaymentByID(id int) (*domain.Payment, error) {
	var payment domain.Payment
	err := repo.DB.QueryRow(
		"SELECT id, order_id, amount, status, payment_method FROM payments WHERE id = ?",
		id,
	).Scan(&payment.ID, &payment.OrderID, &payment.Amount, &payment.Status, &payment.PaymentMethod)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &payment, nil
}

func (repo *MySQLPaymentRepository) GetAllPayments() ([]domain.Payment, error) {
	rows, err := repo.DB.Query("SELECT id, order_id, amount, status, payment_method FROM payments")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	payments := []domain.Payment{}
	for rows.Next() {
		var p domain.Payment
		if err := rows.Scan(&p.ID, &p.OrderID, &p.Amount, &p.Status, &p.PaymentMethod); err != nil {
			return nil, err
		}
		payments = append(payments, p)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return payments, nil
}

func (repo *MySQLPaymentRepository) UpdatePayment(payment domain.Payment) error {
	_, err := repo.DB.Exec(
		"UPDATE payments SET order_id=?, amount=?, status=?, payment_method=? WHERE id=?",
		payment.OrderID, payment.Amount, payment.Status, payment.PaymentMethod, payment.ID,
	)
	return err
}

func (repo *MySQLPaymentRepository) DeletePayment(id int) error {
	res, err := repo.DB.Exec("DELETE FROM payments WHERE id=?", id)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("no se eliminó ningún registro")
	}

	return nil
}
