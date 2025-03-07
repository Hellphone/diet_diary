package domain

const (
	CaloriesInProtein = float32(4)
	CaloriesInFat     = float32(9)
	CaloriesInCarb    = float32(4)
)

type Macro struct {
	Protein      float32 `json:"protein"`
	Fat          float32 `json:"fat"`
	Carbohydrate float32 `json:"carbohydrate"`
	Calories     float32 `json:"calories"`
}
