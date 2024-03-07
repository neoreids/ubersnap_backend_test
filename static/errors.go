package static

import (
	"github.com/rotisserie/eris"
)

var (
	FILE_CONFIG_NOT_FOUND       = eris.New("file config not found")
	CONFIG_NOT_FOUND            = eris.New("config not found for this key")
	UNABLE_DECODE_IMAGE         = eris.New("unable to decode image")
	UNABLE_CONVERT_IMAGE_TO_MAT = eris.New("unable to convert image to gocv.Mat")
	FAILED_PARSE_QUERY_PARAMS   = eris.New("failed to parse query params request")
	FAILED_PARSE_FORM_BODY      = eris.New("failed to parse body form request")
	FAILED_COMPRESS_IMAGE       = eris.New("failed when compressing image")
	FAILED_OPEN_FILE            = eris.New("failed open uploaded file")
	FAILED_SAVE_IMAGE           = eris.New("failed saving file")
)
