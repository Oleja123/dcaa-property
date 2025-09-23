package property

import (
	"database/sql"
	"time"
)

type Property struct {
	Id         int             `json:"id"`
	Name       string          `json:"name"`
	Addr       string          `json:"addr"`
	Price      sql.NullFloat64 `json:"price"`
	Info       sql.NullString  `json:"info"`
	CategoryId int             `json:"category_id"`
	LastUpdate time.Time       `json:"last_update"`
}
