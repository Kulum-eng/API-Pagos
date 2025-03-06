package ports

import "ModaVane/payments/domain"

type PaymentRepository interface {
    CreatePayment(payment domain.Payment) (int, error)
    GetPaymentByID(id int) (*domain.Payment, error)
    GetAllPayments() ([]domain.Payment, error)
    UpdatePayment(payment domain.Payment) error
    DeletePayment(id int) error
}
