// MIT License

// Copyright (c) 2022 Tree Xie

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package charts

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wcharczuk/go-chart/v2"
	"github.com/wcharczuk/go-chart/v2/drawing"
)

func TestEChartsPosition(t *testing.T) {
	assert := assert.New(t)

	var p EChartsPosition
	err := p.UnmarshalJSON([]byte("12"))
	assert.Nil(err)
	assert.Equal("12", string(p))

	err = p.UnmarshalJSON([]byte(`"12%"`))
	assert.Nil(err)
	assert.Equal("12%", string(p))
}
func TestEChartStyle(t *testing.T) {
	assert := assert.New(t)

	s := EChartStyle{
		Color: "#aaa",
	}
	r := drawing.Color{
		R: 170,
		G: 170,
		B: 170,
		A: 255,
	}
	assert.Equal(chart.Style{
		FillColor:   r,
		FontColor:   r,
		StrokeColor: r,
	}, s.ToStyle())
}

func TestEChartsXAxis(t *testing.T) {
	assert := assert.New(t)
	ex := EChartsXAxis{}
	err := ex.UnmarshalJSON([]byte(`{
		"boundaryGap": false,
		"splitNumber": 5,
		"data": [
			"Mon",
			"Tue",
			"Wed",
			"Thu",
			"Fri",
			"Sat",
			"Sun"
		]
	}`))
	assert.Nil(err)
	assert.Equal(EChartsXAxis{
		Data: []EChartsXAxisData{
			{
				BoundaryGap: FalseFlag(),
				SplitNumber: 5,
				Data: []string{
					"Mon",
					"Tue",
					"Wed",
					"Thu",
					"Fri",
					"Sat",
					"Sun",
				},
			},
		},
	}, ex)
}

func TestEChartsYAxis(t *testing.T) {
	assert := assert.New(t)
	ey := EChartsYAxis{}

	err := ey.UnmarshalJSON([]byte(`{
		"min": 1,
		"max": 10,
		"axisLabel": {
			"formatter": "ab"
		}
	}`))
	assert.Nil(err)
	assert.Equal(EChartsYAxis{
		Data: []EChartsYAxisData{
			{
				Min: NewFloatPoint(1),
				Max: NewFloatPoint(10),
				AxisLabel: EChartsAxisLabel{
					Formatter: "ab",
				},
			},
		},
	}, ey)

	ey = EChartsYAxis{}
	err = ey.UnmarshalJSON([]byte(`[
		{
			"min": 1,
			"max": 10,
			"axisLabel": {
				"formatter": "ab"
			}
		},
		{
			"min": 2,
			"max": 20,
			"axisLabel": {
				"formatter": "cd"
			}
		}
	]`))
	assert.Nil(err)
	assert.Equal(EChartsYAxis{
		Data: []EChartsYAxisData{
			{
				Min: NewFloatPoint(1),
				Max: NewFloatPoint(10),
				AxisLabel: EChartsAxisLabel{
					Formatter: "ab",
				},
			},
			{
				Min: NewFloatPoint(2),
				Max: NewFloatPoint(20),
				AxisLabel: EChartsAxisLabel{
					Formatter: "cd",
				},
			},
		},
	}, ey)
}

func TestEChartsPadding(t *testing.T) {
	assert := assert.New(t)

	ep := EChartsPadding{}

	ep.UnmarshalJSON([]byte(`10`))
	assert.Equal(EChartsPadding{
		Box: chart.Box{
			Top:    10,
			Right:  10,
			Bottom: 10,
			Left:   10,
		},
	}, ep)

	ep = EChartsPadding{}
	ep.UnmarshalJSON([]byte(`[10, 20]`))
	assert.Equal(EChartsPadding{
		Box: chart.Box{
			Top:    10,
			Right:  20,
			Bottom: 10,
			Left:   20,
		},
	}, ep)

	ep = EChartsPadding{}
	ep.UnmarshalJSON([]byte(`[10, 20, 30]`))
	assert.Equal(EChartsPadding{
		Box: chart.Box{
			Top:    10,
			Right:  20,
			Bottom: 30,
			Left:   20,
		},
	}, ep)

	ep = EChartsPadding{}
	ep.UnmarshalJSON([]byte(`[10, 20, 30, 40]`))
	assert.Equal(EChartsPadding{
		Box: chart.Box{
			Top:    10,
			Right:  20,
			Bottom: 30,
			Left:   40,
		},
	}, ep)

}
func TestEChartsLegend(t *testing.T) {
	assert := assert.New(t)

	el := EChartsLegend{}

	err := json.Unmarshal([]byte(`{
		"data": ["a", "b", "c"],
		"align": "right",
		"padding": [10],
		"left": "20%",
		"top": 10
	}`), &el)
	assert.Nil(err)
	assert.Equal(EChartsLegend{
		Data: []string{
			"a",
			"b",
			"c",
		},
		Align: "right",
		Padding: EChartsPadding{
			Box: chart.Box{
				Left:   10,
				Top:    10,
				Right:  10,
				Bottom: 10,
			},
		},
		Left: EChartsPosition("20%"),
		Top:  EChartsPosition("10"),
	}, el)
}

func TestEChartsSeriesData(t *testing.T) {
	assert := assert.New(t)

	esd := EChartsSeriesData{}
	err := esd.UnmarshalJSON([]byte(`123`))
	assert.Nil(err)
	assert.Equal(EChartsSeriesData{
		Value: 123,
	}, esd)

	esd = EChartsSeriesData{}
	err = esd.UnmarshalJSON([]byte(`{
		"value": 123.12,
		"name": "test",
		"itemStyle": {
			"color": "#aaa"
		}
	}`))
	assert.Nil(err)
	assert.Equal(EChartsSeriesData{
		Value: 123.12,
		Name:  "test",
		ItemStyle: EChartStyle{
			Color: "#aaa",
		},
	}, esd)
}

func TestEChartsSeries(t *testing.T) {
	assert := assert.New(t)

	esList := make([]EChartsSeries, 0)
	err := json.Unmarshal([]byte(`[
		{
			"name": "Email",
			"data": [
				120,
				132
			]
		},
		{
			"name": "Union Ads",
			"type": "bar",
			"data": [
				220,
				182
			]
		}
	]`), &esList)
	assert.Nil(err)
	assert.Equal([]EChartsSeries{
		{
			Name: "Email",
			Data: []EChartsSeriesData{
				{
					Value: 120,
				},
				{
					Value: 132,
				},
			},
		},
		{
			Name: "Union Ads",
			Type: "bar",
			Data: []EChartsSeriesData{
				{
					Value: 220,
				},
				{
					Value: 182,
				},
			},
		},
	}, esList)
}

func TestEChartsMarkPoint(t *testing.T) {
	assert := assert.New(t)

	p := EChartsMarkPoint{}

	err := json.Unmarshal([]byte(`{
		"symbolSize": 30,
		"data": [
			{
				"type": "max"
			},
			{
				"type": "min"
			}
		]
	}`), &p)
	assert.Nil(err)
	assert.Equal(SeriesMarkPoint{
		SymbolSize: 30,
		Data: []SeriesMarkData{
			{
				Type: "max",
			},
			{
				Type: "min",
			},
		},
	}, p.ToSeriesMarkPoint())
}

func TestEChartsMarkLine(t *testing.T) {
	assert := assert.New(t)
	l := EChartsMarkLine{}

	err := json.Unmarshal([]byte(`{
		"data": [
			{
				"type": "max"
			},
			{
				"type": "min"
			}
		]
	}`), &l)
	assert.Nil(err)
	assert.Equal(SeriesMarkLine{
		Data: []SeriesMarkData{
			{
				Type: "max",
			},
			{
				Type: "min",
			},
		},
	}, l.ToSeriesMarkLine())
}

func TestEChartsTextStyle(t *testing.T) {
	assert := assert.New(t)

	s := EChartsTextStyle{
		Color:      "#aaa",
		FontFamily: "test",
		FontSize:   14,
	}
	assert.Equal(chart.Style{
		FontColor: drawing.Color{
			R: 170,
			G: 170,
			B: 170,
			A: 255,
		},
		FontSize: 14,
	}, s.ToStyle())
}

func TestEChartsSeriesList(t *testing.T) {
	assert := assert.New(t)

	// pie
	es := EChartsSeriesList{
		{
			Type:   ChartTypePie,
			Radius: "30%",
			Data: []EChartsSeriesData{
				{
					Name:  "1",
					Value: 1,
				},
				{
					Name:  "2",
					Value: 2,
				},
			},
		},
	}
	assert.Equal(SeriesList{
		{
			Type: ChartTypePie,
			Name: "1",
			Label: SeriesLabel{
				Show: true,
			},
			Radius: "30%",
			Data: []SeriesData{
				{
					Value: 1,
				},
			},
		},
		{
			Type: ChartTypePie,
			Name: "2",
			Label: SeriesLabel{
				Show: true,
			},
			Radius: "30%",
			Data: []SeriesData{
				{
					Value: 2,
				},
			},
		},
	}, es.ToSeriesList())

	es = EChartsSeriesList{
		{
			Type: ChartTypeBar,
			Data: []EChartsSeriesData{
				{
					Value: 1,
					ItemStyle: EChartStyle{
						Color: "#aaa",
					},
				},
				{
					Value: 2,
				},
			},
			YAxisIndex: 1,
		},
		{
			Data: []EChartsSeriesData{
				{
					Value: 3,
				},
				{
					Value: 4,
				},
			},
			ItemStyle: EChartStyle{
				Color: "#ccc",
			},
			Label: EChartsLabelOption{
				Color:    "#ddd",
				Show:     true,
				Distance: 5,
			},
		},
	}
	assert.Equal(SeriesList{
		{
			Type: ChartTypeBar,
			Data: []SeriesData{
				{
					Value: 1,
					Style: chart.Style{
						FontColor: drawing.Color{
							R: 170,
							G: 170,
							B: 170,
							A: 255,
						},
						StrokeColor: drawing.Color{
							R: 170,
							G: 170,
							B: 170,
							A: 255,
						},
						FillColor: drawing.Color{
							R: 170,
							G: 170,
							B: 170,
							A: 255,
						},
					},
				},
				{
					Value: 2,
				},
			},
			YAxisIndex: 1,
		},
		{
			Data: []SeriesData{
				{
					Value: 3,
				},
				{
					Value: 4,
				},
			},
			Style: chart.Style{
				FontColor: drawing.Color{
					R: 204,
					G: 204,
					B: 204,
					A: 255,
				},
				StrokeColor: drawing.Color{
					R: 204,
					G: 204,
					B: 204,
					A: 255,
				},
				FillColor: drawing.Color{
					R: 204,
					G: 204,
					B: 204,
					A: 255,
				},
			},
			Label: SeriesLabel{
				Color: drawing.Color{
					R: 221,
					G: 221,
					B: 221,
					A: 255,
				},
				Show:     true,
				Distance: 5,
			},
			MarkPoint: SeriesMarkPoint{},
			MarkLine:  SeriesMarkLine{},
		},
	}, es.ToSeriesList())

}
