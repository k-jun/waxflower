package registry

import (
	"reflect"
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/k-jun/waxflower/model"
	"github.com/k-jun/waxflower/util"
)

func TestMySQL_SelectGame(t *testing.T) {
	type fields struct {
		db *sqlx.DB
		g  *model.Game
	}
	type args struct {
		g *model.Game
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.Game
		wantErr bool
	}{
		{
			name: "OK: SelectGame",
			fields: fields{
				db: db,
				g: &model.Game{
					Id:   "50948c39-28dc-4022-b0c7-2eff5fc175df",
					Date: "2017-08-25",
					Time: "19:20:00",
				},
			},
			args: args{
				g: &model.Game{
					Id: "50948c39-28dc-4022-b0c7-2eff5fc175df",
				},
			},
			want: &model.Game{
				Id:   "50948c39-28dc-4022-b0c7-2eff5fc175df",
				Date: "2017-08-25",
				Time: "19:20:00",
			},
		},
		{
			name: "NG: SelectGame",
			fields: fields{
				db: db,
				g: &model.Game{
					Id:   "b789710a-ec2d-4ad6-afab-c5a5403e26a1",
					Date: "2017-08-25",
					Time: "19:20:00",
				},
			},
			args: args{
				g: &model.Game{
					Id: "ae8c0396-c9c1-41e2-8ed8-ee30b26d842c",
				},
			},
			want: &model.Game{
				Id: "ae8c0396-c9c1-41e2-8ed8-ee30b26d842c",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sql := &MySQL{
				db: tt.fields.db,
			}

			_ = util.InsertGame(t, db, tt.fields.g)
			got, err := sql.SelectGame(tt.args.g)
			if (err != nil) != tt.wantErr {
				t.Errorf("MySQL.SelectGame() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MySQL.SelectGame() = %v, want %v", got, tt.want)
			}
		})
	}
}
