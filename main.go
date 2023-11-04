package main

import (
	"gambituser/db"
	"gambituser/models"
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
	lambda "github.com/aws/aws-lambda-go/lambda"
	"github.com/josetitic/gambituser/tree/main/awsgo"
	"github.com/josetitic/gambituser/tree/main/db"
	"github.com/josetitic/gambituser/tree/main/models"

)

func main() {
	lambda.Start(LambdaExecution)
}

func LambdaExecution(ctx context.Context, event events.CognitoEventUserPoolsPostConfirmation) (events.CognitoEventUserPoolsPostConfirmation, error) {
	awsgo.StartAWS()
	if !ValidParameters() {
		fmt.Println("Error en los parámetros en 'SecretName'")
		err := errors.New("Error en los parámetros, debe enviar SecretName")
		return event, err
	}

	var dats models.SignUp

	for row,att := range event.Request.UserAttributes {
		switch row {
		case "email":
			dats.UserEmail = att
			fmt.Println("Email = "+dats.UserEmail)
		case "sub":
			dats.UserUUID = att
			fmt.Println("Email = "+dats.UserUUID)
		}
	}

	err := db.ReadSecret()
	if err != nil {
		fmt.Printf("Error al leer el secret", err)
	}

	err = db.SignUp(dats)

	return event,err
}

func ValidParameters() bool {
	var getParameter bool
	_, getParameter = os.LookupEnv("SecretName")
	return getParameter
}
