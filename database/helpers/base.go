package dbHelpers

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/burabatbold/delivery-auth-service/common"
	"github.com/burabatbold/delivery-auth-service/utils"
	df "github.com/survivorbat/gorm-deep-filtering"
	"gorm.io/gorm"
)

type DBHelper struct {
	*gorm.DB
}

func NewOrm(db *gorm.DB) *DBHelper {
	return &DBHelper{DB: db}
}

func (d *DBHelper) FilterByBase(fields []string, input *common.BaseFilterInput) *DBHelper {
	if input == nil {
		return d
	}

	if input.CreatedAt != nil && input.CreatedAt[0] != nil && input.CreatedAt[1] != nil {
		d.DB = d.DB.Where(fmt.Sprintf("date(%v.created_at) BETWEEN date(?) and date(?)", input.TableName), input.CreatedAt[0], input.CreatedAt[1])
	}

	if input.Search != nil && len(*input.Search) > 0 {
		d = d.Search(fields, input.Search)
	}

	if input.StartDate != nil && input.EndDate != nil {
		d.DB = d.DB.Where(fmt.Sprintf("date(%v.created_at) BETWEEN date(?) and date(?)", input.TableName), input.StartDate, input.EndDate)
	}

	return d
}

func (d *DBHelper) Sort(input *common.PaginateInput, tableName string) *DBHelper {
	if input != nil && input.Sorter != nil {
		for key, value := range *input.Sorter {
			order := "ASC"
			if strings.HasPrefix(strings.ToLower(value), "desc") {
				order = "DESC"
			}
			d.DB = d.DB.Order(fmt.Sprintf("%s.%s %s", tableName, key, order))
		}
	}

	return d
}

func (d *DBHelper) Paginate(input *common.PaginateInput) *DBHelper {
	d.DB = d.DB.Scopes(func(d *gorm.DB) *gorm.DB {
		if input == nil {
			return d
		}

		if input.Limit == 0 {
			input.Limit = 20
		}

		return d.Offset((input.Page) * input.Limit).Limit(input.Limit)
	})

	return d
}

func (d *DBHelper) Filter(entity interface{}, input map[string]any) *DBHelper {
	if input == nil {
		return d
	}
	db, err := df.AddDeepFilters(d.DB, entity, input)
	if err != nil {
		return d
	}

	d.DB = db

	return d
}

func (d *DBHelper) Entity(entity interface{}) *DBHelper {
	d.DB = d.DB.Model(&entity)
	return d
}

// join
func (d *DBHelper) Join(table string, on string) *DBHelper {
	d.DB = d.DB.Joins(fmt.Sprintf("JOIN %s ON %s", table, on))
	return d
}

func (d *DBHelper) LeftJoin(table string, on string) *DBHelper {
	d.DB = d.DB.Joins(fmt.Sprintf("LEFT JOIN %s ON %s", table, on))
	return d
}

func (d *DBHelper) RightJoin(table string, on string) *DBHelper {
	d.DB = d.DB.Joins(fmt.Sprintf("RIGHT JOIN %s ON %s", table, on))
	return d
}

func (d *DBHelper) Search(fields []string, value *string) *DBHelper {
	if value == nil {
		return d
	}
	queryString := ""
	for index, field := range fields {
		if index == 0 {
			queryString = fmt.Sprintf("LOWER(%s) like LOWER(@value)", field)
			continue
		}
		fieldQuery := fmt.Sprintf("LOWER(%s) like LOWER(@value)", field)
		queryString = fmt.Sprintf("%s or %s", queryString, fieldQuery)
	}
	d.DB = d.DB.Where(queryString, sql.Named("value", "%"+*value+"%"))
	return d
}

func (d *DBHelper) In(field string, value interface{}) *DBHelper {
	if value == nil || utils.IsNil(value) {
		return d
	}
	queryString := fmt.Sprintf("%s in (?)", field)
	d.DB = d.DB.Where(queryString, value)
	return d
}

func (d *DBHelper) Total() (int64, error) {
	var total int64
	if err := d.Count(&total).Error; err != nil {
		return 0, err
	}
	return total, nil
}

func (d *DBHelper) Equal(field string, value interface{}) *DBHelper {
	if utils.IsNil(value) {
		return d
	}

	d.DB = d.DB.Where(fmt.Sprintf("%s = ?", field), value)
	return d
}

func (d *DBHelper) NotEqual(field string, value *string) *DBHelper {
	if value == nil {
		return d
	}
	d.DB = d.DB.Where(fmt.Sprintf("%s != ?", field), *value)
	return d
}

func (d *DBHelper) EqualGreaterThan(field string, value *int) *DBHelper {
	if value == nil {
		return d
	}
	d.DB = d.DB.Where(fmt.Sprintf("%s >= ?", field), *value)
	return d
}

func (d *DBHelper) EqualLessThan(field string, value *int) *DBHelper {
	if value == nil {
		return d
	}
	d.DB = d.DB.Where(fmt.Sprintf("%s <= ?", field), *value)
	return d
}

func (d *DBHelper) NotIn(field string, value *string) *DBHelper {
	if value == nil {
		return d
	}
	d.DB = d.DB.Where(fmt.Sprintf("%s not in (?)", field), *value)
	return d
}

func (d *DBHelper) Bool(field string, value *string) *DBHelper {
	if value == nil {
		return d
	}
	d.DB = d.DB.Where(fmt.Sprintf("%s = ?", field), *value == "1")
	return d
}

func (d *DBHelper) Like(field string, value *string) *DBHelper {
	if value == nil {
		return d
	}
	d.DB = d.DB.Where(fmt.Sprintf("%s like ?", field), "%"+*value+"%")
	return d
}

func (d *DBHelper) BetweenDates(field string, times []*string) *DBHelper {
	if len(times) < 2 || len(times) > 2 {
		return d
	}
	d.DB = d.DB.Where(fmt.Sprintf("date(%s) between date(?) and date(?)", field), times[0], times[1])
	return d
}

func (d *DBHelper) BetweenDateTimes(field string, times []*string) *DBHelper {
	if len(times) < 2 || len(times) > 2 {
		return d
	}
	d.DB = d.DB.Where(fmt.Sprintf("%s between ? and ?", field), times[0], times[1])
	return d
}

func (d *DBHelper) EqualDate(field string, value *time.Time) *DBHelper {
	if value == nil {
		return d
	}
	d.DB = d.DB.Where(fmt.Sprintf("date(%s) = date(?)", field), value)
	return d
}
