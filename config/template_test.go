package config

import (
	"image"
	"reflect"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	var sampleFile = "../_template/furnace.toml"

	var answer = Config{
		Type:  "furnace",
		Image: "furnace.png",
		Crop:  image.Rect(7, 6, 168, 79),
		Slot: []image.Point{
			{56, 53},
			{56, 17},
			{116, 35},
		},
		Progress: []Progress{
			{
				Crop:  image.Rect(176, 0, 189, 13),
				Paste: image.Pt(57, 36),
			},
			{
				Crop:  image.Rect(176, 14, 24, 17),
				Paste: image.Pt(199, 30),
			},
		},
	}

	sample := sampleFile
	ans := answer
	if !reflect.DeepEqual(LoadConfig(sample), ans) {
		t.Error("toml convert error: not equal input and output \n", LoadConfig(sample), ans)
	}
}

func TestItop(t *testing.T) {
	var input = [][]int{
		{0, 0},
		{-1, -1},
		{1000, 1000000},
	}
	var output = []image.Point{
		{0, 0},
		{-1, -1},
		{1000, 1000000},
	}

	for i := range input {
		if !reflect.DeepEqual(itop(input[i]), output[i]) {
			t.Error("int to Point convert error")
		}
	}
}

func TestItor(t *testing.T) {
	var input = [][][]int{
		{{0, 0}, {0, 0}},
		{{-10, -10}, {-10, -10}},
		{{10000000000, 10000000000}, {10000000000, 10000000000}},
		{{-10000000000, -10000000000}, {-10000000000, -10000000000}},
	}
	var output = []image.Rectangle{
		{image.Pt(0, 0), image.Pt(0, 0)},
		{image.Pt(-10, -10), image.Pt(-10, -10)},
		{image.Pt(10000000000, 10000000000), image.Pt(10000000000, 10000000000)},
		{image.Pt(-10000000000, -10000000000), image.Pt(-10000000000, -10000000000)},
	}

	for i := range input {
		if !reflect.DeepEqual(itor(input[i]), output[i]) {
			t.Error("int to image.Rectangle convert error")
		}
	}
}
