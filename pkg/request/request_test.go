package request

import "testing"

func TestNewRequest(t *testing.T) {
	requestBody := make(map[string]interface{})
	requestBody["name"] = "Renato"
	method := "GET"
	request, _ := NewRequest(
		method,
		"https://appws.picpay.com/ecommerce/public/payments",
		map[string]string{
			"Content-Type": "application/json",
		},
		requestBody,
	)
	if request.Method != method {
		t.Errorf("Esperado %s - Encontrado %s", method, request.Method)
	}
}

func TestNewRequestInvalidMethod(t *testing.T) {
	requestBody := make(map[string]interface{})
	requestBody["name"] = "Renato"
	method := "bad method"
	_, err := NewRequest(
		method,
		"https://appws.picpay.com/ecommerce/public/payments",
		map[string]string{
			"Content-Type": "application/json",
		},
		requestBody,
	)
	if err == nil {
		t.Error()
	}
}

func TestNewRequestInvalidBody(t *testing.T) {
	requestBody := make(map[string]interface{})
	requestBody["name"] = make(chan int)
	method := "GET"
	_, err := NewRequest(
		method,
		"https://appws.picpay.com/ecommerce/public/payments",
		map[string]string{
			"Content-Type": "application/json",
		},
		requestBody,
	)
	if err == nil {
		t.Error()
	}
}
