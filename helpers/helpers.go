package helpers

import (
	"strconv"
	"strings"
	"time"

	"orion/logger"
	"orion/models"

	"github.com/xuri/excelize/v2"
)

func TimeFormatter(t time.Time) string {
	// That formats the given time to RFC3339 format (2023-01-16T13:50:56.910Z)
	return t.UTC().Format(time.RFC3339)
}

func ArrayToString(array []string) string {
	// That converts the given array to string
	var str string
	for _, v := range array {
		str += v + ","
	}
	return str
}

// Find key value in string json
func FindKeyValueInJson(json string, key string) string {
	// That finds the given key value in string json
	parameterList := strings.Split(json, ",")
	for _, v := range parameterList {
		// if value contains "username" string, split it by equal sign and get the second value
		if strings.Contains(v, "username") {
			username := strings.Split(v, ":")
			return username[1]
		}
	}
	return ""
}

// Filter to Logs by ignored error logs like ("AuthLoginFailed 1002: Geçersiz e-posta / kullanıcı adı veya şifre.")
func FilterChanges(changes []models.Log, ignored []string) []models.Log {
	var filteredChanges []models.Log

	for _, change := range changes {
		if len(ignored) > 0 {
			for _, v := range ignored {
				if change.ErrorMessage != v {
					filteredChanges = append(filteredChanges, change)
				}
			}
		} else {
			filteredChanges = append(filteredChanges, change)
		}
	}
	return filteredChanges
}

// Excel File Creation Function
func SetChangesToExcel(changes []models.Log) *excelize.File {
	// Create a new spreadsheet
	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			logger.ERROR.Println("ERROR: ", err)
		}
	}()

	// Change the name of the worksheet.
	f.SetSheetName("Sheet1", "Logs")

	f.SetCellValue("Logs", "A1", "ID")
	f.SetCellValue("Logs", "B1", "Page ID")
	f.SetCellValue("Logs", "C1", "User ID")
	f.SetCellValue("Logs", "D1", "User IP")
	f.SetCellValue("Logs", "E1", "URL Info")
	f.SetCellValue("Logs", "F1", "Date Time")
	f.SetCellValue("Logs", "G1", "Action Info")
	f.SetCellValue("Logs", "H1", "Method Name")
	f.SetCellValue("Logs", "I1", "Action Detail")
	f.SetCellValue("Logs", "J1", "Error Message")
	f.SetCellValue("Logs", "K1", "User Provider ID")
	f.SetCellValue("Logs", "L1", "Server Host Name")
	f.SetCellValue("Logs", "M1", "Parameter List")

	// Set value of a cell.
	index := 2
	for _, change := range changes {
		f.SetCellValue("Logs", "A"+strconv.Itoa(index), change.ID)
		f.SetCellValue("Logs", "B"+strconv.Itoa(index), change.PageID)
		f.SetCellValue("Logs", "C"+strconv.Itoa(index), change.UserID)
		f.SetCellValue("Logs", "D"+strconv.Itoa(index), change.UserIP)
		f.SetCellValue("Logs", "E"+strconv.Itoa(index), change.UrlInfo)
		f.SetCellValue("Logs", "F"+strconv.Itoa(index), change.DateTime)
		f.SetCellValue("Logs", "G"+strconv.Itoa(index), change.ActionInfo)
		f.SetCellValue("Logs", "H"+strconv.Itoa(index), change.MethodName)
		f.SetCellValue("Logs", "I"+strconv.Itoa(index), change.ActionDetail)
		f.SetCellValue("Logs", "J"+strconv.Itoa(index), change.ErrorMessage)
		f.SetCellValue("Logs", "K"+strconv.Itoa(index), change.UserProviderID)
		f.SetCellValue("Logs", "L"+strconv.Itoa(index), change.ServerHostName)
		f.SetCellValue("Logs", "M"+strconv.Itoa(index), change.ParameterList)
		index++
	}

	// Set active sheet of the workbook.
	f.SetActiveSheet(index)

	// Save spreadsheet to the db
	if err := f.SaveAs("Logs.xlsx"); err != nil {
		logger.ERROR.Println("ERROR: ", err)
		return nil
	}

	// return file for attachment
	return f
}

func UrlToOptions(url string) (string, string, string, string, string, string) {

	options := strings.Split(url, "://")

	// split protocol and info
	protocol := options[0]
	info := options[1]

	// split info to username, password, host, port, db
	infoOptions := strings.Split(info, "@")

	// split username and password
	usernamePassword := infoOptions[0]
	hostPortDb := infoOptions[1]

	usernamePasswordOptions := strings.Split(usernamePassword, ":")
	username := usernamePasswordOptions[0]
	password := usernamePasswordOptions[1]

	// split host, port and db
	hostPortDbOptions := strings.Split(hostPortDb, "/")
	hostPort := hostPortDbOptions[0]
	db := hostPortDbOptions[1]

	// split host and port
	hostPortOptions := strings.Split(hostPort, ":")
	host := hostPortOptions[0]
	port := hostPortOptions[1]

	return protocol, username, password, host, port, db
}
