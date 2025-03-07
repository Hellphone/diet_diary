package domain

type Product struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
	*Macro
}

type Products []*Product

func (p *Product) CalcCalories() {
	if p != nil {
		p.Calories = p.Macro.Protein*CaloriesInProtein + p.Macro.Fat*CaloriesInFat +
			p.Macro.Carbohydrate*CaloriesInCarb
	}
}
