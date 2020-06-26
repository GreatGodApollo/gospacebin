package gospacebin

// NewCreateDocumentOpts creates a new *CreateDocumentOpts for use in the CreateDocument method. The only parameter is
// the content for the *CreateDocumentOpts.
// It returns the new *CreateDocumentOpts.
func NewCreateDocumentOpts(content string) *CreateDocumentOpts {
	return &CreateDocumentOpts{
		Content:   content,
		Extension: "txt",
	}
}

// SetExtension sets the content for the *CreateDocumentOpts.
// It returns the *CreateDocumentOpts
func (cdo *CreateDocumentOpts) SetContent(content string) *CreateDocumentOpts {
	cdo.Content = content
	return cdo
}

// SetExtension sets the extension for the *CreateDocumentOpts.
// It returns the *CreateDocumentOpts.
func (cdo *CreateDocumentOpts) SetExtension(extension string) *CreateDocumentOpts {
	cdo.Extension = extension
	return cdo
}