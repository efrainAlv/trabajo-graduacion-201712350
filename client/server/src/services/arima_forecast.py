import pandas as pd
from statsmodels.tsa.arima.model import ARIMA
from datetime import datetime, date
import json

def get_prediction(request):

    data = pd.read_csv('https://storage.googleapis.com/crypto-forecast-models/arima/btcusd-history-price_daily.csv', delimiter=';', parse_dates=['time_period_end'], usecols=['rate_close', 'time_period_end'])
    data.index = pd.date_range(start='2016-01-01 00:00:00', end='2022-10-22 00:00:00', freq='D')

    return make_forecast(data, request['periods'])



def make_forecast(data, periods):
    
    data_model = data
    forecast = []

    try:

        initial_date = datetime.fromisoformat(data_model[-1:].index.strftime('%Y-%m-%d %H:%M:%S').item()).timestamp()
        initial_date=datetime.fromtimestamp(initial_date+(1*24*60*60))
        final_date = None

        for i in range(periods):
            model = ARIMA(data_model['rate_close'], order=(3, 2, 2))
            fitted = model.fit()
            
            fcst = fitted.forecast(1, alpha=0.05)
            forecast.append(fcst)  # 95% conf

            start = datetime.fromisoformat(data_model[-1:].index.strftime('%Y-%m-%d %H:%M:%S').item()).timestamp()
            end = datetime.fromtimestamp(start+(1*24*60*60))
            final_date = end
            some = pd.DataFrame([{'rate_close': fcst[0]}], columns=['rate_close'], index=pd.date_range(start=end, end=end, freq='D'))
            data_model = pd.DataFrame(pd.concat([data_model, some]))

        return transform_data(forecast)

    except Exception as err:
        return {'msg': "{}".format(err), 'code':400}



def transform_data(forecast):

    data = []

    for i in range(len(forecast)):
        time_period_end = datetime.fromisoformat(forecast[i].index.strftime('%Y-%m-%d %H:%M:%S').item()).strftime('%Y-%m-%dT%H:%M:%S')
        time_period_start = datetime.fromtimestamp(datetime.fromisoformat(time_period_end).timestamp()-(1*24*60*60)).strftime('%Y-%m-%dT%H:%M:%S')
        time_open = datetime.fromtimestamp(datetime.fromisoformat(time_period_start).timestamp()+(1*60)).strftime('%Y-%m-%dT%H:%M:%S')
        time_close = datetime.fromtimestamp(datetime.fromisoformat(time_period_start).timestamp()+(23*60*60+59*60)).strftime('%Y-%m-%dT%H:%M:%S')
        rate_open = forecast[i][0]
        rate_high = forecast[i][0]
        rate_low = forecast[i][0]
        rate_close = forecast[i][0]
        some = pd.DataFrame([{
            'time_period_start': time_period_start,
            'time_period_end': time_period_end, 
            'time_open': time_open,
            'time_close': time_close,
            'rate_open': rate_open,
            'rate_high': rate_high,
            'rate_low': rate_low,
            'rate_close': forecast[i][0],
            }], columns=[ 'time_period_start', 'time_period_end', 'time_open', 'time_close', 'rate_open', 'rate_high', 'rate_low', 'rate_close'], index=forecast[i].index)
        data = pd.DataFrame(data.append(some))


    result = data.to_json(orient="records")
    result = json.loads(result)

    return {
        'msg': result,
        'code':200
    }