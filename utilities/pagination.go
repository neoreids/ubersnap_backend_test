package utilities

import (
	"fmt"
	"github.com/aws/smithy-go/ptr"
	"gorm.io/gorm"
	"math"
)

type Pagination struct {
	Limit         int     `json:"limit,omitempty;query:limit"`
	Page          int     `json:"page,omitempty;query:page"`
	Sort          *string `json:"sort,omitempty;query:sort"`
	SortDirection *string
	TotalRows     int64       `json:"total_rows"`
	TotalPages    int         `json:"total_pages"`
	Rows          interface{} `json:"rows,omitempty"`
}

func (p *Pagination) GetOffset() int {
	return (p.GetPage() - 1) * p.GetLimit()
}

func (p *Pagination) GetLimit() int {
	if p.Limit == 0 {
		p.Limit = 10
	}
	return p.Limit
}

func (p *Pagination) GetPage() int {
	if p.Page == 0 {
		p.Page = 1
	}
	return p.Page
}

func (p *Pagination) GetSort() string {
	if p.Sort == nil {
		p.Sort = ptr.String("id desc")
	}
	return fmt.Sprintf("%s %s", *p.Sort, *p.SortDirection)
}

func NewPagination(page int, limit int, sort string, sortDirection string) *Pagination {
	pagination := &Pagination{Limit: limit, Page: page, Sort: ptr.String(sort), SortDirection: ptr.String(sortDirection)}
	return pagination
}

func (p *Pagination) Paginate(model interface{}, db *gorm.DB) func(dbScope *gorm.DB) *gorm.DB {
	var totalRows int64
	db.Model(model).Count(&totalRows)

	p.TotalRows = totalRows
	totalPages := int(math.Ceil(float64(totalRows) / float64(p.Limit)))
	p.TotalPages = totalPages
	return func(dbScope *gorm.DB) *gorm.DB {
		return db.Model(model).Offset(p.GetOffset()).Limit(p.GetLimit()).Order(p.GetSort())
	}
}

func (p *Pagination) BuildMeta() map[string]interface{} {
	return map[string]interface{}{
		"current_page": p.GetPage(),
		"limit":        p.GetLimit(),
		"total_rows":   p.TotalRows,
		"total_page":   p.TotalPages,
	}
}
