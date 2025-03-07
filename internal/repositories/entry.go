package repositories

import (
	"fmt"
	"log"

	"diet_diary/internal/database"
	"diet_diary/internal/domain"

	sq "github.com/Masterminds/squirrel"
)

const tableEntry = "entries"

func GetEntryById(id int64) (*domain.Entry, error) {
	query := `SELECT * FROM entries WHERE id = $1`

	entry := database.Entry{}
	err := database.DB.Get(&entry, query, id)
	if err != nil {
		log.Println("Error getting entry by id:", err)
		return nil, err
	}

	return database.EntryToDomain(&entry), nil
}

// TODO: separate EntrySet and []*Entry
func GetEntries(filter *database.Filter) (domain.EntrySet, error) {
	var err error
	builder := sq.Select(
		"entries.id AS id",
		"entries.quantity AS quantity",
		"entries.created_at AS created_at",
		"entries.entry_order AS entry_order",
		"products.id AS id",
		"products.name AS name",
		"products.protein AS protein",
		"products.fat AS fat",
		"products.carbohydrate AS carbohydrate",
		"products.calories AS calories",
	).From(tableEntry)
	builder, err = SetFilters(builder, filter)
	if err != nil {
		log.Println("Error applying filters:", err)
		return nil, err
	}

	builder = builder.Join(fmt.Sprintf("%[1]s ON %[1]s.%s = %s.%s", tableProduct, "id", tableEntry, "product"))
	entrySet := database.EntrySet{}
	query, args, err := builder.ToSql()
	if err != nil {
		log.Println("Error building SQL query:", err)
		return nil, err
	}

	err = database.DB.Select(&entrySet, rebind(query), args...)
	if err != nil {
		log.Println("Error getting entries:", err)
		return nil, err
	}

	return database.EntrySetToDomain(entrySet), nil
}

func InsertEntry(entry *database.Entry) (int64, error) {
	query := `INSERT INTO entries (product, quantity, created_at, entry_order)
			VALUES (:product_id, :quantity, :created_at, :entry_order) RETURNING id`

	rows, err := database.DB.NamedQuery(query, entry)
	if err != nil {
		log.Println("Error inserting entry:", err)
		return 0, err
	}
	defer rows.Close()

	var id int64
	if rows.Next() {
		err := rows.Scan(&id)
		if err != nil {
			return 0, err
		}
	}

	return id, nil
}

// TODO: EntrySets should be stored in a separate table and connected
// with products and users
func InsertEntrySet(es domain.EntrySet) error {
	query := `INSERT INTO entries (entry_id, product, quantity, created_at, order)
			VALUES (:entry_id, :product, :quantity, :created_at, :order)`

	_, err := database.DB.NamedExec(query, es)
	if err != nil {
		log.Println("Error inserting entry set:", err)
		return err
	}

	return  nil
}

func UpdateEntry(entry *domain.Entry) (int64, error) {
	query := `UPDATE entries
              SET name = :name, protein = :protein, fat = :fat, carbohydrate = :carbohydrate, calories = :calories
			  WHERE id = :id`

	rows, err := database.DB.NamedQuery(query, entry)
	if err != nil {
		log.Println("Error updating entry:", err)
		return 0, err
	}
	defer rows.Close()

	return entry.ID, nil
}

func DeleteEntry(id int64) (int64, error) {
	query := `DELETE FROM entries
	 		  WHERE id = $1`

	_, err := database.DB.Exec(query, id)
	if err != nil {
		log.Printf("Error deleting entry with id = %d: %v", id, err)
		return 0, err
	}

	return id, nil
}
