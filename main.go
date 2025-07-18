package main

import (
	"fmt"
	"go-study/rulego"
	"go-study/rulego/api/types"
	"log"
	"time"
)

// 只需要执行basicb包下的init函数 cttr+shift+D

func main() {
	//初始化一个配置
	config := rulego.NewConfig()
	config.OnDebug = func(chainId, flowType string, nodeId string, msg types.RuleMsg, relationType string, err error) {
		config.Logger.Printf("flowType=%s,nodeId=%s,msgType=%s,data=%s,metaData=%s,relationType=%s,err=%s", flowType, nodeId, msg.Type, msg.Data, msg.Metadata, relationType, err)
	}

	metaData := types.NewMetadata() // 创建一个key类型和值类型都为string的空的map
	//赋值
	metaData.PutValue("id", "1")
	metaData.PutValue("age", "18")
	metaData.PutValue("name", "test01")
	metaData.PutValue("updateAge", "21")
	//加载规则链
	ruleEngine, err := rulego.New("rule01", []byte(chainJsonFile), rulego.WithConfig(config))
	if err != nil {
		log.Fatal(err)
	}

	msg := types.NewMsg(0, "TEST_MSG_TYPE1", types.JSON, metaData, "{\"temperature\":41}")

	ruleEngine.OnMsg(msg)
	fmt.Println("-----等2s-----")
	time.Sleep(time.Second * 1)
	fmt.Println("-----结束-----")

}

var chainJsonFile = `
 {
	 "ruleChain": {
	 		"id":"rule01",
		  "name": "测试规则链",
	    "root": true
	 },
	 "metadata": {
		 "nodes": [
				{
				 "id": "s1",
				 "type": "dbClient",
				 "name": "插入1条记录",
		     "debugMode":true,
				 "configuration": {
			   "driverName":"mysql",
			   "dsn":"root:root@tcp(127.0.0.1:3306)/test",
			   "poolSize":5,
			   "sql":"insert into users (id,name, age) values (?,?,?)",
			   "params":["${metadata.id}", "${metadata.name}", "${metadata.age}"]
				 }
			 },
			{
				 "id": "s2",
				 "type": "dbClient",
				 "name": "查询1条记录",
		     "debugMode":true,
				 "configuration": {
			   "driverName":"mysql",
			   "dsn":"root:root@tcp(127.0.0.1:3306)/test",
			   "sql":"select * from users where id = ?",
			   "params":["${metadata.id}"],
			   "getOne":true
				 }
			 },
		 {
				 "id": "s3",
				 "type": "dbClient",
				 "name": "查询多条记录，参数不使用占位符",
		     "debugMode":true,
				 "configuration": {
			   "driverName":"mysql",
			   "dsn":"root:root@tcp(127.0.0.1:3306)/test",
			   "sql":"select * from users where age >= 18"
				 }
			 },
		 {
				 "id": "s4",
				 "type": "dbClient",
				 "name": "更新记录，参数使用占位符",
		     "debugMode":true,
				 "configuration": {
			   "driverName":"mysql",
			   "dsn":"root:root@tcp(127.0.0.1:3306)/test",
			   "sql":"update users set age = ? where id = ?",
			   "params":["${metadata.updateAge}","${metadata.id}"]
				 }
			 },
		 {
				 "id": "s5",
				 "type": "dbClient",
				 "name": "删除记录",
		     "debugMode":true,
				 "configuration": {
			   "driverName":"mysql",
			   "dsn":"root:root@tcp(127.0.0.1:3306)/test",
			   "sql":"delete from users"
				 }
			 }
		 ],
		 "connections": [
			{
				 "fromId": "s1",
				 "toId": "s2",
				 "type": "Success"
			 },
		{
		 "fromId": "s2",
		 "toId": "s3",
		 "type": "Success"
		 },
		{
		 "fromId": "s3",
		 "toId": "s4",
		 "type": "Success"
		 },
	 {
		 "fromId": "s4",
		 "toId": "s5",
		 "type": "Success"
		 }
		 ]
	 }
 }
 `
