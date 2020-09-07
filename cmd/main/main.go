package main

import (
	"github.com/czelabueno/go-microservice-sample/internal/database"
	"github.com/czelabueno/go-microservice-sample/internal/logs"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate"
	migration "github.com/golang-migrate/migrate/database/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
)

const (
	migrationRootFolder     = "file://migrations"
	migrationScriptsVersion = 1
)

func main() {
	_ = logs.InitLogger()
	mySQLClient := database.NewMySQLClient("root:root@tcp(localhost:3306)/phones_review")
	doMigration(mySQLClient, "phones_review")
}

func doMigration(client *database.MySqlClient, dbname string) {
	driver, _ := migration.WithInstance(client.DB, &migration.Config{})
	m, err := migrate.NewWithDatabaseInstance(
		migrationRootFolder,
		dbName,
		driver,
	)

	if err != nil {
		logs.Log().Error(err.Error())
		return
	}

	current, _, _ := m.Version()
	logs.Log().Infof("Current migrations version in %d", current)
	err = m.Migrate(migrationScriptsVersion)

	if err != nil && err.Error() == "no change" {
		logs.Log().Info("No migration needed")
	}
}
