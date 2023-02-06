package charts

import (
	"io"
	"log"
	"math/rand"

	// "math/rand"
	"os"

	"github.com/ehsanghaffar/github-stats-golang/api"
	"github.com/ehsanghaffar/github-stats-golang/utils"
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/opts"
)

type void struct{}

// var count int

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

// Generate fake data for pie chart
func genPieItems() []opts.PieData {
	// subjects := []string{"Ehsan", "Ali", "Parsa", "Sara", "Naser", "David"}
	// userData := api.GetUserData("ehsanghaffar")
	v := api.UserRepos("ehsanghaffar")
	var languages []string
	// languages := make(map[string]int)

	for _, c := range v {
		languages = append(languages, c.Language)
	}

	removeDuplicateStr(languages)

	log.Println(languages)

	items := make([]opts.PieData, 0)
	for i := 0; i < itemCount; i++ {
		items = append(items, opts.PieData{
			Name:  languages[i],
			Value: rand.Intn(400)})
	}

	return items
}

func removeDuplicateStr(strSlice []string) []string {
	allKeys := make(map[string]bool)
	list := []string{}
	for _, item := range strSlice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	log.Println("list", list)
	return list
}

func pieChart() *charts.Pie {
	// create instance
	pie := charts.NewPie()
	pie.SetGlobalOptions(
		charts.WithTitleOpts(
			opts.Title{
				Title:    "Pie Chart",
				Subtitle: "This is Pie",
			},
		),
	)
	pie.SetSeriesOptions()
	pie.AddSeries("Skills",
		genPieItems()).
		SetSeriesOptions(
			charts.WithPieChartOpts(
				opts.PieChart{
					Radius: 200,
				},
			),
			charts.WithLabelOpts(
				opts.Label{
					Show:      true,
					Formatter: "{b}: {c}",
				},
			),
		)
	return pie
}

type PieType struct{}

func (PieType) Example() {
	page := components.NewPage()
	page.AddCharts(
		pieChart(),
	)
	f, err := os.Create("html/pieChart.html")
	if err != nil {
		utils.Perror(err)
	}
	page.Render(io.MultiWriter(f))
}
