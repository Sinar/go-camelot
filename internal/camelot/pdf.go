package camelot

import (
	"fmt"
	"strings"

	"github.com/ledongthuc/pdf"
	"github.com/sanity-io/litter"
)

type PDFPage struct {
	PageNo           int
	PDFPlainText     string
	PDFTxtSameLines  []string // combined content with same line .. proxy for changes
	PDFTxtSameStyles []string // combined content with same style .. proxy for changes
}

type PDFDocument struct {
	NumPages   int
	Pages      []PDFPage
	SourcePath string
}

type ExtractPDFOptions struct {
	StartPage int
	NumPages  int
}

const (
	MaxLineProcessed = 1000
)

func NewPDFDocument(pdfPath string, options *ExtractPDFOptions) (*PDFDocument, error) {
	f, r, err := pdf.Open(pdfPath)
	defer func() {
		_ = f.Close()
	}()
	if err != nil {
		return nil, fmt.Errorf("PDFAccessErr: %w", err)
	}

	startPage := 1
	totalPage := r.NumPage()
	// DEBUG
	//totalPage = 3
	if options != nil {
		if options.StartPage > 1 {
			startPage = options.StartPage
		}
		if options.NumPages > 0 {
			totalPage = options.NumPages
		}
	}
	var pdfPages []PDFPage
	// Init it and fill it with the extracted info  earlier ..
	pdfDoc := PDFDocument{
		NumPages:   totalPage,
		Pages:      pdfPages,
		SourcePath: pdfPath,
	}

	for pageIndex := startPage; pageIndex <= totalPage; pageIndex++ {
		p := r.Page(pageIndex)
		if p.V.IsNull() {
			continue
		}
		// New Page to be Processed ..
		newPageProcessed := PDFPage{
			PageNo:           pageIndex,
			PDFTxtSameLines:  []string{},
			PDFTxtSameStyles: []string{},
		}
		// copy over plain text; short form
		pt, pterr := p.GetPlainText(nil)
		if pterr != nil {
			if pterr.Error() == "malformed PDF: reading at offset 0: stream not present" {
				fmt.Println("**WILL IGNORE!!!! *****")
				continue
			}
			return nil, fmt.Errorf(" GetPlainText ERROR: %w", pt)
		}
		//newPageProcessed.PDFPlainText = pt
		// processStyleChanges ..
		//extractTxtSameStyles()
		// DEBUG
		//fmt.Println("LEN: ", p.V.Len())
		//fmt.Println("KEYS", p.V.Keys())
		//fmt.Println("KIND", p.V.Kind())
		// DEBUG
		//fmt.Println("== START CONTENT PAGE ", i)
		//spew.Dump(pt)
		// Top 10 lines for this page by line analysis
		//fmt.Println("== START ANALYZE by LINE")
		//newPageProcessed.PDFTxtSameLines = make([]string, 0, 20)
		extractTxtSameLine(&newPageProcessed.PDFTxtSameLines, p.Content().Text)

		// Top 10
		//fmt.Println("== START ANALYZE by STYLE")
		//newPageProcessed.PDFTxtSameStyles = make([]string, 0, 20)
		extractTxtSameStyles(&newPageProcessed.PDFTxtSameStyles, p.Content().Text)
		//fmt.Println("== END ANALYZE by STYLE")

		// If OK, append them ..
		pdfDoc.Pages = append(pdfDoc.Pages, newPageProcessed)
	}

	// DEBUG
	litter.Dump(pdfDoc)
	return &pdfDoc, nil
}

func extractTxtSameLine(ptrTxtSameLine *[]string, pdfContentTxt []pdf.Text) error {

	var numValidLineCounted int
	var currentLineNumber float64
	var currentContent string

	var pdfTxtSameLine []string

	// DEBUG
	//spew.Dump(pdfContentTxt)

	for _, v := range pdfContentTxt {

		// Guard function .. what is it?
		//if strings.TrimSpace(v.S) == "" {
		//	fmt.Println("Skipping blank line / content ..")
		//	continue
		//}

		if currentLineNumber == 0 {
			currentLineNumber = v.Y
			// DEBUG
			//fmt.Println("Set first line to ", currentLineNumber)
			currentContent += v.S
			continue
		}

		// Happy path ..
		// DEBUG
		//fmt.Println("Append CONTENT: ", currentContent, " X: ", v.X, " Y: ", v.Y)
		// number of valid line increase when new valid line ..
		if currentLineNumber != v.Y {
			if strings.TrimSpace(currentContent) != "" {
				// trim new lines ..
				currentContent = strings.ReplaceAll(currentContent, "\n", "")
				// DEBUG
				//fmt.Println("NEW Line ... collected: ", currentContent)
				pdfTxtSameLine = append(pdfTxtSameLine, currentContent)
				numValidLineCounted++
			}
			currentContent = v.S // reset .. after append
			currentLineNumber = v.Y
		} else {
			// If on the same line, just build up the content ..
			currentContent += v.S
		}

		// NOTE: Only get MaxLineProcessed lines ..
		if numValidLineCounted > MaxLineProcessed {
			break
		}

	}
	// All the left over, do one more final check ...
	if strings.TrimSpace(currentContent) != "" {
		// trim new lines ..
		currentContent = strings.ReplaceAll(currentContent, "\n", "")
		// DEBUG
		//fmt.Println("NEW Line ... collected: ", currentContent)
		pdfTxtSameLine = append(pdfTxtSameLine, currentContent)
	}

	*ptrTxtSameLine = pdfTxtSameLine
	//spew.Dump(ptrTxtSameLine)
	return nil
}

func extractTxtSameStyles(ptrTxtSameStyles *[]string, pdfContentTxt []pdf.Text) error {
	var numValidLineCounted int
	var currentFont string
	var currentContent string

	var pdfTxtSameStyles []string

	for _, v := range pdfContentTxt {

		// Guard function .. what is it?

		if currentFont == "" {
			currentFont = v.Font
			// DEBUG
			//fmt.Println("Set first font to ", currentFont)
			currentContent += v.S
			continue
		}

		// Happy path ..
		if currentFont != v.Font {
			if strings.TrimSpace(currentContent) != "" {
				// trim new lines ..
				currentContent = strings.ReplaceAll(currentContent, "\n", "")
				// DEBUG
				//fmt.Println("NEW Style ... collected: ", currentContent)
				pdfTxtSameStyles = append(pdfTxtSameStyles, currentContent)
				//fmt.Println("CURRENT ,...")
				//spew.Dump(pdfTxtSameStyles)
				numValidLineCounted++
			}
			// reset for next iteraton ..
			currentContent = v.S // reset .. after append
			currentFont = v.Font
		} else {
			// If with the same style, just build up the content ..
			currentContent += v.S
		}

		// NOTE: Only get MaxLineProcessed lines ..
		if numValidLineCounted > MaxLineProcessed {
			break
		}
	}
	// All the left over, do one more final check ...
	if strings.TrimSpace(currentContent) != "" {
		// trim new lines ..
		currentContent = strings.ReplaceAll(currentContent, "\n", "")
		// DEBUG
		//fmt.Println("NEW Style ... collected: ", currentContent)
		pdfTxtSameStyles = append(pdfTxtSameStyles, currentContent)
	}

	*ptrTxtSameStyles = pdfTxtSameStyles
	//spew.Dump(ptrTxtSameStyles)

	return nil
}
