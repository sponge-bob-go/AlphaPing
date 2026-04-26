package main

import (
	"fmt"
	"main/internal/signal_logic"
	bybit "main/internal/signal_logic/ByBit"
	mexc "main/internal/signal_logic/MEXC"
)

func main() {
	Mexc := mexc.CreateMEXCModel()
	Bybit := bybit.CreateBybitModel()
	c, _ := Bybit.GetOHLC("BTCUSDT", "1")
	fmt.Println(c)
	engine := signal_logic.CreateOHLCServces(Mexc)
	engine = signal_logic.CreateOHLCServces(Bybit)
	b, err := engine.GetCandle("BTCUSDT", "1")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(b)
}
