from statsmodels.tsa.stattools import adfuller
import pandas as pd
import numpy as np

df = pd.read_csv('./data/btcusd-history-price.csv', parse_dates=['time_period_end'], delimiter=';', usecols=['rate_close'])
series = df.loc[:, 'value'].values