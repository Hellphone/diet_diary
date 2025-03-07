package repositories

import (
	db "diet_diary/internal/database"
	"fmt"
	"strings"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

func rebind(query string) string {
	return sqlx.Rebind(sqlx.DOLLAR, query)
}

func SetFilters(builder sq.SelectBuilder, filter *db.Filter) (sq.SelectBuilder, error) {
	if filter != nil {
		for _, f := range filter.Filters {
			switch f.Op {
			case db.Eq:
				builder = builder.Where(sq.Eq{f.Field: f.Value})
			case db.Gt:
				builder = builder.Where(sq.Gt{f.Field: f.Value})
			case db.Lt:
				builder = builder.Where(sq.Lt{f.Field: f.Value})
			case db.Like:
				field := fmt.Sprintf("%s(%s)", "LOWER", f.Field)
				value := fmt.Sprintf("%%%s%%", strings.ToLower(f.Value))
				builder = builder.Where(sq.Like{field: value})
			default:
				return sq.SelectBuilder{}, fmt.Errorf("unsupported SQL operator")
			}
		}

		if filter.Limit > 0 {
			builder = builder.Limit(uint64(filter.Limit))
		}

		if filter.Offset > 0 {
			builder = builder.Offset(uint64(filter.Offset))
		}

		if len(filter.OrderBy) > 0 {
			builder = builder.OrderBy(filter.OrderBy)
		}
	}

	return builder, nil
}
