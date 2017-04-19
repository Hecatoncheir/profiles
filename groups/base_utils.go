package groups

import (
	"log"

	rethink "gopkg.in/gorethink/gorethink.v3"
)

func checkDataBase(databaseNameForCheck string, session rethink.QueryExecutor) error {
	dbList := make([]string, 128)

	dbListCursor, err := rethink.DBList().Run(session)
	defer dbListCursor.Close()

	err = dbListCursor.All(&dbList)
	if err != nil {
		return err
	}

	dataBaseNotExist := true
	for _, dbName := range dbList {
		if dbName == databaseNameForCheck {
			dataBaseNotExist = false
		}
	}

	if dataBaseNotExist == true {
		_, err = rethink.DBCreate(databaseNameForCheck).Run(session)
		if err != nil {
			return err
		}
	}

	return nil
}

func checkTable(tableNameForCheck string, session rethink.QueryExecutor) error {
	var err error

	tableListCursor, err := rethink.TableList().Run(session)
	if err != nil {
		log.Fatalln(err.Error())
	}

	defer tableListCursor.Close()

	tableList := make([]string, 128)
	err = tableListCursor.All(&tableList)

	userTableNotExist := true
	for _, tableName := range tableList {
		if tableName == tableNameForCheck {
			userTableNotExist = false
		}
	}

	if userTableNotExist == true {
		_, err = rethink.TableCreate(tableNameForCheck, rethink.TableCreateOpts{PrimaryKey: "ID"}).Run(session)
		if err != nil {
			return err
		}

	}

	return nil
}
