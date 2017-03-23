package shapeshift

import (
	"testing"
)

var newSendToAddress string
var newSendToAddress2 string

func TestPairs(t *testing.T) {

	pair := Pair{"eth_btc"}

	rate := pair.GetRates()

	t.Logf("Pair: %v at %v", rate.Pair, rate.Rate)

}

func TestLimits(t *testing.T) {

	pair := Pair{"eth_btc"}

	limits := pair.GetLimits()

	t.Logf("Limits on Pair: %v LIMIT: %v", limits.Pair, limits.Limit)

}

func TestMarketInfo(t *testing.T) {

	pair := Pair{"btc_eth"}

	info := pair.GetInfo()

	t.Logf("Info Pair: %v | Min: %f | LIMIT: %f", info.Pair, info.Min, info.Limit)

}

func TestRecentTransactions(t *testing.T) {

	recent := RecentTransactions()

	t.Log(recent)

}

func TestDepositStatus(t *testing.T) {

	status := DepositStatus("1L75eRMgeCwAxEjD1oWXjLgud9jxwxm34u")

	t.Log(status.Status)

}

func TestGetSupportedCoins(t *testing.T) {

	coins := Coins()
	t.Log(coins)

}

func TestNewTransaction(t *testing.T) {

	new := New{
		Pair:        "eth_btc",
		ToAddress:   "1L75eRMgeCwAxEjD1oWXjLgud9jxwxm34u",
		FromAddress: "0xcf2f204aC8D7714990912fA422874371c001217D",
	}

	response := new.Shift()

	t.Log("Send To Address: ", response.SendTo, "\n")
	t.Log("Receiving at Address: ", response.ReturnTo, "\n")
	t.Log("Receiving Type: ", response.ReturnType, "\n")
	t.Log("Send Type: ", response.SendType, "\n")
	t.Log("Send Type: ", response.ApiKey, "\n")

	newSendToAddress = response.SendTo

}

func TestEmailReceipt(t *testing.T) {

	info := Receipt{
		Email:         "user@awesome.com",
		TransactionID: "owkdwodkkwokdwdw",
	}

	response := info.Send()

	t.Log(response)

}

func TestNewFixedTransaction(t *testing.T) {

	new := New{
		Pair:        "eth_btc",
		Amount:      0.25,
		ToAddress:   "1L75eRMgeCwAxEjD1oWXjLgud9jxwxm34u",
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

	newSendToAddress2 = response.Deposit
}

func TestTimeRemaining(t *testing.T) {

	status := TimeRemaining(newSendToAddress2)

	t.Log(status.Status)

}

func TestCancelTransaction(t *testing.T) {

	old := Address{
		Id: newSendToAddress,
	}

	response := old.Cancel()

	t.Log(response.Error)
	t.Log(response.Success)

}
