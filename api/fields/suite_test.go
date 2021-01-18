package fields

import (
	"bytes"
	"github.com/ainsleyclark/verbis/api/domain"
	mocks "github.com/ainsleyclark/verbis/api/mocks/models"
	"github.com/ainsleyclark/verbis/api/models"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/suite"
	"testing"
)

type FieldTestSuite struct {
	suite.Suite
	logWriter bytes.Buffer
}

type noStringer struct{}

func TestFields(t *testing.T) {
	suite.Run(t, new(FieldTestSuite))
}

func (t *FieldTestSuite) BeforeTest(suiteName, testName string) {
	b := bytes.Buffer{}
	t.logWriter = b
	log.SetOutput(&t.logWriter)
}

func (t *FieldTestSuite) Reset() {
	t.logWriter.Reset()
}

func (t *FieldTestSuite) GetMockService(fields []domain.PostField, fnc func(f *mocks.FieldsRepository, c *mocks.CategoryRepository)) *Service {
	fieldsMock := &mocks.FieldsRepository{}
	categoryMock := &mocks.CategoryRepository{}

	if fnc != nil {
		fnc(fieldsMock, categoryMock)
	}

	s := t.GetService(fields)
	s.store = &models.Store{
		Categories: categoryMock,
		Fields:     fieldsMock,
	}

	return s
}

func (t *FieldTestSuite) GetPostsMockService(fields []domain.PostField, fnc func(p *mocks.PostsRepository)) *Service {
	postsMocks := &mocks.PostsRepository{}

	if fnc != nil {
		fnc(postsMocks)
	}

	s := t.GetService(fields)
	s.store = &models.Store{
		Posts: postsMocks,
	}

	return s
}

func (t *FieldTestSuite) GetTypeMockService(fnc func(c *mocks.CategoryRepository, m *mocks.MediaRepository, p *mocks.PostsRepository, u *mocks.UserRepository)) *Service {

	categoryMock := &mocks.CategoryRepository{}
	mediaMock := &mocks.MediaRepository{}
	postsMock := &mocks.PostsRepository{}
	userMock := &mocks.UserRepository{}

	if fnc != nil {
		fnc(categoryMock, mediaMock, postsMock, userMock)
	}

	s := t.GetService(nil)
	s.store = &models.Store{
		Categories: categoryMock,
		Media:      mediaMock,
		Posts:      postsMock,
		User:       userMock,
	}

	return s
}

func (t *FieldTestSuite) GetService(fields []domain.PostField) *Service {
	return &Service{
		fields: fields,
		store:  &models.Store{},
	}
}
