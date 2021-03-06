package gospacebin

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

// NewClient creates a new Spacebin API client. The only parameter is the host the client is connecting to.
// The return value of this method is the new *Client.
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
		return nil, fmt.Errorf("%s", "an unexpected response was received")
	}
	resp.Status = int64(httpResp.StatusCode)
	return &resp, nil
}

// CreateDocument allows you to create a new document on the Spacebin. The only parameter is a *CreateDocumentOpts.
// The return value of this method is a *HashDocument.
func (cli *Client) CreateDocument(opts *CreateDocumentOpts) (*HashDocument, error) {
	url := fmt.Sprintf("%s/api/v1/documents", cli.Host)
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
		return nil, fmt.Errorf("%s", resp.Error)
	}
	return &HashDocument{
		ID:          *resp.Payload.ID,
		ContentHash: *resp.Payload.ContentHash,
	}, nil
}

// GetDocument allows you to retrieve a document from the Spacebin. The only parameter is a string, the document ID.
// The return value of this method is a *Document.
func (cli *Client) GetDocument(docID string) (*Document, error) {
	url := fmt.Sprintf("%s/api/v1/documents/%s", cli.Host, docID)
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
			return nil, fmt.Errorf("%s", resp.Error)
		}
	}
	return &Document{
		ID:        *resp.Payload.ID,
		Content:   *resp.Payload.Content,
		Extension: *resp.Payload.Extension,
		UpdatedAt: *resp.Payload.UpdatedAt,
		CreatedAt: *resp.Payload.CreatedAt,
	}, nil
}
