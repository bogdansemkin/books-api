package repository

import (
	"books-api/pkg/model"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	sqlmock "github.com/zhashkevych/go-sqlxmock"
	"testing"
)

func TestBookPostgres_CreateBook(t *testing.T) {
	db, mock, err := sqlmock.Newx()
	if err != nil{
		logrus.Error(err)
	}
	defer db.Close()

	r := NewBookPostgres(db)

	type args struct{
		name string
		author string
	}
	type mockBehavior func(args args, id int)

	testTable := []struct{
		name 		 string
		mockBehavior mockBehavior
		args		 args
		id 			 int
		wantErr 	 bool
	} {
		{
			name: "OK",
			args: args{
				name:   "Test",
				author: "Test",
			},
			id: 2,
			mockBehavior: func(args args, id int) {
				rows := sqlmock.NewRows([]string{"id"}).AddRow(id)
				mock.ExpectQuery("INSERT INTO books").WithArgs(args.name, args.author).WillReturnRows(rows)
			},
		},
	}

	for _, testCase := range testTable{
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mockBehavior(testCase.args, testCase.id)

			book := model.Book{
				Id:     testCase.id,
				Name:   testCase.args.name,
				Author: testCase.args.author,
			}

			got, err := r.CreateBook(book)
			if testCase.wantErr{
				assert.Error(t, err)
			}else{
				assert.NoError(t, err)
				assert.Equal(t, testCase.id, got)
			}
		})
	}
}
