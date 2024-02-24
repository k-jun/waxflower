package registry

import (
	"os"
	"reflect"
	"testing"

	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/k-jun/waxflower/model"
	"github.com/k-jun/waxflower/util"
)

var (
	conf = &mysql.Config{
		Addr:   "127.0.0.1:3306",
		User:   "root",
		Passwd: "root",
		DBName: "mydb",
	}
	db = sqlx.MustConnect("mysql", conf.FormatDSN())
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestMySQL_Reset(t *testing.T) {
	t.Cleanup(func() {
		util.DeleteAll(db)
	})
	type fields struct {
		db *sqlx.DB
		t  *model.Ticket
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "OK: Reset",
			fields: fields{
				db: db,
				t: &model.Ticket{
					Id:    "2b98ca6e-f8ca-4b14-b9c7-ed4286b92cfd",
					Price: 3000,
					Game: &model.Game{
						Id:   "e3f6edb2-f010-46df-a032-396a6d6b18f1",
						Date: "1973-08-25",
						Time: "10:41:43",
					},
					Seat: &model.Seat{
						Id:  "8a352d17-69b6-4092-8130-53fd7e3791b1",
						Sec: 8,
						Row: "ZZ",
						Col: 99,
					},
					User: &model.User{
						Id:    "fcea7d4f-944b-4734-a502-533ee14c130a",
						Email: "zpatel@example.net",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sql := &MySQL{
				db: tt.fields.db,
			}

			util.InsertUser(t, db, tt.fields.t.User)
			util.InsertGame(t, db, tt.fields.t.Game)
			util.InsertSeat(t, db, tt.fields.t.Seat)
			util.InsertTicket(t, db, tt.fields.t)

			if err := sql.Reset(); (err != nil) != tt.wantErr {
				t.Errorf("MySQL.Reset() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMySQL_Search(t *testing.T) {
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
		want    []*model.Ticket
		wantErr bool
	}{
		{
			name: "OK: Search",
			fields: fields{
				db: db,
				t: &model.Ticket{
					Id:    "492ee014-5b84-4d3f-90c1-6a29cd7e0e20",
					Price: 2000,
					Game: &model.Game{
						Id:   "e6b72889-6e1a-4189-8a7a-2212bb386d50",
						Date: "1970-11-11",
						Time: "09:31:53",
					},
					Seat: &model.Seat{
						Id:  "b772c0b0-f2de-4601-932e-fa9a260f0c55",
						Sec: 1,
						Col: 10,
						Row: "AA",
					},
				},
			},
			args: args{
				t: &model.Ticket{
					Game: &model.Game{Date: "1970-11-11"},
					Seat: &model.Seat{Sec: 1},
				},
			},
			want: []*model.Ticket{
				{
					Id:    "492ee014-5b84-4d3f-90c1-6a29cd7e0e20",
					Price: 2000,
					Game: &model.Game{
						Id:   "e6b72889-6e1a-4189-8a7a-2212bb386d50",
						Date: "1970-11-11",
						Time: "09:31:53",
					},
					Seat: &model.Seat{
						Id:  "b772c0b0-f2de-4601-932e-fa9a260f0c55",
						Sec: 1,
						Col: 10,
						Row: "AA",
					},
				},
			},
		},
		{
			name: "NG: Search",
			fields: fields{
				db: db,
				t: &model.Ticket{
					Id:    "6bcc26a7-b323-4cfa-9f27-88c02e41064d",
					Price: 5000,
					Game: &model.Game{
						Id:   "4e523685-8bd1-464b-9863-0cf4b05f5aea",
						Date: "1980-04-13",
						Time: "06:23:50",
					},
					Seat: &model.Seat{
						Id:  "f3b16b3d-2abc-487c-a45a-ec1292d58102",
						Sec: 9,
						Col: 99,
						Row: "ZZ",
					},
				},
			},
			args: args{
				t: &model.Ticket{
					Game: &model.Game{Date: "1970-11-11"},
					Seat: &model.Seat{Sec: 2},
				},
			},
			want: []*model.Ticket{},
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

			got, err := sql.Search(tt.args.t)
			if (err != nil) != tt.wantErr {
				t.Errorf("MySQL.Search() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MySQL.Search() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMySQL_Reserve(t *testing.T) {
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
			name: "OK: Reserve",
			fields: fields{
				db: db,
				t: &model.Ticket{
					Id:    "c3352915-b1f5-4723-b2ab-3f30ec345d84",
					Price: 3000,
					Seat: &model.Seat{
						Id:  "f43af2dc-3af0-4967-a5fb-cc175b501f72",
						Row: "XX",
						Sec: 3,
						Col: 3,
					},
					Game: &model.Game{
						Id:   "de41e19c-6b9d-47b5-ab4a-77bc86fe497a",
						Date: "2000-09-04",
						Time: "03:34:12",
					},
					User: &model.User{
						Id:    "4bc099c8-5a89-4021-acdc-77f3206cc2f3",
						Email: "qjimenez@example.com",
					},
				},
			},
			args: args{
				t: &model.Ticket{
					Id:   "c3352915-b1f5-4723-b2ab-3f30ec345d84",
					User: &model.User{Id: "4bc099c8-5a89-4021-acdc-77f3206cc2f3"},
				},
			},
			want: &model.Ticket{
				Id:   "c3352915-b1f5-4723-b2ab-3f30ec345d84",
				User: &model.User{Id: "4bc099c8-5a89-4021-acdc-77f3206cc2f3"},
			},
		},
		{
			name: "NG: Reserve",
			fields: fields{
				db: db,
				t: &model.Ticket{
					Id:    "32c70594-67d0-4068-a8bf-7a440c9574e8",
					Price: 5000,
					Seat: &model.Seat{
						Id:  "7b4dbab3-9d56-4fc0-9aa0-2fecd5b40c80",
						Row: "XY",
						Sec: 2,
						Col: 69,
					},
					Game: &model.Game{
						Id:   "b5e30713-8e13-4199-a8ce-e1d2b98fa2f0",
						Date: "2020-08-31",
						Time: "16:32:05",
					},
					User: &model.User{
						Id:    "548cd8e1-7d60-4f73-be17-0b75f7bdc669",
						Email: "cooperandrea@example.net",
					},
				},
			},
			args: args{
				t: &model.Ticket{
					Id:   "32c70594-67d0-4068-a8bf-7a440c9574e8",
					User: &model.User{Id: "5119c911-90da-4cf8-9446-d229bdc66082"},
				},
			},
			want: &model.Ticket{
				Id:   "32c70594-67d0-4068-a8bf-7a440c9574e8",
				User: &model.User{Id: "5119c911-90da-4cf8-9446-d229bdc66082"},
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
			_ = util.InsertUser(t, db, tt.fields.t.User)
			_ = util.InsertTicket(t, db, tt.fields.t)

			got, err := sql.Reserve(tt.args.t)
			if (err != nil) != tt.wantErr {
				t.Errorf("MySQL.Reserve() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MySQL.Reserve() = %v, want %v", got, tt.want)
			}
		})
	}
}
