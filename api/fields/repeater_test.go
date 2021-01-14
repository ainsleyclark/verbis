package fields

import (
	"github.com/ainsleyclark/verbis/api/domain"
)

func (t *FieldTestSuite) TestService_GetRepeater() {

	tt := map[string]struct {
		fields []domain.PostField
		key    string
		want   interface{}
	}{
		"Test": {
			fields: []domain.PostField{
				{Id: 1, Type: "repeater", Name: "repeater", OriginalValue: ""},
				{Id: 2, Type: "text", Name: "name", Value: 2, Key: "repeater_0_text"},
				{Id: 3, Type: "text", Name: "name", Value: 3, Key: "repeater_1_text"},
			},
			key: "repeater",
		},
	}
	//	"Success": {
	//		fields: []domain.PostField{
	//			{Id: 1, Type: "repeater", UUID: uniq, Name: "key1", OriginalValue: "", Parent: nil},
	//			{Id: 2, Type: "text", Name: "key2", Value: 2, Parent: &uniq},
	//			{Id: 3, Type: "text", Name: "key3", Value: 3, Parent: &uniq},
	//			{Id: 4, Type: "text", Name: "key4", Value: 4, Parent: &uniq},
	//		},
	//		key: "key1",
	//		want: Repeater{
	//			{Id: 2, Type: "text", Name: "key2", Value: 2, Parent: &uniq},
	//			{Id: 3, Type: "text", Name: "key3", Value: 3, Parent: &uniq},
	//			{Id: 4, Type: "text", Name: "key4", Value: 4, Parent: &uniq},
	//		},
	//	},
	//	"Sorted Index": {
	//		fields: []domain.PostField{
	//			{Id: 1, Type: "repeater", UUID: uniq, Name: "key1", OriginalValue: "", Parent: nil},
	//			{Id: 2, Type: "text", Name: "key2", Value: 2, Parent: &uniq, Index: &two},
	//			{Id: 3, Type: "text", Name: "key3", Value: 3, Parent: &uniq, Index: &zero},
	//			{Id: 4, Type: "text", Name: "key4", Value: 4, Parent: &uniq, Index: &one},
	//		},
	//		key: "key1",
	//		want: Repeater{
	//			{Id: 3, Type: "text", Name: "key3", Value: 3, Parent: &uniq, Index: &zero},
	//			{Id: 4, Type: "text", Name: "key4", Value: 4, Parent: &uniq, Index: &one},
	//			{Id: 2, Type: "text", Name: "key2", Value: 2, Parent: &uniq, Index: &two},
	//		},
	//	},
	//	"Parent": {
	//		fields: []domain.PostField{
	//			{Id: 1, Type: "repeater", UUID: uniq2, Name: "key1", Value: 1, Parent: &uniq},
	//			{Id: 2, Type: "text", Name: "key2", Value: 2, Parent: &uniq2, Index: &two},
	//			{Id: 3, Type: "text", Name: "key3", Value: 3, Parent: &uniq2, Index: &zero},
	//			{Id: 4, Type: "text", Name: "key4", Value: 4, Parent: &uniq2, Index: &one},
	//		},
	//		key: "key1",
	//		want: Repeater{
	//			{Id: 3, Type: "text", Name: "key3", Value: 3, Parent: &uniq2, Index: &zero},
	//			{Id: 4, Type: "text", Name: "key4", Value: 4, Parent: &uniq2, Index: &one},
	//			{Id: 2, Type: "text", Name: "key2", Value: 2, Parent: &uniq2, Index: &two},
	//		},
	//	},
	//	"Not Found": {
	//		fields: []domain.PostField{},
	//		key:    "wrongval",
	//		want:   "no field exists with the name: wrongval",
	//	},
	//	"Invalid Type": {
	//		fields: []domain.PostField{
	//			{Id: 1, Type: "text", UUID: uniq, Name: "key1", Value: 1, Parent: nil},
	//		},
	//		key:  "key1",
	//		want: "field with the name: key1, is not a repeater",
	//	},
	//}

	for name, test := range tt {
		t.Run(name, func() {
			s := t.GetService(test.fields)

			got, err := s.GetRepeater(test.key)
			if err != nil {
				t.Contains(err.Error(), test.want)
				return
			}

			t.Equal(test.want, got)
		})
	}
}

//
//func (t *FieldTestSuite) TestRepeater_HasRows() {
//
//	tt := map[string]struct {
//		repeater Repeater
//		want     interface{}
//	}{
//		"With Rows": {
//			repeater: Repeater{
//				{Id: 1}, {Id: 2}, {Id: 3},
//			},
//			want: true,
//		},
//		"Without Rows": {
//			repeater: Repeater{},
//			want:     false,
//		},
//	}
//
//	for name, test := range tt {
//		t.Run(name, func() {
//			t.Equal(test.want, test.repeater.HasRows())
//		})
//	}
//}
//
//func (t *FieldTestSuite) TestRepeater_SubField() {
//
//	repeater := Repeater{
//		{Id: 1, Name: "test1", Value: 1},
//		{Id: 2, Name: "test2", Value: 2},
//		{Id: 3, Name: "test3", Value: 3},
//	}
//
//	tt := map[string]struct {
//		key  string
//		want interface{}
//	}{
//		"Found": {
//			key:  "test1",
//			want: 1,
//		},
//		"Not Found": {
//			key:  "wrongval",
//			want: nil,
//		},
//	}
//
//	for name, test := range tt {
//		t.Run(name, func() {
//			t.Equal(test.want, repeater.SubField(test.key))
//		})
//	}
//}
//
//func (t *FieldTestSuite) TestRepeater_First() {
//
//	tt := map[string]struct {
//		repeater Repeater
//		want     interface{}
//	}{
//		"Found": {
//			repeater: Repeater{
//				{Id: 1, Name: "test1", Value: 1},
//				{Id: 2, Name: "test2", Value: 2},
//				{Id: 3, Name: "test3", Value: 3},
//			},
//			want: domain.PostField{Id: 1, Name: "test1", Value: 1},
//		},
//		"Not Found": {
//			repeater: Repeater{},
//			want:     nil,
//		},
//	}
//
//	for name, test := range tt {
//		t.Run(name, func() {
//			t.Equal(test.want, test.repeater.First())
//		})
//	}
//}
//
//func (t *FieldTestSuite) TestRepeater_Last() {
//
//	tt := map[string]struct {
//		repeater Repeater
//		want     interface{}
//	}{
//		"Found": {
//			repeater: Repeater{
//				{Id: 1, Name: "test1", Value: 1},
//				{Id: 2, Name: "test2", Value: 2},
//				{Id: 3, Name: "test3", Value: 3},
//			},
//			want: domain.PostField{Id: 3, Name: "test3", Value: 3},
//		},
//		"Not Found": {
//			repeater: Repeater{},
//			want:     nil,
//		},
//	}
//
//	for name, test := range tt {
//		t.Run(name, func() {
//			t.Equal(test.want, test.repeater.Last())
//		})
//	}
//}
