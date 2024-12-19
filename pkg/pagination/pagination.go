package pagination

import (
	"gorm.io/gorm"
)

type Pagination struct {
	Page  *int `form:"page,omitempty" binding:"omitempty,min=1"`
	Limit *int `form:"limit,omitempty" binding:"omitempty,min=1,max=50"`
}

func (p *Pagination) Paginate() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset((p.GetPage() - 1) * p.GetLimit()).Limit(p.GetLimit())
	}
}

func (p *Pagination) GetPage() int {
	if p.Page == nil {
		return 1
	}

	return *p.Page
}

func (p *Pagination) GetLimit() int {
	if p.Limit == nil {
		return 5
	}

	if *p.Limit > 50 {
		return 50
	}

	return *p.Limit
}
