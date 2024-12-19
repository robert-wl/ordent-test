package pagination

import (
	"github.com/go-playground/assert/v2"
	"testing"
)

func TestNormalPagination(t *testing.T) {
	p := &Pagination{}

	assert.Equal(t, p.GetPage(), 1)

	assert.Equal(t, p.GetLimit(), 5)
}

func TestCustomPagination(t *testing.T) {
	page := 6
	limit := 99

	p := &Pagination{
		Page:  &page,
		Limit: &limit,
	}

	assert.Equal(t, p.GetPage(), 6)

	assert.Equal(t, p.GetLimit(), 50)
}

func TestIncorrectPagination(t *testing.T) {
	page := -1
	limit := -1

	p := &Pagination{
		Page:  &page,
		Limit: &limit,
	}

	assert.Equal(t, p.GetPage(), 1)

	assert.Equal(t, p.GetLimit(), 5)
}

func TestIncorrectPagination2(t *testing.T) {
	page := 0
	limit := 999999

	p := &Pagination{
		Page:  &page,
		Limit: &limit,
	}

	assert.Equal(t, p.GetPage(), 1)

	assert.Equal(t, p.GetLimit(), 50)
}
