# go-camelot

Clean room implementation for PDF table detection; inspired by camelot.
Starts with the raw image file; get it from their corresponding go-pardocs + go-dundocs commands.
This is called via simple os.exec for every page; pass back OCR-ed text to sync with calling function.

Assumption: Only target pure strongly separated tables with line; e.g. detect line

## Ideas

Use edge detection (e.g. Canny) + line detection (e.g Hough Lines); available via gocv to create the row/column slices

Use font (via pdfcpu - https://pdfcpu.io/extract/extract_fonts.html) + character-sets to detect text in sliced image - https://github.com/Th1nkK1D/gocr

Alt: Use OCR on the slices made available

    - https://github.com/otiai10/gosseract
    - https://github.com/otiai10/ocrserver

See one implementation: https://github.com/hybridgroup/gocv/tree/master/cmd/find-lines

## Techniques

- Canny - https://opencv-python-tutroals.readthedocs.io/en/latest/py_tutorials/py_imgproc/py_canny/py_canny.html
- HoughLines - https://opencv-python-tutroals.readthedocs.io/en/latest/py_tutorials/py_imgproc/py_houghlines/py_houghlines.html
- Hough Transform Theory - http://homepages.inf.ed.ac.uk/rbf/HIPR2/hough.htm

## Libraries available for use

- Lightweight image + rotation - https://github.com/disintegration/imaging
- Advanced Transforms + Filters - https://github.com/disintegration/gift
- exif info - https://github.com/disintegration/imageorient
