package behavior3go

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"io"
)

//生成32位md5字串
func getMd5String(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

//生成Guid字串
func getGuid() string {
	b := make([]byte, 48)

	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return getMd5String(base64.URLEncoding.EncodeToString(b))
}

/**
 * This function is used to create unique IDs for trees and nodes.
 *
 * (consult http://www.ietf.org/rfc/rfc4122.txt).
 *
 * @class createUUID
 * @construCtor
 * @return {String} A unique ID.
**/
func CreateUUID() string {
	return getGuid()
}
