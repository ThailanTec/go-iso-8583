package service

import (
	"github.com/moov-io/iso8583"
	"github.com/moov-io/iso8583/specs"
	"os"
)

func ReaderISOASCII(file []byte) error {
	isomessage := iso8583.NewMessage(specs.Spec87ASCII)
	isomessage.Unpack(file)
	iso8583.DoNotFilterFields()
	iso8583.Describe(isomessage, os.Stdout)
	return nil
}

func ReaderISOHEX(file []byte) error {
	isomessage := iso8583.NewMessage(specs.Spec87Hex)
	isomessage.Unpack(file)

	iso8583.DoNotFilterFields()
	iso8583.Describe(isomessage, os.Stdout)
	return nil
}
