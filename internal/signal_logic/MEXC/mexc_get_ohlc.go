package mexc

import (
	"encoding/json"
	"errors"
	"fmt"
	"main/internal/signal_logic"
	"net/http"
	"strconv"
	"time"
)

func (m *MEXCModel) GetOHLC(coinSymbol, interval string) ([]signal_logic.OHLCStruct, error) {
	url := fmt.Sprintf("https://api.mexc.com/api/v3/klines?symbol=%s&interval=%s&limit=200", coinSymbol, interval)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var res [][]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, err
	}

	var listOHLC []signal_logic.OHLCStruct

	for _, candle := range res {
		intTimeStamp, ok := candle[0].(float64)
		if !ok {
			return nil, errors.New("Ошибка декодирования")
		}
		Time := time.UnixMilli(int64(intTimeStamp))
		openStr, ok := candle[1].(string)
		if !ok {
			return nil, errors.New("Ошибка декодирования")
		}
		Open, _ := strconv.ParseFloat(openStr, 64)
		highStr, _ := candle[2].(string)
		High, _ := strconv.ParseFloat(highStr, 64)
		lowStr, _ := candle[3].(string)
		Low, _ := strconv.ParseFloat(lowStr, 64)
		closeStr, _ := candle[4].(string)
		Close, _ := strconv.ParseFloat(closeStr, 64)
		volumeStr, _ := candle[5].(string)
		Volume, _ := strconv.ParseFloat(volumeStr, 64)

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

	return listOHLC, nil
}
