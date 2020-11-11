package _des

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
)

//该函数用于实现3des算法加密
func TripleDesEncrypt(origintext []byte, key []byte) ([]byte, error) {
	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		return nil, err
	}
	//2、对明文尾部进行填充
	cryptData := PKCS5Endpadding(origintext,block.BlockSize())
	//3、实例化加密模式mode
	mode := cipher.NewCBCEncrypter(block, key)
	//4、对填充后的明文进行分组加密
	cipherText := make([]byte, len(cryptData))
	mode.CryptBlocks(cipherText, cryptData)
	return cipherText, nil
}

//该函数用于对密文进行解密，并返回明文
func TripleDesDecrypt(ciphertext []byte,key []byte) []byte {
	//三元素： key、data、mode
	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		return err
	}
	//不需要对密文进行尾部填充
	blockMode := cipher.NewCBCDecrypter(block, key)
	originText := make([]byte, len(ciphertext))
	block.Encrypt(originText, )
}


//该函数将对明文进行尾部填充，采用PKCS5方式
func PKCS5Endpadding(text []byte,size int) []byte {
	paddingSize := size - len(text)%size
	paddingText := bytes.Repeat([]byte{byte(paddingSize)},paddingSize)
	//fmt.Println("填充的内容：",paddingText)
	return append(text,paddingText...)
}

//
