import os
import json
import csv
import re

BASE_DIR = os.path.dirname(os.path.abspath(__file__))
DATA_DIR = os.path.join(BASE_DIR, "..", "dataset", "BTCUSDT")
OUTPUT_FILE = os.path.join(BASE_DIR, "..", "dataset", "build_dataset", "dataset.csv")

def trend_to_num(t):
    return {"Up": 1, "Down": -1}.get(t, 0)

def volume_to_num(v):
    return 1 if v == "Volume Up" else 0

def pattern_features(p):
    bullish = p.get("hammer", 0) + (p.get("builish_engulfing", 0) * 2) + p.get("white_solders", 0)
    bearish = p.get("gravestone_doji", 0) + (p.get("bearish_engulfing", 0) * 2) + p.get("black_solders", 0)
    neutral = p.get("doji", 0) + p.get("dragonfly_doji", 0)
    return bullish, bearish, neutral


all_files = [f for f in os.listdir(DATA_DIR) if f.endswith(".json")]
all_files.sort(key=lambda f: int(re.search(r'\d+', f).group()))

rows = []

for file in all_files:
    with open(os.path.join(DATA_DIR, file), "r", encoding="utf-8") as f:
        data = json.load(f)

    ind = data.get("indicators", {})
    if not ind: continue 

    macd = ind.get("MACD", {})
    patterns = ind.get("Patterns", {})
    

    sma = ind["SMA"]
    ema = ind["EMA"]
    
   
    ema_sma_diff = (ema - sma) / sma if sma != 0 else 0
    
    bullish, bearish, neutral = pattern_features(patterns)

    row = [
        ema_sma_diff,         
        ind["RSI"],
        macd.get("MACDLine", 0),
        macd.get("SignalLine", 0),
        macd.get("Histogram", 0),
        ind.get("ATRPercent", 0),
        ind.get("ADX", 0),
        bullish,
        bearish,
        neutral,
        trend_to_num(ind.get("Trend15")),
        volume_to_num(ind.get("Volume")),
        data["label_return"]  
    ]
    rows.append(row)


os.makedirs(os.path.dirname(OUTPUT_FILE), exist_ok=True)
with open(OUTPUT_FILE, "w", newline="", encoding="utf-8") as f:
    writer = csv.writer(f)
    writer.writerow([
        "EMA_SMA_Diff", "RSI", "MACDLine", "SignalLine", "Histogram",
        "ATRPercent", "ADX", "bullish", "bearish", "neutral",
        "trend15", "volume", "label"
    ])
    writer.writerows(rows)

print(f"Dataset created: {len(rows)} rows. ")
