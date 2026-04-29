package main

import (
	"fmt"
	coreapi "main/cmd/core_api"
	"main/cmd/core_api/handlers"
	internal "main/internal"
	basecryptodata "main/internal/base_crypto_data"
	"net/http"
)

func main() {
	coinMap := internal.MakeCoinMap()
	//ctx := context.Background()
	//conn, err := database.ConnetcToBD(ctx)
	//fmt.Println(conn)
	//if err != nil {
	//panic(err)
	//}

	mux := coreapi.MakeMux()

	mux.HandleFunc("/core-api/coins_base_data", handlers.HandlerDefaultData(coinMap))
	mux.HandleFunc("/core-api/coins_tech_data", handlers.HandlerTechData)

	go func() {

		if err := basecryptodata.StartUpdateCoinsPrice(coinMap); err != nil {
			fmt.Println(err)
		}
	}()

	go func() {

		if err := basecryptodata.StartUpdateCoinsInfo(coinMap); err != nil {
			fmt.Println(err)
		}
	}()
	http.ListenAndServe(":8080", mux)
	//sa;ldkasl;d;alsdkl
}
