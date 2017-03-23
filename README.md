<p align="center">
  <img src="https://cdn.pbrd.co/images/MK2f4akQc.jpg" alt="shapeshift api golang"/>
</p>

# ShapeShift in Go Language
[![Build Status](https://travis-ci.org/hunterlong/shapeshift.svg?branch=master)](https://travis-ci.org/hunterlong/shapeshift)  [![Coverage Status](https://coveralls.io/repos/github/hunterlong/shapeshift/badge.svg?branch=master)](https://coveralls.io/github/hunterlong/shapeshift?branch=master) [![GoDoc](https://godoc.org/github.com/hunterlong/shapeshift?status.svg)](https://godoc.org/github.com/hunterlong/shapeshift) [![Go Report Card](https://goreportcard.com/badge/github.com/hunterlong/shapeshift)](https://goreportcard.com/report/github.com/hunterlong/shapeshift)

This Go Language Package will allow you to use the [ShapeShift API](https://info.shapeshift.io/) and convert your cryptocurrencies in your very own application. It includes most of the ShapeShift API requests listed on their references website. Below you'll find a perfect example of a new ShapeShift transaction.

```go
go get github.com/hunterlong/shapeshift
```
```go
import "github.com/hunterlong/shapeshift"
```

### Simple ShapeShift Transaction
I want to convert Ethereum to Bitcoin. The 'ToAddress' is my Bitcoin address. Once I run this, I'll get a Ethereum address from ShapeShift.
```go
new := shapeshift.New{
	Pair: "eth_btc",
	ToAddress: "16FdfRFVPUwiKAceRSqgEfn1tmB4sVUmLh",
	// FromAddress: "0xcf2f204aC8D7714990912fA422874371c001217D",  (Optional Return To Ethereum Address)
       }

response := new.Shift()

fmt.Println("Send Ethereum to Address: ",response.SendTo)
fmt.Println("Receiving Coin: ",response.ReturnType)
```

### Get Status of Transaction
Once I sent some Ethereum to the given Ethereum address, I want to check the status of my ShapeShift transaction by inserting my 'ToAddress' above. (Address I want the Bitcoin to go)
```go
status := shapeshift.DepositStatus("16FdfRFVPUwiKAceRSqgEfn1tmB4sVUmLh")

fmt.Println(status.Status)
// no_deposits
// received
// complete
// failed
```

# Functions
The other ShapeShift API requests are available for you to use. 

### :white_check_mark: Get Rate
Gets the current rate offered by Shapeshift. This is an estimate because the rate can occasionally change rapidly depending on the markets. The rate is also a 'use-able' rate not a direct market rate. Meaning multiplying your input coin amount times the rate should give you a close approximation of what will be sent out. This rate does not include the transaction (miner) fee taken off every transaction.
```go
pair := Pair{"eth_btc"}

rate := pair.GetRates()

t.Logf("Pair: %v at %v",rate.Pair, rate.Rate)
```

### :white_check_mark: Deposit Limits
Gets the current deposit limit set by Shapeshift. Amounts deposited over this limit will be sent to the return address if one was entered, otherwise the user will need to contact ShapeShift support to retrieve their coins. This is an estimate because a sudden market swing could move the limit.
```go
pair := Pair{"eth_btc"}

limits := pair.GetLimits()

t.Logf("Limits on Pair: %v LIMIT: %v",limits.Pair, limits.Limit)
```

### :white_check_mark: Market Info
This gets the market info (pair, rate, limit, minimum limit, miner fee)
```go
pair := Pair{"btc_eth"}

info := pair.GetInfo()

t.Logf("Info Pair: %v | Min: %f | LIMIT: %f",info.Pair, info.Min, info.Limit)
```

### :white_check_mark: Recent Transactions
```go
recent := RecentTransactions()

t.Log(recent)
```
### :white_check_mark: Deposit Address Status
This returns the status of the most recent deposit transaction to the address.
```go
status := DepositStatus("16FdfRFVPUwiKAceRSqgEfn1tmB4sVUmLh")

t.Log(status.Status)
```

### :white_check_mark: Time Remaining on Fixed Transaction Amount
Get a list of the most recent transactions.
```go
status := TimeRemaining("16FdfRFVPUwiKAceRSqgEfn1tmB4sVUmLh")

t.Log(status.Status)
```
### :white_check_mark: Get Available Coins
Allows anyone to get a list of all the currencies that Shapeshift currently supports at any given time. The list will include the name, symbol, availability status, and an icon link for each.
```go
coins := Coins()
t.Log(coins)
```

### :white_large_square: Validate Address with Coin Symbol
Allows user to verify that their receiving address is a valid address according to a given wallet daemon. If isvalid returns true, this address is valid according to the coin daemon indicated by the currency symbol.

# Primary Requests

### :white_check_mark: Create New Transaction
This is the primary data input into ShapeShift.
```go
new := New{
		Pair: "eth_btc",
		ToAddress: "16FdfRFVPUwiKAceRSqgEfn1tmB4sVUmLh",
		FromAddress: "0xcf2f204aC8D7714990912fA422874371c001217D",
	}

response := new.Shift()

t.Log("Send To Address: ", response.SendTo, "\n")
t.Log("Send Type: ", response.SendType, "\n")
t.Log("Send Type: ", response.ApiKey, "\n")
```

### :white_check_mark: Request Email Receipt
This call requests a receipt for a transaction. The email address will be added to the conduit associated with that transaction as well. (Soon it will also send receipts to subsequent transactions on that conduit)
```go
info := Receipt{
		Email: "user@awesome.com",
		TransactionID: "owkdwodkkwokdwdw",
	}

response := info.Send();

t.Log(response)
```

### :white_check_mark: Fixed Amount Transaction
When a transaction is created with a fixed amount requested there is a 10 minute window for the deposit. After the 10 minute window if the deposit has not been received the transaction expires and a new one must be created. This api call returns how many seconds are left before the transaction expires. Please note that if the address is a ripple address, it will include the "?dt=destTagNUM" appended on the end, and you will need to use the URIEncodeComponent() function on the address before sending it in as a param, to get a successful response.
```go
new := New{
		Pair: "eth_btc",
		Amount: 0.25,
		ToAddress: "16FdfRFVPUwiKAceRSqgEfn1tmB4sVUmLh",
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

### :white_check_mark: Cancel Pending Transaction
This call allows you to request for canceling a pending transaction by the deposit address. If there is fund sent to the deposit address, this pending transaction cannot be canceled.
```go
old := Address{
		Id: newSendToAddress,
	}

	response := old.Cancel()
```


# API Key Required Requests

### :white_large_square: Get Transactions from Private API Key
Allows vendors to get a list of all transactions that have ever been done using a specific API key. Transactions are created with an affilliate PUBLIC KEY, but they are looked up using the linked PRIVATE KEY, to protect the privacy of our affiliates' account details.

### :white_large_square: Get Transactions from Output Address
Allows vendors to get a list of all transactions that have ever been sent to one of their addresses. The affilliate's PRIVATE KEY must be provided, and will only return transactions that were sent to output address AND were created using / linked to the affiliate's PUBLIC KEY. Please note that if the address is a ripple address and it includes the "?dt=destTagNUM" appended on the end, you will need to use the URIEncodeComponent() function on the address before sending it in as a param, to get a successful response.
