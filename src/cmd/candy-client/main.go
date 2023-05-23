package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"day_4/client"
	operations "day_4/client/operations"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

type Flag struct {
	K string
	C int64
	M int64
}

func newFlag() *Flag {
	candy := flag.String("k", "", "for the candy type")
	count := flag.Int64("c", 0, "count of candy to buy")
	money := flag.Int64("m", 0, "amount of money")
	flag.Parse()
	return &Flag{
		K: *candy,
		C: *count,
		M: *money,
	}
}

func main() {
	data, _ := ioutil.ReadFile("minica.pem")

	cp, _ := x509.SystemCertPool()
	cp.AppendCertsFromPEM([]byte(data))
	tlsConf := &tls.Config{
		RootCAs: cp,
	}

	cl := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: tlsConf,
		},
	}

	client := client.NewHTTPClientWithConfig(strfmt.Default, &client.TransportConfig{
		Host:    "localhost:8080",
		Schemes: []string{"https"},
	})

	flag := newFlag()
	money := swag.Int64(flag.M)
	count := swag.Int64(flag.C)
	candyType := swag.String(flag.K)

	resp, err := client.Operations.BuyCandy(&operations.BuyCandyParams{
		Order: operations.BuyCandyBody{
			Money:      money,
			CandyCount: count,
			CandyType:  candyType,
		},
		Context:    context.Background(),
		HTTPClient: cl,
	}, nil)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp.Payload.Thanks, "Your change is", resp.Payload.Change)
}
