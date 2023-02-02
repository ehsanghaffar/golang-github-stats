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

var (
	itemCount = 7
	day       = []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
)

func genItems() []opts.BarData {
	items := make([]opts.BarData, 0)
	for i := 0; i < itemCount; i++ {
		items = append(items, opts.BarData{Value: rand.Intn(420)})
	}
	return items
}

func barChart() *charts.Bar {
	bar := charts.NewBar()
	bar.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title:    "Bar Chart",
			Subtitle: "This is bar Chart",
			Link:     "https://ehsanghaffarii.ir/chart/bar",
			Right:    "10%",
		}),
		charts.WithToolboxOpts(opts.Toolbox{Show: true}),
	)

	bar.SetXAxis(day).
		AddSeries("Series 1", genItems()).
		AddSeries("Series 2", genItems())
	return bar
}

type Exampler interface {
	Example()
}

type BarType struct{}

func (BarType) Example() {
	page := components.NewPage()
	page.AddCharts(
		barChart(),
	)
	f, err := os.Create("html/barChart.html")
	if err != nil {
		utils.Perror(err)
	}
	page.Render(io.MultiWriter(f))
}
