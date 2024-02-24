package registry

import (
	"reflect"
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/k-jun/waxflower/model"
	"github.com/k-jun/waxflower/util"
)

func TestMySQL_SelectTicket(t *testing.T) {
	t.Cleanup(func() {
		util.DeleteAll(db)
	})
	type fields struct {
		db *sqlx.DB
		t  *model.Ticket
	}
	type args struct {
		t *model.Ticket
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.Ticket
		wantErr bool
	}{
		{
			name: "OK: SelectTicket",
			fields: fields{
				db: db,
				t: &model.Ticket{
					Id:    "2b98ca6e-f8ca-4b14-b9c7-ed4286b92cfd",
					Price: 3000,
					Game: &model.Game{
						Id:   "c20e2ee6-948e-4ea4-a9af-b875d68b1f7c",
						Date: "2009-08-19",
						Time: "14:26:40",
					},
					Seat: &model.Seat{
						Id:  "07b63277-ad69-484f-8618-ebfaa8893163",
						Sec: 1,
						Row: "AA",
						Col: 19,
					},
				},
			},
			args: args{t: &model.Ticket{
				Id: "2b98ca6e-f8ca-4b14-b9c7-ed4286b92cfd",
			}},
			want: &model.Ticket{
				Id:    "2b98ca6e-f8ca-4b14-b9c7-ed4286b92cfd",
				Price: 3000,
				Game: &model.Game{
					Id:   "c20e2ee6-948e-4ea4-a9af-b875d68b1f7c",
					Date: "2009-08-19",
					Time: "14:26:40",
				},
				Seat: &model.Seat{
					Id:  "07b63277-ad69-484f-8618-ebfaa8893163",
					Sec: 1,
					Row: "AA",
					Col: 19,
				},
			},
		},
		{
			name: "NG: SelectTicket",
			fields: fields{
				db: db,
				t: &model.Ticket{
					Id:    "57806b1b-dfc2-4dec-b864-2e888db4219e",
					Price: 5000,
					Game: &model.Game{
						Id:   "d5f46724-f189-4946-924d-19fcd51331b7",
						Date: "1994-02-18",
						Time: "04:19:41",
					},
					Seat: &model.Seat{
						Id:  "1ca028ef-3d22-4f2b-9178-24c881abcaae",
						Sec: 7,
						Row: "BB",
						Col: 29,
					},
				},
			},
			args: args{t: &model.Ticket{
				Id: "d281e8e5-b004-40cb-9a5a-b1c6f8aac422",
			}},
			want: &model.Ticket{
				Id:   "d281e8e5-b004-40cb-9a5a-b1c6f8aac422",
				Game: &model.Game{},
				Seat: &model.Seat{},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sql := &MySQL{
				db: tt.fields.db,
			}
			_ = util.InsertGame(t, db, tt.fields.t.Game)
			_ = util.InsertSeat(t, db, tt.fields.t.Seat)
			_ = util.InsertTicket(t, db, tt.fields.t)

			got, err := sql.SelectTicket(tt.args.t)
			if (err != nil) != tt.wantErr {
				t.Errorf("MySQL.SelectTicket() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MySQL.SelectTicket() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMySQL_InsertTicket(t *testing.T) {
	t.Cleanup(func() {
		util.DeleteAll(db)
	})
	type fields struct {
		db *sqlx.DB
		t  *model.Ticket
	}
	type args struct {
		t *model.Ticket
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.Ticket
		wantErr bool
	}{
		{
			name: "OK: InsertTicket",
			fields: fields{
				db: db,
				t: &model.Ticket{
					Game: &model.Game{
						Id:   "c6e2375e-e585-4771-913d-b5f5b4d870fe",
						Date: "2017-01-14",
						Time: "18:14:06",
					},
					Seat: &model.Seat{
						Id:  "06c83a6f-afe0-43df-ae6b-387aaa57d935",
						Sec: 3,
						Row: "AX",
						Col: 20,
					},
				},
			},
			args: args{t: &model.Ticket{
				Id:    "a92155fc-63d1-4e8b-b442-32435b2d3e56",
				Price: 9000,
				Game:  &model.Game{Id: "c6e2375e-e585-4771-913d-b5f5b4d870fe"},
				Seat:  &model.Seat{Id: "06c83a6f-afe0-43df-ae6b-387aaa57d935"},
			}},
			want: &model.Ticket{
				Id:    "a92155fc-63d1-4e8b-b442-32435b2d3e56",
				Price: 9000,
				Game: &model.Game{
					Id:   "c6e2375e-e585-4771-913d-b5f5b4d870fe",
					Date: "2017-01-14",
					Time: "18:14:06",
				},
				Seat: &model.Seat{
					Id:  "06c83a6f-afe0-43df-ae6b-387aaa57d935",
					Sec: 3,
					Row: "AX",
					Col: 20,
				},
			},
		},
		{
			name: "NG: InsertTicket",
			fields: fields{
				db: db,
				t: &model.Ticket{
					Game: &model.Game{
						Id:   "9ebc7d83-c4db-48a7-a93b-bb97e669f3ea",
						Date: "1987-07-25",
						Time: "14:44:06",
					},
					Seat: &model.Seat{
						Id:  "ef3bdc8c-b7d9-4cb9-a0d7-1701c717d3b3",
						Sec: 5,
						Row: "XY",
						Col: 89,
					},
				},
			},
			args: args{t: &model.Ticket{
				Id:    "96ea1bcd-f1ed-4f0b-aaee-f80893cb6f59",
				Price: 10000,
				Game:  &model.Game{Id: "5be44684-4f4c-4d24-bbf6-b60deab57d1f"},
				Seat:  &model.Seat{Id: "0901a863-8048-497e-ae7f-a34266585f4f"},
			}},
			want: &model.Ticket{
				Id:    "96ea1bcd-f1ed-4f0b-aaee-f80893cb6f59",
				Price: 10000,
				Game:  &model.Game{Id: "5be44684-4f4c-4d24-bbf6-b60deab57d1f"},
				Seat:  &model.Seat{Id: "0901a863-8048-497e-ae7f-a34266585f4f"},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sql := &MySQL{
				db: tt.fields.db,
			}
			_ = util.InsertGame(t, db, tt.fields.t.Game)
			_ = util.InsertSeat(t, db, tt.fields.t.Seat)

			_, err := sql.InsertTicket(tt.args.t)
			if (err != nil) != tt.wantErr {
				t.Errorf("MySQL.InsertTicket() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			got := util.SelectTicket(t, db, tt.args.t)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MySQL.InsertTicket() = %v, want %v", got, tt.want)
			}
		})
	}
}
