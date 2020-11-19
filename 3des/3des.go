package _des

import (
	"DesCode/utils"
	"crypto/cipher"
	"crypto/des"
)

/**
 * 该函数用于实现3des算法的加密
 */
func TripleDESEncrypt(origintext []byte, key []byte) ([]byte, error) {

	//三要素：key、data、mode
	//1、实例化一个cipher
	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		return nil, err
	}
	//2、对明文进行尾部填充
	cryptData := utils.PKCS5EndPadding(origintext, block.BlockSize())
	//cryptData := utils.ZerosEndPadding(origintext, block.BlockSize())
	//3、实例化加密模式mode
	mode := cipher.NewCBCEncrypter(block, key[:block.BlockSize()])

	//4、对填充后的明文进行分组加密
	cipherText := make([]byte, len(cryptData))
	mode.CryptBlocks(cipherText, cryptData)
	return cipherText, nil
}

/**
 * 该函数用于对密文进行解密，并返回明文
 */
func TripleDESDecrypt(ciphertext []byte, key []byte) ([]byte, error) {
	//三元素：key、data、mode
	//1、实例化一个cipher
	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		return nil, err
	}
	//2、不需要对密文进行尾部填充，可以直接使用, 实例化mode
	blockMode := cipher.NewCBCDecrypter(block, key[:block.BlockSize()])
	originText := make([]byte, len(ciphertext))
	blockMode.CryptBlocks(originText, ciphertext)
	//清除明文尾部填充
	utils.ClearPKCS5Padding(originText, block.BlockSize())
	return originText, nil
}
