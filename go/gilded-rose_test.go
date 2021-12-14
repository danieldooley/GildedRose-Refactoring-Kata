package main

import "testing"

var UpdateQualityTestData = map[string]struct {
	name                                                       string
	inputQuality, inputSellIn, expectedQuality, expectedSellIn int
}{
	"sellIn and quality should be decremented for a normal item": {
		name:            "Shield of Sadness",
		inputQuality:    10,
		inputSellIn:     10,
		expectedQuality: 9,
		expectedSellIn:  9,
	},
	"sellIn and quality should be decremented twice for a normal item when sellIn is negative": {
		name:            "Shield of Sadness",
		inputQuality:    10,
		inputSellIn:     0,
		expectedQuality: 8,
		expectedSellIn:  -1,
	},
	"quality should never be less than 0": {
		name:            "Shield of Sadness",
		inputQuality:    0,
		inputSellIn:     10,
		expectedQuality: 0,
		expectedSellIn:  9,
	},
	"quality should never be less than 0 even when reducing by 2": {
		name:            "Shield of Sadness",
		inputQuality:    1,
		inputSellIn:     0,
		expectedQuality: 0,
		expectedSellIn:  -1,
	},
	"Aged Brie should increase in quality as it ages": {
		name:            "Aged Brie",
		inputQuality:    10,
		inputSellIn:     10,
		expectedQuality: 11,
		expectedSellIn:  9,
	},
	"Aged Brie should increase in quality twice as fast once sellIn becomes negative": {
		name:            "Aged Brie",
		inputQuality:    10,
		inputSellIn:     0,
		expectedQuality: 12,
		expectedSellIn:  -1,
	},
	"Sulfuras should not have its sellIn or quality changed": {
		name:            "Sulfuras, Hand of Ragnaros",
		inputQuality:    20,
		inputSellIn:     30,
		expectedQuality: 20,
		expectedSellIn:  30,
	},
	"Sulfuras should not have its sellIn or quality changed even when values are negative": {
		name:            "Sulfuras, Hand of Ragnaros",
		inputQuality:    -10,
		inputSellIn:     -30,
		expectedQuality: -10,
		expectedSellIn:  -30,
	},
	"Sulfuras should not have its quality changed even when it exceeds 50": {
		name:            "Sulfuras, Hand of Ragnaros",
		inputQuality:    80,
		inputSellIn:     50,
		expectedQuality: 80,
		expectedSellIn:  50,
	},
	"Backstage passes should have their quality increased by one when sellIn > 10 (15)": {
		name:            "Backstage passes to a TAFKAL80ETC concert",
		inputQuality:    10,
		inputSellIn:     15,
		expectedQuality: 11,
		expectedSellIn:  14,
	},
	"Backstage passes should have their quality increased by one when sellIn > 10 (11)": {
		name:            "Backstage passes to a TAFKAL80ETC concert",
		inputQuality:    10,
		inputSellIn:     11,
		expectedQuality: 11,
		expectedSellIn:  10,
	},
	"Backstage passes should have their quality increased by 2 when sellIn < 10 (10)": {
		name:            "Backstage passes to a TAFKAL80ETC concert",
		inputQuality:    10,
		inputSellIn:     10,
		expectedQuality: 12,
		expectedSellIn:  9,
	},
	"Backstage passes should have their quality increased by 2 when sellIn < 10 (6)": {
		name:            "Backstage passes to a TAFKAL80ETC concert",
		inputQuality:    10,
		inputSellIn:     6,
		expectedQuality: 12,
		expectedSellIn:  5,
	},
	"Backstage passes should have their quality increased by 3 when sellIn < 5 (5)": {
		name:            "Backstage passes to a TAFKAL80ETC concert",
		inputQuality:    10,
		inputSellIn:     5,
		expectedQuality: 13,
		expectedSellIn:  4,
	},
	"Backstage passes should have their quality increased by 3 when sellIn < 5 (1)": {
		name:            "Backstage passes to a TAFKAL80ETC concert",
		inputQuality:    10,
		inputSellIn:     1,
		expectedQuality: 13,
		expectedSellIn:  0,
	},
	"Backstage passes should have their quality set to 0 when sellIn is negative (0)": {
		name:            "Backstage passes to a TAFKAL80ETC concert",
		inputQuality:    10,
		inputSellIn:     0,
		expectedQuality: 0,
		expectedSellIn:  -1,
	},
	"Backstage passes should have their quality set to 0 when sellIn is negative (-5)": {
		name:            "Backstage passes to a TAFKAL80ETC concert",
		inputQuality:    10,
		inputSellIn:     -5,
		expectedQuality: 0,
		expectedSellIn:  -6,
	},
	"Backstage passes quality should not exceed 50": {
		name:            "Backstage passes to a TAFKAL80ETC concert",
		inputQuality:    48,
		inputSellIn:     2,
		expectedQuality: 50,
		expectedSellIn:  1,
	},
	"Conjured items should degrade twice": {
		name:            "Conjured Manage Cake",
		inputQuality:    10,
		inputSellIn:     10,
		expectedQuality: 8,
		expectedSellIn:  9,
	},
	"Conjured items should degrade 4x when sellIn has passed": {
		name:            "Conjured Manage Cake",
		inputQuality:    10,
		inputSellIn:     0,
		expectedQuality: 6,
		expectedSellIn:  -1,
	},
	"Conjured items should not be less than 0 quality": {
		name:            "Conjured Dressing Gown",
		inputQuality:    3,
		inputSellIn:     0,
		expectedQuality: 0,
		expectedSellIn:  -1,
	},
}

func TestUpdateQuality(t *testing.T) {
	for name, data := range UpdateQualityTestData {
		t.Run(name, func(t *testing.T) {
			item := &Item{
				name:    data.name,
				sellIn:  data.inputSellIn,
				quality: data.inputQuality,
			}
			UpdateQuality([]*Item{item})

			if item.quality != data.expectedQuality {
				t.Errorf("Expected Quality of %d got %d", data.expectedQuality, item.quality)
			}

			if item.sellIn != data.expectedSellIn {
				t.Errorf("Expected SellIn of %d got %d", data.expectedSellIn, item.sellIn)
			}
		})
	}
}
