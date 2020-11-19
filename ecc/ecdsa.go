package ecc

import (
	"DesCode/utils"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"math/big"
)

//--------------------生成私钥和公钥的秘钥对-------------------------------//
/**
 * 生成椭圆曲线加密算法中的私钥
 */
func GenerateECDSAKey() (*ecdsa.PrivateKey, error) {
	//1、实例化一个椭圆曲线方程实例
	curve := elliptic.P256()
	pri, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		return nil, err
	}
	return pri, nil
}

//--------------------私钥签名，公钥验签----------------------------------//
/**
 * 使用ECDSA算法中的私钥对数据进行签名
 */
func ECDSASign(pri *ecdsa.PrivateKey, data []byte) (*big.Int, *big.Int, error) {
	hash := utils.Sha256Hash(data)
	r, s, err := ecdsa.Sign(rand.Reader, pri, hash)
	if err != nil {
		return nil, nil, err
	}
	return r, s, nil
}

/**
 * 椭圆曲线数字签名算法中的签名验证函数
 */
func ECDSAVerify(pub ecdsa.PublicKey, data []byte, r *big.Int, s *big.Int) (bool) {
	hash := utils.Sha256Hash(data)
	return ecdsa.Verify(&pub, hash, r, s)
}
