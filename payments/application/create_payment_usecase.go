package aplication

import (
	"ModaVane/payments/domain"
	"ModaVane/payments/domain/ports"
)

type CreatePaymentUseCase struct {
	repo ports.PaymentRepository
}

func NewCreatePaymentUseCase(repo ports.PaymentRepository) *CreatePaymentUseCase {
	return &CreatePaymentUseCase{repo: repo}
}

func (uc *CreatePaymentUseCase) Execute(payment domain.Payment) (int, error) {
	return uc.repo.CreatePayment(payment)
}
