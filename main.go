package main

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
	lambda "github.com/aws/aws-lambda-go/lambda"
	"github.com/josetitic/gambituser/awsgo"
	"github.com/josetitic/gambituser/db"
	"github.com/josetitic/gambituser/models")

func main() {
	lambda.Start(LambdaExecution)
}

func LambdaExecution(ctx context.Context, event events.CognitoEventUserPoolsPostConfirmation) (events.CognitoEventUserPoolsPostConfirmation, error) {
	fmt.Println("> Antes de StartAWS'")
	awsgo.StartAWS()

	fmt.Println("> Despues de StartAWS'")

	if !ValidParameters() {
		fmt.Println("Error en los parámetros en 'SecretName'")
		err := errors.New("Error en los parámetros, debe enviar SecretName")
		return event, err
	}

	fmt.Printf("entrando a models.SignUp")

	var dats models.SignUp

	err := db.ReadSecret()

	fmt.Printf("resultado de ReadScret: ", err)

	fmt.Println("Email = " + dats.UserUUID)

	if err != nil {
		fmt.Printf("Error al leer el secret", err)
	}

	for row, att := range event.Request.UserAttributes {
		switch row {
		case "email":
			dats.UserEmail = att
			fmt.Println("Email = " + dats.UserEmail)
		case "sub":
			dats.UserUUID = att
			fmt.Println("Email = " + dats.UserUUID)
		}
	}


	err = db.SignUp(dats)

	return event, err
}

func ValidParameters() bool {
	var getParameter bool
	_, getParameter = os.LookupEnv("SecretName")
	fmt.Printf("Parameters: ", getParameter)
	return getParameter
}
