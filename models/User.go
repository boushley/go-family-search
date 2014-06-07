package models

type User struct {
	AlternateEmail    string `json:"alternateEmail"`
	BirthDate         string `json:"birthDate"`
	ContactName       string `json:"contactName"`
	Country           string `json:"country"`
	DisplayName       string `json:"displayName"`
	Email             string `json:"email"`
	FamilyName        string `json:"familyName"`
	FullName          string `json:"fullName"`
	Gender            string `json:"gender"`
	GivenName         string `json:"givenName"`
	HelperAccessPin   string `json:"helperAccessPin"`
	LdsMemberAccount  bool   `json:"ldsMemberAccount"`
	MailingAddress    string `json:"mailingAddress"`
	PersonId          string `json:"personId"`
	PhoneNumber       string `json:"phoneNumber"`
	PreferredLanguage string `json:"preferredLanguage"`
	TreeUserId        string `json:"treeUserId"`
	Links             []Link `json:"links"`
	Id                ID     `json:"id"`
}
