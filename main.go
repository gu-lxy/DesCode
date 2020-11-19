package main

import (
	"DesCode/3des"
	"DesCode/des"
	"DesCode/ecc"
	"DesCode/rsa"
	"bytes"
	"fmt"
)

func main() {
	//des: 块加密
	//des: key、data、mode
	/**
	 * key：密钥
	 * data：明文，即将加密的明文
	 * mode：两种模式，加密和解密
	 */
	//key := []byte("c1906041")
	//
	//data := "遇贵人先立业，遇良人先成家，遇阿姨成家立业!"
	//
	//////加密：crypto
	//block, err := des.NewCipher(key)
	//
	//if err != nil {
	//	panic("初始化密钥错误，请重试！")
	//}
	////dst, src
	//dst := make([]byte, len([]byte(data)))
	////加密过程
	//block.Encrypt(dst, []byte(data))
	//
	//fmt.Println("加密后的内容：", dst)
	//
	////解密
	//deBlock, err := des.NewCipher(key)
	//if err != nil {
	//	panic("初始化密钥错误，请重试！")
	//}
	//deData := make([]byte, len(dst))
	//deBlock.Decrypt(deData, dst)
	//
	//fmt.Println(string(deData))

	//一、对数据进行加密  DES秘钥为8字节、3DES秘钥长度为24字节
	//key := []byte("c1906041")
	//
	//data := "遇贵人先立业，遇良人先成家"

	//1、得到cipher
	//block, err := des.NewCipher(key)
	//if err != nil {
	//	fmt.Println(err.Error())
	//	return
	//}
	//
	////2、对数据明文进行结尾块填充
	//paddingData := PKCS5Padding([]byte(data), block.BlockSize())
	////3、选择模式
	//mode := cipher.NewCBCEncrypter(block, key)
	////4、加密
	//dstData := make([]byte, len(paddingData))
	//mode.CryptBlocks(dstData, paddingData)
	//
	//fmt.Println("加密后的内容：", string(dstData))

	//二、对数据进行解密
	//DES三元素：key、data、mode
	//key1 := []byte("c1906041")
	//data1 := dstData //待解密的数据
	//block1, err := des.NewCipher(key1)
	//if err != nil {
	//	panic(err.Error())
	//}
	//deMode := cipher.NewCBCDecrypter(block1, key1)
	//originalData := make([]byte, len(data1))
	////分组解密
	//deMode.CryptBlocks(originalData, data1)
	//originalData = utils.ClearPKCS5Padding(originalData, block1.BlockSize())
	//fmt.Println("解密后的内容:", string(originalData))

	//一、使用DES进行加解密
	fmt.Println("DES算法：")
	key := []byte("20201112") //des秘钥长度：8字节
	data := "窗含西岭千秋雪，门泊东吴万里船"
	cipherText, err := des.DESEnCrypt([]byte(data), key)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	originText, err := des.DESDeCrypt(cipherText, key)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("DES解密后：", string(originText))

	//二、3DES加解密
	fmt.Println("3DES算法：")
	key1 := []byte("202011122020111220201112") //3des的秘钥长度是24字节
	data1 := "穷在闹市无人问，富在深山有远亲"
	cipherText1, err := _des.TripleDESEncrypt([]byte(data1), key1)
	if err != nil {
		fmt.Println("3DES加密失败:", err.Error())
		return
	}
	originalText1, err := _des.TripleDESDecrypt(cipherText1, key1)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("3DES解密后的内容：", string(originalText1))

	//三、AES算法:
	//fmt.Println("AES算法：")
	//key2 := []byte("2020111212345678") //8
	//data2 := "人生在世不称意，明朝散发弄扁舟"
	//
	//cipherText2, err := aes.AESEnCrypt([]byte(data2), key2)
	//if err != nil {
	//	fmt.Println(err.Error())
	//	return
	//}
	//fmt.Println(cipherText2)
	//cipherText2, err := aes.

	//四、RSA算法的使用
	//fmt.Println("RSA算法：")
	//data4 := "把我的悲伤留给自己"
	////4.1 生成秘钥对
	//priv, err := rsa.CreatePairKeys()
	//if err != nil {
	//	fmt.Println("生成秘钥对出错：", err.Error())
	//	return
	//}
	//
	////4.2 对数据进行加密
	//cipherText4, err := rsa.RSAEncrypt(priv.PublicKey, []byte(data4))
	//if err != nil {
	//	fmt.Println("rsa算法加密失败:", err.Error())
	//	return
	//}
	//originalText4, err := rsa.RSADecrypt(priv, cipherText4)
	//if err != nil {
	//	fmt.Println("rsa算法解密失败:", err.Error())
	//	return
	//}
	//fmt.Println("rsa解密成功，结果是：", string(originalText4))
	//
	////4.3 对原文数据进行签名
	//signText4, err := rsa.RSASign(priv, []byte(data4))
	//if err != nil {
	//	fmt.Println("rsa算法签名失败:", err.Error())
	//	return
	//}
	//
	////4.4 对签名数据进行验证
	//verifyResult, err := rsa.RSAVerify(priv.PublicKey, []byte(data4), signText4)
	//if err != nil {
	//	fmt.Println("rsa算法签名真正失败:", err.Error())
	//}
	//if verifyResult {
	//	fmt.Println("rsa签名验证成功")
	//} else {
	//	fmt.Println("rsa签名验证失败")
	//}

	//五、生成私钥公钥证书文件
	//将生成的私钥保存到硬盘上一个pem文件中，持久化存储下来
	//err = rsa.GenerateKeys("huang")
	//if err != nil {
	//	fmt.Println("生成私钥证书失败：", err.Error())
	//}

	//五.五 从pem文件中读取私钥和公钥
	fmt.Println("==============使用pem文件格式的秘钥进行加解密=========")
	priKey, err := rsa.ReadPemPriKey("rsa_pri_huang.pem")
	if err != nil {
		fmt.Println("读取私钥文件出现错误:", err.Error())
		return
	}
	pubKey, err := rsa.ReadPemPubKey("rsa_pub_huang.pem")
	if err != nil {
		fmt.Println("读取公钥文件出现错误:", err.Error())
		return
	}

	//用读取的公钥进行加密
	data5_5 := "两个黄鹂鸣翠柳，一行白鹭上青天"
	cipherText5_5, err := rsa.RSAEncrypt(*pubKey,[]byte(data5_5))
	//用读取到的私钥进行解密
	originalText5_5, err := rsa.RSADecrypt(priKey,cipherText5_5)
	fmt.Println("解密后的原文是:",string(originalText5_5))


	//六、椭圆曲线数字签名算法
	fmt.Println("==================椭圆曲线数字签名算法===================")
	//① 生成秘钥
	pri, err := ecc.GenerateECDSAKey()
	if err != nil {
		fmt.Println("生成ECDSA秘钥对失败:", err.Error())
		return
	}
	//② 准备数据
	data6 := "张华考上了北京大学,李萍进了软件职业技术学校,我在百货公司当售货员,我们都有光明的前途和未来"

	//③ 数字签名
	r, s, err := ecc.ECDSASign(pri, []byte(data6))
	fmt.Printf("%x\n", r)
	fmt.Printf("%x\n", s)
	fmt.Println("r:",r)
	fmt.Println("s:",s)
	//pem格式：密钥证书文件的格式
	//der格式：

	//④ 数字签名验证
	//data6 = "他开着比克，你提了林肯，我在拼多多砍单自行车成功，我们都有的未来"
	verifyResult := ecc.ECDSAVerify(pri.PublicKey, []byte(data6), r, s)
	if verifyResult {
		fmt.Println("ECDSA数字签名验证成功")
	} else {
		fmt.Println("ECDSA数字签名验证失败")
		return
	}

}

/**
 * 明文数据尾部填充
 */
func PKCS5Padding(text []byte, blockSize int) []byte {
	//计算要填充的快内容的大小
	paddingSize := blockSize - len(text)%blockSize
	paddingText := bytes.Repeat([]byte{byte(paddingSize)}, paddingSize)
	//fmt.Println("明文尾部的内容：",paddingText)
	return append(text, paddingText...)
}
