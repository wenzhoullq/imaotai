package lib

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"github.com/google/uuid"
	"regexp"
	"time"
	"zuoxingtao/constant"
)

func Signature(mobile string) string {
	text := constant.SALT + mobile + fmt.Sprintf("%d", time.Now().UnixNano()/int64(time.Millisecond))
	hash := md5.Sum([]byte(text))
	md5Hash := hex.EncodeToString(hash[:])
	return md5Hash
}

func GetUuID() string {
	uuid := uuid.New()
	return uuid.String()
}

func GetMTDeviceID(deviceID string) string {
	var ret []byte
	i10 := 72

	for _, char := range deviceID {
		i10 ^= int(char)
		ret = append(ret, byte(i10))
	}

	retBase64 := base64.StdEncoding.EncodeToString(ret)
	return "clips_" + retBase64
}
func Md5Hash(deviceID string) string {
	if deviceID == "" {
		return ""
	}

	data := []byte(deviceID)
	hash := md5.Sum(data)
	hashString := hex.EncodeToString(hash[:])

	return hashString
}

func GetDeviceID() string {
	uuid := GetUuID()
	deviceidMd5 := Md5Hash(uuid)
	return GetMTDeviceID(deviceidMd5)
}

// CheckMobileNumber 检查手机号码是否符合规范
func CheckMobileNumber(number string) bool {
	// 定义手机号码的正则表达式模式
	pattern := `^1[345789]\d{9}$`
	re := regexp.MustCompile(pattern)
	return re.MatchString(number)
}

func CheckVerifyCode(code string) bool {
	pattern := `^\d{6}$`
	re := regexp.MustCompile(pattern)
	return re.MatchString(code)
}

func AesEncrypt(params string) (string, error) {
	key := []byte(constant.AESKEY)
	iv := []byte(constant.AESIV)
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	plaintext := []byte(params)
	// 补全填充
	paddingLength := aes.BlockSize - len(plaintext)%aes.BlockSize
	padding := []byte{byte(paddingLength)}
	plaintext = append(plaintext, bytes.Repeat(padding, paddingLength)...)
	// 加密
	ciphertext := make([]byte, len(plaintext))
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext, plaintext)
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}
