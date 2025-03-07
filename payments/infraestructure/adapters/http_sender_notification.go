package adapters

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

)

type HTTPSenderNotification struct {
	Host string
	Port int
}

func NewHTTPSenderNotification(host string, port int) *HTTPSenderNotification {
	return &HTTPSenderNotification{
		Host: host,
		Port: port,
	}
}

func (s *HTTPSenderNotification) SendNotification(data map[string]interface{}) error {
	endpoint := fmt.Sprintf("http://%s:%d", s.Host, s.Port)
	
	requestByte, err := json.Marshal(data)
	if err != nil {
		return err
	}

	responseHttp, err := http.Post(endpoint, "application/json", bytes.NewBuffer(requestByte))
	if err != nil {
		return err
	}

	if responseHttp.StatusCode != http.StatusOK {
		return errors.New("error sending notification")
	}

	return nil
}