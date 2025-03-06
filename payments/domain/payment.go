package domain

type Payment struct {
    ID            int     `json:"id"`
    OrderID       int     `json:"order_id"`
    Amount        float64 `json:"amount"`
    Status        string  `json:"status"`
    PaymentMethod string  `json:"payment_method"`
}
