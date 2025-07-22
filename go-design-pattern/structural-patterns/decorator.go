package structuralpatterns

type Car struct {
	Brand string
	Price float64
}

type PriceDecorator interface {
	DecoratePrice(car Car) Car
}

type ExtraPriceDecorator struct {
	ExtraPrice float64
}

func (d ExtraPriceDecorator) DecoratePrice(car Car) Car {
	car.Price += d.ExtraPrice
	return car
}
