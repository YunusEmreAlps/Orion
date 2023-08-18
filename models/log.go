package models

// Log is a struct that represents the kodeks_log_info table
type Log struct {
	ID             int    `json:"id" gorm:"id"`
	PageID         int    `json:"page_id" gorm:"page_id"`
	UserID         int    `json:"user_id" gorm:"user_id"`
	UserIP         string `json:"user_ip" gorm:"user_ip"`
	UrlInfo        string `json:"url_info" gorm:"url_info"`
	DateTime       string `json:"date_time" gorm:"date_time"`
	ActionInfo     string `json:"action_info" gorm:"action_info"`
	MethodName     string `json:"method_name" gorm:"method_name"`
	ActionDetail   string `json:"action_detail" gorm:"action_detail"`
	ErrorMessage   string `json:"error_message" gorm:"error_message"`
	UserProviderID string `json:"user_provider_id" gorm:"user_provider_id"`
	ServerHostName string `json:"server_host_name" gorm:"server_host_name"`
	ParameterList  string `json:"parameter_list" gorm:"parameter_list"`
	// username not getting from db, it's getting from logs.parameter_list string and add it to logs array
	Username string `json:"username" gorm:"-"`
}

// Table Name
func (Log) TableName() string {
	return "kodeks_log_info"
}
