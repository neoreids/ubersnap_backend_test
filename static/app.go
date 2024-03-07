package static

import "fmt"

const (
	PUBLIC_DIR                   = "./public-static"
	APP_NAME                     = "ubersnap"
	MIN_LENGTH_SECRET_KEY        = 32
	PREFIX_JWT_TOKEN             = "bearer"
	STATIC_ASSET_PATH            = "/static"
	CAPACITY_SHOULD_COMPRESS     = 1000000
)

var (
	ONLY_IMAGE = []string{
		"image/png", "image/jpeg", "image/svg+xml", "image/webp",
	}
	INDONESIAN_MONTH = map[int]string{
		1:  "januari",
		2:  "februari",
		3:  "maret",
		4:  "april",
		5:  "mei",
		6:  "juni",
		7:  "juli",
		8:  "agustus",
		9:  "september",
		10: "oktober",
		11: "november",
		12: "desember",
	}
	CONTENT_DIR_IMAGE = fmt.Sprintf("%s/%s", PUBLIC_DIR, "content")
	CONTENT_URL_IMAGE = fmt.Sprintf("%s/%s", STATIC_ASSET_PATH, "content")
)
