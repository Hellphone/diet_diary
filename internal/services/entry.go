package services

import (
	"time"

	"diet_diary/internal/domain"
	"diet_diary/internal/database"
	"diet_diary/internal/repositories"
)

func GetEntries(filter *database.Filter) (domain.EntrySet, error) {
	return repositories.GetEntries(filter)
}

func GetEntryById(id int64) (*domain.Entry, error) {
	return repositories.GetEntryById(id)
}

func InsertEntry(entry *domain.Entry) (int64, error) {
	if entry.CreatedAt == nil || entry.CreatedAt.IsZero() {
		now := time.Now()
		entry.CreatedAt = &now
	}

	// TODO: if entry_order is not defined, set it to max(entry_order) of the chosen day + 1

	return repositories.InsertEntry(database.DomainToEntry(entry))
}

func InsertEntrySet(es domain.EntrySet) error {
	for _, e := range es {
		if e.CreatedAt.IsZero() {
			now := time.Now()
			e.CreatedAt = &now
		}
	}

	return repositories.InsertEntrySet(es)
}

func UpdateEntry(entry *domain.Entry) (int64, error) {
	_, err := GetEntryById(entry.ID)
	if err != nil {
		return 0, err
	}

	if entry.Product.Calories == 0 {
		entry.CalcCalories()
	}

	return repositories.UpdateEntry(entry)
}

func DeleteEntry(id int64) (int64, error) {
	_, err := GetEntryById(id)
	if err != nil {
		return 0, err
	}

	return repositories.DeleteEntry(id)
}
