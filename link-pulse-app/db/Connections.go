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
	if err != nil {
		log.Fatalf("failed to load AWS config: %v", err)
	}

	svc := secretsmanager.NewFromConfig(cfg)

	input := &secretsmanager.GetSecretValueInput{
		SecretId:     aws.String(secretName),
		VersionStage: aws.String("AWSCURRENT"),
	}

	result, err := svc.GetSecretValue(context.TODO(), input)
	if err != nil {
		log.Fatalf("failed to get secret value: %v", err)
	}

	var creds RDSCredentials
	if err := json.Unmarshal([]byte(*result.SecretString), &creds); err != nil {
		log.Fatalf("failed to parse secret JSON: %v", err)
	}

	// PostgreSQL DSN
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		creds.Host, creds.Port, creds.Username, creds.Password, creds.DBName)

	// log.Println("The dsn was below ------------------------")
	// log.Println(dsn)
	// log.Println("The dsn was above -------------------------")
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("failed to open DB: %v", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatalf("failed to ping DB: %v", err)
	}

	log.Println("🔗 Connected to RDS successfully!")
	return db
}
