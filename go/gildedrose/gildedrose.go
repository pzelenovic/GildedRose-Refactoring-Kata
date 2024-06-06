package gildedrose

type Item struct {
	Name    string
	SellIn  int
	Quality int
}

type SulfurasItem struct {
	Item
}

func (item *Item) Process() bool {
	item.Quality += 1
	return true
}

func (item *SulfurasItem) Process() bool {
	item.Quality += 1
	return true
}

func UpdateQuality(items []*Item) {
	for i := 0; i < len(items); i++ {

		if notAgedBrie(items[i]) && notBackPasses(items[i]) && items[i].Quality > 0 && notSulfuras(items[i]) {
			changeQuality(items[i], -1)
		} else {
			if belowMaxQuality(items[i]) {
				changeQuality(items[i], 1)
				if !notBackPasses(items[i]) {
					if items[i].SellIn < 11 {
						if belowMaxQuality(items[i]) {
							changeQuality(items[i], 1)
						}
					}
					if items[i].SellIn < 6 {
						if belowMaxQuality(items[i]) {
							changeQuality(items[i], 1)
						}
					}
				}
			}
		}

		if notSulfuras(items[i]) {
			items[i].SellIn = items[i].SellIn - 1
		}

		if items[i].SellIn < 0 {
			if notAgedBrie(items[i]) {
				if notBackPasses(items[i]) {
					if items[i].Quality > 0 {
						if notSulfuras(items[i]) {
							changeQuality(items[i], -1)
						}
					}
				} else {
					items[i].Quality = items[i].Quality - items[i].Quality
				}
			} else {
				if belowMaxQuality(items[i]) {
					changeQuality(items[i], 1)
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

func changeQuality(item *Item, quality int) {
	item.Quality += quality
}

func notSulfuras(item *Item) bool {
	return item.Name != "Sulfuras, Hand of Ragnaros"
}
