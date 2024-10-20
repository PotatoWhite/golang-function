package factory

type Discount func(price float64) float64

func GetDiscountFunction(rate float64) Discount {
	return func(price float64) float64 {
		return price * (1 - rate/100) // return discounted price
	}
}
