// DO NOT EDIT.
package main

import (
	"fmt"
	"log"

	_ "github.com/100steps/gin-layout/dao/migration"
)

var bootPaint string = `
         .IBBBBBBQBBB5:
       iBBBBBBBBBBQBBBBBi
      QBBQBIr...r5BBBQQQBg
     BBBP.         .PBQRRBB
    QBB:   .IQBBDi   :BQMMBb
   .BB:  :QBBBIEQBB.  rBggQB
   .BQ  iBBB.    iBB   BQDRB:   华南理工大学百步梯学生创新中心
    BZ  BBBi RBY  5B   QQgQB.   
    7Q  BBB. PBB  Bq   BQgBQ    技术部制造 
     2: JBBQ  .gBQr   gBMBQ.
      .  BQQB.      .BBQBBi
       :QBQBBBBZ1UqBBBBBB.
      vBBQB1PBBQBQBQBBg:
     SBBBB7   :vu5s7.
    :BBQB.
    ::               
`

func init() {
	fmt.Println(bootPaint)
}

func main() {
	app, err := initApp()
	app.crontab.Excute()
	if err != nil {
		panic(err)
	}
	if err = app.serve(); err != nil {
		log.Fatalln(err.Error())
	}
}
