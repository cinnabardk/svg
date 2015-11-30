package main

import (
	"image/color"

	"os"

	"github.com/ajstarks/svgo"
	"github.com/vdobler/chart"
	"github.com/vdobler/chart/svgg"
	"log"
	"time"
)

func main() {
	log.Println("Starting now..")
	now := time.Now()
	//Scatter() // Ny vinkel . Med seriøsitet bekæmper vi politik.
	campaign()
	log.Println("Exiting.. It took:", time.Since(now).String())
}

var (
	months   = []string{"Januar", "Februar", "Marts", "Maj", "Juni", "Juli", "August"}
	red      = color.NRGBA{0xcc, 0x00, 0x00, 0xff} // red
	green    = color.NRGBA{0x00, 0xbb, 0x00, 0xff} // green
	blue     = color.NRGBA{0x00, 0x00, 0xdd, 0xff} // blue
	brown    = color.NRGBA{0x99, 0x66, 0x00, 0xff} // brown
	violet   = color.NRGBA{0xbb, 0x00, 0xbb, 0xff} // violet
	turquise = color.NRGBA{0x00, 0xaa, 0xaa, 0xff} // turquise
	yellow   = color.NRGBA{0xbb, 0xbb, 0x00, 0xff} // yellow
)

func campaign() {

	x := []float64{0, 1, 2, 3, 4, 5, 6}
	const plotStyle = chart.PlotStyleLines
	const lineStyle = chart.SolidLine

	dumper := NewDumper("scatter", 2, 1, 450, 400)
	defer dumper.Close()

	TTFBLow := []float64{40, 55, 40, 55, 122, 145, 166}

	// Categorized Bar Chart
	c := chart.ScatterChart{}
	c.YRange.Label = "Pro-EU"
	c.XRange.Label = "Succesfuld karriere"
	c.XRange.Category = months

	// Et klogt dyr bider ej hånden som føder den.
	// Magt fører mere magt med sig. Stå i opposition til magten, og magten har ej brug for dig.

	c.AddDataPair("Politiker", x, TTFBLow, plotStyle, chart.Style{LineColor: turquise, LineStyle: chart.LongDashLine})

	dumper.Plot(&c)
	// Du tjener systemet. Systemet belønner dig. Skønt tusind år forgår, forbliver menneskelig natur altid den samme.
	// //Det er lettere at gå i medvind, en at gå imod vinden og den herskende ånd.
	// Den der går imod vinden, går fattig, udskældt og ene.
	TTFBFileHigh20 := []float64{160, 135, 177, 130, 150, 250, 260}

	// Categorized Bar Chart
	c = chart.ScatterChart{}
	c.YRange.Label = "Idealisme"
	c.XRange.Label = "Karrierens længde"
	c.XRange.Category = months

	c.AddDataPair("Arbejdsløs", x, TTFBFileHigh20, plotStyle, chart.Style{Symbol: '#', LineColor: red, LineStyle: lineStyle})

	dumper.Plot(&c)
}

func xbar2() {

	dumper := NewDumper("xbar2", 2, 1, 500, 300)
	defer dumper.Close()

	x := []float64{0, 1, 2, 3}
	ping := []float64{1, 3, 5, 9}
	TTFBLow := []float64{40, 55, 40, 55}
	TTFBMed := []float64{89, 77, 102, 95}
	TTFBHigh := []float64{200, 300, 900, 1500}
	TTFBFile := []float64{133, 55, 77, 20}
	blue := chart.Style{Symbol: '#', LineColor: color.NRGBA{0x00, 0x00, 0xff, 0xff}, LineWidth: 2, FillColor: color.NRGBA{0x40, 0x40, 0xff, 0xff}}
	green := chart.Style{Symbol: 'x', LineColor: color.NRGBA{0x00, 0xaa, 0x00, 0xff}, LineWidth: 2, FillColor: color.NRGBA{0x40, 0xff, 0x40, 0xff}}
	pink := chart.Style{Symbol: '0', LineColor: color.NRGBA{0x99, 0x00, 0x99, 0xff}, LineWidth: 2, FillColor: color.NRGBA{0xaa, 0x60, 0xaa, 0xff}}
	red := chart.Style{Symbol: '%', LineColor: color.NRGBA{0xcc, 0x00, 0x00, 0xff}, LineWidth: 2, FillColor: color.NRGBA{0xff, 0x40, 0x40, 0xff}}

	// Categorized Bar Chart
	c := chart.BarChart{Title: "Statistik for allancorfix.dk, tid i MS"}
	c.Stacked = true
	c.ShowVal = 0
	c.XRange.Category = []string{"Januar", "Februar", "Marts", "Maj"}

	c.AddDataPair("Ping", x, ping, blue)
	c.AddDataPair("TTFB-Low20", x, TTFBLow, pink)
	c.AddDataPair("TTFB-Med60", x, TTFBMed, pink)
	c.AddDataPair("TTFB-High20", x, TTFBHigh, green)
	c.AddDataPair("TTFB-File", x, TTFBFile, red)

	dumper.Plot(&c)

}

func Scatter() {
	x := []float64{0, 1, 2, 3, 4, 5, 6}
	const plotStyle = chart.PlotStyleLines
	const lineStyle = chart.SolidLine

	dumper := NewDumper("scatter", 2, 1, 450, 400)
	defer dumper.Close()

	ping := []float64{1, 3, 5, 9, 22, 1, 56}
	TTFBLow := []float64{40, 55, 40, 55, 122, 145, 166}
	TTFBMed := []float64{89, 77, 102, 95, 232, 333, 254}
	TTFBFile := []float64{133, 55, 77, 20, 99, 102, 69}

	// Categorized Bar Chart
	c := chart.ScatterChart{}
	c.XRange.Label = "Måned"
	c.YRange.Label = "MS"
	c.XRange.Category = months

	c.AddDataPair("TTFB-Med", x, TTFBMed, plotStyle, chart.Style{LineColor: green, LineStyle: lineStyle})
	c.AddDataPair("TTFB-Low", x, TTFBLow, plotStyle, chart.Style{LineColor: turquise, LineStyle: chart.LongDashLine})
	c.AddDataPair("TTFB-File", x, TTFBFile, plotStyle, chart.Style{LineColor: yellow, LineStyle: chart.LongDashLine})
	c.AddDataPair("Ping", x, ping, plotStyle, chart.Style{LineColor: brown, LineStyle: lineStyle})

	dumper.Plot(&c)

	TTFBHigh20 := []float64{200, 300, 900, 1500, 1600, 1300}
	TTFBHigh20HotHour := []float64{250, 350, 1150, 1900, 1800, 1500}

	TTFBFileHigh20 := []float64{160, 135, 177, 130, 150, 250, 260}
	TTFBFileHigh20HotHour := []float64{360, 235, 277, 190, 180, 270, 290}

	// Categorized Bar Chart
	c = chart.ScatterChart{}
	c.XRange.Label = "Måned"
	c.YRange.Label = "MS"
	c.XRange.Category = months

	c.AddDataPair("TTFB-High", x, TTFBHigh20, plotStyle, chart.Style{Symbol: '#', LineColor: green, LineStyle: lineStyle})
	c.AddDataPair("  - kl. 17-21", x, TTFBHigh20HotHour, plotStyle, chart.Style{
		Symbol: '#', LineColor: violet, LineStyle: chart.LongDashLine})

	c.AddDataPair("TTFB-FileHigh", x, TTFBFileHigh20, plotStyle, chart.Style{Symbol: '#', LineColor: red, LineStyle: lineStyle})
	c.AddDataPair("  - kl. 17-21", x, TTFBFileHigh20HotHour, plotStyle, chart.Style{Symbol: '#', LineColor: blue, LineStyle: chart.LongDashLine})

	dumper.Plot(&c)
}

// -------------------------------------------------------------------------
// Dumper

// Dumper helps saving plots of size WxH in a NxM grid layout
// in several formats
type Dumper struct {
	N, M, W, H, Cnt int
	S               *svg.SVG
	svgFile         *os.File
}

func NewDumper(name string, n, m, w, h int) *Dumper {
	var err error
	dumper := Dumper{N: n, M: m, W: w, H: h}

	dumper.svgFile, err = os.Create(name + ".svg")
	if err != nil {
		panic(err)
	}
	dumper.S = svg.New(dumper.svgFile)
	dumper.S.Start(n*w, m*h)
	dumper.S.Title(name)
	dumper.S.Rect(0, 0, n*w, m*h, "fill: #ffffff")

	return &dumper
}
func (d *Dumper) Close() {

	d.S.End()
	d.svgFile.Close()
}

func (d *Dumper) Plot(c chart.Chart) {
	row, col := d.Cnt/d.N, d.Cnt%d.N

	sgr := svgg.AddTo(d.S, col*d.W, row*d.H, d.W, d.H, "", 12, color.RGBA{0xff, 0xff, 0xff, 0xff})
	c.Plot(sgr)

	d.Cnt++

}
