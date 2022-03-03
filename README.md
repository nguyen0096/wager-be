# wager-be

* To run server as docker container: `./start.sh bootstrap`
* To build and run server on local machine `./start.sh start`

## Curl examples

* Create new wager
```
curl --location -g --request POST 'http://localhost:8080/wagers' \
--header 'Content-Type: application/json' \
--data-raw '{
    "total_wager_value": 100,
    "odds": 150,
    "selling_percentage": 50,
    "selling_price": 1.5
}'
```

* List wagers
```
curl --location -g --request GET 'http://localhost:8080/wagers?page=1&limit=1'
```

* Buy
```
curl --location -g --request POST 'http://localhost:8080/buy/1' \
--header 'Content-Type: application/json' \
--data-raw '{
    "buying_price": 1.3
}'
```