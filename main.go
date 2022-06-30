package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/rickar/cal/v2"
	"github.com/rickar/cal/v2/us"
)

type myHoliday struct {
	Holiday  *cal.Holiday
	Selected bool
}

type Workdays struct {
	Year     string   `json: "year"`
	Month    string   `json:"month"`
	Holidays []string `json:"holidays"`
	Days     int      `json: "days"`
}

var mapHolidays = make(map[string]*myHoliday)
var selectedHolidays []string

func main() {
	port:= os.Getenv("PORT")
	if port == ""{
		port = "8080"
	}
	r := gin.Default()
	r.Use(cors.Default())
	r.GET("/api", returnWorkdays)
	r.Run(":"+port)
}

func getWorkDays(year string, month string, holidays []string) Workdays {
	selectedHolidays = holidays

	c := cal.NewBusinessCalendar()

	mapHolidays["NewYear"] = &myHoliday{us.NewYear, false}
	mapHolidays["MLK"] = &myHoliday{us.MlkDay, false}
	mapHolidays["Presidents"] = &myHoliday{us.PresidentsDay, false}
	mapHolidays["Memorial"] = &myHoliday{us.MemorialDay, false}
	mapHolidays["Juneteenth"] = &myHoliday{us.Juneteenth, false}
	mapHolidays["Indenpendence"] = &myHoliday{us.IndependenceDay, false}
	mapHolidays["Labor"] = &myHoliday{us.LaborDay, false}
	mapHolidays["Columbus"] = &myHoliday{us.ColumbusDay, false}
	mapHolidays["Veterans"] = &myHoliday{us.VeteransDay, false}
	mapHolidays["Thanksgiving"] = &myHoliday{us.ThanksgivingDay, false}
	mapHolidays["Christmas"] = &myHoliday{us.ChristmasDay, false}

	for k, _ := range mapHolidays {
		for _, h := range selectedHolidays {
			if h == k {
				mapHolidays[k].Selected = true
			}
		}
	}

	for _, v := range mapHolidays {
		if v.Selected {
			c.AddHoliday(v.Holiday)
		}
	}

	inputYear := year
	inputMonth := month
	inputString := inputYear + "-" + inputMonth + "-" + "01 12:01:01"

	t, err := time.Parse("2006-01-02 03:04:05", inputString)
	if err != nil {
		fmt.Println("could not parse: ", err)
	}

	var workdays int

	if intYear, err := strconv.Atoi(inputYear); err == nil {
		workdays = c.WorkdaysInMonth(intYear, t.Month())
	}

	result := Workdays{Year: inputYear, Month: inputMonth, Holidays: selectedHolidays, Days: workdays}

	return result
}

func returnWorkdays(context *gin.Context) {
	h := context.QueryArray("holidays")
	y := context.Query("year")
	m := context.Query("month")
	// http://localhost:8080/api?year=2023&month=11&holidays=Veterans&holidays=Thanksgiving
	result := getWorkDays(y, m, h)
	context.IndentedJSON(http.StatusOK, result)
}
