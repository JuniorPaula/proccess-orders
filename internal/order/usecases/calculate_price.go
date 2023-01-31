package usecases

import (
	"gorabbitmq/internal/order/entity"
	"gorabbitmq/internal/order/infra/database"
)

type OrderInputDTO struct {
	ID    string
	Price float64
	Tax   float64
}

type OrderOutputDTO struct {
	ID         string
	Price      float64
	Tax        float64
	FinalPrice float64
}

type CalculateFinalPriceUsecase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewCalculateFinalPriceUsecase(orderRepository database.OrderRepository) *CalculateFinalPriceUsecase {
	return &CalculateFinalPriceUsecase{
		OrderRepository: &orderRepository,
	}
}

func (c *CalculateFinalPriceUsecase) Execute(input OrderInputDTO) (*OrderOutputDTO, error) {
	order, err := entity.NewOrder(input.ID, input.Price, input.Tax)
	if err != nil {
		return nil, err
	}

	err = order.CalculateFinalPrice()
	if err != nil {
		return nil, err
	}

	err = c.OrderRepository.Save(order)
	if err != nil {
		return nil, err
	}

	return &OrderOutputDTO{
		ID:         order.ID,
		Price:      order.Price,
		Tax:        order.Tax,
		FinalPrice: order.FinalPrice,
	}, nil
}
