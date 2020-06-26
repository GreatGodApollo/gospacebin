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
	Error   Error   `json:"error"`   // The error, if one occured
	Payload Payload `json:"payload"` // The main body of the response
	Status  int64   `json:"status"`  // The status code of the response
}

// The error, if one occured
type Error struct {
	FileName   *string `json:"fileName,omitempty"`  // The file it occurred in
	LineNumber *int64  `json:"lineNumber,omitempty"`// The line it occurred on
	Message    *string `json:"message,omitempty"`   // The error message
}

// The main body of the response
type Payload struct {
	Content     *string `json:"content,omitempty"`    // The document content
	ContentHash *string `json:"contentHash,omitempty"`// The hash of the document content
	DateCreated *string `json:"dateCreated,omitempty"`// The date & time of which the document was created
	Extension   *string `json:"extension,omitempty"`  // The file extension of the document
	ID          *string `json:"id,omitempty"`         // The document ID
	Exists      *bool   `json:"exists,omitempty"`     // If the document exists or not
}

// Represents the Spacebin API Client
type Client struct {
	Host string // The host that the client is connecting to.
}

// Information needed to create a document
type CreateDocumentOpts struct {
	Content   string `json:"content"` 	// The content of the document
	Extension string `json:"extension"` // The file extension of the document
}

// Turn a CreateDocumentOpts into JSON
func (r *CreateDocumentOpts) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

// Returned when POSTing a document
type HashDocument struct {
	ID 			string // The document ID
	ContentHash string // The hash of the document content
	Extension 	string // The file extension of the document
}

// Returned when GETting a document
type Document struct {
	ID 			string // The document ID
	Content 	string // The document content
	Extension 	string // The file extension of the document
	DateCreated string // The date & time of which the document was created
}