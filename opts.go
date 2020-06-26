package gospacebin

func NewCreateDocumentOpts(content string) *CreateDocumentOpts {
	return &CreateDocumentOpts{
		Content:   content,
		Extension: "txt",
	}
}

func (cdo *CreateDocumentOpts) SetContent(content string) *CreateDocumentOpts {
	cdo.Content = content
	return cdo
}

func (cdo *CreateDocumentOpts) SetExtension(extension string) *CreateDocumentOpts {
	cdo.Extension = extension
	return cdo
}