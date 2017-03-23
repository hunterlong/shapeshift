<p align="center">
  <img src="https://cdn.pbrd.co/images/MK2f4akQc.jpg" alt="shapeshift api golang"/>
</p>

# ShapeShift in Go Language
[![Build Status](https://travis-ci.org/hunterlong/shapeshift.svg?branch=master)](https://travis-ci.org/hunterlong/shapeshift)  [![Coverage Status](https://coveralls.io/repos/github/hunterlong/shapeshift/badge.svg?branch=master)](https://coveralls.io/github/hunterlong/shapeshift?branch=master) [![GoDoc](https://godoc.org/github.com/hunterlong/shapeshift?status.svg)](https://godoc.org/github.com/hunterlong/shapeshift) [![Go Report Card](https://goreportcard.com/badge/github.com/hunterlong/shapeshift)](https://goreportcard.com/report/github.com/hunterlong/shapeshift)

### Functions

:white_check_mark: Get Rate
```go
pair := Pair{"eth_btc"}

rate := pair.GetRates()

t.Logf("Pair: %v at %v",rate.Pair, rate.Rate)
```

:white_check_mark: Deposit Limits
```go
pair := Pair{"eth_btc"}

limits := pair.GetLimits()

t.Logf("Limits on Pair: %v LIMIT: %v",limits.Pair, limits.Limit)
```

:white_check_mark: Market Info
```go
pair := Pair{"btc_eth"}

info := pair.GetInfo()

t.Logf("Info Pair: %v | Min: %f | LIMIT: %f",info.Pair, info.Min, info.Limit)
```

:white_check_mark: Recent Transactions
```go
recent := RecentTransactions()

t.Log(recent)
```
:white_check_mark: Deposit Address Status
```go
status := DepositStatus("1L75eRMgeCwAxEjD1oWXjLgud9jxwxm34u")

t.Log(status.Status)
```

:white_check_mark: Time Remaining on Fixed Transaction Amount
```go
status := TimeRemaining("1L75eRMgeCwAxEjD1oWXjLgud9jxwxm34u")

t.Log(status.Status)
```
:white_check_mark: Get Available Coins
```go
coins := Coins()
t.Log(coins)
```

:white_large_square: Validate Address with Coin Symbol

### Primary Requests

:white_check_mark: Create New Transaction
```go
new := New{
		Pair: "eth_btc",
		ToAddress: "1L75eRMgeCwAxEjD1oWXjLgud9jxwxm34u",
		FromAddress: "0xcf2f204aC8D7714990912fA422874371c001217D",
	}

response := new.Shift()

t.Log("Send To Address: ", response.SendTo, "\n")
t.Log("Send Type: ", response.SendType, "\n")
t.Log("Send Type: ", response.ApiKey, "\n")
```

:white_check_mark: Request Email Receipt
```go
info := Receipt{
		Email: "user@awesome.com",
		TransactionID: "owkdwodkkwokdwdw",
	}

response := info.Send();

t.Log(response)
```

:white_check_mark: Fixed Amount Transaction
```go
new := New{
		Pair: "eth_btc",
		Amount: 0.25,
		ToAddress: "1L75eRMgeCwAxEjD1oWXjLgud9jxwxm34u",
		FromAddress: "0xcf2f204aC8D7714990912fA422874371c001217D",
	}

response := new.FixedShift()

t.Log("Pair: ", response.Pair)
t.Log("Quoted Rate: ", response.QuotedRate)
t.Log("Deposit Address: ", response.Deposit)
t.Log("Deposit Amount: ", response.DepositAmount)
t.Log("Withdraw Amount: ", response.WithdrawalAmount)
t.Log("Withdraw Address: ", response.Withdrawal)
t.Log("Expiration: ", response.Expiration)
```
:white_check_mark: Cancel Pending Transaction
```go
old := Address{
		Id: newSendToAddress,
	}

	response := old.Cancel()
```


### API Key Required Requests

:white_large_square: Get Transactions from Private API Key

:white_large_square: Get Transactions from Output Address
