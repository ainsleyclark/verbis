package users

import (
	"fmt"
	"github.com/ainsleyclark/verbis/api/deps"
	"github.com/ainsleyclark/verbis/api/domain"
	"github.com/ainsleyclark/verbis/api/errors"
	"github.com/ainsleyclark/verbis/api/models"
	"github.com/ainsleyclark/verbis/api/tpl/params"
	"github.com/stretchr/testify/assert"
	"testing"

	//"github.com/ainsleyclark/verbis/api/errors"
	vhttp "github.com/ainsleyclark/verbis/api/http"
	mocks "github.com/ainsleyclark/verbis/api/mocks/models"
	//"time"
)

type noStringer struct{}

func Setup() (*Namespace, *mocks.UserRepository) {
	mock := &mocks.UserRepository{}
	return &Namespace{deps: &deps.Deps{
		Store: &models.Store{
			User: mock,
		},
	}}, mock
}

func TestNamespace_Find(t *testing.T) {

	user := domain.User{
		UserPart: domain.UserPart{Id: 1, FirstName: "verbis"},
	}

	tt := map[string]struct {
		input interface{}
		mock  func(m *mocks.UserRepository)
		want  interface{}
	}{
		"Success": {
			1,
			func(m *mocks.UserRepository) {
				m.On("GetById", 1).Return(user, nil).Once()
			},
			user.HideCredentials(),
		},
		"Not Found": {
			1,
			func(m *mocks.UserRepository) {
				m.On("GetById", 1).Return(domain.User{}, fmt.Errorf("error")).Once()
			},
			nil,
		},
		"No Stringer": {
			noStringer{},
			func(m *mocks.UserRepository) {
				m.On("GetById", 1).Return(user, nil).Once()
			},
			nil,
		},
		"Nil": {
			nil,
			func(m *mocks.UserRepository) {
				m.On("GetById", 1).Return(user, nil).Once()
			},
			nil,
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			ns, mock := Setup()
			test.mock(mock)
			got := ns.Find(test.input)
			assert.Equal(t, test.want, got)
		})
	}
}

func TestNamespace_List(t *testing.T) {

	users := domain.Users{
		domain.User{UserPart: domain.UserPart{Id: 1, FirstName: "verbis"}},
		domain.User{UserPart: domain.UserPart{Id: 1, FirstName: "cms"}},
	}

	p := vhttp.Params{
		Page:           1,
		Limit:          15,
		LimitAll:       false,
		OrderBy:        OrderBy,
		OrderDirection: OrderDirection,
	}

	tt := map[string]struct {
		input params.Query
		mock  func(m *mocks.UserRepository)
		want  interface{}
	}{
		"Success": {
			params.Query{"limit": 15},
			func(m *mocks.UserRepository) {
				m.On("Get", p).Return(users, 2, nil).Once()
			},
			Users{
				Users: users.HideCredentials(),
				Pagination: &vhttp.Pagination{
					Page:  1,
					Pages: 1,
					Limit: 15,
					Total: 2,
					Next:  false,
					Prev:  false,
				},
			},
		},
		"Nil": {
			nil,
			func(m *mocks.UserRepository) {
				m.On("Get", p).Return(users, 2, nil).Once()
			},
			Users{
				Users: users.HideCredentials(),
				Pagination: &vhttp.Pagination{
					Page:  1,
					Pages: 1,
					Limit: 15,
					Total: 2,
					Next:  false,
					Prev:  false,
				},
			},
		},
		"Not Found": {
			params.Query{"limit": 15},
			func(m *mocks.UserRepository) {
				m.On("Get", p).Return(nil, 0, &errors.Error{Code: errors.NOTFOUND, Message: "no users found"}).Once()
			},
			nil,
		},
		"Internal Error": {
			params.Query{"limit": 15},
			func(m *mocks.UserRepository) {
				m.On("Get", p).Return(nil, 0, &errors.Error{Code: errors.INTERNAL, Message: "internal error"}).Once()
			},
			"internal error",
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			ns, mock := Setup()
			test.mock(mock)
			got, err := ns.List(test.input)
			if err != nil {
				assert.Contains(t, err.Error(), test.want)
				return
			}
			assert.Equal(t, test.want, got)
		})
	}
}
