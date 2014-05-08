package cryptsy

import (
	"net/http"
)

type client struct {
	pubKey  string
	privKey string
	http.Client
}

func NewClient(pubKey, privKey string) (c *client) {
	return &client{pubKey, privKey}
}

func (c *client) doTimeoutRequest(timer *time.Timer, req *http.Request) (*http.Response, error) {
	// Do the request in the background so we can check the timeout
	type result struct {
		resp *http.Response
		err  error
	}
	done := make(chan result, 1)
	go func() {
		resp, err := c.Do(req)
		done <- result{resp, err}
	}()
	// Wait for the read or the timeout
	select {
	case r := <-done:
		return r.resp, r.err
	case <-timer.C:
		return nil, errors.New("timeout on reading data from Cryspy API")
	}
}

func (c *client) do(method string, ressource string, payload string, isPrivate bool) (response []byte, err error) {
	connectTimer := time.NewTimer(DEFAULT_HTTPCLIENT_TIMEOUT * time.Second)

	if !isPrivate {
		query := fmt.Sprintf("%s?%s", API_BASE_PUB, ressource)

	} else {
		query := fmt.Sprintf("%s?%s", API_BASE_PRIV, ressource)
	}

	//fmt.Println(query)
	req, err := http.NewRequest(method, query, strings.NewReader(payload))
	if err != nil {
		return
	}
	if method == "POST" || method == "PUT" {
		req.Header.Add("Content-Type", "application/json;charset=utf-8")
	}
	req.Header.Add("Accept", "application/json")
	if isPrivate {
		req.Header.Add("Key", c.pubKey)
		req.Header.Add("Sign", "toto")
		req.Header.Add("nonce", 1)
	}

	//fmt.Println(req)
	resp, err := c.doTimeoutRequest(connectTimer, req)
	if err != nil {
		return
	}

	defer resp.Body.Close()
	response, err = ioutil.ReadAll(resp.Body)
	//fmt.Println(fmt.Sprintf("reponse %s", response), err)
	if err != nil {
		return response, err
	}
	if resp.StatusCode != 200 {
		err = errors.New(resp.Status)
	}
	return response, err
}
