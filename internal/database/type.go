package database

import (
	"time"
)

type Product struct {
	ID    int64  `db:"product_id"`
	Name  string `db:"name"`
	Macro *Macro `db:",inline"`
}

type Products []*Product

type Macro struct {
	Protein      float32 `db:"protein"`
	Fat          float32 `db:"fat"`
	Carbohydrate float32 `db:"carbohydrate"`
	Calories     float32 `db:"calories"`
}

type Entry struct {
	ID        int64      `db:"id"`
	Product   *Product   `db:",inline"`
	Quantity  int        `db:"quantity"`
	CreatedAt *time.Time `db:"created_at"`
	Order     int        `db:"entry_order"`
}

type EntrySet []*Entry
