package domain

import (
	"encoding/json"
	"fmt"
	"time"
)

type Entry struct {
	ID       int64    `json:"id"`
	Product  *Product `json:"product,inline"`
	Quantity int      `json:"quantity"` // in grams
	// User    int64 // later
	CreatedAt *time.Time `json:"createdAt"`
	Order     int        `json:"order"`
}

func (e *Entry) UnmarshalJSON(data []byte) error {
	type Alias Entry
	aux := struct {
		Product json.RawMessage `json:"product"`
		*Alias
	}{
		Alias: (*Alias)(e),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	var productID int64
	if err := json.Unmarshal(aux.Product, &productID); err == nil {
		if e.Product == nil {
			e.Product = &Product{}
		}

		e.Product.ID = productID
		return nil
	}

	var product Product
	if err := json.Unmarshal(aux.Product, &product); err != nil {
		return fmt.Errorf("invalid product format: %w", err)
	}

	if e.Product == nil {
		e.Product = &product
	} else {
		*e.Product = product
	}

	return nil
}

func (e *Entry) CalcProtein() float32 {
	if e == nil || e.Product == nil || e.Product.Macro == nil {
		return 0
	}

	return e.Product.Macro.Protein * float32(e.Quantity) / 100
}

func (e *Entry) CalcFat() float32 {
	if e == nil || e.Product == nil || e.Product.Macro == nil {
		return 0
	}

	return e.Product.Macro.Fat * float32(e.Quantity) / 100
}

func (e *Entry) CalcCarbohydrate() float32 {
	if e == nil || e.Product == nil || e.Product.Macro == nil {
		return 0
	}

	return e.Product.Macro.Carbohydrate * float32(e.Quantity) / 100
}

func (e *Entry) CalcCalories() float32 {
	if e == nil || e.Product == nil || e.Product.Macro == nil {
		return 0
	}

	return float32(e.Quantity) * (e.Product.Macro.Protein*CaloriesInProtein +
		e.Product.Macro.Fat*CaloriesInFat + e.Product.Macro.Carbohydrate*CaloriesInCarb) / 100
}

type EntrySet []*Entry

type EntryTotal struct {
	*Macro
	Date *time.Time `json:"date"`
}

func (es EntrySet) Total() *EntryTotal {
	if es == nil {
		return nil
	}

	entryTotal := &EntryTotal{&Macro{}, &time.Time{}}
	for _, e := range es {
		entryTotal.Protein += e.CalcProtein()
		entryTotal.Fat += e.CalcFat()
		entryTotal.Carbohydrate += e.CalcCarbohydrate()
		entryTotal.Calories += e.CalcCalories()
	}

	entryTotal.Date = es[0].CreatedAt

	return entryTotal
}
