package simple

type Database struct {
	Name string
}

type DatabaseMysql Database
type DatabaseMongo Database

func NewDatabaseMysql() DatabaseMysql {
	return (DatabaseMysql)(Database{Name: "Mysql"})
}

func NewDatabaseMongo() DatabaseMongo {
	return (DatabaseMongo)(Database{Name: "Mongo"})
}

type DatabaseRepository struct {
	DatabaseMysql *DatabaseMysql
	DatabaseMongo *DatabaseMongo
}

func NewDatabaseRepository(databaseMysql DatabaseMysql, databaseMongo DatabaseMongo) *DatabaseRepository {
	return &DatabaseRepository{
		DatabaseMysql: &databaseMysql,
		DatabaseMongo: &databaseMongo,
	}
}