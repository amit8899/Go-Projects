package main

import (
	"os"
	"sync"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	log "github.com/sirupsen/logrus"
)

var (
	db       *gorm.DB
	file     *os.File
	ch1, ch2 chan Player
	players  []Player
	wg       sync.WaitGroup
)

type Player struct {
	gorm.Model
	JerseyNo      int    `gorm:"primaryKey"`
	Name          string `gorm:"name"`
	Age           string `gorm:"age"`
	NoOfCenturies int
}

func init() {
	dsn := "root:toor@tcp(127.0.0.1:3306)/gormdb?charset=utf8mb4&parseTime=True&loc=Local"
	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	db = d
	db.AutoMigrate(&Player{})

	f, e := os.Create("file.log")
	if e != nil {
		panic(e)
	}
	file = f

	players = []Player{
		{JerseyNo: 1, Name: "KL Rahul", Age: "28", NoOfCenturies: 34},
		{JerseyNo: 2, Name: "Virat", Age: "35", NoOfCenturies: 49},
		{JerseyNo: 3, Name: "Rohit", Age: "33", NoOfCenturies: 30},
		{JerseyNo: 4, Name: "Shubhman", Age: "25", NoOfCenturies: 5},
		{JerseyNo: 5, Name: "Surya", Age: "33", NoOfCenturies: 4},
	}

}

func main() {
	// log.SetLevel(log.DebugLevel)
	log.SetOutput(file)

	ch1 = make(chan Player, 5)
	ch2 = make(chan Player, 5)

	defer wg.Wait()
	wg.Add(3)

	//1
	go func() {
		for i, player := range players {
			ch1 <- player
			log.Println("player added", i, player)
			db.Create(&player)
		}
		wg.Done()
	}()

	//2 update
	go func() {
		for i := 1; i < 6; i++ {
			time.Sleep(time.Second)

			var p Player
			db.Where("jersey_no = ?", i).First(&p)
			UpdatePlayerCenturies(&p)
			ch2 <- p
			log.Println("player updated", i, p)
			db.Save(&p)

		}
		wg.Done()
	}()

	//3 print
	go func() {
		for {
			select {
			case p := <-ch1:
				log.Info(p)

			case p := <-ch2:
				log.Info(p)
			}
		}
	}()

	//4
	go func() {
		defer wg.Done()
		time.Sleep(time.Second * 6)

		res := db.Delete(&Player{}, []int{1, 2, 3, 4, 5})
		log.Info(res.RowsAffected)

		// for i := 0; i < 5; i++ {
		// 	db.Delete(&Player{}, i+1)
		// 	log.Info("Deleted:", i+1)
		// }
	}()
}

func UpdatePlayerCenturies(p *Player) {
	switch p.JerseyNo {
	case 1:
		p.NoOfCenturies = 36
	case 2:
		p.NoOfCenturies = 50
	case 3:
		p.NoOfCenturies = 31
	case 4:
		p.NoOfCenturies = 7
	case 5:
		p.NoOfCenturies = 6
	}
}
