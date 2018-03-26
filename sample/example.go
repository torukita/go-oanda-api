package main

import(
	apiclient "github.com/torukita/go-oanda-api/client"
	"github.com/torukita/go-oanda-api/client/oanda"
	"log"
	"fmt"
	"os"
)

/*
Set environment varaiables OANDA_HOST, OANDA_ID, OANDA_TOKEN
exp.) 
export OANDA_HOST="api-fxpractice.oanda.com"  this is demo api server
export OANDA_ID="XXX-YYY-ZZZZZZZ-AAA"  your registered id
export OANDA_TOKEN="xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx-yyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyy" get your token to use api
*/

func main () {

	host := os.Getenv("OANDA_HOST")
	id := os.Getenv("OANDA_ID")
	token := os.Getenv("OANDA_TOKEN")

	auth := fmt.Sprintf("Bearer %s", token)
	config := &apiclient.TransportConfig{
		Host: host,
		BasePath: "/v3",
		Schemes: []string{"https"},
	}
	
	api := apiclient.NewHTTPClientWithConfig(nil, config).Oanda
	param := oanda.NewGetPricesParams()
	param.SetAccountID(id)
	param.SetAuthorization(auth)
	param.SetInstruments([]string{"AUD_JPY,EUR_USD"})

	response, err := api.GetPrices(param)
	if err != nil {
		log.Fatal(err)
	}

	for _, price := range response.Payload.Prices {
		// price is model.Price
		// you can use struct model.Price in you program.
		fmt.Printf("%s bid=%s ask=%s\n", price.Instrument, price.Bids[0].Price, price.Asks[0].Price)
	}		
	fmt.Println()
	prices, err := response.Payload.MarshalBinary()
	if  err != nil {
		log.Fatal(err)
	}
	fmt.Println("-- json string --")
	fmt.Println(string(prices))
}

