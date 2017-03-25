package shapeshift

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
)

var apiUrl string = "https://shapeshift.io"

func ToFloat(s string) float64 {
	f, _ := strconv.ParseFloat(s, 64)
	return f
}

type Pair struct {
	Name string `json:"pair,omitempty"`
}

type RateResponse struct {
	Pair string `json:"pair,omitempty"`
	Rate string `json:"rate"`
	Error
}

type LimitResponse struct {
	Pair  string `json:"pair,omitempty"`
	Limit string `json:"limit"`
	Error
}

type MarketInfoResponse struct {
	Pair     string  `json:"pair,omitempty"`
	Rate     float64 `json:"rate,omitempty"`
	Limit    float64 `json:"limit,omitempty"`
	Min      float64 `json:"min,omitempty"`
	MinerFee float64 `json:"minerFee,omitempty"`
	Error
}

type RecentTranxResponse []struct {
	CurIn     string  `json:"curIn"`
	CurOut    string  `json:"curOut"`
	Timestamp float64 `json:"timestamp"`
	Amount    float64 `json:"amount"`
	Error
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
	Error
}

type Receipt struct {
	Email         string `json:"email"`
	TransactionID string `json:"txid"`
}

type ValidateResponse struct {
	Valid bool `json:"isValid"`
	Error
}

type CancelResponse struct {
	Success string `json:"success,omitempty"`
	Error
}

type Address struct {
	Id string `json:"address"`
}

type New struct {
	Pair        string  `json:"pair,omitempty"`
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
	Error
}

type FixedTransactionResponse struct {
	Response NewFixedTransactionResponse `json:"success"`
	Error
}

type NewFixedTransactionResponse struct {
	OrderID          string  `json:"orderId"`
	Pair             string  `json:"pair,omitempty"`
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
	Error
}

type ListTransactionsAPIResponse struct {
	Transactions []Transaction
	Error
}

type ErrorMsg interface {
	ErrorMsg() string
	isOk() bool
}

func (e Error) ErrorMsg() string {
	return e.Message
}

func (e Error) isOk() bool {
	if e.Message == "" {
		return true
	}
	return false
}

type Error struct {
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
	Error
}

func (p Pair) GetRates() (float64, error) {
	r, err := DoHttp("GET", "rate", p.Name)
	if err != nil {
		panic(err)
	}
	var g RateResponse
	json.Unmarshal(r, &g)
	return ToFloat(g.Rate), err
}

func (p Pair) GetLimits() (float64, error) {
	r, err := DoHttp("GET", "limit", p.Name)
	if err != nil {
		panic(err)
	}
	var g LimitResponse
	json.Unmarshal(r, &g)
	return ToFloat(g.Limit), err
}

func (p Pair) GetInfo() MarketInfoResponse {
	r, err := DoHttp("GET", "marketinfo", p.Name)
	if err != nil {
		panic(err)
	}
	var g MarketInfoResponse
	json.Unmarshal(r, &g)
	return g
}

func RecentTransactions(count string) RecentTranxResponse {
	r, err := DoHttp("GET", "recenttx", count)
	if err != nil {
		panic(err)
	}
	var g RecentTranxResponse
	json.Unmarshal(r, &g)
	return g
}

func DepositStatus(addr string) DepositStatusResponse {
	r, err := DoHttp("GET", "txStat", addr)
	if err != nil {
		panic(err)
	}
	var g DepositStatusResponse
	json.Unmarshal(r, &g)
	return g
}

func TimeRemaining(addr string) TimeRemainingResponse {
	r, err := DoHttp("GET", "timeremaining", addr)
	if err != nil {
		panic(err)
	}
	var g TimeRemainingResponse
	json.Unmarshal(r, &g)
	return g
}

type ReceiptResponse struct {
	Email struct {
		Status  string `json:"status"`
		Message string `json:"message"`
	} `json:"email"`
	Error
}

func Coins() CoinsResponse {
	r, err := DoHttp("GET", "getcoins", "")
	if err != nil {
		panic(err)
	}
	var g CoinsResponse
	json.Unmarshal(r, &g)
	return g
}

func (r Receipt) Send() ReceiptResponse {
	q, err := DoPostHttp("POST", "mail", r)
	if err != nil {
		panic(err)
	}
	var g ReceiptResponse
	json.Unmarshal(q, &g)
	return g
}

func (n New) Shift() NewTransactionResponse {
	r, err := DoPostHttp("POST", "shift", n)
	if err != nil {
		panic(err)
	}
	var g NewTransactionResponse
	json.Unmarshal(r, &g)
	return g
}

func (n New) FixedShift() NewFixedTransactionResponse {
	r, err := DoPostHttp("POST", "sendamount", n)
	if err != nil {
		panic(err)
	}
	var g FixedTransactionResponse
	json.Unmarshal(r, &g)
	return g.Response
}

func (n Address) Cancel() CancelResponse {
	r, err := DoPostHttp("POST", "cancelpending", n)
	if err != nil {
		panic(err)
	}
	var g CancelResponse
	json.Unmarshal(r, &g)
	return g
}

func Validate(addr string, coin string) ValidateResponse {
	r, err := DoHttp("GET", "validateAddress/"+addr, coin)
	if err != nil {
		panic(err)
	}
	var g ValidateResponse
	json.Unmarshal(r, &g)
	return g
}

func (i API) ListTransactions() ListTransactionsAPIResponse {
	var r []byte
	//var err error
	var g ListTransactionsAPIResponse
	if i.Address != "" {
		r, _ = DoHttp("GET", "txbyaddress/"+i.Address, i.Key)
	} else {
		r, _ = DoHttp("GET", "txbyapikey", i.Key)
	}
	json.Unmarshal(r, &g)
	return g
}

func DoPostHttp(method string, apimethod string, data interface{}) ([]byte, error) {
	new, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	req, err := http.NewRequest(method, apiUrl+"/"+apimethod, bytes.NewBuffer(new))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	body = bytes.TrimPrefix(body, []byte("\xef\xbb\xbf"))
	return body, err
}

func DoHttp(method string, apimethod string, url string) ([]byte, error) {
	req, err := http.NewRequest(method, apiUrl+"/"+apimethod+"/"+url, bytes.NewBuffer([]byte("")))
	if err != nil {
		panic(err)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	body = bytes.TrimPrefix(body, []byte("\xef\xbb\xbf"))
	return body, err
}
