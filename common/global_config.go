package common

type GlobalConfig struct {
	Debug   bool           `json:"debug"`
	Http    *HttpConfig    `json:"http"`
	MailGun *MailGunConfig `json:"mailgun"`
}
