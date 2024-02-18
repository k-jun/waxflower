package registry

import (
	"os"
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
		g  *model.Game
		s  *model.Seat
		u  *model.User
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
					Price:  3000,
					Id:     "b2ee77b5-6cbf-4079-b53d-274c5cbb8b61",
					UserId: "42e00bac-85d4-495c-98dd-b2424a6c3ecb",
					GameId: "598654a5-d503-44cc-a8b8-69526dd4f3b5",
					SeatId: "11e1ae6b-0d6e-4c80-9b82-3c2cbd3aa6a5",
				},
				u: &model.User{
					Id:    "42e00bac-85d4-495c-98dd-b2424a6c3ecb",
					Email: "larry98@example.com",
				},
				g: &model.Game{
					Id:   "598654a5-d503-44cc-a8b8-69526dd4f3b5",
					Date: "1971-06-07",
				},
				s: &model.Seat{
					Id: "11e1ae6b-0d6e-4c80-9b82-3c2cbd3aa6a5",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sql := &MySQL{
				db: tt.fields.db,
			}

			util.InsertUser(t, db, tt.fields.u)
			util.InsertGame(t, db, tt.fields.g)
			util.InsertSeat(t, db, tt.fields.s)
			util.InsertTicket(t, db, tt.fields.t)

			if err := sql.Reset(); (err != nil) != tt.wantErr {
				t.Errorf("MySQL.Reset() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
