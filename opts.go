package gospacebin

// NewCreateDocumentOpts creates a new *CreateDocumentOpts for use in the CreateDocument method. The only parameter is
// the content for the *CreateDocumentOpts.
// It returns the new *CreateDocumentOpts.
func NewCreateDocumentOpts(content string) *CreateDocumentOpts {
	return &CreateDocumentOpts{
		Content:   content,
		Extension: "none",
	}
}

// SetContent sets the content for the *CreateDocumentOpts.
// It returns the *CreateDocumentOpts
func (opts *CreateDocumentOpts) SetContent(content string) *CreateDocumentOpts {
	opts.Content = content
	return opts
}

// SetExtension sets the extension for the *CreateDocumentOpts.
// It returns the *CreateDocumentOpts.
func (opts *CreateDocumentOpts) SetExtension(extension string) *CreateDocumentOpts {
	opts.Extension = extension
	return opts
}
