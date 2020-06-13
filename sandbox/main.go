package main

import (
	"fmt"
	"log"
	"os"

	"github.com/asciiu/appa/lib/config"
	"github.com/kelseyhightower/envconfig"
	coinbasepro "github.com/preichenberger/go-coinbasepro/v2"
)

//type Bet struct {
//	wagerAmount uint
//}
//
//func findUser(db string) func(string) string {
//	return func(userID string) string {
//		user := fmt.Sprintf("%s %s", db, userID)
//		return user
//	}
//}
//
//func WithStatus(db string) func(status string) string {
//	return func(status string) string {
//		str := fmt.Sprintf("%s %s", db, status)
//		return str
//	}
//}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type CoinbaseConfig struct {
	Key        string `envconfig:"COINBASE_KEY" required:"true"`
	Secret     string `envconfig:"COINBASE_SECRET" required:"true"`
	Passphrase string `envconfig:"COINBASE_PASSPHRASE" required:"true"`
}

func main() {
	argsWithoutProg := os.Args[1:]
	if len(argsWithoutProg) == 0 {
		log.Fatal("command line env file not found in command args")
	}

	envfile := argsWithoutProg[0]
	config.LoadEnv(envfile)

	var cfg CoinbaseConfig
	err := envconfig.Process("myapp", &cfg)
	check(err)

	// coinbase test
	client := coinbasepro.NewClient()

	// optional, configuration can be updated with ClientConfig
	client.UpdateConfig(&coinbasepro.ClientConfig{
		BaseURL:    "https://api.pro.coinbase.com",
		Key:        cfg.Key,
		Passphrase: cfg.Passphrase,
		Secret:     cfg.Secret,
	})

	accounts, err := client.GetAccounts()
	if err != nil {
		fmt.Println(err.Error())
	}

	for _, a := range accounts {
		fmt.Printf("%s %s\n", a.Currency, a.Balance)
	}
}
