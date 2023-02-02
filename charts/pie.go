package charts

import (
	"io"
	"math/rand"
	"os"

	"github.com/ehsanghaffar/github-stats-golang/utils"
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/opts"
)

// Generate fake data for pie chart
func genPieItems() []opts.PieData {
	subjects := []string{"Ehsan", "Ali", "Parsa", "Sara", "Naser", "David"}
	items := make([]opts.PieData, 0)
	for i := 0; i < 6; i++ {
		items = append(items, opts.PieData{
			Name:  subjects[i],
			Value: rand.Intn(400)})
	}
	return items
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
