package aes

import (
	"CryptCode/utils"
	"crypto/aes"
	"crypto/cipher"
)

/**
 * 使用AES算法对明文进行加密
 */
func AESEnCrypt(origin []byte, key []byte) ([]byte, error) {
	//3元素：key、data、mode
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	//2、对明文数据进行尾部填充
	cryptData := utils.PKCS5EndPadding(origin, block.BlockSize())
	//3、实例化加密mode
	blockMode := cipher.NewCBCEncrypter(block, key)
	//4、加密
	cipherData := make([]byte, len(cryptData))
	blockMode.CryptBlocks(cipherData, cryptData)
	return cipherData, nil
}
