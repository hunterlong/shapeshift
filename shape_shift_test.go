package shapeshift

import (
	"testing"
)

func TestPairs(t *testing.T) {

	pair := Pair{"btc_eth"}

	rate := pair.GetRates()

	t.Logf("Pair: %v at %v",rate.Pair, rate.Rate)

}

func TestLimits(t *testing.T) {

	pair := Pair{"btc_eth"}

	limits := pair.GetLimits()

	t.Logf("Limits on Pair: %v LIMIT: %v",limits.Pair, limits.Limit)

}

func TestMarketInfo(t *testing.T) {

	pair := Pair{"btc_eth"}

	info := pair.GetInfo()

	t.Logf("Info Pair: %v | Min: %f | LIMIT: %f",info.Pair, info.Min, info.Limit)

}

func TestRecentTransactions(t *testing.T) {

	recent := RecentTransactions()

	t.Log(recent)

}


func TestDepositStatus(t *testing.T) {

	status := DepositStatus("1L75eRMgeCwAxEjD1oWXjLgud9jxwxm34u")

	t.Log(status.Status)

}

func TestTimeRemaining(t *testing.T) {

	status := TimeRemaining("1L75eRMgeCwAxEjD1oWXjLgud9jxwxm34u")

	t.Log(status.Status)

}

func TestGetSupportedCoins(t *testing.T) {

	coins := Coins()
	t.Log(coins)

}

func TestNewTransaction(t *testing.T) {

	new := New{
		Pair: "eth_btc",
		ToAddress: "1L75eRMgeCwAxEjD1oWXjLgud9jxwxm34u",
		FromAddress: "0xcf2f204aC8D7714990912fA422874371c001217D",
	}

	response := new.Shift()

	t.Log("Send To Address: ", response.SendTo, "\n")
	t.Log("Send Type: ", response.SendType, "\n")

}