// Nameday generator for calendar util
package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

//Namedays type
type Namedays map[string][]string

var months = map[int]string{
	1:  "Jan",
	2:  "Feb",
	3:  "Mar",
	4:  "Apr",
	5:  "May",
	6:  "Jun",
	7:  "Jul",
	8:  "Aug",
	9:  "Sep",
	10: "Oct",
	11: "Nov",
	12: "Dec",
}

//NameStruc blah
type NameStruc struct {
	Month int
	Day   int
	Names []string
}

//NameStruct blah
type NameStruct []NameStruc

func (n NameStruct) Len() int {
	return len(n)
}

func (n NameStruct) Less(i, j int) bool {
	if n[i].Month == n[j].Month {
		return n[i].Day < n[j].Day
	}
	return n[i].Month < n[j].Month
}

func (n NameStruct) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}

func main() {
	file, err := os.Open("namedays.json")
	if err != nil {
		panic(err)
	}
	names := make(Namedays)
	json.NewDecoder(file).Decode(&names)
	r := convert(names)
	sort.Sort(r)
	for _, v := range r {
		if v.Month == 2 && v.Day == 29 {
			continue
		}
		daystr := ""
		if v.Day < 10 {
			daystr = fmt.Sprintf("0%d", v.Day)
		} else {
			daystr = fmt.Sprintf("%d", v.Day)
		}
		fmt.Printf("%s %s\tVārdadienu svin: %s\n", daystr, months[v.Month], strings.Join(v.Names, ", "))
	}
}

func convert(n Namedays) (r NameStruct) {
	r = make(NameStruct, 0)
	for k, v := range n {
		arr := strings.Split(k, "-")
		arr[0] = strings.TrimLeft(arr[0], "0")
		arr[1] = strings.TrimLeft(arr[1], "0")
		mon, _ := strconv.Atoi(arr[0])
		day, _ := strconv.Atoi(arr[1])
		r = append(r, NameStruc{
			Month: mon,
			Day:   day,
			Names: v,
		})
	}
	return r
}
