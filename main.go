package main

import (
	"flag"
	"fmt"
	"github.com/google/uuid"
	"log"
	"termdb/structs"
)

func main() {
	titleArg := flag.String("title", "", "title of message")
	msgArg := flag.String("save", "", "The text that will be stored in the db")
	getArg := flag.Int("get", -1, "the record to get by number")
	flag.Parse()

	PrepareSql()

	ProcessArgs(msgArg, titleArg, getArg)
}

func ProcessArgs(msgArg *string, titleArg *string, getArg *int) {
	if len(*msgArg) > 0 {
		if len(*titleArg) > 0 {
			StoreMessageTitle(msgArg, titleArg)
		} else {
			StoreMessageNoTitle(msgArg)
		}
		GetAllMessages()
		return
	}
	if *getArg == -1 {
		GetAllMessages()
		return
	}
	GetMessageByIdNum(getArg)
}

func StoreMessageTitle(arg *string, title *string) {
	InsertMessage(arg, title)
}

func StoreMessageNoTitle(arg *string) {
	uuid, _ := uuid.NewUUID()
	title := uuid.String()
	InsertMessage(arg, &title)
}

func GetAllMessages() {
	results, err := GetAllRecords()
	if err != nil {
		log.Fatal(err)
	}
	r := *results
	for i := 0; i < len(r); i++ {
		a := r[i]
		PrintMessageInfo(&a)
	}
}

func GetMessageByIdNum(id *int) {
	result, err := GetMessageById(id)
	if err != nil {
		return
	}
	PrintFullMessageInfo(result)
}

func PrintMessageInfo(message *structs.MessageInfo) {
	a := fmt.Sprintf("%d)  %s \r", message.Id, message.Title)
	fmt.Println(a)
}

func PrintFullMessageInfo(result *structs.MessageInfo) {
	a := fmt.Sprintf("%s \n %s", result.Title, result.Message)
	fmt.Println(a)
}
