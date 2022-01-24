package main

import (
	"io"
	"os"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"
	"github.com/go-echarts/go-echarts/v2/opts"
)

// generate random data for bar chart
func addBarItems(values []float64) []opts.BarData {
	items := make([]opts.BarData, 0)
	for _, value := range values {
		items = append(items, opts.BarData{Value: value})
	}

	return items
}

func bar() *charts.Bar {
	// create a new bar instance
	bar := charts.NewBar()
	// set some global options like Title/Legend/ToolTip or anything else
	bar.SetGlobalOptions(
		charts.WithColorsOpts(opts.Colors{"#FFCC00", "#D34E2B", "#EF923A", "#F8E046"}),
		charts.WithTitleOpts(opts.Title{
			Title:    "Disminución",
			Subtitle: "Número de colmenas por región(en millones)",
		}),
		charts.WithToolboxOpts(opts.Toolbox{
			Show:  true,
			Right: "20%",
			Feature: &opts.ToolBoxFeature{
				SaveAsImage: &opts.ToolBoxFeatureSaveAsImage{
					Show:  true,
					Type:  "png",
					Title: "Guardar como imagen",
				},
			}},
		),
		charts.WithLegendOpts(opts.Legend{
			Show: true,
		}),
		charts.WithTooltipOpts(opts.Tooltip{
			Show: true,
			AxisPointer: &opts.AxisPointer{
				Type: "shadow",
			},
		}),
	)

	// Put data into instance
	bar.SetXAxis([]string{"1969", "2019"}).
		AddSeries("África", addBarItems([]float64{7.9, 17.4})).
		AddSeries("América", addBarItems([]float64{8.7, 11.6})).
		AddSeries("Europa", addBarItems([]float64{21.7, 16.2})).
		AddSeries("Asia", addBarItems([]float64{14.7, 43.6})).
		SetSeriesOptions(
			charts.WithLabelOpts(opts.Label{
				Show:     true,
				Position: "top",
				Color:    "black",
			}),
		)

	return bar
}

// var (
// 	itemCntPie = 4
// 	seasons    = []string{"Spring", "Summer", "Autumn ", "Winter"}
// )

// func generatePieItems() []opts.PieData {
// 	items := make([]opts.PieData, 0)
// 	items = append(items, opts.PieData{Value: 75,
// 		Label: &opts.Label{Show: false}})
// 	items = append(items, opts.PieData{Value: 100 - 75, Label: &opts.Label{Show: false}})
// 	return items
// }

// func pieBase() *charts.Pie {
// 	pie := charts.NewPie()
// 	pie.SetGlobalOptions(
// 		charts.WithTitleOpts(opts.Title{Title: "basic pie example"}),
// 	)

// 	pie.AddSeries("pie", generatePieItems()).SetSeriesOptions()

// 	return pie
// }
var (
	baseMapData = []opts.MapData{
		{Name: "Canada", Value: 12},
		{Name: "United States", Value: 12},

		{Name: "Mexico", Value: 18},
		{Name: "Venezuela", Value: 45},
		{Name: "Colombia", Value: 45},
		{Name: "Ecuador", Value: 13},
		{Name: "Chile", Value: 53},
		{Name: "Brazil", Value: 37},
		{Name: "Peru", Value: 13},
		{Name: "Bolivia", Value: 30},
		{Name: "Paraguay", Value: 30},
		{Name: "Uruguay", Value: 27},
		{Name: "Argentina", Value: 34},

		{Name: "Russia", Value: 12},
		{Name: "China", Value: 36},
		{Name: "India", Value: 7},
		{Name: "Mongolia", Value: 6},
		{Name: "Saudi Arabia", Value: 25},
		{Name: "Yemen", Value: 27},
		{Name: "Oman", Value: 14},
		{Name: "Kazakhstan", Value: 25},
		{Name: "Turkmenistan", Value: 47},
		{Name: "Uzbekistan", Value: 52},
		{Name: "Afghanistan", Value: 22},
		{Name: "Pakistan", Value: 15},
		{Name: "Turkey", Value: 21},
		{Name: "Iran", Value: 21},
		{Name: "Syria", Value: 25},

		{Name: "Libya", Value: 60},
		{Name: "Algeria", Value: 36},
		{Name: "Egypt", Value: 12},
		{Name: "Madagascar", Value: 12},
		{Name: "Angola", Value: 10},
		{Name: "Sudan", Value: 23},
		{Name: "S. Sudan", Value: 23},
		{Name: "Kenya", Value: 22},
		{Name: "Mali", Value: 16},
		{Name: "Dem. Rep. Congo", Value: 3},

		{Name: "Italy", Value: 24},
		{Name: "Spain", Value: 26},
		{Name: "France", Value: 11},
		{Name: "Poland", Value: 11},
		{Name: "Germany", Value: 2},
		{Name: "United Kingdom", Value: 1},

		{Name: "Australia", Value: 6},
		{Name: "Thailand", Value: 2},
		{Name: "Indonesia", Value: 4},
		{Name: "Papua New Guinea", Value: 21},
	}
)

func mapVisualMap() *charts.Map {
	mc := charts.NewMap()
	mc.RegisterMapType("world")
	mc.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: "Perdidas de colmenas por país",
		}),
		charts.WithColorsOpts(opts.Colors{"#ffffff"}),
		charts.WithVisualMapOpts(opts.VisualMap{
			Calculable: true,
			Max:        60,
		}),
		charts.WithTooltipOpts(opts.Tooltip{
			Show: true,
			AxisPointer: &opts.AxisPointer{
				Type: "shadow",
			},
		}),
	)

	mc.AddSeries("map", baseMapData)
	return mc
}

func main() {
	page := components.NewPage()
	//page.SetLayout(components.PageFlexLayout)
	page.AddCharts(
		bar(),
		mapVisualMap(),
	)

	f, err := os.Create("data.html")
	if err != nil {
		panic(err)
	}
	page.Render(io.MultiWriter(f))
}
