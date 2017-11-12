package cmd

import (
	"fmt"
	"os"

	consul "github.com/hashicorp/consul/api"
)

func config() *consul.Config {

	clientConfig := consul.DefaultConfig()

	URLenv, URLenvExists := os.LookupEnv("CEX_CONSUL_URL")

	if URLenvExists && (URL == "http://localhost:8500") {
		clientConfig.Address = URLenv
	} else {
		clientConfig.Address = URL
	}

	if DC != "" {
		clientConfig.Datacenter = DC
	}

	return clientConfig

}

func client() *consul.Client {

	client, err := consul.NewClient(config())

	if err != nil {
		fmt.Println("something wrong. unable to initiate connection.")
		fmt.Println(err)
		return nil
	}

	return client
}
