package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"go.learn/ch4/github"
)

type bucket int

const (
	Day bucket = iota
	Week
	Month
	Year
	MoreThanAYear
)

func getTimeBucket(t time.Time) bucket {
	d := time.Since(t).Truncate(time.Minute)
	//h := float64(time.Duration.Hours(d))
	day := 24 * time.Hour
	if d < day {
		return Day
	} else if d < 7*day {
		return Week
	} else if d < 30*day {
		return Month
	} else if d < 365*day {
		return Year
	} else {
		return MoreThanAYear
	}
}
func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatalf("%s", err)

	}

	fmt.Printf("%d issues:\n", result.TotalCount)
	fmt.Println(result)
	for _, item := range result.Items {
		//fmt.Println(item)
		fmt.Printf("#%-5d %d %s %9.9s %.55s\n",
			item.Number, getTimeBucket(item.CreatedAt),
			item.CreatedAt, item.User.Login, item.Title)
	}
}
