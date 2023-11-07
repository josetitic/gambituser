package secretm

import (
	"encoding/json"
	"fmt"


	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/josetitic/gambituser/awsgo"
	"github.com/josetitic/gambituser/models")

func GetSecret(nameSecret string) (models.SecretRDSJson, error) {
	var secretData models.SecretRDSJson
	fmt.Println(" > Pido Secreto " + nameSecret)

	svc := secretsmanager.NewFromConfig(awsgo.Cfg)

	key, err := svc.GetSecretValue(awsgo.Ctx, &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(nameSecret),
	})

	if err != nil {
		fmt.Println(err.Error())
		return secretData, err
	}

	json.Unmarshal([]byte(*key.SecretString), &secretData)

	fmt.Println(" > Lectura Secret OK " + nameSecret)

	return secretData, nil
}
