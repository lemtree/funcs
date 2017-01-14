/*
 * golang common funcs
 * auth quiincy
 */
package tools

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"strconv"
	"strings"
)

// cal md5 string
func MD5(data []byte) string {
	md5Raw := md5.Sum(data)
	return hex.EncodeToString([]byte(md5Raw[:]))
}

// cal md5 string other func
func MD52(data []byte) string {
	md5Ctx := md5.New()
	md5Ctx.Write(data)
	cipherStr := md5Ctx.Sum(nil)
	return hex.EncodeToString(cipherStr)
}

// ip to int
func Ip2Long(ip string) int {
	list := strings.Split(ip, ".")
	ip0, _ := strconv.Atoi(list[0])
	ip1, _ := strconv.Atoi(list[1])
	ip2, _ := strconv.Atoi(list[2])
	ip3, _ := strconv.Atoi(list[3])
	return ip0<<24 + ip1<<16 + ip2<<8 + ip3
}

// int to ip
func Long2Ip(ipInt int) string {
	buff := bytes.Buffer{}
	var ip0 = (ipInt >> 24) & 0xff
	var ip1 = (ipInt >> 16) & 0xff
	var ip2 = (ipInt >> 8) & 0xff
	var ip3 = ipInt & 0xff
	buff.WriteString(strconv.Itoa(ip0))
	buff.WriteByte('.')
	buff.WriteString(strconv.Itoa(ip1))
	buff.WriteByte('.')
	buff.WriteString(strconv.Itoa(ip2))
	buff.WriteByte('.')
	buff.WriteString(strconv.Itoa(ip3))
	return buff.String()
}
