package mysql

import (
	"AllenSo/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

func TestInsertPostList(t *testing.T) {
	dsn := "root:9012@tcp(localhost:3306)/AllenSo?charset=utf8mb4&parseTime=True"
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to connect database: %v", err)
	}

	type args struct {
		p []*model.Post
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "test", args: args{p: []*model.Post{
			{Title: "Post 1", Content: "Content 1"},
			{Title: "Post 2", Content: "Content 2"},
		},
		}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := InsertPostList(tt.args.p); (err != nil) != tt.wantErr {
				t.Errorf("InsertPostList() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
