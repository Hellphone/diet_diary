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

func (e *Entry) CalcCalories() float32 {
	if e == nil {
		return 0
	}

	return float32(e.Quantity) * (e.Product.Macro.Protein*CaloriesInProtein +
		e.Product.Macro.Fat*CaloriesInFat + e.Product.Macro.Carbohydrate*CaloriesInCarb)
}

type EntrySet []*Entry

func (es EntrySet) Total() *Macro {
	macro := &Macro{}
	for _, e := range es {
		macro.Protein += e.Product.Macro.Protein
		macro.Fat += e.Product.Macro.Fat
		macro.Carbohydrate += e.Product.Macro.Carbohydrate
		macro.Calories += e.CalcCalories()
	}

	return macro
}
