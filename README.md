# wager-be

## Prequisites
* golang migrate `brew install golang-migrate`


## Concerns

* `selling_price` must be greater than `total_wager_value` *  ( `selling_percentage` / 100 ) => why compare price with amount?
* wager has amount sold, but buy action don't have amount. Instead, buy action has wager id => buy 100% amount of wager?
