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
	fmt.Println("Valores del Secreto ", ModelSecret)
	return err
}

func DbConnect() error {
	Db, err = sql.Open("mysql", ConnStr(ModelSecret))
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	err = Db.Ping()
	if err != nil {
		fmt.Println(err.Error())
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
	fmt.Println(dsn)
	return dsn
}
