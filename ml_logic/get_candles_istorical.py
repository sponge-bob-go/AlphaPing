import asyncio
from datetime import datetime, timedelta
import aiohttp
import json
import os

now = datetime.now()
start_time = now - timedelta(days=365)

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


def make_label(candles_100):
    start_price = float(candles_100[60][4])
    end_price = float(candles_100[-1][4])
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
                f"&end={convert_to_ts(day + timedelta(days=1))}"
                "&limit=300"
            )

            urls = {
                "1": base_url("1"),
                "5": base_url("5"),
                "15": base_url("15")
            }

            tasks = {k: asyncio.create_task(fetch(session, v)) for k, v in urls.items()}
            results = await asyncio.gather(*tasks.values())

            data_1m = results[0]["result"]["list"]

            indicators = await post_to_go(session, {
                "candles": data_1m
            })

            last_100 = data_1m[:100]

            label = make_label(last_100)

            dataset_item = {
                "symbol": SYMBOL,
                "time": str(day),
                "label_return": label,
                "indicators": indicators,
            }

            save_file(SYMBOL, day_index, dataset_item)

            print(f"Saved {SYMBOL} day_{day_index}")

            day_index += 1
            day += timedelta(days=1)


asyncio.run(pipeline())