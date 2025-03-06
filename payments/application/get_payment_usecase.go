package aplication

import (
	"ModaVane/payments/domain"
	"ModaVane/payments/domain/ports"
)

type GetPaymentUseCase struct {
	repo ports.PaymentRepository
}

func NewGetPaymentUseCase(repo ports.PaymentRepository) *GetPaymentUseCase {
	return &GetPaymentUseCase{repo: repo}
}

func (uc *GetPaymentUseCase) ExecuteByID(id int) (*domain.Payment, error) {
	return uc.repo.GetPaymentByID(id)
}

func (uc *GetPaymentUseCase) ExecuteAll() ([]domain.Payment, error) {
	return uc.repo.GetAllPayments()
}
