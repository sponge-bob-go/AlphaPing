package main

import (
	"context"
	"fmt"
	"main/database"
	internal "main/internal"
	basecryptodata "main/internal/base_crypto_data"
	"sync"
)

func main() {
	coinMap := internal.MakeCoinMap()
	var wg sync.WaitGroup

	wg.Add(1)

	go func() {
		defer wg.Done()

		if err := basecryptodata.StartUpdateCoinsPrice(coinMap); err != nil {
			fmt.Println(err)
		}
	}()

	wg.Add(1)

	go func() {
		defer wg.Done()

		if err := basecryptodata.StartUpdateCoinsInfo(coinMap); err != nil {
			fmt.Println(err)
		}
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		ctx := context.Background()
		conn, err := database.ConnetcToBD(ctx)
		if err != nil {
			panic(err)
		}
		fmt.Println(conn)
	}()

	wg.Wait()
}
