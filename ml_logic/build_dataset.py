import os
import json
import csv

BASE_DIR = os.path.dirname(os.path.abspath(__file__))
DATA_DIR = os.path.join(BASE_DIR, "..", "dataset", "BTCUSDT")
OUTPUT_FILE = os.path.join(BASE_DIR,"..", "dataset","build_dataset","dataset.csv")



def trend_to_num(t):
    if t == "Up":
        return 1
    if t == "Down":
        return -1
    return 0


def volume_to_num(v):
    return 1 if v == "Volume Up" else 0


def pattern_features(p):
    bullish = 0
    bearish = 0
    neutral = 0

    if p.get("hammer", False):
        bullish += 1

    if p.get("builish_engulfing", False):
        bullish += 2

    if p.get("white_solders", False):
        bullish += 1

    if p.get("gravestone_doji", False):
        bearish += 1

    if p.get("bearish_engulfing", False):
        bearish += 2

    if p.get("black_solders", False):
        bearish += 1

    if p.get("doji", False):
        neutral += 1

    if p.get("dragonfly_doji", False):
        neutral += 1

    return bullish, bearish, neutral


rows = []

for file in os.listdir(DATA_DIR):
    if not file.endswith(".json"):
        continue

    with open(os.path.join(DATA_DIR, file), "r", encoding="utf-8") as f:
        data = json.load(f)

    ind = data["indicators"]
    macd = ind["MACD"]
    patterns = ind["Patterns"]

    bullish, bearish, neutral = pattern_features(patterns)

    row = [
        ind["SMA"],
        ind["EMA"],
        ind["RSI"],

        macd["MACDLine"],
        macd["SignalLine"],
        macd["Histogram"],

        ind["ATRPercent"],
        ind["ADX"],

        bullish,
        bearish,
        neutral,

        trend_to_num(ind["Trend15"]),
        volume_to_num(ind["Volume"]),

        data["label_return"]
    ]

    rows.append(row)


with open(OUTPUT_FILE, "w", newline="", encoding="utf-8") as f:
    writer = csv.writer(f)

    writer.writerow([
        "SMA", "EMA", "RSI",
        "MACDLine", "SignalLine", "Histogram",
        "ATRPercent", "ADX",
        "bullish", "bearish", "neutral",
        "trend15", "volume",
        "label"
    ])

    writer.writerows(rows)

print("Dataset created:", len(rows))