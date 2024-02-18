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
		db    *sqlx.DB
		t     *model.Ticket
		email string
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
					Id:     "2b98ca6e-f8ca-4b14-b9c7-ed4286b92cfd",
					Price:  3000,
					UserId: "4860e3a8-c77d-4e74-84bb-cb4522de8f31",
					GameId: "c20e2ee6-948e-4ea4-a9af-b875d68b1f7c",
					SeatId: "07b63277-ad69-484f-8618-ebfaa8893163",
				},
				email: "christopherdavis@example.net",
			},
			args: args{t: &model.Ticket{
				Id: "2b98ca6e-f8ca-4b14-b9c7-ed4286b92cfd",
			}},
			want: &model.Ticket{
				Id:     "2b98ca6e-f8ca-4b14-b9c7-ed4286b92cfd",
				Price:  3000,
				UserId: "4860e3a8-c77d-4e74-84bb-cb4522de8f31",
				GameId: "c20e2ee6-948e-4ea4-a9af-b875d68b1f7c",
				SeatId: "07b63277-ad69-484f-8618-ebfaa8893163",
			},
		},
		{
			name: "NG: SelectTicket",
			fields: fields{
				db: db,
				t: &model.Ticket{
					Id:     "570cba82-295f-4c1e-b9af-a545377991ef",
					Price:  5000,
					UserId: "be706b7e-17a1-4275-ae41-f831e0e47d7b",
					GameId: "9c825a30-f598-499b-a2a8-6a8eec5042d4",
					SeatId: "454214e9-2d02-4581-8590-b2869d418233",
				},
				email: "teresasmith@example.net",
			},
			args: args{t: &model.Ticket{
				Id: "9881c659-4095-4a32-b888-288105ed92cf",
			}},
			want: &model.Ticket{
				Id: "9881c659-4095-4a32-b888-288105ed92cf",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sql := &MySQL{
				db: tt.fields.db,
			}
			_ = util.InsertUser(t, db, &model.User{Id: tt.fields.t.UserId, Email: tt.fields.email})
			_ = util.InsertGame(t, db, &model.Game{Id: tt.fields.t.GameId, Date: "1999-05-18"})
			_ = util.InsertSeat(t, db, &model.Seat{Id: tt.fields.t.SeatId})
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
		db    *sqlx.DB
		t     *model.Ticket
		email string
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
					UserId: "bb7e67f7-acda-447a-92c8-8f347c2d4020",
					GameId: "25c955c9-9465-4083-bbc9-bdba4cda62a9",
					SeatId: "bb3fa16b-237f-4c07-9b28-8548e46bf536",
				},
				email: "michelle48@example.com",
			},
			args: args{t: &model.Ticket{
				Id:     "14a15b3e-1937-4d7d-a7fc-b3e1948dfc76",
				Price:  5000,
				UserId: "bb7e67f7-acda-447a-92c8-8f347c2d4020",
				GameId: "25c955c9-9465-4083-bbc9-bdba4cda62a9",
				SeatId: "bb3fa16b-237f-4c07-9b28-8548e46bf536",
			}},
			want: &model.Ticket{
				Id:     "14a15b3e-1937-4d7d-a7fc-b3e1948dfc76",
				Price:  5000,
				UserId: "bb7e67f7-acda-447a-92c8-8f347c2d4020",
				GameId: "25c955c9-9465-4083-bbc9-bdba4cda62a9",
				SeatId: "bb3fa16b-237f-4c07-9b28-8548e46bf536",
			},
		},
		{
			name: "NG: InsertTicket",
			fields: fields{
				db: db,
				t: &model.Ticket{
					UserId: "3cb4baab-5d7f-42be-a3e2-98eeb9840575",
					GameId: "c7d46e1e-ee93-4518-a359-db8f2f131276",
					SeatId: "e28cbd18-95da-47da-939c-c96c0651308c",
				},
				email: "kmartinez@example.com",
			},
			args: args{t: &model.Ticket{
				Id:     "14a15b3e-1937-4d7d-a7fc-b3e1948dfc76",
				Price:  5000,
				UserId: "bb7e67f7-acda-447a-92c8-8f347c2d4020",
				GameId: "25c955c9-9465-4083-bbc9-bdba4cda62a9",
				SeatId: "bb3fa16b-237f-4c07-9b28-8548e46bf536",
			}},
			want: &model.Ticket{
				Id:     "14a15b3e-1937-4d7d-a7fc-b3e1948dfc76",
				Price:  5000,
				UserId: "bb7e67f7-acda-447a-92c8-8f347c2d4020",
				GameId: "25c955c9-9465-4083-bbc9-bdba4cda62a9",
				SeatId: "bb3fa16b-237f-4c07-9b28-8548e46bf536",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sql := &MySQL{
				db: tt.fields.db,
			}
			_ = util.InsertUser(t, db, &model.User{Id: tt.fields.t.UserId, Email: tt.fields.email})
			_ = util.InsertGame(t, db, &model.Game{Id: tt.fields.t.GameId, Date: "1999-05-18"})
			_ = util.InsertSeat(t, db, &model.Seat{Id: tt.fields.t.SeatId})
			got, err := sql.InsertTicket(tt.args.t)
			if (err != nil) != tt.wantErr {
				t.Errorf("MySQL.InsertTicket() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MySQL.InsertTicket() = %v, want %v", got, tt.want)
			}
		})
	}
}
