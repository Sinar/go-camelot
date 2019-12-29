package camelot

// Here we'll just use the super simple and stupid PDF method ..
// with a clustering mechanism ..
func getMetadata(pdfSourcePath string) {

	NewPDFDocument(pdfSourcePath, &ExtractPDFOptions{
		StartPage: 2,
		NumPages:  1,
	})
}
