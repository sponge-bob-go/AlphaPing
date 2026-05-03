import asyncio
from datetime import datetime, timedelta
import aiohttp



now = datetime.now()

year2025 = now - timedelta(days=365)

def convert_to_ts(time: datetime):
    return int(time.timestamp()*1000)


async def fetch(ses, url):
    async with ses.get(url) as resp:
        return await resp.json()


async def get_candles():
    while year2025 < now:
        urls = [
            f"https://api.bybit.com/v5/market/kline?category=linear&symbol=BTCUSDT&interval=1&start={convert_to_ts(year2025)}&end={convert_to_ts(year2025+timedelta(days=1))}&limit=1000",
            f"https://api.bybit.com/v5/market/kline?category=linear&symbol=TONUSDT&interval=1&start={convert_to_ts(year2025)}&end={convert_to_ts(year2025 + timedelta(days=1))}&limit=1000",
            f"https://api.bybit.com/v5/market/kline?category=linear&symbol=ETHUSDT&interval=1&start={convert_to_ts(year2025)}&end={convert_to_ts(year2025 + timedelta(days=1))}&limit=1000"
        ]

        async with aiohttp.ClientSession() as ses:
            tas = [fetch(ses, url) for url in urls]
            results = await asyncio.gather(*tas)

        print(results)
        
asyncio.run(get_candles())