# Prediccion de precios de criptomonedas a traves de modelos predictivos y streaming de datos


# LINKS

## COINAPI

Time unit 	Period identifiers
Second 	1SEC, 2SEC, 3SEC, 4SEC, 5SEC, 6SEC, 10SEC, 15SEC, 20SEC, 30SEC
Minute 	1MIN, 2MIN, 3MIN, 4MIN, 5MIN, 6MIN, 10MIN, 15MIN, 20MIN, 30MIN
Hour 	1HRS, 2HRS, 3HRS, 4HRS, 6HRS, 8HRS, 12HRS
Day 	1DAY, 2DAY, 3DAY, 5DAY, 7DAY, 10DAY
Month 	1MTH, 2MTH, 3MTH, 4MTH, 6MTH
Year 	1YRS, 2YRS, 3YRS, 4YRS, 5YRS

https://rest.coinapi.io/v1/exchangerate/BTC/USD/history?period_id=1HRS&time_start=2016-01-01T00:00:00&time_end=2022-10-01T00:00:00&output_format=csv&limit=100000&apikey=9C6EF219-FBCD-4B72-93BF-D182AA91F44A

https://rest.coinapi.io/v1/exchangerate/BTC/USD/history?period_id=30MIN&time_start=2019-10-20T21:00:00&time_end=2022-08-01T00:00:00&apikey=1D70A894-4DF0-4872-A7E5-1DBB330CF364&output_format=csv&limit=100000

## BINANCE

https://api.binance.com/api/v3/klines?symbol=BTCUSDT&interval=30m&startTime=1588636800000&endTime=1588723200000
https://api.binance.com/api/v3/klines?symbol=BTCUSDT&interval=30m&startTime=1588636800000&endTime=1588723200000


2020-05-05T09:00:00.0000000Z;2020-05-05T10:00:00.0000000Z;2020-05-05T09:00:00.0000000Z;2020-05-05T09:59:00.0000000Z;9828.697785866061;2274324.155965172;8990.516862038741;17769.376185582314

## KEYS

CC0504E9-05C2-431B-95D8-95052BB378B5
9C6EF219-FBCD-4B72-93BF-D182AA91F44A
036A1654-3C7D-4442-A766-F279709A58F5
F967CEB6-34B9-40FE-BB57-D9F22F535656
831C5749-72F3-469B-92D1-FB3A886918D5
DD80510C-FBEE-40B7-B637-0E6E061B5F74



## Diferenciador (d) 

ADF Statistic: 0.4398155777739321
n_lags: 0.9829243699467475
p-value: 0.9829243699467475


## Versiones Pythnon

matplotlib==3.6.0
numpy==1.23.3
pandas==1.5.0
pmdarima==2.0.1
prophet==1.1.1
statsmodels==0.13.2

## Create kafka topic

cd /bin
kafka-topics --create --topic ada-topic --partitions 1 --replication-factor 1 --if-not-exists --bootstrap-server kafka-broker1:29092