package database

import (
	"diet_diary/internal/domain"
)

func ProductToDomain(p *Product) *domain.Product {
	return &domain.Product{
		ID:   p.ID,
		Name: p.Name,
		Macro: &domain.Macro{
			Protein:      p.Macro.Protein,
			Fat:          p.Macro.Fat,
			Carbohydrate: p.Macro.Carbohydrate,
			Calories:     p.Macro.Calories,
		},
	}
}

func ProductsToDomain(ps Products) domain.Products {
	products := make(domain.Products, 0, len(ps))

	for _, p := range ps {
		products = append(products, ProductToDomain(p))
	}

	return products
}

func EntryToDomain(e *Entry) *domain.Entry {
	return &domain.Entry{
		ID: e.ID,
		Product: &domain.Product{
			ID:   e.Product.ID,
			Name: e.Product.Name,
			Macro: &domain.Macro{
				Protein:      e.Product.Macro.Protein,
				Fat:          e.Product.Macro.Fat,
				Carbohydrate: e.Product.Macro.Carbohydrate,
			},
		},
		Quantity:  e.Quantity,
		CreatedAt: e.CreatedAt,
		Order:     e.Order,
	}
}

func EntrySetToDomain(es EntrySet) domain.EntrySet {
	set := make(domain.EntrySet, 0, len(es))

	for _, p := range es {
		set = append(set, EntryToDomain(p))
	}

	return set
}

func DomainToEntry(e *domain.Entry) *Entry {
	entry := &Entry{
		ID: e.ID,
		Product: &Product{
			ID:   e.Product.ID,
			Name: e.Product.Name,
		},
		Quantity:  e.Quantity,
		CreatedAt: e.CreatedAt,
		Order:     e.Order,
	}

	if e.Product.Macro != nil {
		entry.Product.Macro = &Macro{
			Protein:      e.Product.Macro.Protein,
			Fat:          e.Product.Macro.Fat,
			Carbohydrate: e.Product.Macro.Carbohydrate,
		}
	}

	return entry
}

func DomainToEntrySet(es domain.EntrySet) EntrySet {
	set := make(EntrySet, 0, len(es))

	for _, p := range es {
		set = append(set, DomainToEntry(p))
	}

	return set
}
