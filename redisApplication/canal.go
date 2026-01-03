package redisApplication

// canal 缓存双写
import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/withlin/canal-go/client"
	pbe "github.com/withlin/canal-go/protocol/entry"
)

var hashKey = "canal_redis.baseCreditInfo"

func Run() {
	// 192.168.199.17 替换成你的canal server的地址
	// example 替换成-e canal.destinations=example 你自己定义的名字
	connector := client.NewSimpleCanalConnector("192.168.0.103", 11111, "", "", "example", 60000, 60*60*1000)
	err := connector.Connect()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	// https://github.com/alibaba/canal/wiki/AdminGuide
	//mysql 数据解析关注的表，Perl正则表达式.
	//
	//多个正则之间以逗号(,)分隔，转义符需要双斜杠(\\)
	//
	//常见例子：
	//
	//  1.  所有表：.*   or  .*\\..*
	//	2.  canal schema下所有表： canal\\..*
	//	3.  canal下的以canal打头的表：canal\\.canal.*
	//	4.  canal schema下的一张表：canal\\.test1
	//  5.  多个规则组合使用：canal\\..*,mysql.test1,mysql.test2 (逗号分隔)

	err = connector.Subscribe("canal_redis\\.baseCreditInfo")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	for {
		message, err := connector.Get(100, nil, nil)
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
		batchId := message.Id
		if batchId == -1 || len(message.Entries) <= 0 {
			time.Sleep(300 * time.Millisecond)
			continue
		}

		handleEntries(message.Entries)
	}
}

func handleEntries(entrys []pbe.Entry) {

	for _, entry := range entrys {

		entryType := entry.GetEntryType()
		if entryType == pbe.EntryType_TRANSACTIONBEGIN || entryType == pbe.EntryType_TRANSACTIONEND {
			continue
		}

		rowChange := new(pbe.RowChange)
		err := proto.Unmarshal(entry.GetStoreValue(), rowChange)
		checkError(err)
		if rowChange == nil {
			continue
		}

		eventType := rowChange.GetEventType()
		header := entry.GetHeader()
		fmt.Println(fmt.Sprintf("================> binlog[%s : %d],name[%s,%s], eventType: %s", header.GetLogfileName(), header.GetLogfileOffset(), header.GetSchemaName(), header.GetTableName(), header.GetEventType()))

		ctx := context.Background()
		for _, rowData := range rowChange.GetRowDatas() {
			if eventType == pbe.EventType_DELETE {
				card, _ := rowDataToJson(rowData.GetBeforeColumns())
				RedisCluster().HDel(ctx, hashKey, card.Id)
			} else if eventType == pbe.EventType_INSERT {
				fmt.Println("-------> after")
				card, s := rowDataToJson(rowData.GetAfterColumns())
				fmt.Println(s)
				RedisCluster().HSet(ctx, hashKey, card.Id, s)
			} else if eventType == pbe.EventType_UPDATE {
				fmt.Println("-------> before")
				fmt.Println(rowDataToJson(rowData.GetBeforeColumns()))
				fmt.Println("-------> after")
				card, s := rowDataToJson(rowData.GetAfterColumns())
				fmt.Println(card)
				fmt.Println(s)
				RedisCluster().HSet(ctx, hashKey, card.Id, s)
			} else {
				fmt.Println("-------> before")
				printColumn(rowData.GetBeforeColumns())
				fmt.Println("-------> after")
				printColumn(rowData.GetAfterColumns())
			}
		}

	}
}

func printColumn(columns []*pbe.Column) {
	for _, col := range columns {
		fmt.Println(fmt.Sprintf("%s : %s  update= %t", col.GetName(), col.GetValue(), col.GetUpdated()))
	}
}

func rowDataToJson(columns []*pbe.Column) (BankCard, string) {
	rowDataMap := make(map[string]string)
	for _, col := range columns {
		rowDataMap[col.GetName()] = col.GetValue()
		//fmt.Println(fmt.Sprintf("%s : %s  update= %t", col.GetName(), col.GetValue(), col.GetUpdated()))
	}
	bytes, _ := json.Marshal(rowDataMap)
	var card BankCard
	json.Unmarshal(bytes, &card)
	fmt.Println(card)
	return card, string(bytes)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

type BankCard struct {
	BankName   string `json:"bankName"`
	BillDate   string `json:"billDate"`
	CardNo     string `json:"cardNo"`
	CardType   string `json:"cardType"`
	CreateTime string `json:"createTime"`
	DueDate    string `json:"dueDate"`
	Id         string `json:"id"`
	Score      string `json:"score"`
	UpdateTIme string `json:"updateTIme"`
}
