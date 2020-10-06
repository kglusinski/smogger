package smogger

type ApiClient interface {
	Cities(country string, v interface{}) error
}

type Service struct {
	client ApiClient
}

func NewService(c ApiClient) *Service {
	return &Service{
		client: c,
	}
}
