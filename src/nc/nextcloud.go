package nc

import (
	"gitlab.bertha.cloud/partitio/Nextcloud-Partitio/gonextcloud"
	"go-starter/config"
)

type Login struct {
	Url      string
	Username string
	Password string
}

var login = Login{
	Url:      "",
	Username: "",
	Password: "",
}

func init() {
	login = Login{
		Url:      config.GetConfig("nextcloud.url").(string),
		Username: config.GetConfig("nextcloud.username").(string),
		Password: config.GetConfig("nextcloud.password").(string),
	}
}

func SetUser(url string, username string, password string) {
	login = Login{
		Url:      config.GetConfig("nextcloud.url").(string),
		Username: config.GetConfig("nextcloud.username").(string),
		Password: config.GetConfig("nextcloud.password").(string),
	}
}

func LoginCloud() *gonextcloud.Client {
	g, err := gonextcloud.NewClient(login.Url)
	if err != nil {
		panic(err)
	}
	if err := g.Login(login.Username, login.Password); err != nil {
		panic(err)
	}
	return g
}
