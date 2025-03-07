package ports

type SenderNotification interface {
	SendNotification(data map[string]interface{})error
}