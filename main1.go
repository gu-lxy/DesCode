package main

import "fmt"

func main() {
	/**
	 * 接口和机构体之间的联系和使用规范
	 */
	//接口：是一套标准,适合于定义共性，抽离并定义出一套标准
	//结构体：实体的描述和定义
	//Block interface
	/**
	 * BlockSize int
	 * Encrypt
	 * Decrypt
	 */

	 /**
	  * 对称加密：des,3des,aes ==> 共性：块加密
	  */

	person1 := NewChinese()
	age := person1.Age()
	if age < 80 {
		fmt.Println("在年龄上符合择偶标准")
	} else {
		fmt.Println("年龄不符合择偶标准")
	}
	//person1.IsMajiang

	person2 := NewJapanese()
	person2.Age()
	//person2.Php

}

/**
 *  择偶标准:
		善良
		为人处世
		身高
		体重
		年龄
		月薪

	*接口：有统一的共性进行判断时，往往使用的是接口
 */
type Person interface {
	Shanliang() bool
	WeiRenChuShi()
	Height() int //身高
	Weight() int //体重
	Age() int    //年龄
	Salary() int //薪水
}

type Chinese struct {
	Name        string
	Sex         string
	IsShanliang bool
	High        int  //身高
	Wei         int  //体重
	AgeNum      int  //年龄
	Money       int  //挣钱能力
	IsMajiang   bool //打或者不打麻将
}

func NewChinese() *Chinese {
	c := &Chinese{
		Name:        "名媛",
		Sex:         "女",
		IsShanliang: true,
		High:        170,
		Wei:         98,
		AgeNum:      21,
		Money:       300000,
		IsMajiang:   true,
	}
	return c
}

func (c *Chinese) Shanliang() bool {
	return c.IsShanliang
}
func (c *Chinese) WeiRenChuShi() {
	fmt.Println(c.Name + "伟人处理能力很好")
}
func (c *Chinese) Height() int {
	return c.High
}
func (c *Chinese) Weight() int {
	return c.Wei
}
func (c *Chinese) Age() int {
	return c.AgeNum
}
func (c *Chinese) Salary() int {
	return c.Money
}

func (c *Chinese) Majiang() {
	if c.IsMajiang {
		fmt.Println(c.Name + "会打麻将")
	} else {
		fmt.Println(c.Name + "不会打麻将")
	}
}

type Japanese struct {
	Name        string
	AgeNum      int
	IsShanliang bool
	High        int
	Wei         int
	Money       int
	Php         bool //是否会编程语言
}

func NewJapanese() *Japanese {
	j := &Japanese{
		Name:        "小仓优子",
		AgeNum:      23,
		IsShanliang: true,
		High:        165,
		Wei:         90,
		Money:       2000000,
		Php:         true,
	}
	return j
}

func (j *Japanese) Shanliang() bool {
	return j.IsShanliang
}
func (j *Japanese) WeiRenChuShi() {
	fmt.Println(j.Name + "伟人处理能力很好")
}
func (j *Japanese) Height() int {
	return j.High
}
func (j *Japanese) Weight() int {
	return j.Wei
}
func (j *Japanese) Age() int {
	return j.AgeNum
}
func (j *Japanese) Salary() int {
	return j.Money
}
func (j *Japanese) Eat(food string) {
	fmt.Println(j.Name+"喜欢吃:", food)
}

func (j *Japanese) Phps() bool {
	return j.Php
}
