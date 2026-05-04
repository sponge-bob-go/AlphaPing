import pandas as pd
import numpy as np
from sklearn.model_selection import train_test_split
from xgboost import XGBRegressor
import joblib

df = pd.read_csv("dataset/build_dataset/dataset.csv")

X = df.drop(columns=["label"])
y = df["label"]

X_train, X_test, y_train, y_test = train_test_split(
    X, y, test_size=0.2, shuffle=False
)

model = XGBRegressor(
    n_estimators=300,
    max_depth=6,
    learning_rate=0.05,
    subsample=0.8,
    colsample_bytree=0.8
)

model.fit(X_train, y_train)

preds = model.predict(X_test)

print("Train done")
print("Sample prediction:", preds[:5])

joblib.dump(model, "model_btc_xgboost.pkl")
print("Model saved: model_btc_xgboost")

y_test_direction = np.sign(y_test)
preds_direction = np.sign(preds)

accuracy = (y_test_direction == preds_direction).mean()

print(f"\n--- Результаты анализа направления ---")
print(f"Точность (Direction Accuracy): {accuracy:.2%}")


print(f"Предсказано 'Вверх': {sum(preds_direction == 1)}")
print(f"Предсказано 'Вниз': {sum(preds_direction == -1)}")


mae = np.mean(np.abs(y_test - preds))
print(f"Средняя ошибка (MAE): {mae:.5f}")