package common

import (
	"crypto/md5"
	"encoding/hex"
	"net/http"
	"strconv"
	"strings"
)

func PanicIf(err error) {
	if err != nil {
		panic(err)
	}
}

func StringToInt(value string) int {
	if value == "" {
		return 0
	}

	val, _ := strconv.Atoi(value)
	return val
}

func IntToString(value int) string {
	return strconv.Itoa(value)
}

func Md5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

func GetRemoteIp(r *http.Request) (ip string) {
	ip = r.Header.Get("X-Real-Ip")
	if ip == "" {
		ip = r.RemoteAddr
	}

	ip = strings.Split(ip, ":")[0]
	if len(ip) < 7 || ip == "127.0.0.1" {
		ip = "localhost"
	}

	return
}
