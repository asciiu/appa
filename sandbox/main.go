package main

import (
	"fmt"
	"log"
	"os"
	"time"

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

	printFills := coinbaseFills()
	printFills("BTC-USD")
	printFills("ETH-USD")
	printFills("LTC-USD")
}

func coinbaseFills() func(string) {
	var cfg CoinbaseConfig
	err := envconfig.Process("myapp", &cfg)
	check(err)

	client := coinbasepro.NewClient()

	// optional, configuration can be updated with ClientConfig
	client.UpdateConfig(&coinbasepro.ClientConfig{
		BaseURL:    "https://api.pro.coinbase.com",
		Key:        cfg.Key,
		Passphrase: cfg.Passphrase,
		Secret:     cfg.Secret,
	})

	return func(productID string) {
		fmt.Println(productID)

		btcSearch := coinbasepro.ListFillsParams{
			ProductID: productID,
		}

		dateFormat := "2006-Jan-02"

		var fills []coinbasepro.Fill
		cursor := client.ListFills(btcSearch)
		for cursor.HasMore {
			if err := cursor.NextPage(&fills); err != nil {
				fmt.Println(err.Error())
			}

			for _, f := range fills {

				transactionTime := f.CreatedAt.Time()
				jan2019, _ := time.Parse(dateFormat, "2019-Jan-01")
				jan2020, _ := time.Parse(dateFormat, "2020-Jan-01")

				if transactionTime.After(jan2019) && transactionTime.Before(jan2020) {
					fmt.Printf("\t%s price:%s size:%s date: %s\n", f.Side, f.Price, f.Size, f.CreatedAt.Time().Format(dateFormat))
				}
			}
		}
	}
}
