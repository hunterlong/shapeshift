package shapeshift

import (
	"net/http"
	"bytes"
	"io/ioutil"
	"encoding/json"
	"fmt"
)

var apiUrl string = "https://shapeshift.io";


type Pair struct {
	Name string
}

type RateResponse struct {
	Pair string `json:"pair"`
	Rate string `json:"rate"`
}

type LimitResponse struct {
	Pair string `json:"pair"`
	Limit string `json:"limit"`
}

type MarketInfoResponse struct {
	Pair string `json:"pair"`
	Rate float64 `json:"rate,omitempty"`
	Limit float64 `json:"limit,omitempty"`
	Min float64 `json:"min,omitempty"`
	MinerFee float64 `json:"minerFee,omitempty"`
}


type RecentTranxResponse []struct {
	CurIn string `json:"curIn"`
	CurOut string `json:"curOut"`
	Timestamp float64 `json:"timestamp"`
	Amount float64 `json:"amount"`
}


type DepositStatusResponse struct {
	Status string `json:"status"`
	Address string `json:"address"`
	Withdraw string `json:"withdraw,omitempty"`
	IncomingCoin string `json:"incomingCoin,omitempty"`
	IncomingType string `json:"incomingType,omitempty"`
	OutgoingCoin string `json:"outgoingCoin,omitempty"`
	OutgoingType string `json:"outgoingType,omitempty"`
	Transaction string `json:"transaction,omitempty"`
	Error string `json:"error,omitempty"`
}

type CoinsResponse struct {

}

type TimeRemainingResponse struct {
	Status string `json:"status"`
	Seconds int `json:"seconds_remaining"`
}

func (p Pair) GetRates() RateResponse {
	r := DoHttp("GET", "rate", p.Name)
	var g RateResponse
	json.Unmarshal(r, &g)
	return g
}

func (p Pair) GetLimits() LimitResponse {
	r := DoHttp("GET", "limit", p.Name)
	var g LimitResponse
	json.Unmarshal(r, &g)
	return g
}

func (p Pair) GetInfo() MarketInfoResponse {
	r := DoHttp("GET", "marketinfo", p.Name)
	var g MarketInfoResponse
	json.Unmarshal(r, &g)
	return g
}

func RecentTransactions() RecentTranxResponse {
	r := DoHttp("GET", "recenttx", "5")
	var g RecentTranxResponse
	json.Unmarshal(r, &g)
	return g
}

func DepositStatus(addr string) DepositStatusResponse {
	r := DoHttp("GET", "txStat", addr)
	var g DepositStatusResponse
	json.Unmarshal(r, &g)
	return g
}

func TimeRemaining(addr string) TimeRemainingResponse {
	r := DoHttp("GET", "timeremaining", addr)
	var g TimeRemainingResponse
	json.Unmarshal(r, &g)
	return g
}


func Coins() CoinsResponse {
	r := DoHttp("GET", "getcoins", "")
	fmt.Println(string(r))
	var g CoinsResponse
	json.Unmarshal(r, &g)
	return g
}


type New struct {
	Pair string `json:"pair"`
	ToAddress string `json:"withdrawal"`
	FromAddress string `json:"returnAddress,omitempty"`
	DestTag string `json:"destTag,omitempty"`
	rsAddress string `json:"rsAddress,omitempty"`
	ApiKey string `json:"apiKey,omitempty"`
}

type NewTransactionResponse struct {
	SendTo string `json:"deposit"`
	SendType string `json:"depositType"`
	ReturnTo string `json:"withdrawal"`
	ReturnType string `json:"withdrawalType"`
	Public string `json:"public"`
	XrpDestTag string `json:"xrpDestTag"`
	ApiKey string `json:"apiPubKey"`
}

func (n New) Shift() NewTransactionResponse {
	r := DoPostHttp("POST", "shift", n)
	fmt.Println(string(r))
	var g NewTransactionResponse
	json.Unmarshal(r, &g)
	return g
}


func DoPostHttp(method string, apimethod string, data New) []byte {
	new, _ := json.Marshal(data)
	fmt.Println("Sending ", string(new))
	req, err := http.NewRequest(method, apiUrl+"/"+apimethod, bytes.NewBuffer(new))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	body = bytes.TrimPrefix(body, []byte("\xef\xbb\xbf"))
	return body
}


func DoHttp(method string, apimethod string, url string) []byte {
	req, err := http.NewRequest(method, apiUrl+"/"+apimethod+"/"+url, bytes.NewBuffer([]byte("")))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	body = bytes.TrimPrefix(body, []byte("\xef\xbb\xbf"))
	return body
}
