package main

import (
	"encoding/json"

	"log"
	"net/url"
	"os"
)

//判断路径是否存在
func pathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

//创建路径
func makePath(path string) (err error) {
	x, err := pathExists(path)
	if err != nil {
		return
	}
	if !x {
		err = os.Mkdir(path, os.ModePerm)
		if err != nil {
			return
		}
	}
	return
}

//保存内容到文件
func saveFile(path string, data []byte) (err error) {
	f, err := os.Create(path)
	if err != nil {
		return
	}
	defer func() {
		err := f.Close()
		if err != nil {
			log.Println(err)
		}
	}()
	_, err = f.Write(data)
	if err != nil {
		return
	}
	return
}

//把对象保存为JSON格式文件
func saveJSONFile(path string, v any) (err error) {
	f, err := os.Create(path)
	if err != nil {
		return
	}

	defer func() {
		err := f.Close()
		if err != nil {
			log.Println(err)
		}
	}()

	encoder := json.NewEncoder(f)
	encoder.SetIndent("", "\t")
	err = encoder.Encode(v)
	if err != nil {
		return
	}

	return
}

//把对象转换为JSON格式的字符串
func jsonMarshal(v any) (s string, err error) {
	b, err := json.Marshal(v)
	if err != nil {
		return
	}
	s = string(b)
	return
}

//提取网址中的网站信息
func getUrlHost(u string) (host string, err error) {
	a, err := url.Parse(u)
	if err != nil {
		return
	}
	host = a.Host
	return
}
