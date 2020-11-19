package utils

import (
	"crypto/md5"
	"crypto/sha256"
)

/**
 * 使用MD5哈希函数进行hash
 */
func Md5Hash(data []byte) ([]byte) {
	md5Hash := md5.New()
	md5Hash.Write(data)
	return md5Hash.Sum(nil)
}

/**
 * 使用SHA256算法对数据进行Hash计算
 */
func Sha256Hash(data []byte) ([]byte) {
	sha256Hash := sha256.New()
	sha256Hash.Write(data)
	return sha256Hash.Sum(nil)
}
