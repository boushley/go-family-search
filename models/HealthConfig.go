package models

type HealthConfig struct {
	BuildDate       string `json:"buildDate"`
	BuildVersion    string `json:"buildVersion"`
	DatabaseVersion string `json:"databaseVersion"`
	PlatformVersion string `json:"platformVersion"`
}
