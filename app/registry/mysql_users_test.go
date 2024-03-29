package registry

import (
	"reflect"
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/k-jun/waxflower/model"
	"github.com/k-jun/waxflower/util"
)

func TestMySQL_SelectUser(t *testing.T) {
	t.Cleanup(func() {
		util.DeleteAll(db)
	})
	type fields struct {
		db   *sqlx.DB
		user *model.User
	}
	type args struct {
		u *model.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.User
		wantErr bool
	}{
		{
			name: "OK: SelectUser",
			fields: fields{
				db: db,
				user: &model.User{
					Id:    "b6c8955f-1514-499f-af7a-5b934f3676ed",
					Email: "dicksondeanna@example.com",
				},
			},
			args: args{u: &model.User{
				Id: "b6c8955f-1514-499f-af7a-5b934f3676ed",
			}},
			want: &model.User{
				Id:    "b6c8955f-1514-499f-af7a-5b934f3676ed",
				Email: "dicksondeanna@example.com",
			},
		},
		{
			name: "NG: SelectUser",
			fields: fields{
				db: db,
				user: &model.User{
					Id:    "85fca55e-c8f1-4cfc-a99a-d9abf0812126",
					Email: "david04@example.net",
				},
			},
			args: args{u: &model.User{
				Id: "1bafb830-40b4-49cf-93ad-70bf7c2f4681",
			}},
			want: &model.User{
				Id: "1bafb830-40b4-49cf-93ad-70bf7c2f4681",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sql := &MySQL{
				db: tt.fields.db,
			}
			_ = util.InsertUser(t, tt.fields.db, tt.fields.user)
			got, err := sql.SelectUser(tt.args.u)
			if (err != nil) != tt.wantErr {
				t.Errorf("MySQL.SelectUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MySQL.SelectUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMySQL_InsertUser(t *testing.T) {
	t.Cleanup(func() {
		util.DeleteAll(db)
	})
	type fields struct {
		db *sqlx.DB
	}
	type args struct {
		u *model.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.User
		wantErr bool
	}{
		{
			name:   "OK: InsertUser",
			fields: fields{db: db},
			args: args{
				u: &model.User{
					Id:    "683ec742-2c3a-462e-bd1a-9e43878b294a",
					Email: "zacharycarter@example.com",
				},
			},
			want: &model.User{
				Id:    "683ec742-2c3a-462e-bd1a-9e43878b294a",
				Email: "zacharycarter@example.com",
			},
		},
		{
			name:   "NG: InsertUser",
			fields: fields{db: db},
			args: args{
				u: &model.User{
					Id:    "683ec742-2c3a-462e-bd1a-9e43878b294a",
					Email: "zacharycarter@example.com",
				},
			},
			want: &model.User{
				Id:    "683ec742-2c3a-462e-bd1a-9e43878b294a",
				Email: "zacharycarter@example.com",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sql := &MySQL{
				db: tt.fields.db,
			}
			_, err := sql.InsertUser(tt.args.u)
			if (err != nil) != tt.wantErr {
				t.Errorf("MySQL.InsertUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			got := util.SelectUser(t, db, tt.args.u)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MySQL.InsertUser() = %v, want %v", got, tt.want)
			}
		})
	}
}
