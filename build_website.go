package main

import (
	"encoding/json"
	"html/template"
	"io"
	"io/fs"
	"os"
	"path"
	"strings"
)

type GroupAllotment struct {
	Name string
	Path string
}

type Group struct {
	Name       string
	CountyName string
	Content    map[string]GroupAllotment
}

type County struct {
	Name    string
	Content map[string]map[string]string
}

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
	County          string                      `json:"county"`
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

var dt *template.Template
var gt *template.Template
var ct *template.Template
var it *template.Template

func init() {
	var err error

	dt, err = template.ParseFiles("templates/document.tmpl")
	CheckErr(err)

	gt, err = template.ParseFiles("templates/group.tmpl")
	CheckErr(err)

	ct, err = template.ParseFiles("templates/county.tmpl")
	CheckErr(err)

	it, err = template.ParseFiles("templates/index.tmpl")
	CheckErr(err)
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

func docToHTML(input string, output string) {
	outpath := "web/" + output

	f, err := os.Create(outpath)
	CheckErr(err)

	defer f.Close()
	d := parseDocs(input)

	d.FinalDecision--
	d.ScheduleYears = generateRange(first(d.Schedule))

	dt.Execute(f, d)
}

func groupToHTML(input map[string]string, output string, group_name string, county_name string) {

	var d map[string]GroupAllotment = make(map[string]GroupAllotment)
	outpath := "web/" + output

	for k, v := range input {
		out := strings.Split(k, ".")[0]
		name := strings.ReplaceAll(out, "_", " ")
		out = out + ".html"
		path := strings.Join(strings.Split(v, "/")[1:3], "/")
		out = "/" + path + "/" + out
		d[k] = GroupAllotment{Path: out, Name: name}
	}

	f, err := os.Create(outpath)
	CheckErr(err)

	defer f.Close()

	gt.Execute(f, Group{Name: strings.Title(group_name), CountyName: strings.Title(county_name), Content: d})
}

func countyToHTML(input map[string]map[string]string, output string, county_name string) {
	//var c map[string]map[string]string = make(map[string]map[string]string)
	outpath := "web/" + output

	//for _, _ := range input {
	// out := strings.Split(k, ".")[0]
	// out = out + ".html"
	// path := strings.Join(strings.Split(v, "/")[1:3], "/")
	// out = "/" + path + "/" + out
	// c[k] = out
	//}

	f, err := os.Create(outpath)
	CheckErr(err)

	defer f.Close()

	ct.Execute(f, County{Name: strings.Title(county_name), Content: input})
}

func ParseGroup(p string) (GroupTree map[string]string) {

	// Initialized data
	var err error
	var directory_entries []fs.DirEntry

	// Read the directory entries
	directory_entries, err = os.ReadDir(p)
	CheckErr(err)

	// Allocate a tree
	GroupTree = make(map[string]string)

	// Iterate through the directory entries
	for _, v := range directory_entries {
		if !v.IsDir() {
			GroupTree[v.Name()] = path.Join(p, v.Name())
		}

	}

	return
}

func ParseCounty(p string) (CountyTree map[string]map[string]string) {

	// Initialized data
	var err error
	var directory_entries []fs.DirEntry

	// Read the directory entries
	directory_entries, err = os.ReadDir(p)
	CheckErr(err)

	// Allocate a tree
	CountyTree = make(map[string]map[string]string)

	// Iterate through the directory entries
	for _, v := range directory_entries {
		if v.IsDir() {
			CountyTree[v.Name()] = ParseGroup(path.Join(p, v.Name()))
		}
	}

	return
}

func ParseInput(p string) (InputTree map[string]map[string]map[string]string) {

	// Initialized data
	var err error
	var directory_entries []fs.DirEntry

	// Read the directory entries
	directory_entries, err = os.ReadDir(p)
	CheckErr(err)

	// Allocate a tree
	InputTree = make(map[string]map[string]map[string]string)

	// Iterate through the directory entries
	for _, v := range directory_entries {
		if v == nil {
			continue
		}
		if v.IsDir() {
			InputTree[v.Name()] = ParseCounty(path.Join(p, v.Name()))
		}

	}

	return
}

func main() {

	// // Initialized data
	// var input string

	// // Add the input variable to the argument parser
	// flag.StringVar(&input, "input", "", "Input directory")

	// // Usage message
	// flag.Usage = func() {
	// 	fmt.Fprintf(flag.CommandLine.Output(), "Usage of %s:\n", os.Args[0])
	// 	fmt.Fprintf(flag.CommandLine.Output(), "Build the website.\n")
	// 	flag.PrintDefaults()
	// }

	// // Parse the command line arguments
	// flag.Parse()

	// // Error check
	// if input == "" {

	// 	// Print usage
	// 	flag.Usage()

	// 	// Error
	// 	os.Exit(1)
	// }

	// Parse the input tree
	tree := ParseInput("finished")

	delete(tree, "all")

	for k1, v1 := range tree {

		for k2, v2 := range v1 {
			for k3, v3 := range v2 {
				out := strings.Split(k3, ".")[0]
				out = out + ".html"
				path := strings.Join(strings.Split(v3, "/")[1:3], "/")
				out = path + "/" + out
				docToHTML(v3, out)
			}
			groupToHTML(v2, k1+"/"+k2+"/"+"index.html", k2, k1)
		}
		countyToHTML(v1, k1+"/index.html", k1)
	}

	{

		f, err := os.Create("web/index.html")
		CheckErr(err)

		defer f.Close()

		it.Execute(f, tree)
	}
}
