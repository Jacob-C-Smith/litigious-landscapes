package main

import (
	"encoding/json"
	"html/template"
	"io"
	"os"
)

type AnnualInterval struct {
	Start string `json:"start"`
	End   string `json:"end"`
}

type Livestock struct {
	Type     string `json:"type"`
	Quantity int    `json:"quantity"`
}

type AUM struct {
	Active    int `json:"active"`
	Adaptive  int `json:"adaptive"`
	Custodial int `json:"custodial"`
	Suspended int `json:"suspended"`
}

type AllotmentPermit struct {
	Name          string         `json:"name"`
	Number        int            `json:"number"`
	Livestock     Livestock      `json:"livestock"`
	GrazingPeriod AnnualInterval `json:"grazing period"`
	PublicLand    int            `json:"public land"`
	Aum           AUM            `json:"AUM"`
}

type Document struct {
	Name            string                      `json:"name"`
	Group           string                      `json:"group"`
	Date            string                      `json:"date"`
	CurrentPermit   []AllotmentPermit           `json:"current permit"`
	NextPermit      []AllotmentPermit           `json:"next permit"`
	Issues          []string                    `json:"issues"`
	Alternatives    []string                    `json:"alternatives"`
	FinalDecision   int                         `json:"final decision"`
	ProtestQuantity int                         `json:"protest quantity"`
	Schedule        map[string][]AnnualInterval `json:"schedule"`
	ScheduleYears   []int
	Conditions      map[string]map[string][]string `json:"resource conditions"`
}

func CheckErr(e error) {
	if e != nil {
		panic(e)
	}
}

func parseDocs(path string) (document Document) {
	file, err := os.Open(path)
	CheckErr(err)

	data, err := io.ReadAll(file)
	CheckErr(err)

	CheckErr(json.Unmarshal(data, &document))
	file.Close()
	return
}

func first(d map[string][]AnnualInterval) int {
	for _, v := range d {
		return len(v)
	}
	return 0
}

func generateRange(n int) []int {
	nums := make([]int, n)
	for i := range nums {
		nums[i] = i + 1
	}
	return nums
}

func main() {
	t, err := template.ParseFiles("document.tmpl")
	CheckErr(err)

	f, err := os.Create("web/out.html")
	CheckErr(err)

	defer f.Close()
	d := parseDocs(os.Args[1])

	d.FinalDecision--
	d.ScheduleYears = generateRange(first(d.Schedule))

	t.Execute(f, d)
}
