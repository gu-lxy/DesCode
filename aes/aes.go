package aes

//使用AES算法对明文进行加密
//func AesEnCrypt(origin []byte, key  []byte) ([]byte, error) {
//	//三元素：key、data、mode
//	block, err := aes.NewCipher(key)
//	if err != nil {
//		return nil, err
//	}
//	//2、对明文数据进行填充
//	//2、对明文尾部进行填充
//	cryptData := aes.(origin,block.BlockSize())
//	//3、实例化加密模式mode
//	mode := cipher.NewCBCEncrypter(block, key)
//	//4、对填充后的明文进行分组加密
//	cipherText := make([]byte, len(cryptData))
//	mode.CryptBlocks(cipherText, cryptData)
//	return cipherText, nil
//}
