package relevance

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io"
	"strconv"

	"github.com/nhood-org/nhood-engine-utils/pkg/model"
)

type Matrix struct {
	tags         []string
	tagsIndices  map[string]int
	maxTagsCount int
	vectors      []Vector
}

func NewMatrix(tags *model.Tags) *Matrix {
	mSize := len(*tags)

	m := Matrix{
		tags:        make([]string, mSize),
		tagsIndices: make(map[string]int),
		vectors:     make([]Vector, mSize),
	}

	i := 0
	for t := range *tags {
		m.tags[i] = t
		m.tagsIndices[t] = i
		m.vectors[i] = *NewVector(mSize)
		i++
	}

	return &m
}

func (t *Matrix) Tags() []string {
	return t.tags
}

func (t *Matrix) SetTags(tags []string) {
	t.tags = tags
}

func (t *Matrix) Increment(t1 string, t2 string) {
	if t1 == t2 {
		return
	}

	t1Idx := t.tagsIndices[t1]
	t2Idx := t.tagsIndices[t2]
	t.vectors[t1Idx].increment(t2Idx)

	relevance := t.vectors[t1Idx].vector[t2Idx]
	if t.maxTagsCount < relevance {
		t.maxTagsCount = relevance
	}
}

func (t *Matrix) Relevance(t1 string, t2 string) int {
	t1Idx := t.tagsIndices[t1]
	t2Idx := t.tagsIndices[t2]
	return t.vectors[t1Idx].vector[t2Idx]
}

func (t *Matrix) Weight(t1 string) int {
	t1Idx := t.tagsIndices[t1]
	return t.vectors[t1Idx].weight
}

func (t *Matrix) MaxTagsCount() int {
	return t.maxTagsCount
}

func (t *Matrix) Len() int {
	return len(t.tags)
}

func (t *Matrix) PrintAsText(out io.Writer) {
	t.printHeaderAsText(out)
	for _, r := range t.tags {
		t.printAsText(out, r)
	}
}

func (t *Matrix) PrintAsPNG(out io.Writer) {
	width := len(t.tags)
	height := width

	canvasUpLeft := image.Point{0, 0}
	canvasBottomRight := image.Point{width, height}
	canvas := image.Rectangle{canvasUpLeft, canvasBottomRight}

	img := image.NewRGBA(canvas)

	y := 0
	for _, t1 := range t.tags {
		t1Idx := t.tagsIndices[t1]
		weights := t.vectors[t1Idx]
		x := 0
		for _, t2 := range t.tags {
			t2Idx := t.tagsIndices[t2]
			weight := weights.vector[t2Idx]
			v := (weight * 255 * 1000) / t.maxTagsCount
			c := color.RGBA{uint8(v), 0, 0, 0xff}
			img.Set(x, y, c)
			x++
		}
		y++
	}

	png.Encode(out, img)
}

func (t *Matrix) printHeaderAsText(out io.Writer) {
	s := ""
	for _, r := range t.tags {
		s = s + "\t" + shortenTagName(r)
	}
	fmt.Fprintln(out, s)
}

func (t *Matrix) printAsText(out io.Writer, tag string) {
	s := shortenTagName(tag) + ":"

	idx := t.tagsIndices[tag]
	weights := t.vectors[idx]
	for _, r := range t.tags {

		rIdx := t.tagsIndices[r]
		w := strconv.Itoa(weights.vector[rIdx])
		s = s + "\t" + w
	}

	w := strconv.Itoa(weights.weight)
	s = s + "\t[" + w + "]"

	fmt.Fprintln(out, s)
}

const tagNameLength = 6

func shortenTagName(tag string) string {
	tagName := tag
	if len(tagName) > tagNameLength {
		tagName = tagName[:6]
		tagName = tagName + "*"
	}
	return tagName
}
