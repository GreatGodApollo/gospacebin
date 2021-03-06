package gospacebin

import (
	"encoding/json"
)

// Turn JSON into a response
func UnmarshalResponse(data []byte) (Response, error) {
	var r Response
	err := json.Unmarshal(data, &r)
	return r, err
}

// Turn a Response into JSON
func (r *Response) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

// A Spacebin API response
type Response struct {
	Error   string  `json:"error,omitempty"` // The error, if one occurred
	Payload Payload `json:"payload"`         // The main body of the response
	Status  int64   `json:"status"`          // The status code of the response
}

// The main body of the response
type Payload struct {
	Content     *string `json:"content,omitempty"`      // The document content
	ContentHash *string `json:"content_hash,omitempty"` // The hash of the document content
	Extension   *string `json:"extension,omitempty"`    // The file extension of the document
	ID          *string `json:"id,omitempty"`           // The document ID
	CreatedAt   *int    `json:"created_at,omitempty"`   // The time the document was created
	UpdatedAt   *int    `json:"updated_at,omitempty"`   // The time the document was last modified/updated
}

// Represents the Spacebin API Client
type Client struct {
	Host string // The host that the client is connecting to.
}

// Information needed to create a document
type CreateDocumentOpts struct {
	Content   string `json:"content"`   // The content of the document
	Extension string `json:"extension"` // The file extension of the document
}

// Turn a CreateDocumentOpts into JSON
func (opts *CreateDocumentOpts) Marshal() ([]byte, error) {
	return json.Marshal(opts)
}

// Returned when POSTing a document
type HashDocument struct {
	ID          string // The document ID
	ContentHash string // The hash of the document content
}

// Returned when GETting a document
type Document struct {
	ID        string // The document ID
	Content   string // The document content
	Extension string // The file extension of the document
	CreatedAt int    // The time the document was created
	UpdatedAt int    // The time the document was last modified
}
