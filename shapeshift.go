package shapeshift

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

var apiUrl string = "https://shapeshift.io"

// ShapeShift's API responds in float and string for decimals for different functions.
// Since we arn't really using 'big numbers' I think it's ok to be using this.
// This golang package is not doing any math, just responding back from ShapeShift API.
func ToFloat(s string) float64 {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0.0
	}
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

type ReceiptResponse struct {
	Email struct {
		Status  string `json:"status"`
		Message string `json:"message"`
	} `json:"email"`
	Error
}

type API struct {
	Key     string
	Address string
}

type TimeRemainingResponse struct {
	Status  string `json:"status"`
	Seconds string `json:"seconds_remaining"`
	Error
}

func (p Pair) GetRates() (float64, error) {
	r, err := DoHttp("GET", "rate", p.Name)
	if err != nil {
		return 0.0, err
	}
	var g RateResponse
	err = json.Unmarshal(r, &g)
	return ToFloat(g.Rate), err
}

func (p Pair) GetLimits() (float64, error) {
	r, err := DoHttp("GET", "limit", p.Name)
	if err != nil {
		return 0.0, err
	}
	var g LimitResponse
	err = json.Unmarshal(r, &g)
	return ToFloat(g.Limit), err
}

func (p Pair) GetInfo() (*MarketInfoResponse, error) {
	r, err := DoHttp("GET", "marketinfo", p.Name)
	if err != nil {
		return nil, err
	}
	var g MarketInfoResponse
	err = json.Unmarshal(r, &g)
	return &g, err
}

func RecentTransactions(count string) (*RecentTranxResponse, error) {
	r, err := DoHttp("GET", "recenttx", count)
	if err != nil {
		return nil, err
	}
	var g RecentTranxResponse
	err = json.Unmarshal(r, &g)
	return &g, err
}

func DepositStatus(addr string) (*DepositStatusResponse, error) {
	r, err := DoHttp("GET", "txStat", addr)
	if err != nil {
		return nil, err
	}
	var g DepositStatusResponse
	err = json.Unmarshal(r, &g)
	return &g, err
}

func TimeRemaining(addr string) (*TimeRemainingResponse, error) {
	r, err := DoHttp("GET", "timeremaining", addr)
	if err != nil {
		return nil, err
	}
	var g TimeRemainingResponse
	err = json.Unmarshal(r, &g)
	return &g, err
}

func CoinsAsList() ([]Coin, error) {
	var coins []Coin
	r, err := DoHttp("GET", "getcoins", "")
	if err != nil {
		return nil, err
	}

	// User json.RawMessage to delay marshalling to support arbitrary top level keys
	var coinmap map[string]*json.RawMessage
	if err := json.Unmarshal(r, &coinmap); err != nil {
		return coins, err
	}

	for _, coinJSON := range coinmap {
		var c Coin
		err := json.Unmarshal([]byte(*coinJSON), &c)
		if err != nil {
			log.Println("Error unmarshalling coin:", err)
			continue
		}
		coins = append(coins, c)
	}

	return coins, nil
}

func Coins() (*CoinsResponse, error) {
	r, err := DoHttp("GET", "getcoins", "")
	if err != nil {
		return nil, err
	}
	var g CoinsResponse
	err = json.Unmarshal(r, &g)
	return &g, err
}

func (r Receipt) Send() (*ReceiptResponse, error) {
	q, err := DoPostHttp("POST", "mail", r)
	if err != nil {
		return nil, err
	}
	var g ReceiptResponse
	err = json.Unmarshal(q, &g)
	return &g, err
}

func (n New) Shift() (*NewTransactionResponse, error) {
	r, err := DoPostHttp("POST", "shift", n)
	if err != nil {
		return nil, err
	}
	var g NewTransactionResponse
	err = json.Unmarshal(r, &g)
	return &g, err
}

func (n New) FixedShift() (*NewFixedTransactionResponse, error) {
	r, err := DoPostHttp("POST", "sendamount", n)
	if err != nil {
		return nil, err
	}
	var g FixedTransactionResponse
	err = json.Unmarshal(r, &g)
	return &g.Response, err
}

func (n Address) Cancel() (*CancelResponse, error) {
	r, err := DoPostHttp("POST", "cancelpending", n)
	if err != nil {
		return nil, err
	}
	var g CancelResponse
	err = json.Unmarshal(r, &g)
	return &g, err
}

func Validate(addr string, coin string) (*ValidateResponse, error) {
	r, err := DoHttp("GET", "validateAddress/"+addr, coin)
	if err != nil {
		return nil, err
	}
	var g ValidateResponse
	err = json.Unmarshal(r, &g)
	return &g, err
}

func (i API) ListTransactions() ([]Transaction, error) {
	var r []byte
	var err error
	var g []Transaction
	if i.Address != "" {
		r, err = DoHttp("GET", "txbyaddress/"+i.Address, i.Key)
	} else {
		r, err = DoHttp("GET", "txbyapikey", i.Key)
	}
	err = json.Unmarshal(r, &g)
	return g, err
}

func DoPostHttp(method string, apimethod string, data interface{}) ([]byte, error) {
	new, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(method, apiUrl+"/"+apimethod, bytes.NewBuffer(new))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	body = bytes.TrimPrefix(body, []byte("\xef\xbb\xbf"))
	return body, err
}

func DoHttp(method string, apimethod string, url string) ([]byte, error) {
	req, err := http.NewRequest(method, apiUrl+"/"+apimethod+"/"+url, bytes.NewBuffer([]byte("")))
	if err != nil {
		return nil, err
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	body = bytes.TrimPrefix(body, []byte("\xef\xbb\xbf"))
	return body, err
}
