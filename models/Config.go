package models

type Config struct {
	PublicAddressName       string `json:"PublicAddressName"`
	Email                   string `json:"Email"`
	Password                string `json:"Password"`
	Port                    int    `json:"Port"`
	CallbackURL             string `json:"CallbackURL"`
	TokenStorageEndpoint    string `json:"TokenStorageEndpoint"`
	MessageCallbackEndpoint string `json:"MessageCallbackEndpoint"`
	KahlaServer 			string `json:"KahlaServer"`
}
