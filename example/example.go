package main

import (
	"fmt"
	"os"

	"github.com/taxjar/taxjar-go"
)

func main() {
	client := taxjar.NewClient(taxjar.Config{
		APIKey: os.Getenv("TAXJAR_API_KEY"),
	})
	// or
	// client := taxjar.NewClient()
	// client.APIKey = os.Getenv("TAXJAR_API_KEY")

	res1, err := client.Categories()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res1.Categories)

	res2, err := client.TaxForOrder(taxjar.TaxForOrderParams{
		ToCountry:     "US",
		ToState:       "CA",
		ToZip:         "94043",
		Shipping:      5,
		Amount:        10,
		ExemptionType: "non_exempt",
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res2.Tax)

	res3, err := client.ListOrders(taxjar.ListOrdersParams{
		FromTransactionDate: "2019/05/15",
		ToTransactionDate:   "2019/08/19",
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res3.Orders)

	res4, err := client.ShowOrder("13579-246810")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res4.Order)

	res5, err := client.DeleteOrder("13579-246810")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(res5.Order)
}
