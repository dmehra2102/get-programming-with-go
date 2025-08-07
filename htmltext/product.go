package main

type Product struct {
	Name, Category string
	Price          float64
}

var kayak = Product{
	Name:     "Kayak",
	Category: "WaterSports",
	Price:    279,
}

var Products = []Product{
	{"Kayak", "WaterSports", 279},
	{"Lifejacket", "Watersports", 49.95},
	{"Soccer Ball", "Soccer", 19.50},
	{"Corner Flags", "Soccer", 34.95},
	{"Stadium", "Soccer", 79500},
	{"Thinking Cap", "Chess", 16},
	{"Unsteady Chair", "Chess", 75},
	{"Bling-Bling King", "Chess", 1200},
}

func (p *Product) AddTax() float64 {
	return p.Price * 1.2
}

func (p *Product) ApplyDiscount(amount float64) float64 {
	return p.Price - amount
}