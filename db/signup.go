package db

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/josetitic/gambituser/models"
	"github.com/josetitic/gambituser/tools")

func SignUp(sig models.SignUp) error {
	fmt.Println("Comienza Registro")

	err := DbConnect()
	if err != nil {
		return err
	}

	defer Db.Close()

	sentence := "INSERT INTO users (User_Email,User_UUID,User_DateAdd) VALUES ('" + sig.UserEmail + "','" + sig.UserUUID + "','" + tools.MysqlDate() + "')"

	fmt.Println(sentence)

	_, err = Db.Exec(sentence)

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("Ejecución Exitosa")
	return nil
}
