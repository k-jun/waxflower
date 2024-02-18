package registry

import (
	"reflect"
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/k-jun/waxflower/model"
	"github.com/k-jun/waxflower/util"
)

func TestMySQL_SelectSeat(t *testing.T) {
	t.Cleanup(func() {
		util.DeleteAll(db)
	})
	type fields struct {
		db *sqlx.DB
		s  *model.Seat
	}
	type args struct {
		s *model.Seat
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.Seat
		wantErr bool
	}{
		{
			name: "OK: SelectSeat",
			fields: fields{
				db: db,
				s: &model.Seat{
					Id:  "92d0f7d5-2a0b-4c9b-87b6-8abd9f189a80",
					Col: 19,
					Row: "AA",
					Sec: 9,
				},
			},
			args: args{s: &model.Seat{
				Id: "92d0f7d5-2a0b-4c9b-87b6-8abd9f189a80",
			}},
			want: &model.Seat{
				Id:  "92d0f7d5-2a0b-4c9b-87b6-8abd9f189a80",
				Col: 19,
				Row: "AA",
				Sec: 9,
			},
		},
		{
			name: "NG: SelectSeat",
			fields: fields{
				db: db,
				s: &model.Seat{
					Id:  "d355343f-44e1-4927-a4ea-dae1026f3ab5",
					Col: 19,
					Row: "AA",
					Sec: 9,
				},
			},
			args: args{s: &model.Seat{
				Id: "4e56bf86-db2a-48c8-997b-4af823300b42",
			}},
			want: &model.Seat{
				Id: "4e56bf86-db2a-48c8-997b-4af823300b42",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sql := &MySQL{
				db: tt.fields.db,
			}
			_ = util.InsertSeat(t, db, tt.fields.s)

			got, err := sql.SelectSeat(tt.args.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("MySQL.SelectSeat() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MySQL.SelectSeat() = %v, want %v", got, tt.want)
			}
		})
	}
}
