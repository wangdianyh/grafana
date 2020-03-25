package migrations

import . "github.com/grafana/grafana/pkg/services/sqlstore/migrator"

func addFCMTokenMigrations(mg *Migrator) {
	fcmToken := Table{
		Name: "fcm_token",
		Columns: []*Column{
			{Name: "id", Type: DB_BigInt, IsPrimaryKey: true, IsAutoIncrement: true},
			{Name: "token", Type: DB_NVarchar, Length: 300, Nullable: false},
			{Name: "channel_id", Type: DB_NVarchar, Length: 190, Nullable: false},
			{Name: "user_id", Type: DB_NVarchar, Length: 190, Nullable: true},
			{Name: "created", Type: DB_DateTime, Nullable: false},
			{Name: "updated", Type: DB_DateTime, Nullable: false},
		},
	}

	// create table
	mg.AddMigration("create fcm_token table", NewAddTableMigration(fcmToken))

}
