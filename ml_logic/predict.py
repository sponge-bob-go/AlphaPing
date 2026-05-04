import joblib
import pandas as pd
import numpy as np


model = joblib.load("model_btc_xgboost.pkl")


data_from_go = {
    "SMA": 79595.12,
    "EMA": 79255.95808209402,
    "Trend15": "Up",
    "RSI": 63.50016572754414,
    "MACD": {
        "MACDLine": -187.2975836123078,
        "SignalLine": -224.92066899136503,
        "Histogram": 37.623085379057244
    },
    "ATR": 242.72857142857018,
    "ATRPercent": 0.30756943754309224,
    "ADX": 31.193196191592083,
    "Volume": "Volume Down",
    "Supports": [
        78180.4, 78368.3, 78421.5, 78312.8, 78421.6, 78414.7, 78439.3, 78614.3, 
        78500.7, 78145.2, 78102.7, 78109.1, 78112.9, 78142.6, 78080, 78355.4, 
        78380.5, 78351.2, 78427, 78380, 78700, 78601.7, 78620, 78659.1, 78599.2, 
        78563.5, 78627.5, 78677.7, 78732.4, 78845.8, 78660.1, 78280, 79975.2, 
        79929.6, 80245.8, 79957.8, 79500.1, 79593.7, 79624.3, 79685.2, 78204, 
        78700, 78711.6, 78576.4
    ],
    "Resistances": [
        78269.3, 78435.5, 78395.3, 78487.8, 78485, 78598.5, 78488.7, 78509, 
        78522.3, 78474.5, 78514.8, 79165, 78878.7, 78819.3, 78765.4, 78680.8, 
        78317, 78197.5, 78211.7, 78255.8, 78253.5, 78454.2, 78500, 78455.8, 
        78600.9, 78562.6, 78887.9, 78839.5, 78784.7, 78814.4, 78806.1, 78779.1, 
        78826.5, 78761.9, 78769.1, 78861.6, 79341.1, 79072, 79450, 79383.5, 
        80132.5, 80460.3, 80632.5, 80408.2, 79973.3, 79911, 79855, 79922.4, 
        79108.4, 79124.7, 79020.3, 79010.7, 79120.7
    ],
    "Patterns": {
        "bearish_engulfing": False,
        "black_solders": False,
        "builish_engulfing": False,
        "doji": False,
        "dragonfly_doji": False,
        "gravestone_doji": False,
        "hammer": False,
        "white_solders": False
    }
}


def prepare_features(ind):
    ema_sma_diff = (ind["EMA"] - ind["SMA"]) / ind["SMA"] if ind["SMA"] != 0 else 0
    

    trend_num = 1 if ind["Trend15"] == "Up" else (-1 if ind["Trend15"] == "Down" else 0)
    vol_num = 1 if ind["Volume"] == "Volume Up" else 0
    

    p = ind["Patterns"]
    bullish = p.get("hammer", 0) + (p.get("builish_engulfing", 0) * 2) + p.get("white_solders", 0)
    bearish = p.get("gravestone_doji", 0) + (p.get("bearish_engulfing", 0) * 2) + p.get("black_solders", 0)
    neutral = p.get("doji", 0) + p.get("dragonfly_doji", 0)


    features = [
        ema_sma_diff,
        ind["RSI"],
        ind["MACD"]["MACDLine"],
        ind["MACD"]["SignalLine"],
        ind["MACD"]["Histogram"],
        ind["ATRPercent"],
        ind["ADX"],
        bullish,
        bearish,
        neutral,
        trend_num,
        vol_num
    ]
    return np.array(features).reshape(1, -1)


X_actual = prepare_features(data_from_go)
prediction = model.predict(X_actual)[0]


current_price = 78892.7
price_change_usd = current_price * prediction

print(f"Прогноз изменения: {prediction:.4%}")
print(f"В долларах: {price_change_usd:.2f}$")
print(f"Ожидаемая цена через 1ч 15мин: {current_price + price_change_usd:.2f}$")