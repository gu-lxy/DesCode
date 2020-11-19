package rsa

import (
	"DesCode/utils"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"flag"
	"io/ioutil"
	"os"
)

const RSA_PRIVATE = "RSA PRIVATE KEY"
const RSA_PUBLIC = "RSA PUBLIC KEY"

/**
 * 私钥：
 * 公钥：
 * 汉森堡
 */
func CreatePairKeys() (*rsa.PrivateKey, error) {
	//1、先生成私钥
	var bits int
	flag.IntVar(&bits, "b", 1024, "密钥长度")
	//fmt.Println(bits)
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return nil, err
	}

	//2、根据私钥生成公钥
	//publicKey := privateKey.Public()
	//3、将私钥和公钥进行返回
	return privateKey, nil
}

//---------------------关于pem证书文件的生成和读取------------------------
/**
 * 根据用户传入的内容，自动创建公私钥，并生成相应格式的证书文件
 */
func GenerateKeys(file_name string) error {
	//1、生成私钥
	pri, err := CreatePairKeys()
	if err != nil {
		return err
	}
	//2.创建私钥文件
	err = generatePriFileByPrivateKey(pri, file_name)
	if err != nil {
		return err
	}
	//3、公钥文件
	err = generatePubFileByPubKey(pri.PublicKey, file_name)
	if err != nil {
		return err
	}
	return nil
}

/**
 * 读取pem文件格式的私钥数据
 */
func ReadPemPriKey(file_name string) (*rsa.PrivateKey, error) {
	blockBytes, err := ioutil.ReadFile(file_name)
	if err != nil {
		return nil, err
	}
	//pem.decode:将byte数据解码为内存中的实例对象
	block, _ := pem.Decode(blockBytes)

	priBytes := block.Bytes
	priKey, err := x509.ParsePKCS1PrivateKey(priBytes)
	return priKey, err
}

/**
 * 读取pem文件格式的公钥数据
 */
func ReadPemPubKey(file_name string) (*rsa.PublicKey, error) {
	blockBytes, err := ioutil.ReadFile(file_name)
	if err != nil {
		return nil, err
	}
	block, _ := pem.Decode(blockBytes)
	pubKey, err := x509.ParsePKCS1PublicKey(block.Bytes)
	return pubKey, err
}

/**
 * 根据给定的私钥数据，生成对应的pem文件
 */
func generatePriFileByPrivateKey(pri *rsa.PrivateKey, file_name string) (error) {

	//根据PKCS1规则，序列化后的私钥
	priStream := x509.MarshalPKCS1PrivateKey(pri)

	//pem文件,此时，privateFile文件为空
	privatFile, err := os.Create("rsa_pri_" + file_name + ".pem") //存私钥的生成的文件
	if err != nil {
		return err
	}

	//pem文件中的格式 结构体
	block := &pem.Block{
		Type:  RSA_PRIVATE,
		Bytes: priStream,
	}

	//将准备好的格式内容写入到pem文件中
	err = pem.Encode(privatFile, block)
	if err != nil {
		return err
	}
	return nil
}

/**
 * 根据公钥生成对应的pem文件，持久化存储
 */
func generatePubFileByPubKey(pub rsa.PublicKey, file_name string) error {
	stream := x509.MarshalPKCS1PublicKey(&pub)

	block := pem.Block{
		Type:  RSA_PUBLIC,
		Bytes: stream,
	}

	pubFile, err := os.Create("rsa_pub_" + file_name + ".pem")
	if err != nil {
		return err
	}
	return pem.Encode(pubFile, &block)
}

//=========================第一种组合：公钥加密，私钥解密==============================//

/**
 * 使用RSA算法对数据进行加密,返回加密后的密文
 */
func RSAEncrypt(key rsa.PublicKey, data []byte) ([]byte, error) {

	return rsa.EncryptPKCS1v15(rand.Reader, &key, data)
}

/**
 * 使用RSA算法对密文数据进行解密，返回解密后的明文
 */
func RSADecrypt(private *rsa.PrivateKey, cipher []byte) ([]byte, error) {
	return rsa.DecryptPKCS1v15(rand.Reader, private, cipher)
}

//=========================第二种组合：私钥签名，公钥验签==============================//

/**
 * 使用RSA算法对数据进行数字签名,并返回签名信息
 */
func RSASign(private *rsa.PrivateKey, data []byte) ([]byte, error) {
	hashed := utils.Md5Hash(data)
	return rsa.SignPKCS1v15(rand.Reader, private, crypto.MD5, hashed)
}

/**
 * 使用RSA算法对数据进行签名验证，并返回签名验证的结果
 * 验证通过，返回true
 * 验证不通过，返回false, 同时error中有错误信息
 */
func RSAVerify(pub rsa.PublicKey, data []byte, signText []byte) (bool, error) {
	hashed := utils.Md5Hash(data)
	err := rsa.VerifyPKCS1v15(&pub, crypto.MD5, hashed, signText)
	return err == nil, err
}
