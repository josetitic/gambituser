package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/josetitic/gambituser/models"
	"github.com/josetitic/gambituser/secretm")

var ModelSecret models.SecretRDSJson
var err error
var Db *sql.DB

func ReadSecret() error {
	fmt.Println("Va a leer el secreto ")
	ModelSecret, err = secretm.GetSecret(os.Getenv("nameSecret"))
	return err
}

func DbConnect() error {
	Db, err = sql.Open("mysql", ConnStr(ModelSecret))
	fmt.Println("> DB: ")
	fmt.Println(Db)
	if err != nil {
		fmt.Println("Ocurrió un error al conectarse a la BD: "+err.Error())
		return err
	}

	err = Db.Ping()

	if err != nil {
		fmt.Println("Ocurrió un error al conectarse al hacer ping a la BD: " +err.Error())
		return err
	}

	fmt.Println("Conexión exitosa de la BD")

	return nil
}

func ConnStr(keys models.SecretRDSJson) string {
	var dbUser, authToken, dbEndpoint, dbName string
	dbUser = keys.Username
	authToken = keys.Password
	dbEndpoint = keys.Host
	dbName = "gambit"
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?allowCleartextPasswords=true", dbUser, authToken, dbEndpoint, dbName)
	fmt.Println("DSN: " +dsn)
	return dsn
}
