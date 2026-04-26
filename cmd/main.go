package main

import (
	"fmt"
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

	wg.Wait()
}
