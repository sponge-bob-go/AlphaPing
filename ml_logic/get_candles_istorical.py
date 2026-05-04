import asyncio
from datetime import datetime, timedelta
import aiohttp
import json
import os

now = datetime.now()
start_time = now - timedelta(days=2190)

GO_URL = "http://localhost:8080/ml-api/return_indicators"

SYMBOL = "BTCUSDT"


def convert_to_ts(dt: datetime):
    return int(dt.timestamp() * 1000)


async def fetch(session, url):
    async with session.get(url) as resp:
        return await resp.json()


async def post_to_go(session, payload):
    async with session.post(GO_URL, json=payload) as resp:
        return await resp.json()


def make_label(candles_5):
    start_price = float(candles_5[-1][4])
    end_price = float(candles_5[0][4])
    return (end_price - start_price) / start_price


def save_file(symbol, day_index, data):
    path = f"dataset/{symbol}"
    os.makedirs(path, exist_ok=True)

    file_path = f"{path}/day_{day_index}.json"

    with open(file_path, "w", encoding="utf-8") as f:
        json.dump(data, f, indent=2)


async def pipeline():
    day = start_time
    day_index = 0

    async with aiohttp.ClientSession() as session:

        while day < now:

            base_url = lambda interval: (
                "https://api.bybit.com/v5/market/kline"
                f"?category=linear&symbol={SYMBOL}&interval={interval}"
                f"&start={convert_to_ts(day)}"
                f"&end={convert_to_ts(day + timedelta(days=3))}"
                "&limit=304"
            )


            url15 = base_url("15")

            result = await fetch(session=session, url=url15)

            data_15m = result["result"]["list"]
            
            if len(data_15m) < 204:
                day += timedelta(days=3)
                day_index += 3
                continue 
            else:
                indicators = await post_to_go(session, {
                    "candles": data_15m[5:] 
                })
                print(len(data_15m))
                last_5 = data_15m[:5]

                label = make_label(last_5)

                dataset_item = {
                    "symbol": SYMBOL,
                    "time": str(day),
                    "label_return": label,
                    "indicators": indicators,
                }

                save_file(SYMBOL, day_index, dataset_item)

                print(f"Saved {SYMBOL} day_{day_index}")

                day_index += 3
                day += timedelta(days=3)


asyncio.run(pipeline())