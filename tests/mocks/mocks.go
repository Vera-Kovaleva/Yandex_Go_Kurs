package main

type APIClient interface {
	GetData(query string) (Response, error)
}

type Response struct {
	Text       string
	StatusCode int
}

type MockAPIClient interface {
	APIClient
	SetResponse(resp Response, err error)
}

type Mock struct {
	Response
	Error error
}

func (m *Mock) GetData(query string) (Response, error) {
    return m.Response, m.Error
}

func (m *Mock) 	SetResponse(resp Response, err error) {
    m.Response = resp
    m.Error = err
}
