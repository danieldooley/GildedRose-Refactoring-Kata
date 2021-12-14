package main

import "strings"

const (
	AgedBrie        = "Aged Brie"
	BackstagePasses = "Backstage passes to a TAFKAL80ETC concert"
	Sulfuras        = "Sulfuras, Hand of Ragnaros"
	ConjuredPrefix  = "Conjured"
	MaxQuality      = 50
	MinQuality      = 0
)

type Item struct {
	name            string
	sellIn, quality int
}

func clamp(i int) int {
	if i < MinQuality {
		return MinQuality
	}
	if i > MaxQuality {
		return MaxQuality
	}
	return i
}

func UpdateQuality(items []*Item) {
	for _, item := range items {

		if item.name == Sulfuras {
			continue
		}

		switch {
		case item.name == AgedBrie:
			item.quality = clamp(item.quality + 1)
		case item.name == BackstagePasses:
			item.quality = clamp(item.quality + 1)

			if item.sellIn < 11 {
				item.quality = clamp(item.quality + 1)
			}
			if item.sellIn < 6 {
				item.quality = clamp(item.quality + 1)
			}
		case strings.HasPrefix(item.name, ConjuredPrefix):
			item.quality = clamp(item.quality - 2)
		default:
			item.quality = clamp(item.quality - 1)
		}

		item.sellIn = item.sellIn - 1

		if item.sellIn < 0 {
			switch {
			case item.name == AgedBrie:
				item.quality = clamp(item.quality + 1)
			case item.name == BackstagePasses:
				item.quality = clamp(item.quality - item.quality)
			case strings.HasPrefix(item.name, ConjuredPrefix):
				item.quality = clamp(item.quality - 2)
			default:
				item.quality = clamp(item.quality - 1)
			}
		}
	}
}
