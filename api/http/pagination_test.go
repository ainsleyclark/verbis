package http

import (
	params2 "github.com/ainsleyclark/verbis/api/helpers/params"
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestNewPagination - Test construct
func TestNewPagination(t *testing.T) {
	want := &Pagination{}
	got := NewPagination()
	assert.Equal(t, got, want)
}

// // TestPagination_Get - Test get Pagination
func TestPagination_Get(t *testing.T) {

	tt := map[string]struct {
		params params2.Params
		total  int
		want   *Pagination
	}{
		"Normal": {
			total:  10,
			params: params2.Params{Page: 1, Limit: 10, OrderBy: "order", OrderDirection: "asc", Filters: nil},
			want:   &Pagination{Page: 1, Pages: 1, Limit: 10, Total: 10, Next: false, Prev: false},
		},
		"Next": {
			total:  10,
			params: params2.Params{Page: 1, Limit: 5, OrderBy: "order", OrderDirection: "asc", Filters: nil},
			want:   &Pagination{Page: 1, Pages: 2, Limit: 5, Total: 10, Next: 2, Prev: false},
		},
		"Prev": {
			total:  10,
			params: params2.Params{Page: 2, Limit: 5, OrderBy: "order", OrderDirection: "asc", Filters: nil},
			want:   &Pagination{Page: 2, Pages: 2, Limit: 5, Total: 10, Next: false, Prev: 1},
		},
		"Prev & Next": {
			total:  100,
			params: params2.Params{Page: 5, Limit: 5, OrderBy: "order", OrderDirection: "asc", Filters: nil},
			want:   &Pagination{Page: 5, Pages: 20, Limit: 5, Total: 100, Next: 6, Prev: 4},
		},
		"Odd": {
			total:  600,
			params: params2.Params{Page: 99, Limit: 4, OrderBy: "order", OrderDirection: "asc", Filters: nil},
			want:   &Pagination{Page: 99, Pages: 150, Limit: 4, Total: 600, Next: 100, Prev: 98},
		},
		"Limit All": {
			total:  100,
			params: params2.Params{Page: 1, Limit: 0, LimitAll: true, OrderBy: "order", OrderDirection: "asc", Filters: nil},
			want:   &Pagination{Page: 1, Pages: 1, Limit: "all", Total: 100, Next: false, Prev: false},
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			p := NewPagination()
			got := p.Get(test.params, test.total)
			assert.Equal(t, test.want, got)
		})
	}
}
