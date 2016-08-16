package common

type MailGunConfig struct {
	Domain string `json:"domain"`
	Key    string `json:"key"`
	Pubkey string `json:"pubkey"`
	To     string `json:"to"`
}
