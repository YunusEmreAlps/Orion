package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"orion/config"
	"orion/helpers"
	"orion/logger"
	"orion/mail"
	"orion/models"

	"github.com/go-redis/redis/v8"
	_ "github.com/lib/pq"
	"github.com/roylee0704/gron"
	"github.com/xuri/excelize/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbConn *gorm.DB
var isConfigSuccess = false
var equals string = strings.Repeat("=", 50)

// Every 1 minutes
var repeatTime = 1 * time.Minute

// ignored error messages array
var ignoredErrorMessages = []string{
	// ignored error messages here
}

// To Users
var toUsers = []string{
	// team members here
}

// CC Users
var ccUsers = []string{
	// team members here
}

func main() {
	// Connect to the database
	dbConn = dbConnection()

	// Run the task every minute using the gron library
	c := gron.New()
	c.AddFunc(gron.Every(repeatTime), func() {
		// Query the TARGET table and retrieve changes
		changes, err := getTableChanges(dbConn)
		if err != nil {
			panic(err)
		}

		// Handle the changes
		fmt.Println(equals)
		if len(changes) > 0 {
			for _, change := range changes {
				logger.INFO.Println("INFO: ", strconv.Itoa(change.ID)+" - "+change.ErrorMessage)
			}

			// Filter the changes
			filteredChanges := helpers.FilterChanges(changes, ignoredErrorMessages)

			if len(filteredChanges) > 0 {
				for _, v := range filteredChanges {
					logger.INFO.Println("TRACE: ", strconv.Itoa(v.ID)+" - "+v.ErrorMessage)
				}
				f := helpers.SetChangesToExcel(filteredChanges)
				sendMailWithAttachment(filteredChanges, f)
			}
		} else {
			logger.INFO.Println("INFO: No changes in the last minute.")
		}
		fmt.Println(equals)
	})
	c.Start()

	// Infinite loop to keep the program running
	select {}
}

// Initialize Application
func init() {
	logger.InitLogger(os.Stdout, os.Stdout, os.Stdout, os.Stdout, os.Stderr, os.Stderr)
	isConfigSuccess = configureApplication()
	if !isConfigSuccess {
		logger.ERROR.Println("INIT: Application configuration failed. Please check your config file.")
		os.Exit(1)
	}
}

// Configure Application
func configureApplication() bool {
	// Clear the terminal screen
	fmt.Println(equals)
	dir, err := os.Getwd()
	if err != nil {
		logger.ERROR.Println("INIT: Cannot get current working directory os.Getwd()")
		return false
	} else {
		config.ReadConfig(dir)
		logger.INFO.Println("INIT: Application configuration file read success.")
		return true
	}
}

// DB Connection
func dbConnection() *gorm.DB {
	env := config.C.DB

	// String to Int
	port, err := strconv.Atoi(env.Port)
	if err != nil {
		logger.ERROR.Println("ERROR: ", err)
		os.Exit(1)
	}

	// Connect to the "postgres" database
	dbInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s", env.Host, port, env.Username, env.Password, env.DBName, env.SSLMode)
	db, err := gorm.Open(postgres.Open(dbInfo), &gorm.Config{})
	if err != nil {
		logger.ERROR.Println("ERROR: ", err)
		os.Exit(1)
	}

	// Connection Success
	logger.SUCCESS.Println("PostgreSQL Database Connection Success")
	return db
}

// Redis Connection
func redisConnection(redisUrl string) (*redis.Client, context.Context) {
	// redis://username:password@host:port/db
	_, username, password, host, port, db := helpers.UrlToOptions(redisUrl)

	// convert db to int
	dbInt, err := strconv.Atoi(db)
	if err != nil {
		logger.ERROR.Println("INIT: redis connection database is not integer ", err)
	}

	rContext := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr:     host + ":" + port,
		Username: username,
		Password: password,
		DB:       dbInt,
	})
	// ping redis for check connection
	_, err = rdb.Ping(rContext).Result()
	if err != nil {
		logger.ERROR.Println("INIT: redis ping request failed ", err)
	}

	return rdb, rContext
}

func getTableChanges(db *gorm.DB) ([]models.Log, error) {
	// Get the changes in the last minute and not null error_message from TARGET table
	var logs []models.Log

	// Last Minute Changes
	if err := db.Where("date_time >= ?", helpers.TimeFormatter(time.Now().Add(-repeatTime))).Where("error_message IS NOT NULL").Find(&logs).Error; !errors.Is(err, nil) {
		logger.ERROR.Println("ERROR: ", err)
		return nil, err
	}

	return logs, nil
}

// Send Mail with Excel File
func sendMailWithAttachment(logs []models.Log, f *excelize.File) {
	mailContent := &models.Mail{
		Sender:  config.C.Mail.FromMail,
		To:      toUsers,
		Cc:      ccUsers,
		Bcc:     []string{},
		Subject: config.C.App.TargetApp + " Error Logs",
	}

	mail.SendMail(mailContent, logs, f)
}
