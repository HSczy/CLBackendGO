package utils

import "os"

//@author: [HSczy](https://github.com/HSczy)
//@function: PathExists
//@description: 文件目录是否存在
//@param: path string
//@return: bool, error

func PathExists(path string) (ok bool ,err error) {
	_, err = os.Stat(path)
	if err != nil{
		if os.IsNotExist(err){
			return false, nil
		}
		return false, err
	}
	return true,nil
}


func CreateDir(dirs ...string) (err error){
	for _,dir := range dirs{
		exist, err := PathExists(dir)
		if err != nil{
			return err
		}
		if !exist{
			if err = os.MkdirAll(dir,os.ModePerm);err!=nil{
				return err
			}
		}
	}
	return err
}