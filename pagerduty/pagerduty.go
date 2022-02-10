package pagerduty

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/url"
	"os"
	"syscall"
	"time"

	"github.com/google/go-querystring/query"
)

const (
	defaultBaseURL = "https://api.pagerduty.com"
)

type service struct {
	client *Client
}

// Config represents the configuration for a PagerDuty client
type Config struct {
	BaseURL    string
	HTTPClient *http.Client
	Token      string
	UserAgent  string
	Debug      bool
}

// Logger internal logger
type Logger struct {
	logger *log.Logger
	f      *os.File
}

// Client manages the communication with the PagerDuty API
type Client struct {
	baseURL                    *url.URL
	client                     *http.Client
	FileLogger                 *Logger
	Config                     *Config
	Abilities                  *AbilityService
	Addons                     *AddonService
	EscalationPolicies         *EscalationPolicyService
	Extensions                 *ExtensionService
	MaintenanceWindows         *MaintenanceWindowService
	Rulesets                   *RulesetService
	Schedules                  *ScheduleService
	Services                   *ServicesService
	Teams                      *TeamService
	ExtensionSchemas           *ExtensionSchemaService
	Users                      *UserService
	Vendors                    *VendorService
	EventRules                 *EventRuleService
	BusinessServices           *BusinessServiceService
	ServiceDependencies        *ServiceDependencyService
	Priorities                 *PriorityService
	ResponsePlays              *ResponsePlayService
	SlackConnections           *SlackConnectionService
	Tags                       *TagService
	WebhookSubscriptions       *WebhookSubscriptionService
	BusinessServiceSubscribers *BusinessServiceSubscriberService
}

// Response is a wrapper around http.Response
type Response struct {
	Response  *http.Response
	BodyBytes []byte
}

// RequestOptions is an object to setting options for HTTP requests
type RequestOptions struct {
	Type  string
	Label string
	Value string
}

type ErrType int

const (
	ERR_NETWORK ErrType = iota
	ERR_BODY_LEN
	ERR_API
	ERR_JSON_DATA
)

func (l *Logger) openFile() error {
	f, err := os.OpenFile("go-pagerduty.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	l.f = f
	l.logger = log.New(
		l.f,
		fmt.Sprintf("go-pagerduty (%d) ", os.Getpid()),
		log.Ldate|log.Ltime|log.Lmicroseconds|log.Llongfile,
	)
	return nil
}

func (l *Logger) closeFile() error {
	return l.f.Close()
}

// Print a string log value into a file
func (l *Logger) Print(s string) {
	err := l.openFile()
	if err != nil {
		return
	}

	l.logger.Println(s)

	err = l.closeFile()
}

// NewClient returns a new PagerDuty API client.
func NewClient(config *Config) (*Client, error) {
	config.HTTPClient = &http.Client{
		Transport: &http.Transport{
			DialContext: (&net.Dialer{
				Timeout:   45 * time.Second,
				KeepAlive: 40 * time.Second,
			}).DialContext,
			TLSClientConfig: &tls.Config{
				CipherSuites: []uint16{
					tls.TLS_RSA_WITH_RC4_128_SHA,
					tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
				},
			},
			TLSHandshakeTimeout:   120 * time.Second,
			MaxIdleConns:          0,
			MaxIdleConnsPerHost:   500,
			MaxConnsPerHost:       0,
			IdleConnTimeout:       0,
			ResponseHeaderTimeout: 0,
		},
	}

	if config.BaseURL == "" {
		config.BaseURL = defaultBaseURL
	}

	config.UserAgent = "nordcloud/go-pagerduty(terraform)"

	baseURL, err := url.Parse(config.BaseURL)
	if err != nil {
		return nil, err
	}

	c := &Client{
		baseURL:    baseURL,
		client:     config.HTTPClient,
		Config:     config,
		FileLogger: &Logger{},
	}

	c.FileLogger.Print("go-pagerduty initialization")

	c.Abilities = &AbilityService{c}
	c.Addons = &AddonService{c}
	c.EscalationPolicies = &EscalationPolicyService{c}
	c.MaintenanceWindows = &MaintenanceWindowService{c}
	c.Rulesets = &RulesetService{c}
	c.Schedules = &ScheduleService{c}
	c.Services = &ServicesService{c}
	c.Teams = &TeamService{c}
	c.Users = &UserService{c}
	c.Vendors = &VendorService{c}
	c.Extensions = &ExtensionService{c}
	c.ExtensionSchemas = &ExtensionSchemaService{c}
	c.EventRules = &EventRuleService{c}
	c.BusinessServices = &BusinessServiceService{c}
	c.ServiceDependencies = &ServiceDependencyService{c}
	c.Priorities = &PriorityService{c}
	c.ResponsePlays = &ResponsePlayService{c}
	c.SlackConnections = &SlackConnectionService{c}
	c.Tags = &TagService{c}
	c.WebhookSubscriptions = &WebhookSubscriptionService{c}
	c.BusinessServiceSubscribers = &BusinessServiceSubscriberService{c}

	InitCache(c)
	PopulateCache()

	return c, nil
}

func (c *Client) newRequest(method, url string, body interface{}, options ...RequestOptions) (*http.Request, error) {
	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	if c.Config.Debug {
		log.Printf("[DEBUG] PagerDuty - Preparing %s request to %s with body: %s", method, url, buf)
		c.FileLogger.Print(fmt.Sprintf("[DEBUG] PagerDuty - Preparing %s request to %s with body: %s", method, url, buf))
	}

	u := c.baseURL.String() + url

	req, err := http.NewRequest(method, u, buf)
	if err != nil {
		return nil, err
	}

	if len(options) > 0 {
		for _, o := range options {
			if o.Type == "header" {
				req.Header.Add(o.Label, o.Value)
			}
		}
	}
	req.Header.Add("Accept", "application/vnd.pagerduty+json;version=2")
	req.Header.Add("Authorization", fmt.Sprintf("Token token=%s", c.Config.Token))
	req.Header.Add("Content-Type", "application/json")

	if c.Config.UserAgent != "" {
		req.Header.Add("User-Agent", c.Config.UserAgent)
	}
	return req, nil
}

func (c *Client) newRequestDo(method, url string, qryOptions, body, v interface{}) (*Response, error) {
	if qryOptions != nil {
		values, err := query.Values(qryOptions)
		if err != nil {
			return nil, err
		}

		if v := values.Encode(); v != "" {
			url = fmt.Sprintf("%s?%s", url, v)
		}
	}
	req, err := c.newRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	return c.do(req, v)
}

func (c *Client) newRequestDoOptions(method, url string, qryOptions, body, v interface{}, reqOptions ...RequestOptions) (*Response, error) {
	if qryOptions != nil {
		values, err := query.Values(qryOptions)
		if err != nil {
			return nil, err
		}

		if v := values.Encode(); v != "" {
			url = fmt.Sprintf("%s?%s", url, v)
		}
	}
	req, err := c.newRequest(method, url, body, reqOptions...)
	if err != nil {
		return nil, err
	}

	return c.do(req, v)
}

func (c *Client) handleErrorResponse(err error, resp *http.Response, errType ErrType, tryNum int) (*Response, error, bool) {
	c.FileLogger.Print(fmt.Sprintf("[ERROR] API Error [%d] (try %d): %#v\n\n%#v\n\n%#v\n", errType, tryNum, err, err.Error(), resp))

	if errors.Is(err, syscall.ECONNRESET) {
		return nil, fmt.Errorf("%w (Retryable connection reset error [%d])", err, errType), true
	}
	if errors.Is(err, syscall.EPIPE) {
		return nil, fmt.Errorf("%w (Retryable broken pipe error [%d])", err, errType), true
	}
	if errType == ERR_NETWORK && err.(net.Error).Temporary() {
		return nil, fmt.Errorf("%w (Retryable network error [%d])", err, errType), true
	}
	if errType == ERR_BODY_LEN {
		return nil, fmt.Errorf("%w (Retryable body length error [%d])", err, errType), true
	}
	if errType == ERR_API && (resp.StatusCode == 429 || resp.StatusCode >= 500) {
		return nil, fmt.Errorf("%w (Retryable HTTP error %d [%d])", err, resp.StatusCode, errType), true
	}
	if errType == ERR_JSON_DATA {
		return nil, fmt.Errorf("%w (Retryable error: invalid JSON body [%d])", err, errType), true
	}

	return nil, fmt.Errorf("%w (Non-retryable connection error [%d], try %d)", err, errType, tryNum), false
}

func (c *Client) do(req *http.Request, v interface{}) (*Response, error) {
	maxtries := 4
	var resp *Response
	var err error = nil
	var trynum = 0
	var retry = false
	for trynum <= maxtries {
		resp, err, retry = c.doRetry(req.Clone(context.Background()), v, trynum)
		if err == nil {
			break
		}
		if retry == false {
			c.FileLogger.Print(fmt.Sprintf("[ERROR] Non-retryable error returned: %s", err.Error()))
			return nil, fmt.Errorf("%w (Non-retryable API error)", err)
		}
		c.FileLogger.Print("[INFO] Sleeping between retries")
		time.Sleep(30 * time.Second)
		trynum++
	}

	if trynum > maxtries {
		return nil, fmt.Errorf("%w (API error despite %d low-level retries)", err, maxtries)
	}

	if c.Config.Debug {
		c.FileLogger.Print(fmt.Sprintf("[DEBUG] Returning %#v %#v", resp, err))
	}

	return resp, err
}

func (c *Client) doRetry(req *http.Request, v interface{}, trynum int) (*Response, error, bool) {
	if c.Config.Debug {
		c.FileLogger.Print(fmt.Sprintf("[DEBUG] Executing %s request to %s (try %d)", req.Method, req.URL, trynum))
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return c.handleErrorResponse(err, resp, ERR_NETWORK, trynum)
	}
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return c.handleErrorResponse(err, resp, ERR_BODY_LEN, trynum)
	}
	response := &Response{
		Response:  resp,
		BodyBytes: bodyBytes,
	}

	if err := c.checkResponse(response); err != nil {
		return c.handleErrorResponse(err, resp, ERR_API, trynum)
	}

	if v != nil {
		if err := c.DecodeJSON(response, v); err != nil {
			return c.handleErrorResponse(err, resp, ERR_JSON_DATA, trynum)
		}
	}

	if c.Config.Debug {
		c.FileLogger.Print(fmt.Sprintf("[DEBUG] Returning value for %s request to %s (try %d)", req.Method, req.URL, trynum))
	}

	if response != nil {
		return response, nil, false
	}
	return nil, fmt.Errorf("[FATAL] No error was thrown but no valid response received"), true
}

// ListResp represents a list response from the PagerDuty API
type ListResp struct {
	Offset int  `json:"offset,omitempty"`
	Limit  int  `json:"limit,omitempty"`
	More   bool `json:"more,omitempty"`
	Total  int  `json:"total,omitempty"`
}

// responseHandler is capable of parsing a response. At a minimum it must
// extract the page information for the current page. It can also execute
// additional necessary handling; for example, if a closure, it has access
// to the scope in which it was defined, and can be used to append data to
// a specific slice. The responseHandler is responsible for closing the response.
type responseHandler func(response *Response) (ListResp, *Response, error)

func (c *Client) newRequestPagedGetDo(basePath string, handler responseHandler, reqOptions ...RequestOptions) error {
	// Indicates whether there are still additional pages associated with request.
	var stillMore bool

	// Offset to set for the next page request.
	var nextOffset int

	// While there are more pages, keep adjusting the offset to get all results.
	for stillMore, nextOffset = true, 0; stillMore; {
		response, err := c.newRequestDoOptions("GET", fmt.Sprintf("%s?offset=%d", basePath, nextOffset), nil, nil, nil, reqOptions...)
		if err != nil {
			return err
		}

		// Call handler to extract page information and execute additional necessary handling.
		pageInfo, _, err := handler(response)
		if err != nil {
			return err
		}

		// Bump the offset as necessary and set whether more results exist.
		nextOffset = pageInfo.Offset + pageInfo.Limit
		stillMore = pageInfo.More
	}

	return nil
}

// ValidateAuth validates a token against the PagerDuty API
func (c *Client) ValidateAuth() error {
	_, _, err := c.Abilities.List()
	return err
}

// DecodeJSON decodes json body to given interface
func (c *Client) DecodeJSON(res *Response, v interface{}) error {
	return json.Unmarshal(res.BodyBytes, v)
}

func (c *Client) checkResponse(res *Response) error {
	if res.Response.StatusCode >= 200 && res.Response.StatusCode <= 299 {
		return nil
	}

	return c.decodeErrorResponse(res)
}

func (c *Client) decodeErrorResponse(res *Response) error {
	// Try to decode error response or fallback with standard error
	v := &errorResponse{Error: &Error{ErrorResponse: res}}
	if err := c.DecodeJSON(res, v); err != nil {
		return fmt.Errorf("%s API call to %s failed: %v", res.Response.Request.Method, res.Response.Request.URL.String(), res.Response.Status)
	}

	return v.Error
}
