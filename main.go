package main

import (
	"fmt"
	"github.com/ThailanTec/go-iso-8583/service"
	"os"
)

func main() {
	arq1, _ := os.ReadFile("./data/financial_transaction_message.dat")
	arq2, _ := os.ReadFile("./data/message_with_hex_bcd.dat")

	err := service.ReaderISOASCII(arq1)
	fmt.Println("----------Separação dos Dados----------")
	err = service.ReaderISOHEX(arq2)
	if err != nil {
		return
	}
}
