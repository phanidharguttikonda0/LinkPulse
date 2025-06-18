package db

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"

	"database/sql"
	_ "github.com/lib/pq"
)

type RDSCredentials struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	DBName   string `json:"dbname"`
}

func RdbsConnection() *sql.DB {
	secretName := "rds/psql/main"
	region := "ap-south-1"

	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(region))
	/*
		the access id and secret keys that we pass while running the code will be loaded
		here and , returns an configuration object, such that it will be useful to access
		the aws services. so below we are accessing the aws secrets manager , so we can use
		this cfg to sign the request and the signed headers are send to the aws. this is
		how it works under the hood.
	*/
	if err != nil {
		log.Fatalf("failed to load AWS config: %v", err)
	}

	svc := secretsmanager.NewFromConfig(cfg)
	// here we have passed the cfg which hold the keys and sending a request to access the
	// aws secrets manager, if details were correct it returns the service to get what we want

	input := &secretsmanager.GetSecretValueInput{
		SecretId:     aws.String(secretName),
		VersionStage: aws.String("AWSCURRENT"),
	}
	// here we are mentioning the secret id to get the specific details

	result, err := svc.GetSecretValue(context.TODO(), input)
	if err != nil {
		log.Fatalf("failed to get secret value: %v", err)
	}

	var creds RDSCredentials
	if err := json.Unmarshal([]byte(*result.SecretString), &creds); err != nil {
		log.Fatalf("failed to parse secret JSON: %v", err)
	} // deserializing the data

	// PostgreSQL DSN
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=require",
		creds.Host, creds.Port, creds.Username, creds.Password, "linkpulse")

	log.Println("The dsn was below ------------------------")
	log.Println(dsn)
	log.Println("The dsn was above -------------------------")
	db, err := sql.Open("postgres", dsn)
	// here a connection was created not yet established
	if err != nil {
		log.Fatalf("failed to open DB: %v", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatalf("failed to ping DB: %v", err)
	}

	log.Println("🔗 Connected to RDS successfully!")
	return db
}

/*

🔐 What is AWS_ACCESS_KEY_ID and AWS_SECRET_ACCESS_KEY?
These are your AWS credentials, just like a username and password, that allow you (or your program) to authenticate with AWS services.

➤ AWS_ACCESS_KEY_ID
Like a username.

Public identifier for your AWS account or IAM user.

➤ AWS_SECRET_ACCESS_KEY
Like a password.

Used to sign requests — this is confidential and should never be shared.

*/
