package rsa

import (
	"DesCode/utils"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"flag"
)

//私钥：
//公钥：
func GratePairKeys() (*rsa.PrivateKey, *rsa.PublicKey, error) {
	//1、先生成私钥
	var bits int
	flag.IntVar(&bits, "b",2048,"秘钥长度")
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return nil, nil, err
	}
	//2、根据私钥生成公钥
	publicKey := privateKey.Public()
	//3、蒋公要和私钥返回
	return privateKey, publicKey, nil
}

//使用RSA算法对数据进行加密，返回加密后的密文
func RSAEncrypt(key rsa.PublicKey,data []byte) ([]byte,error) {
	return rsa.EncryptPKCS1v15(rand.Reader, &key, data)
}

//使用RSA算法对数据进行解密，返回解密后的明文
func RSADecrypt(private *rsa.PublicKey, cipher []byte) ([]byte, error) {
	return rsa.EncryptPKCS1v15(rand.Reader, private, cipher)
}

//=========第二种组合：私钥签名， 公钥验签

//使用RSA算法对数据进行签名，并返回签名信息
func RSASian(private *rsa.PublicKey, data []byte) ([]byte, error) {
	hashed := utils.Md5Hash(data)
	return rsa.SignPKCS1v15(rand.Reader, private, crypto.MD5, hashed)
}

func RSAVerify(pub rsa.PublicKey, data []byte, signText []byte) (bool, error) {
	hashed := utils.Md5Hash(data)
	err := rsa.VerifyPKCS1v15(&pub, crypto.MD5, hashed, signText )
	return err == nil, err
}