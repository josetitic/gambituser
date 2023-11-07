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

	awsgo.StartAWS()

	var dats models.SignUp

	err := db.ReadSecret()

	if err != nil {
		fmt.Printf("Error al leer el secret", err)
	}

	if !ValidParameters() {
		fmt.Println("Error en los parámetros en 'SecretName'")
		err := errors.New("Error en los parámetros, debe enviar SecretName")
		return event, err
	}

	fmt.Printf("entrando a models.SignUp")

	for row, att := range event.Request.UserAttributes {
		switch row {
		case "email":
			dats.UserEmail = att
		case "sub":
			dats.UserUUID = att
		}
	}


	err = db.SignUp(dats)

	return event, err
}

func ValidParameters() bool {
	var getParameter bool
	os.Environ()
	_, getParameter = os.LookupEnv("nameSecret")
	fmt.Printf("Parameters: ", getParameter)
	return getParameter
}
