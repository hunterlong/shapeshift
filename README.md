<p align="center">
  <img src="https://cdn.pbrd.co/images/MK2f4akQc.jpg" alt="shapeshift api golang"/>
</p>

# ShapeShift in Go Language
[![Build Status](https://travis-ci.org/hunterlong/shapeshift.svg?branch=master)](https://travis-ci.org/hunterlong/shapeshift)  [![Coverage Status](https://coveralls.io/repos/github/hunterlong/shapeshift/badge.svg?branch=master)](https://coveralls.io/github/hunterlong/shapeshift?branch=master) [![GoDoc](https://godoc.org/github.com/hunterlong/shapeshift?status.svg)](https://godoc.org/github.com/hunterlong/shapeshift) [![Go Report Card](https://goreportcard.com/badge/github.com/hunterlong/shapeshift)](https://goreportcard.com/report/github.com/hunterlong/shapeshift)

This Go Language Package will allow you to use the [ShapeShift API](https://info.shapeshift.io/) and convert your cryptocurrencies in your very own application. It includes most of the ShapeShift API requests listed on their references website. Below you'll find a perfect example of a new ShapeShift transaction.

```go
go get -u github.com/hunterlong/shapeshift
```
###### get the most up to date version
```go
import "github.com/hunterlong/shapeshift"
```
Once you've imported shapeshift into your golang project, you can use any of the requests below. Checkout the Travis CI test logs for responses of each function. See an issue? PR it!

# :new: New ShapeShift Transaction
I want to convert Ethereum to Bitcoin. The 'ToAddress' is my Bitcoin address. Once I run this, I'll get a Ethereum address from ShapeShift.
```go
new := shapeshift.New{
	Pair: "eth_btc",
	ToAddress: "16FdfRFVPUwiKAceRSqgEfn1tmB4sVUmLh",
	// FromAddress: "0xcf2f204aC8D7714990912fA422874371c001217D",  (Optional Return To Ethereum Address)
       }

response := new.Shift()

if response.isOk() {

    sendToAddress := response.SendTo
    // i will send Ether to this address

    fmt.Println("Send To Address: ", sendToAddress)
    fmt.Println("Send Type: ", response.SendType)
    fmt.Println("Receiving at Address: ", response.ReturnTo)
    fmt.Println("Receiving Type: ", response.ReturnType)
    fmt.Println("Send Type: ", response.SendType)
    fmt.Println("API Key: ", response.ApiKey)
    fmt.Println("Public Data: ", response.Public)
    fmt.Println("XrpDestTag: ", response.XrpDestTag)

} else {
    fmt.Println(response.ErrorMsg())
}
```

# :repeat: Get Status of Transaction
Once I sent some Ethereum to the given Ethereum address, I want to check the status of my ShapeShift transaction by inserting the Etheruem address 'sendToAddress' that ShapeShift gave me in previous function.
```go
var newTransactionId string

status := shapeshift.DepositStatus(sendToAddress)

if !response.isOk() {
    fmt.Println(status.ErrorMsg())
}

fmt.Println(status.Status)
// no_deposits
// received
// complete
// failed

if status.Status == "complete" {
	fmt.Println("Incoming Coin: ", status.IncomingCoin)
	fmt.Println("Incoming Type: ", status.IncomingType)
	fmt.Println("Outgoing Coin: ", status.OutgoingCoin)
	fmt.Println("Outgoing Type: ", status.OutgoingType)
	fmt.Println("Address: ", status.Address)
	fmt.Println("Transaction ID: ", status.Transaction)
	fmt.Println("Withdraw: ", status.Withdraw)
	
	newTransactionId = status.Transaction
	// saving transaction ID so i can send a receipt
}
```

# :arrow_double_up: Send an Email Receipt
Want to send a receipt of this transaction? Just include an email address and the transaction ID affiliated with the ShapeShift transaction. 
```go
receipt := shapeshift.Receipt{
	Email:         "user@myemailer.com",
	TransactionID: newTransactionId,
     }

response := receipt.Send()

if response.isOk() {
    fmt.Println("Receipt was sent to user")
} else {
    fmt.Println(status.ErrorMsg())
}

```

# Additional Functions
The other ShapeShift API requests are available for you to use. 

### :white_check_mark: Get Rate
Gets the current rate offered by Shapeshift. This is an estimate because the rate can occasionally change rapidly depending on the markets. The rate is also a 'use-able' rate not a direct market rate. Meaning multiplying your input coin amount times the rate should give you a close approximation of what will be sent out. This rate does not include the transaction (miner) fee taken off every transaction.
```go
pair := shapeshift.Pair{"eth_btc"}
rate := pair.GetRates()

fmt.Println("Rate: ", rate)
```

### :white_check_mark: Deposit Limits
Gets the current deposit limit set by Shapeshift. Amounts deposited over this limit will be sent to the return address if one was entered, otherwise the user will need to contact ShapeShift support to retrieve their coins. This is an estimate because a sudden market swing could move the limit.
```go
pair := shapeshift.Pair{"eth_btc"}
limits := pair.GetLimits()

fmt.Println("Limit: ", limits)
```

### :white_check_mark: Market Info
This gets the market info (pair, rate, limit, minimum limit, miner fee)
```go
pair := shapeshift.Pair{"btc_eth"}
info := pair.GetInfo()

fmt.Println("Pair: ", info.Pair)
fmt.Println("Min: ", info.Min)
fmt.Println("Miner Fee: ", info.MinerFee)
fmt.Println("Limit: ", info.Limit)
fmt.Println("Rate: ", info.Rate)
```

### :white_check_mark: Recent Transactions
```go
recent := shapeshift.RecentTransactions("5")

for _, v := range recent {
    fmt.Println("In: ", v.CurIn)
    fmt.Println("Out: ", v.CurOut)
    fmt.Println("Amount: ", v.Amount)
    fmt.Println("Timestamp: ", v.Timestamp)
    fmt.Println("-------------------------------")
}
```
### :white_check_mark: Deposit Address Status
This returns the status of the most recent deposit transaction to the address.
```go
status := shapeshift.DepositStatus("1JP7QWC9GbpKEHSvefygWk5woFy9xeQHKc")

fmt.Println("Deposit Status: ", status.Status)
fmt.Println("Incoming Coin: ", status.IncomingCoin)
fmt.Println("Incoming Type: ", status.IncomingType)
fmt.Println("Outgoing Coin: ", status.OutgoingCoin)
fmt.Println("Outgoing Type: ", status.OutgoingType)
fmt.Println("Address: ", status.Address)
fmt.Println("Transaction ID: ", status.Transaction)
fmt.Println("Withdraw: ", status.Withdraw)
```

### :white_check_mark: Time Remaining on Fixed Transaction Amount
Get a list of the most recent transactions.
```go
status := shapeshift.TimeRemaining("16FdfRFVPUwiKAceRSqgEfn1tmB4sVUmLh")

fmt.Println(status.Status)
```
### :white_check_mark: Get Available Coins
Allows anyone to get a list of all the currencies that Shapeshift currently supports at any given time. The list will include the name, symbol, availability status, and an icon link for each.
```go
coins := shapeshift.Coins()
eth := coins.ETH
fmt.Println("Coin: ", eth.Name)
fmt.Println("Status: ", eth.Status)
```

### :white_check_mark: Validate Address with Coin Symbol
Allows user to verify that their receiving address is a valid address according to a given wallet daemon. If isvalid returns true, this address is valid according to the coin daemon indicated by the currency symbol.
```go
address := shapeshift.Validate("16FdfRFVPUwiKAceRSqgEfn1tmB4sVUmLh", "btc")

fmt.Println("Address is: ", address.Valid)
fmt.Println("Error: ",address.Error)
```
# Primary Requests

### :white_check_mark: Create New Transaction
This is the primary data input into ShapeShift.
```go
new := shapeshift.New{
		Pair:        "eth_btc",
		ToAddress:   "16FdfRFVPUwiKAceRSqgEfn1tmB4sVUmLh",
		FromAddress: "0xcf2f204aC8D7714990912fA422874371c001217D",
	}

response := new.Shift()

fmt.Println("Send To Address: ", response.SendTo)
fmt.Println("Send Type: ", response.SendType)
fmt.Println("Receiving at Address: ", response.ReturnTo)
fmt.Println("Receiving Type: ", response.ReturnType)
fmt.Println("Send Type: ", response.SendType)
fmt.Println("API Key: ", response.ApiKey)
fmt.Println("Public Data: ", response.Public)
fmt.Println("XrpDestTag: ", response.XrpDestTag)
```

### :white_check_mark: Request Email Receipt
This call requests a receipt for a transaction. The email address will be added to the conduit associated with that transaction as well. (Soon it will also send receipts to subsequent transactions on that conduit)
```go
info := shapeshift.Receipt{
		Email: "user@awesome.com",
		TransactionID: "owkdwodkkwokdwdw",
	}

response := info.Send();

fmt.Println(response)
```

### :white_check_mark: Fixed Amount Transaction
When a transaction is created with a fixed amount requested there is a 10 minute window for the deposit. After the 10 minute window if the deposit has not been received the transaction expires and a new one must be created. This api call returns how many seconds are left before the transaction expires. Please note that if the address is a ripple address, it will include the "?dt=destTagNUM" appended on the end, and you will need to use the URIEncodeComponent() function on the address before sending it in as a param, to get a successful response.
```go
new := shapeshift.New{
		Pair: "eth_btc",
		Amount: 0.25,
		ToAddress: "16FdfRFVPUwiKAceRSqgEfn1tmB4sVUmLh",
		FromAddress: "0xcf2f204aC8D7714990912fA422874371c001217D",
	}

response := new.FixedShift()

fmt.Println("Pair: ", response.Pair)
fmt.Println("Quoted Rate: ", response.QuotedRate)
fmt.Println("Deposit Address: ", response.Deposit)
fmt.Println("Deposit Amount: ", response.DepositAmount)
fmt.Println("Withdraw Amount: ", response.WithdrawalAmount)
fmt.Println("Withdraw Address: ", response.Withdrawal)
fmt.Println("Expiration: ", response.Expiration)
```

### :white_check_mark: Cancel Pending Transaction
This call allows you to request for canceling a pending transaction by the deposit address. If there is fund sent to the deposit address, this pending transaction cannot be canceled.
```go
old := shapeshift.Address{
		Id: newSendToAddress,
	}

	response := old.Cancel()
```


# API Key Required Requests

### :white_check_mark: Get Transactions from Private API Key
Allows vendors to get a list of all transactions that have ever been done using a specific API key. Transactions are created with an affilliate PUBLIC KEY, but they are looked up using the linked PRIVATE KEY, to protect the privacy of our affiliates' account details.
```go
api := shapeshift.API{
		Key: "oskdfoijsfuhsdhufhewhuf",
	   }

list := api.ListTransactions()

for _,v := range list.Transactions {
    fmt.Println("Input: ",v.InputAddress)
    fmt.Println("Amount: ",v.InputAmount)
}
```
###### there was no way for me to test this transaction since i'm not a vendor

### :white_check_mark: Get Transactions from Output Address
Allows vendors to get a list of all transactions that have ever been sent to one of their addresses. The affilliate's PRIVATE KEY must be provided, and will only return transactions that were sent to output address AND were created using / linked to the affiliate's PUBLIC KEY. Please note that if the address is a ripple address and it includes the "?dt=destTagNUM" appended on the end, you will need to use the URIEncodeComponent() function on the address before sending it in as a param, to get a successful response.
```go
api := shapeshift.API{
		Key: "oskdfoijsfuhsdhufhewhuf",
		Address: "1JP7QWC9GbpKEHSvefygWk5woFy9xeQHKc",
	}

list := api.ListTransactions()

for _,v := range list.Transactions {
    fmt.Println("Input: ",v.InputAddress)
    fmt.Println("Amount: ",v.InputAmount)
}
```
###### there was no way for me to test this transaction since i'm not a vendor

# Coin Pairs
Many of the requests require a 'coin pair'. A coin pair is of the format deposit_withdrawal. Example: 'btc_ltc'. Valid pairs are any combination of the below listed valid coins.* The list will grow as we add more:
```
btc, ltc, ppc, drk, doge, nmc, ftc, blk, nxt, btcd, qrk, rdd, nbt, bts, bitusd, xcp, xmr
```
[ShapeShift Coins](https://shapeshift.io/#/coins)
* If a particular coin goes offline any pairs using it will return a message stating that pair is temporarily unavailable.

All requests are only available via HTTPS, in the interest of security best practices we do not support API calls over HTTP.

### Package Useful? :beer: :bug:
If this package saved you some time, or if you're excited to make that next crypto-bot, feel free to throw some coins my way. If you see an issue with this golang package, please submit a pull request. 
```
ETH: 0x9741C5522B85E195B92C71CE29B54A4C99D76c13
BTC: 16FdfRFVPUwiKAceRSqgEfn1tmB4sVUmLh
```

# License
This golang package is built for the cryptocurrency community and is released with MIT license. 

:thumbsup: :thumbsup: [ShapeShift.io](https://shapeshift.io) :thumbsup: :thumbsup:
