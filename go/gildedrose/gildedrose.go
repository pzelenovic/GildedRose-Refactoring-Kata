package gildedrose

type Item struct {
	Name       string
	SellIn     int
	Quality    int
}


func UpdateQuality(items []*Item) {
	for i := 0; i < len(items); i++ {

		if notAgedBrie(items[i]) && notBackPasses(items[i]) {
			if items[i].Quality > 0 {
				if items[i].Name != "Sulfuras, Hand of Ragnaros" {
					decrementQuality(items[i])
				}
			}
		} else {
			if belowMaxQuality(items[i]) {
				items[i].Quality = items[i].Quality + 1
				if items[i].Name == "Backstage passes to a TAFKAL80ETC concert" {
					if items[i].SellIn < 11 {
						if belowMaxQuality(items[i]) {
							incrementQuality(items[i])
						}
					}
					if items[i].SellIn < 6 {
						if belowMaxQuality(items[i]) {
							incrementQuality(items[i])
						}
					}
				}
			}
		}

		if items[i].Name != "Sulfuras, Hand of Ragnaros" {
			items[i].SellIn = items[i].SellIn - 1
		}

		if items[i].SellIn < 0 {
			if items[i].Name != "Aged Brie" {
				if items[i].Name != "Backstage passes to a TAFKAL80ETC concert" {
					if items[i].Quality > 0 {
						if items[i].Name != "Sulfuras, Hand of Ragnaros" {
							items[i].Quality = items[i].Quality - 1
						}
					}
				} else {
					items[i].Quality = items[i].Quality - items[i].Quality
				}
			} else {
				if belowMaxQuality(items[i]) {
					incrementQuality(items[i])
				}
			}
		}
	}

}

func belowMaxQuality(item *Item) bool {
	return item.Quality < 50
}

func notBackPasses(item *Item) bool {
	return item.Name != "Backstage passes to a TAFKAL80ETC concert"
}

func notAgedBrie(items *Item) bool {
	return items.Name != "Aged Brie"
}

func decrementQuality(item *Item) {
	item.Quality -= 1
}

func incrementQuality(item *Item) {
	item.Quality += 1
}
