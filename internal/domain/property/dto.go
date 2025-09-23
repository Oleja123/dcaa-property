package property

type CreatePropertyDTO struct {
	Addr       string  `json:"addr"`
	Price      float64 `json:"price"`
	Info       string  `json:"info"`
	CategoryId string  `json:"category_id"`
}
