package main

import (
	"database/sql"

	// "github.com/giancarlobastos/loteca-backend/domain"
	// "github.com/giancarlobastos/loteca-backend/scraper"
	"github.com/giancarlobastos/loteca-backend/image"
	_ "github.com/go-sql-driver/mysql"
)

var (
	database *sql.DB
)

func main() {
	// scraper := scraper.NewTransferMarktLeagueScraper()
	// scraper := scraper.NewTransferMarktCupScraper()
	// scraper.GetMatchList(domain.Season{
	// 	Code: "2020",
	// 	Competition: domain.Competition{
	// 		Code:     "CL",
	// 		CodeName: "uefa-champions-league",
	// 	},
	// })
	// scraper.GetMatchList(domain.Round{
	// 	Code: "30",
	// 	Season: domain.Season{
	// 		Code: "2020",
	// 		Competition: domain.Competition{
	// 			Code:     "BRA1",
	// 			CodeName: "campeonato-brasileiro-serie-a",
	// 		},
	// 	},
	// })
	// defer destroy()
	image.ConvertSvgToPngWithChrome("https://s.glbimg.com/es/sde/f/organizacoes/2020/02/12/botsvg.svg", "./assets/test.png")
}

// func init() {
// 	var err error
// 	database, err = sql.Open("mysql", "root:secret@tcp(mysql:3306)/loteca")

// 	if err != nil {
// 		panic(err.Error())
// 	}
// }

// func destroy() {
// 	err := database.Close()

// 	if err != nil {
// 		panic(err.Error())
// 	}
// }
