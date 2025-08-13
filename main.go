package main

import (
	"fmt"
	"go-study/utils"
	"os"
)

// 只需要执行basicb包下的init函数 ctr+shift+D

func main() {
	wd, err := os.Getwd() // 获取当前工作目录的绝对路径
	if err != nil {
		fmt.Println("Error:", err)
	}
	filepath := wd + "/utils/test.docx" // 相对路径
	utils.Execute(filepath)
	//初始化一个配置
	// config := rulego.NewConfig()
	// config.OnDebug = func(chainId, flowType string, nodeId string, msg types.RuleMsg, relationType string, err error) {
	// 	config.Logger.Printf("flowType=%s,nodeId=%s,msgType=%s,data=%s,metaData=%s,relationType=%s,err=%s", flowType, nodeId, msg.Type, msg.Data, msg.Metadata, relationType, err)
	// }

	// metaData := types.NewMetadata() // 创建一个key类型和值类型都为string的空的map
	// //赋值
	// metaData.PutValue("id", "8")
	// metaData.PutValue("phone", "19998881238")
	// metaData.PutValue("username", "胡八一")
	// metaData.PutValue("updatePhone", "21")
	// //加载规则链
	// ruleEngine, err := rulego.New("rule01", []byte(chainJsonFile), rulego.WithConfig(config))
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// msg := types.NewMsg(0, "TEST_MSG_TYPE1", types.JSON, metaData, "{\"temperature\":41}")

	// ruleEngine.OnMsg(msg)
	// fmt.Println("-----等2s-----")
	// time.Sleep(time.Second * 2)
	// fmt.Println("-----结束-----")

}

// var chainJsonFile = `
//  {
// 	 "ruleChain": {
// 	 		"id":"rule01",
// 		  "name": "测试规则链",
// 	    "root": true
// 	 },
// 	 "metadata": {
// 		 "nodes": [
// 				{
// 				 "id": "s1",
// 				 "type": "dbClient",
// 				 "name": "插入1条记录",
// 		     "debugMode":true,
// 				 "configuration": {
// 			   "driverName":"mysql",
// 			   "dsn":"root:admin@tcp(127.0.0.1:3306)/go-admin",
// 			   "poolSize":5,
// 			   "sql":"insert into sys_user (user_id,username, phone) values (?,?,?)",
// 			   "params":["${metadata.id}", "${metadata.username}", "${metadata.phone}"]
// 				 }
// 			 }
// 		 ],
// 		 "connections": [
// 			{
// 				 "fromId": "s1",
// 				 "toId": "s2",
// 				 "type": "Success"
// 			 },
// 		{
// 		 "fromId": "s2",
// 		 "toId": "s3",
// 		 "type": "Success"
// 		 },
// 		{
// 		 "fromId": "s3",
// 		 "toId": "s4",
// 		 "type": "Success"
// 		 },
// 	 {
// 		 "fromId": "s4",
// 		 "toId": "s5",
// 		 "type": "Success"
// 		 }
// 		 ]
// 	 }
//  }
//  `
