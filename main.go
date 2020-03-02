package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"path"

	"github.com/coderguang/GameEngine_go/sgfile"
	"github.com/coderguang/GameEngine_go/sglog"
	"github.com/coderguang/GameEngine_go/sgserver"
	"github.com/coderguang/GameEngine_go/sgthread"
)

func main() {
	sgserver.StartServer(sgserver.ServerTypeLog, "debug", "./log/", log.LstdFlags, true)
	sglog.Info("start check all json file in current dir.")

	files, _ := sgfile.GetAllFile("./")
	jsonFiles := []string{}
	for k, v := range files {
		sglog.Debug("file list:[", k, "]", v)
		fileSuffix := path.Ext(v) //获取文件后缀
		if fileSuffix == ".json" {
			jsonFiles = append(jsonFiles, v)
		}
	}

	errFiles := []string{}
	for k, v := range jsonFiles {
		sglog.Debug("开始解析 json file:[", k, "]", v)
		jsonBytes, err := ioutil.ReadFile(v)
		if err != nil {
			sglog.Error("读取文件失败,file:", v, ",err:", err)
			errFiles = append(errFiles, v)
			continue
		}
		var target interface{}
		if err := json.Unmarshal(jsonBytes, &target); err != nil {
			sglog.Error("解析 ", v, "失败，不是有效的json格式,err:", err)
			errFiles = append(errFiles, v)
			continue
		}
		sglog.Debug("解析 ", v, "成功!")
	}

	if len(errFiles) > 0 {
		sglog.Debug("以下为失败的json文件")
		for _, v := range errFiles {
			sglog.Error(v)
		}
	} else {
		sglog.Info("所有json文件解析成功")
	}

	sgthread.SleepBySecond(2)
	sgserver.StopAllServer()
}
