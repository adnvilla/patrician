package main

import (
	"github.com/adnvilla/patrician/patrician/domain"
)

func main() {
	for _, city := range domain.Cities {
		city.SetCommodities(domain.GetCommodities())
	}
}
