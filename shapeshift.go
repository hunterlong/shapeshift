package shapeshift

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var apiUrl string = "https://shapeshift.io"

type Pair struct {
	Name string
}

type RateResponse struct {
	Pair string `json:"pair"`
	Rate string `json:"rate"`
}

type LimitResponse struct {
	Pair  string `json:"pair"`
	Limit string `json:"limit"`
}

type MarketInfoResponse struct {
	Pair     string  `json:"pair"`
	Rate     float64 `json:"rate,omitempty"`
	Limit    float64 `json:"limit,omitempty"`
	Min      float64 `json:"min,omitempty"`
	MinerFee float64 `json:"minerFee,omitempty"`
}

type RecentTranxResponse []struct {
	CurIn     string  `json:"curIn"`
	CurOut    string  `json:"curOut"`
	Timestamp float64 `json:"timestamp"`
	Amount    float64 `json:"amount"`
}

type DepositStatusResponse struct {
	Status       string  `json:"status"`
	Address      string  `json:"address"`
	Withdraw     string  `json:"withdraw,omitempty"`
	IncomingCoin float64 `json:"incomingCoin,omitempty"`
	IncomingType string  `json:"incomingType,omitempty"`
	OutgoingCoin string  `json:"outgoingCoin,omitempty"`
	OutgoingType string  `json:"outgoingType,omitempty"`
	Transaction  string  `json:"transaction,omitempty"`
	Error        string  `json:"error,omitempty"`
}

type Receipt struct {
	Email         string `json:"email"`
	TransactionID string `json:"txid"`
}

type ValidateResponse struct {
	Valid bool   `json:"isValid"`
	Error string `json:"error"`
}

type CancelResponse struct {
	Success string `json:"success,omitempty"`
	Error   string `json:"error,omitempty"`
}

type Address struct {
	Id string `json:"address"`
}

type New struct {
	Pair        string  `json:"pair"`
	ToAddress   string  `json:"withdrawal"`
	FromAddress string  `json:"returnAddress,omitempty"`
	DestTag     string  `json:"destTag,omitempty"`
	RsAddress   string  `json:"rsAddress,omitempty"`
	ApiKey      string  `json:"apiKey,omitempty"`
	Amount      float64 `json:"amount,omitempty"`
}

type NewTransactionResponse struct {
	SendTo     string `json:"deposit"`
	SendType   string `json:"depositType"`
	ReturnTo   string `json:"withdrawal"`
	ReturnType string `json:"withdrawalType"`
	Public     string `json:"public"`
	XrpDestTag string `json:"xrpDestTag"`
	ApiKey     string `json:"apiPubKey"`
}

type FixedTransactionResponse struct {
	Response NewFixedTransactionResponse `json:"success"`
}

type NewFixedTransactionResponse struct {
	OrderID          string  `json:"orderId"`
	Pair             string  `json:"pair"`
	Withdrawal       string  `json:"withdrawal"`
	WithdrawalAmount string  `json:"withdrawalAmount"`
	Deposit          string  `json:"deposit"`
	DepositAmount    string  `json:"depositAmount"`
	Expiration       int64   `json:"expiration"`
	QuotedRate       string  `json:"quotedRate"`
	MaxLimit         float64 `json:"maxLimit"`
	ReturnAddress    string  `json:"returnAddress"`
	APIPubKey        string  `json:"apiPubKey"`
	MinerFee         string  `json:"minerFee"`
}

type ListTransactionsAPIResponse struct {
	Transactions []Transaction
	ErrorResponse
}

type ErrorResponse struct {
	Message string `json:"error,omitempty"`
}

type Transaction struct {
	InputTXID      string  `json:"inputTXID"`
	InputAddress   string  `json:"inputAddress"`
	InputCurrency  string  `json:"inputCurrency,omitempty"`
	InputAmount    float64 `json:"inputAmount,omitempty"`
	OutputTXID     string  `json:"outputTXID,omitempty"`
	OutputAddress  string  `json:"outputAddress,omitempty"`
	OutputCurrency string  `json:"outputCurrency,omitempty"`
	OutputAmount   string  `json:"outputAmount,omitempty"`
	ShiftRate      string  `json:"shiftRate,omitempty"`
	Status         string  `json:"status,omitempty"`
}

type API struct {
	Key     string
	Address string
}

type TimeRemainingResponse struct {
	Status  string `json:"status"`
	Seconds int    `json:"seconds_remaining"`
}

func (p Pair) GetRates() string {
	r := DoHttp("GET", "rate", p.Name)
	var g RateResponse
	json.Unmarshal(r, &g)
	return g.Rate
}

func (p Pair) GetLimits() string {
	r := DoHttp("GET", "limit", p.Name)
	var g LimitResponse
	json.Unmarshal(r, &g)
	return g.Limit
}

func (p Pair) GetInfo() MarketInfoResponse {
	r := DoHttp("GET", "marketinfo", p.Name)
	var g MarketInfoResponse
	json.Unmarshal(r, &g)
	return g
}

func RecentTransactions(count string) RecentTranxResponse {
	r := DoHttp("GET", "recenttx", count)
	fmt.Println(string(r))
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

type ReceiptResponse struct {
	Email struct {
		Status  string `json:"status"`
		Message string `json:"message"`
	} `json:"email"`
}

func Coins() CoinsResponse {
	r := DoHttp("GET", "getcoins", "")
	fmt.Println(string(r))
	var g CoinsResponse
	json.Unmarshal(r, &g)
	return g
}

func (r Receipt) Send() ReceiptResponse {
	q := DoPostHttp("POST", "mail", r)
	var g ReceiptResponse
	json.Unmarshal(q, &g)
	return g
}

func (n New) Shift() NewTransactionResponse {
	r := DoPostHttp("POST", "shift", n)
	var g NewTransactionResponse
	json.Unmarshal(r, &g)
	return g
}

func (n New) FixedShift() NewFixedTransactionResponse {
	r := DoPostHttp("POST", "sendamount", n)
	var g FixedTransactionResponse
	json.Unmarshal(r, &g)
	return g.Response
}

func (n Address) Cancel() CancelResponse {
	r := DoPostHttp("POST", "cancelpending", n)
	var g CancelResponse
	json.Unmarshal(r, &g)
	return g
}

func Validate(addr string, coin string) ValidateResponse {
	r := DoHttp("GET", "validateAddress/"+addr, coin)
	var g ValidateResponse
	json.Unmarshal(r, &g)
	return g
}

func (i API) ListTransactions() ListTransactionsAPIResponse {
	var r []byte
	var g ListTransactionsAPIResponse
	if i.Address != "" {
		r = DoHttp("GET", "txbyaddress/"+i.Address, i.Key)
	} else {
		r = DoHttp("GET", "txbyapikey", i.Key)
	}
	json.Unmarshal(r, &g)
	return g
}

func DoPostHttp(method string, apimethod string, data interface{}) []byte {
	new, _ := json.Marshal(data)
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
