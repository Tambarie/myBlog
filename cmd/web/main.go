package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

type application struct {
	errorLog *log.Logger
	infoLog *log.Logger
}

func main()  {
	addr := flag.String("addr",":8000","Pass the network address")
	flag.Parse()

	infoLog := log.New(os.Stdout,"INFO\t",log.Ltime|log.Ldate)
	errLog := log.New(os.Stderr,"ERROR\t",log.Ltime|log.Ldate|log.Lshortfile)

	app := &application{
		errorLog: errLog,
		infoLog:  infoLog,
	}

	server := &http.Server{
		Addr:              *addr,
		Handler:           app.routes(),
		ErrorLog:          errLog,

	}


	infoLog.Printf("Starting server on port %s\n",*addr)

	if err := server.ListenAndServe(); err != nil{
		errLog.Fatal(err)
	}

}