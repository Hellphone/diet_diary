package services

import (
	"database/sql"
	"fmt"
	"time"

	"diet_diary/internal/database"
	"diet_diary/internal/domain"
	"diet_diary/internal/repositories"
)

func GetEntries(filter *database.Filter) (domain.EntrySet, error) {
	return repositories.GetEntries(filter)
}

func GetEntryById(id int64) (*domain.Entry, error) {
	return repositories.GetEntryById(id)
}

func GetEntryTotalByDate(date *time.Time) (*domain.EntryTotal, error) {
	filter := &database.Filter{
		Filters: []database.SQLFilter{
			{Op: "gt", Field: "created_at", Value: date.Format(time.DateOnly)},
			{Op: "lt", Field: "created_at", Value: date.Add(24*time.Hour).Format(time.DateOnly)},
		},
	}

	entrySet, err := repositories.GetEntries(filter)
	if err != nil {
		return nil, fmt.Errorf("repositories.GetEntries: %v", err)
	}

	if len(entrySet) == 0 {
		return nil, sql.ErrNoRows
	}

	return entrySet.Total(), nil
}

func InsertEntry(entry *domain.Entry) (int64, error) {
	if entry.CreatedAt == nil || entry.CreatedAt.IsZero() {
		now := time.Now()
		entry.CreatedAt = &now
	}

	if entry.Order == 0 {
		maxOrder, err := repositories.GetMaxEntryOrderByDate(entry.CreatedAt)
		if err != nil && err != sql.ErrNoRows {
			return 0, err
		}

		entry.Order = maxOrder + 1
	}

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
