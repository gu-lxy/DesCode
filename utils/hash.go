package utils

import "crypto/md5"

func Md5Hash(data []byte) ([]byte) {
	md5Hash := md5.New()
	md5Hash.Write(data)
	return md5Hash.Sum(nil)
}

func Sha256Hash(data)  {

}
