package hanling_sensitve_data

import (
	"crypto/sha256"
	"encoding/base64"
)

var peppper = "peppper"

func Hashing(original string) string {
	hasher := sha256.New()
	hasher.Write([]byte(original + peppper))
	return base64.URLEncoding.EncodeToString(hasher.Sum(nil))

}
