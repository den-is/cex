package cmd

import (
	"encoding/json"
	"fmt"

	consul "github.com/hashicorp/consul/api"
)

func getCatalog() *consul.Catalog {
	return client().Catalog()
}

func output(s []string) {
	fmt.Println("Total: ", len(s))
	fmt.Print("---------------\n")
	if JSON {
		jsonout, _ := json.Marshal(s)
		fmt.Println(string(jsonout))
	} else {
		for _, item := range s {
			fmt.Println(item)
		}
	}
}

// GetDatacenters prints all available datacenters
func GetDatacenters() {

	cat := getCatalog()

	items, err := cat.Datacenters()

	if err == nil {
		output(items)
	} else {
		fmt.Println("something wrong. unable to fetch list of datacenters.")
		fmt.Println(err)
	}

}

// GetServices prints all available services in given datacenter
// Default datacenter is one discovered via consul API url
func GetServices() {

	cat := getCatalog()

	items, _, err := cat.Services(nil)

	if err == nil {
		services := make([]string, 0, len(items))
		for k := range items {
			services = append(services, k)
		}
		output(services)
	} else {
		fmt.Println("something wrong. unable to fetch list of services.")
		fmt.Println(err)
	}
}

// GetService prints nodes of given service
func GetService(svc string) {

	cat := getCatalog()

	tag := "" // to implement as argument

	count := 0

	items, _, err := cat.Service(svc, tag, nil)
	fmt.Println("Service:", svc)

	if err == nil {
		nodes := make([]string, 0, len(items))
		for _, item := range items {
			count++
			if HOSTNAME {
				nodes = append(nodes, item.Node)
			} else {
				nodes = append(nodes, item.Address)
			}
		}
		output(nodes)
	} else {
		fmt.Println("something wrong. unable to fetch list of nodes for", svc)
		fmt.Println(err)
	}

}
