package templates

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	vhttp "github.com/ainsleyclark/verbis/api/http"
	mocks "github.com/ainsleyclark/verbis/api/mocks/models"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	//"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_GetPost(t *testing.T) {

	post := domain.Post{
		Id:     1,
		Title:  "test title",
		UserId: 1,
	}

	author := &domain.PostAuthor{}
	category := &domain.PostCategory{}

	viewData := ViewPost{
		Author:   author,
		Category: category,
		Post:     post,
	}

	tt := map[string]struct {
		input interface{}
		mock  func(m *mocks.PostsRepository)
		want  interface{}
	}{
		"Success": {
			input: 1,
			mock: func(m *mocks.PostsRepository) {
				m.On("GetById", 1).Return(post, nil)
				m.On("Format", post).Return(domain.PostData{Post: post, Author: author, Category: category}, nil)
			},
			want: viewData,
		},
		"Format Error": {
			input: 1,
			mock: func(m *mocks.PostsRepository) {
				m.On("GetById", 1).Return(post, nil)
				m.On("Format", post).Return(domain.PostData{}, fmt.Errorf("error"))
			},
			want: nil,
		},
		"Not Found": {
			input: 1,
			mock: func(m *mocks.PostsRepository) {
				m.On("GetById", 1).Return(domain.Post{}, fmt.Errorf("error"))
				m.On("Format", post).Return(domain.PostData{Post: post, Author: author, Category: category}, nil)
			},
			want: nil,
		},
		"No Stringer": {
			input: noStringer{},
			mock: func(m *mocks.PostsRepository) {
				m.On("GetById", 1).Return(post, nil)
				m.On("Format", post).Return(domain.PostData{Post: post, Author: author, Category: category}, nil)
			},
			want: nil,
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			f := newTestSuite()
			postsMock := mocks.PostsRepository{}

			test.mock(&postsMock)
			f.store.Posts = &postsMock

			tpl := `{{ post .PostId }}`

			runtv(t, f, tpl, test.want, map[string]interface{}{"PostId": test.input})
		})
	}
}

func Test_GetPosts(t *testing.T) {

	post := domain.Post{Id: 1, Title: "Title"}
	posts := []domain.Post{
		post, post,
	}

	author := &domain.PostAuthor{}
	category := &domain.PostCategory{}
	viewData := []ViewPost{
		{
			Author:   author,
			Category: category,
			Post:     post,
		},
		{
			Author:   author,
			Category: category,
			Post:     post,
		},
	}

	categoryTest := &domain.PostCategory{
		Name: "cat",
	}
	viewDataCategory := []ViewPost{
		{
			Author:   author,
			Category: categoryTest,
			Post:     post,
		},
		{
			Author:   author,
			Category: categoryTest,
			Post:     post,
		},
	}

	tt := map[string]struct {
		input map[string]interface{}
		mock  func(m *mocks.PostsRepository)
		want  interface{}
	}{
		"Success": {
			input: map[string]interface{}{"limit": 15},
			mock: func(m *mocks.PostsRepository) {
				m.On("Get", vhttp.Params{Page: 1, Limit: 15, LimitAll: false, OrderBy: "published_at", OrderDirection: "desc"}, "all", "published").Return(posts, 5, nil)
				m.On("Format", post).Return(domain.PostData{Post: post, Author: author, Category: category}, nil)
			},
			want: map[string]interface{}{
				"Posts": viewData,
				"Pagination": &vhttp.Pagination{
					Page:  1,
					Pages: 1,
					Limit: 15,
					Total: 5,
					Next:  false,
					Prev:  false,
				},
			},
		},
		"Failed Params": {
			input: map[string]interface{}{"limit": "wrongval"},
			mock: func(m *mocks.PostsRepository) {
				m.On("Get", vhttp.Params{Page: 1, Limit: 15, LimitAll: false, OrderBy: "published_at", OrderDirection: "desc"}, "all", "published").Return(posts, 5, nil)
				m.On("Format", post).Return(domain.PostData{Post: post, Author: author, Category: category}, nil)
			},
			want: "cannot unmarshal string into Go struct field TemplateParams.limit",
		},
		"Not Found": {
			input: map[string]interface{}{"limit": 15},
			mock: func(m *mocks.PostsRepository) {
				m.On("Get", vhttp.Params{Page: 1, Limit: 15, LimitAll: false, OrderBy: "published_at", OrderDirection: "desc"}, "all", "published").Return(nil, 0, &errors.Error{Code: errors.NOTFOUND, Message: "no posts found"})
				m.On("Format", post).Return(domain.PostData{Post: post, Author: author, Category: category}, nil)
			},
			want: map[string]interface{}(nil),
		},
		"Internal Error": {
			input: map[string]interface{}{"limit": 15},
			mock: func(m *mocks.PostsRepository) {
				m.On("Get", vhttp.Params{Page: 1, Limit: 15, LimitAll: false, OrderBy: "published_at", OrderDirection: "desc"}, "all", "published").Return(nil, 0, &errors.Error{Code: errors.INTERNAL, Message: "internal error"})
				m.On("Format", post).Return(domain.PostData{Post: post, Author: author, Category: category}, nil)
			},
			want: "internal error",
		},
		"Format Error": {
			input: map[string]interface{}{"limit": 15},
			mock: func(m *mocks.PostsRepository) {
				m.On("Get", vhttp.Params{Page: 1, Limit: 15, LimitAll: false, OrderBy: "published_at", OrderDirection: "desc"}, "all", "published").Return(posts, 5, nil)
				m.On("Format", post).Return(domain.PostData{Post: post, Author: author, Category: category}, fmt.Errorf("error"))
			},
			want: map[string]interface{}{
				"Posts": []ViewPost(nil),
				"Pagination": &vhttp.Pagination{
					Page:  1,
					Pages: 1,
					Limit: 15,
					Total: 5,
					Next:  false,
					Prev:  false,
				},
			},
		},
		"Category": {
			input: map[string]interface{}{"limit": 15, "category": "cat"},
			mock: func(m *mocks.PostsRepository) {
				m.On("Get", vhttp.Params{Page: 1, Limit: 15, LimitAll: false, OrderBy: "published_at", OrderDirection: "desc"}, "all", "published").Return(posts, 2, nil)
				m.On("Format", post).Return(domain.PostData{Post: post, Author: author, Category: categoryTest}, nil)
			},
			want: map[string]interface{}{
				"Posts": viewDataCategory,
				"Pagination": &vhttp.Pagination{
					Page:  1,
					Pages: 1,
					Limit: 15,
					Total: 2,
					Next:  false,
					Prev:  false,
				},
			},
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			f := newTestSuite()
			postsMock := mocks.PostsRepository{}

			test.mock(&postsMock)
			f.store.Posts = &postsMock

			p, err := f.getPosts(test.input)
			if err != nil {
				assert.Contains(t, err.Error(), test.want)
				return
			}

			assert.EqualValues(t, test.want, p)
		})
	}
}

func TestGetPagination(t *testing.T) {
	f := newTestSuite()
	gin.SetMode(gin.ReleaseMode)
	g, _ := gin.CreateTestContext(httptest.NewRecorder())
	g.Request, _ = http.NewRequest("GET", "/get?page=123", nil)
	f.gin = g
	tpl := "{{ paginationPage }}"
	runt(t, f, tpl, 123)
}

func TestGetPagination_NoPage(t *testing.T) {
	f := newTestSuite()
	tpl := "{{ paginationPage }}"
	runt(t, f, tpl, 1)
}

func TestGetPagination_ConvertString(t *testing.T) {
	f := newTestSuite()
	gin.SetMode(gin.ReleaseMode)
	g, _ := gin.CreateTestContext(httptest.NewRecorder())
	g.Request, _ = http.NewRequest("GET", "/get?page=wrongval", nil)
	f.gin = g
	tpl := "{{ paginationPage }}"
	runt(t, f, tpl, "1")
}
