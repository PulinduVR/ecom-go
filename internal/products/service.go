// This contains the business logic behind the routers related to products
package products

import "context"

type Service interface {
	ListProducts(ctx context.Context) error
}

type svc struct {
	//Dependency will be the repository
}

// This is similar to the constructor we created at the handler
func NewService() Service {
	return &svc{}
}

func (s *svc) ListProducts(ctx context.Context) error {
	return nil
}
