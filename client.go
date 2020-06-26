package gospacebin

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func NewClient(host string) *Client {
	return &Client{
		Host: host,
	}
}

func (cli *Client) makeRequest(req *http.Request) (*Response, error) {
	req.Header.Set("User-Agent", "spacebin-go")
	httpCli := &http.Client{}
	httpResp, err := httpCli.Do(req)
	if err != nil {
		return nil, err
	}
	defer httpResp.Body.Close()
	body, err := ioutil.ReadAll(httpResp.Body)
	if err != nil {
		return nil, err
	}
	resp, err := UnmarshalResponse(body)
	if err != nil {
		fmt.Println(string(body))
		return nil, fmt.Errorf("%s", "an unexpected response was received")
	}
	resp.Status = int64(httpResp.StatusCode)
	return &resp, nil
}

func (cli *Client) CreateDocument(opts *CreateDocumentOpts) (*HashDocument, error) {
	url := fmt.Sprintf("%s/api/v1/document", cli.Host)
	j, err := opts.Marshal()
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(j))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := cli.makeRequest(req)
	if err != nil {
		return nil, err
	}
	if resp.Status != 201 {
		return nil, fmt.Errorf("%s", resp.Error.Message)
	}
	return &HashDocument{
		ID:          *resp.Payload.ID,
		ContentHash: *resp.Payload.ContentHash,
		Extension:   *resp.Payload.Extension,
	}, nil
}

func (cli *Client) GetDocument(docID string) (*Document, error) {
	url := fmt.Sprintf("%s/api/v1/document/%s", cli.Host, docID)
	req, err := http.NewRequest("GET", url, bytes.NewBuffer([]byte{}))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := cli.makeRequest(req)
	if err != nil {
		return nil, err
	}
	if resp.Status != 200 {
		if resp.Status == 404 {
			return nil, fmt.Errorf("%s", "document not found")
		} else {
			return nil, fmt.Errorf("%s", resp.Error.Message)
		}
	}
	return &Document{
		ID:          *resp.Payload.ID,
		Content:     *resp.Payload.Content,
		Extension:   *resp.Payload.Extension,
		DateCreated: *resp.Payload.DateCreated,
	}, nil
}

func (cli *Client) DocumentExists(docID string) (bool, error) {
	url := fmt.Sprintf("%s/api/v1/document/%s/verify", cli.Host, docID)
	req, err := http.NewRequest("GET", url, bytes.NewBuffer([]byte{}))
	if err != nil {
		return false, err
	}
	resp, err := cli.makeRequest(req)
	if err != nil {
		return false, err
	}
	if resp.Status != 200 && resp.Status != 404 {
		return false, fmt.Errorf("%s", resp.Error.Message)
	}
	return *resp.Payload.Exists, nil
}