package migrations

import . "github.com/grafana/grafana/pkg/services/sqlstore/migrator"

func addFCMTokenMigration(mg *Migrator) {

	// create a new table without channel_id column
	fcmToken := Table{
		Name: "fcm_token",
		Columns: []*Column{
			{Name: "id", Type: DB_BigInt, IsPrimaryKey: true, IsAutoIncrement: true},
			{Name: "token", Type: DB_NVarchar, Length: 300, Nullable: false},
			{Name: "user_id", Type: DB_NVarchar, Length: 190, Nullable: false},
			{Name: "created", Type: DB_DateTime, Nullable: false},
			{Name: "updated", Type: DB_DateTime, Nullable: false},
		},
		Indices: []*Index{
			{Cols: []string{"token"}, Type: UniqueIndex},
		},
	}

	mg.AddMigration("create fcm_token_v2 table", NewAddTableMigration(fcmToken))
	//-------  indexes ------------------
	mg.AddMigration("add index fcm_token.token", NewAddIndexMigration(fcmToken, fcmToken.Indices[0]))

}
