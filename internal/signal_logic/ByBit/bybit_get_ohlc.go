package bybit

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"main/internal/signal_logic"
	"net/http"
	"strconv"
	"time"
)

type bybitOhlcResp struct {
	RetCode int    `json:"retCode"`
	RetMsg  string `json:"retMsg"`
	Result  struct {
		Category string     `json:"category"`
		Symbol   string     `json:"symbol"`
		List     [][]string `json:"list"`
	} `json:"result"`
	Time int64 `json:"time"`
}

func (b *ByBitModel) GetOHLC(coinSymbol string, interval string) ([]signal_logic.OHLCStruct, error) {
	var ohlcResp bybitOhlcResp
	url := fmt.Sprintf("https://api.bybit.com/v5/market/kline?category=spot&symbol=%s&interval=%s&limit=%d", coinSymbol, interval, 200)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(body, &ohlcResp); err != nil {
		return nil, err
	}

	if ohlcResp.RetCode != 0 {
		return nil, errors.New("Ошибка получения OHlC")
	}

	listOHLC := make([]signal_logic.OHLCStruct, 0, len(ohlcResp.Result.List))

	for _, i := range ohlcResp.Result.List {
		intTimeStamp, _ := strconv.ParseInt(i[0], 10, 64)
		Time := time.UnixMilli(intTimeStamp)
		Open, _ := strconv.ParseFloat(i[1], 64)
		High, _ := strconv.ParseFloat(i[2], 64)
		Low, _ := strconv.ParseFloat(i[3], 64)
		Close, _ := strconv.ParseFloat(i[4], 64)
		Volume, _ := strconv.ParseFloat(i[5], 64)
		newCandle := signal_logic.OHLCStruct{
			TimeOpen: Time,
			Open:     Open,
			High:     High,
			Low:      Low,
			Close:    Close,
			Volume:   Volume,
		}
		listOHLC = append(listOHLC, newCandle)
	}
	for i, j := 0, len(listOHLC)-1; i < j; i, j = i+1, j-1 {
		listOHLC[i], listOHLC[j] = listOHLC[j], listOHLC[i]
	}

	return listOHLC, nil
}
