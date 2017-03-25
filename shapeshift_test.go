package shapeshift

import (
	"testing"
)

var newSendToAddress string
var newSendToAddress2 string
var newTranxId string

func TestPairs(t *testing.T) {

	pair := Pair{"eth_btc"}

	rate, err := pair.GetRates()
	if err != nil {
		t.Fail()
	}

	t.Log("Rate: ", rate)

}

func TestErrorPairs(t *testing.T) {

	pair := Pair{"xxx_btc"}

	rate, err := pair.GetRates()
	if err != nil {
		t.Fail()
	}

	t.Log("Rate: ", rate)

}

func TestLimits(t *testing.T) {

	pair := Pair{"eth_btc"}

	limits, err := pair.GetLimits()

	if err != nil {
		t.Fail()
	}

	t.Log("Limit: ", limits)

}

func TestMarketInfo(t *testing.T) {

	pair := Pair{"btc_eth"}

	info := pair.GetInfo()

	if !info.isOk() {
		t.Log(info.ErrorMsg())
	}

	t.Log("Pair: ", info.Pair)
	t.Log("Min: ", info.Min)
	t.Log("Miner Fee: ", info.MinerFee)
	t.Log("Limit: ", info.Limit)
	t.Log("Rate: ", info.Rate)

}

func TestRecentTransactions(t *testing.T) {

	recent := RecentTransactions("5")

	for _, v := range recent {
		t.Log("In: ", v.CurIn)
		t.Log("Out: ", v.CurOut)
		t.Log("Amount: ", v.Amount)
		t.Log("Timestamp: ", v.Timestamp)
		t.Log("-------------------------------")
	}

}

func TestValidateAddress(t *testing.T) {

	address := Validate("16FdfRFVPUwiKAceRSqgEfn1tmB4sVUmLh", "btc")
	t.Log("Address is: ", address.Valid)

	address2 := Validate("1JP7QWC9GbpKEHSvefygWk5woFy9xeQHKc", "btc")
	t.Log("Second Address is: ", address2.Valid)
	t.Log("Second Error: ", address2.ErrorMsg())

}

func TestDepositStatus(t *testing.T) {

	status := DepositStatus("1JP7QWC9GbpKEHSvefygWk5woFy9xeQHKc")

	if status.isOk() {

		t.Log("Deposit Status: ", status.Status)

		t.Log("Incoming Coin: ", status.IncomingCoin)
		t.Log("Incoming Type: ", status.IncomingType)
		t.Log("Outgoing Coin: ", status.OutgoingCoin)
		t.Log("Outgoing Type: ", status.OutgoingType)
		t.Log("Address: ", status.Address)
		t.Log("Transaction ID: ", status.Transaction)
		t.Log("Withdraw: ", status.Withdraw)

	}

	if status.Status != "complete" {
		t.Fail()
	}

	newTranxId = status.Transaction

}

func TestGetSupportedCoins(t *testing.T) {

	coins := Coins()
	eth := coins.ETH
	t.Log("Coin: ", eth.Name)
	t.Log("Status: ", eth.Status)

}

func TestNewTransaction(t *testing.T) {

	new := New{
		Pair:        "eth_btc",
		ToAddress:   "16FdfRFVPUwiKAceRSqgEfn1tmB4sVUmLh",
		FromAddress: "0xcf2f204aC8D7714990912fA422874371c001217D",
	}

	response := new.Shift()

	if response.isOk() {

		t.Log("Send To Address: ", response.SendTo)
		t.Log("Send Type: ", response.SendType)
		t.Log("Receiving at Address: ", response.ReturnTo)
		t.Log("Receiving Type: ", response.ReturnType)
		t.Log("Send Type: ", response.SendType)
		t.Log("API Key: ", response.ApiKey)
		t.Log("Public Data: ", response.Public)
		t.Log("XrpDestTag: ", response.XrpDestTag)

		if response.SendType != "ETH" || response.ReturnType != "BTC" {
			t.Fail()
		}

	}

	newSendToAddress = response.SendTo

}

func TestEmailReceipt(t *testing.T) {

	info := Receipt{
		Email:         "info@socialeck.com",
		TransactionID: newTranxId,
	}

	response := info.Send()

	if response.isOk() {
		t.Log("Response was good!")
	}

	t.Log(response)

}

func TestNewFixedTransaction(t *testing.T) {

	new := New{
		Pair:        "eth_btc",
		Amount:      0.25,
		ToAddress:   "16FdfRFVPUwiKAceRSqgEfn1tmB4sVUmLh",
		FromAddress: "0xcf2f204aC8D7714990912fA422874371c001217D",
	}

	response := new.FixedShift()

	if response.isOk() {

		t.Log("Pair: ", response.Pair)
		t.Log("Quoted Rate: ", response.QuotedRate)
		t.Log("Deposit Address: ", response.Deposit)
		t.Log("Deposit Amount: ", response.DepositAmount)
		t.Log("Withdraw Amount: ", response.WithdrawalAmount)
		t.Log("Withdraw Address: ", response.Withdrawal)
		t.Log("Expiration: ", response.Expiration)

	} else {
		t.Log(response.ErrorMsg())
	}

	newSendToAddress2 = response.Deposit

	if response.Withdrawal != "16FdfRFVPUwiKAceRSqgEfn1tmB4sVUmLh" {
		t.Fail()
	}

	if response.WithdrawalAmount != "0.25" {
		t.Fail()
	}

	if response.Pair != "eth_btc" {
		t.Fail()
	}
}

func TestTimeRemaining(t *testing.T) {

	status := TimeRemaining("1JP7QWC9GbpKEHSvefygWk5woFy9xeQHKc")

	if status.isOk() {
		t.Log("Seconds Remaining: ", status.Seconds)
	} else {
		t.Log(status.ErrorMsg())
	}

}

func TestCancelTransaction(t *testing.T) {

	old := Address{
		Id: newSendToAddress,
	}

	response := old.Cancel()

	if response.isOk() {
		t.Log(response.Success)
	} else {
		t.Log(response.ErrorMsg())
	}

}

func TestListTransactionsFromAPI(t *testing.T) {

	api := API{
		Key: "oskdfoijsfuhsdhufhewhuf",
	}

	list := api.ListTransactions()

	if !list.isOk() {
		t.Log(list.ErrorMsg())
	}

	for _, v := range list.Transactions {
		t.Log("Input: ", v.InputAddress)
		t.Log("Amount: ", v.InputAmount)
	}

	t.Log(list)

}

func TestListAddressTransactionsFromAPI(t *testing.T) {

	api := API{
		Key:     "oskdfoijsfuhsdhufhewhuf",
		Address: "1JP7QWC9GbpKEHSvefygWk5woFy9xeQHKc",
	}

	list := api.ListTransactions()

	if !list.isOk() {
		t.Log(list.ErrorMsg())
	}

	for _, v := range list.Transactions {
		t.Log("Input: ", v.InputAddress)
		t.Log("Amount: ", v.InputAmount)
	}

	t.Log(list)
}
