package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"log"
)

const APP_VERSION = "3.0"

func main() {
	/*
	err := initConfig()
	if err != nil {
		log.Fatalf("err : %s", err)
	}
	Clean(&service.S)
	*/
	fmt.Println("Hello Lux")
	//fmt.Println("\n1111...")
	tcpArgs := TCPArgs{}

	//heroku
	SYSPORT := os.Getenv("PORT")
	laddr :=":"+SYSPORT
	//laddr :="1880:"
	ptype := "tcp"
	paddr := "18.182.150.118:16000"//
	timeout := 5000
	isTLS := false
	poolSize :=0
	interval :=0

	tcpArgs.Local = &laddr
	tcpArgs.Parent = &paddr
	tcpArgs.Timeout = &timeout
	tcpArgs.ParentType = &ptype
	tcpArgs.IsTLS = &isTLS
	tcpArgs.PoolSize = &poolSize
	tcpArgs.CheckParentInterval = &interval

	Regist("tcp", NewTCP(), tcpArgs)

	var serviceName string = "tcp"

	service, err := Run(serviceName)
	if err != nil {
		log.Fatalf("run service [%s] fail, ERR:%s", service, err)
	}

	Clean(&service.S)
}

func Clean(s *Service) {
	signalChan := make(chan os.Signal, 1)
	cleanupDone := make(chan bool)
	signal.Notify(signalChan,
		os.Interrupt,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)
	go func() {
		for _ = range signalChan {
			fmt.Println("\nReceived an interrupt, stopping ..")
			(*s).Clean()
			cleanupDone <- true
		}
	}()
	<-cleanupDone
}

