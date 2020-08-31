package main

import (
    "fmt"
    "time"
)

func main() {
    //const layoutISO = "2006-01-02"
    //date := "2020-08-28"
    const layoutISO = "2006-Jan-02"
    //date := "2020-Mar-28"
    //month := "Mar"
    //year := "2020"
    //date := fmt.Sprintf("%s-%s-01", year, month)
    //t, _ := time.Parse(layoutISO, date)
    var month time.Month = 2
    year := 2020
    t := time.Date(year, month, 1, 0,0,0,0,time.UTC)
    //t_str := t.Format(layoutISO)

    //Display year-month-day in ISO format
    fmt.Println("ISO date: ", t.String()[:10])
    //fmt.Println("ISO date: ", t_str[:10])

    //Display last day of the month
    lastDayOfMonth := time.Date(year, month+1, 0, 0,0,0,0,time.UTC).Day()
    fmt.Println("Last day of month: ", lastDayOfMonth)

    //Display day of week
    fmt.Println("Day of week: ", t.Weekday())

    //Add a day to 't'
    t = t.AddDate(0, 0, 1)
    fmt.Println("Day of week after added 1 day: ", t.Weekday())
    fmt.Printf("t.Weekday() type: %T\n", t.Weekday())
    fmt.Println("is t.Weekday()==6 : ", t.Weekday()==6)
}
