package birc

import (
	"crypto/rand"
	"fmt"
	"strings"
	"time"

	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/encoding/korean"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/encoding/traditionalchinese"
	"golang.org/x/text/encoding/unicode"
)

var encoders = map[string]encoding.Encoding{
	"utf-8":       unicode.UTF8,
	"iso-2022-jp": japanese.ISO2022JP,
	"big5":        traditionalchinese.Big5,
	"gbk":         simplifiedchinese.GBK,
	"euc-kr":      korean.EUCKR,
	"gb2312":      simplifiedchinese.HZGB2312,
	"shift-jis":   japanese.ShiftJIS,
	"euc-jp":      japanese.EUCJP,
	"gb18030":     simplifiedchinese.GB18030,
}

func toUTF8(from string, input string) string {
	enc, ok := encoders[from]
	if !ok {
		return input
	}

	res, _ := enc.NewDecoder().String(input)

	return res
}

func escapeTagValue(tag string) string {
	tag = strings.ReplaceAll(tag, `\`, `\\`)
	tag = strings.ReplaceAll(tag, `;`, `\:`)
	tag = strings.ReplaceAll(tag, ` `, `\s`)
	tag = strings.ReplaceAll(tag, "\r", "")
	tag = strings.ReplaceAll(tag, "\n", "")

	return tag
}

func newMsgID() string { // Adding 4 random hex characters for safety
	suffix := make([]byte, 2)

	_, err := rand.Read(suffix)
	if err != nil {
		// Fallback to a simple counter or just the timestamp if entropy fails
		return fmt.Sprintf("mbirc-%d-fb", time.Now().UnixNano())
	}

	return fmt.Sprintf("mbirc-%d%x", time.Now().UnixNano(), suffix)
}
