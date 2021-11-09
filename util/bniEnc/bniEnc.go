package bniEnc

import (
	"encoding/base64"
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
	"sync"
	"time"
)

var lock = &sync.Mutex{}

const timeDiffLimit = 480

func Encrypt(json string, client_id string, secret_key string) string {
	lock.Lock()
	defer lock.Unlock()
	
	return doubleEncrypt(reverse(fmt.Sprintf("%v", time.Now().Unix()))+"."+json, client_id, secret_key)
}

func Decrypt(encrypted string, client_id string, secret_key string) (string, error) {
	lock.Lock()
	defer lock.Unlock()

	parsed_string := doubleDecrypt(encrypted, client_id, secret_key)
	var lst = strings.SplitN(parsed_string, ".", 2)
	if len(lst) < 2 {
		return "", errors.New("bniEnc: parsing error, wrong cid or sck or invalid data.")
	}
	if tsDiff(reverse(lst[0])) == false {
		return "", errors.New("bniEnc: data has been expired.")
	}
	return lst[1], nil
}

func tsDiff(ts string) bool {
	_ts, err := strconv.ParseInt(ts, 10, 64)
	if err != nil {
		return false
	}
	return math.Abs(float64(_ts-time.Now().Unix())) <= timeDiffLimit
}

func doubleEncrypt(str string, cid string, sck string) string {
	arr := []byte(str)
	result := encrypt(arr, cid)
	result = encrypt(result, sck)
	return strings.Replace(strings.Replace(strings.TrimRight(base64.StdEncoding.EncodeToString(result), "="), "+", "-", -1), "/", "_", -1)
}

func encrypt(str []byte, k string) []byte {
	var result []byte
	strls := len(str)
	strlk := len(k)
	for i := 0; i < strls; i++ {
		char := str[i]
		keychar := k[(i+strlk-1)%strlk]
		char = byte((int(char) + int(keychar)) % 128)
		result = append(result, char)
	}
	return result
}

func doubleDecrypt(str string, cid string, sck string) string {
	if i := len(str) % 4; i != 0 {
		str += strings.Repeat("=", 4-i)
	}
	result, err := base64.StdEncoding.DecodeString(strings.Replace(strings.Replace(str, "-", "+", -1), "_", "/", -1))
	if err != nil {
		return ""
	}
	result = decrypt(result, cid)
	result = decrypt(result, sck)
	return string(result[:])
}

func decrypt(str []byte, k string) []byte {
	var result []byte
	strls := len(str)
	strlk := len(k)
	for i := 0; i < strls; i++ {
		char := str[i]
		keychar := k[(i+strlk-1)%strlk]
		char = byte(((int(char) - int(keychar)) + 256) % 128)
		result = append(result, char)
	}
	return result
}

func reverse(s string) string {
	chars := []rune(s)
	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}
	return string(chars)
}