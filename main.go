package main

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"fmt"
)

func main() {
	//des：块加密
	/*des: key, data, mode
	  key: 秘钥
	  data: 明文，即将加密的明文
	   mode： 两种模式，加密和解密
	*/
	//key := []byte("C190604a")
	//
	//data := "遇良人先成家，遇贵人先立业，遇阿姨成家立业"
	////加密：crypto
	//block, err := des.NewCipher(key)
	//if err != nil {
	//	panic("初始化错误，请重试")
	//}
	////dst, src
	//dst := make([]byte, len([]byte(data)))
	////加密的过程
	//block.Encrypt(dst, []byte(data))
	//
	//fmt.Println("加密后的内容：",dst)
	//
	////解密
	//deBlock, err := des.NewCipher(key)
	//if err != nil {
	//	panic("初始化错误，请重试")
	//}
	//deData := make([]byte, len(dst))
	//deBlock.Decrypt(deData,dst)
	//fmt.Println("解密后的数据", string(deData))
	//一、 对数据进行加密
	key := []byte("C190604a")

	data := "遇小刘先成家，遇小胡先立业，遇小丁成家立业"
	//1、 得到cipher
	block, _ := des.NewCipher(key)
	//2、 对数据明文进行结尾块填充
	paddingData := EndPadding([]byte(data), block.BlockSize())
	//3、 选择模式
	mode := cipher.NewCBCEncrypter(block, key)
	//4、 加密
	dstData := make([]byte, len(paddingData))
	mode.CryptBlocks(dstData, paddingData)
	fmt.Println("加密后的内容：",dstData)

	//二、 对数据进行解密
	key1 := []byte("C190604a")
	data1 := dstData //待解密的数据
	block1, err := des.NewCipher(key1)
	if err != nil {
		panic(err.Error())
	}
	deMode := cipher.NewCBCDecrypter(block1,key1)
	originalData := make([]byte, len(data1))
	deMode.CryptBlocks(originalData,data1)
	fmt.Println("解密后的内容：",string(originalData))

}


//明文数据尾部填充
func EndPadding(text []byte, blockSize int) []byte{
	//计算要填充的块的大小
	paddingSize := blockSize - len(text)%blockSize
    paddingText := bytes.Repeat([]byte{byte(paddingSize)},paddingSize)
    fmt.Println("填充的内容：",paddingText)
    return append(text,paddingText...)
}

