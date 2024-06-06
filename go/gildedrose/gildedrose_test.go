package gildedrose_test

import (
	"testing"

	"github.com/emilybache/gildedrose-refactoring-kata/gildedrose"
	"github.com/stretchr/testify/require"
)

const (
	BackstagePasses = "Backstage passes to a TAFKAL80ETC concert"
	AgedBrie        = "Aged Brie"
	Sulfuras        = "Sulfuras, Hand of Ragnaros"
	Pera            = "Pera Proizvod"
)

func Test_SellInAndQualityDecreases(t *testing.T) {
	items := []*gildedrose.Item{
		{
			Name:    AgedBrie,
			SellIn:  24,
			Quality: 30,
		},
	}
	gildedrose.UpdateQuality(items)
	require.Equal(t, 31, items[0].Quality)
	require.Equal(t, 23, items[0].SellIn)
}

func Test_SulfurasSellInDecreasesAndQualityRemainsTheSame(t *testing.T) {
	items := []*gildedrose.Item{
		{
			Name:    Sulfuras,
			SellIn:  24,
			Quality: 30,
		},
	}
	gildedrose.UpdateQuality(items)
	require.Equal(t, 30, items[0].Quality)
	require.Equal(t, 24, items[0].SellIn)
}

func Test_QualityOfAnyItemIsNeverMoreThanFifty(t *testing.T) {
	items := []*gildedrose.Item{
		{
			Name:    AgedBrie,
			SellIn:  10,
			Quality: 50,
		},
	}

	gildedrose.UpdateQuality(items)
	require.Equal(t, 50, items[0].Quality)
	require.Equal(t, 9, items[0].SellIn)
}

func Test_QualityOfAnItemIsNeverNegative(t *testing.T) {
	items := []*gildedrose.Item{
		{
			Name:    BackstagePasses,
			SellIn:  0,
			Quality: 0,
		},
	}

	gildedrose.UpdateQuality(items)
	require.Equal(t, 0, items[0].Quality)
	require.Equal(t, -1, items[0].SellIn)
}

func Test_AfterSellInExpiryQualityDecreasesOnTheDouble(t *testing.T) {
	items := []*gildedrose.Item{
		{
			Name:    Pera,
			SellIn:  0,
			Quality: 10,
		},
	}

	gildedrose.UpdateQuality(items)
	require.Equal(t, -1, items[0].SellIn)
	require.Equal(t, 8, items[0].Quality)
}
func Test_BackStagePasses(t *testing.T) {
	t.Run("Between ten and five days before concert", func(t *testing.T) {
		items := []*gildedrose.Item{
			{
				Name:    BackstagePasses,
				SellIn:  9,
				Quality: 5,
			},
		}

		gildedrose.UpdateQuality(items)
		require.Equal(t, 7, items[0].Quality)
		require.Equal(t, 8, items[0].SellIn)
	})

	t.Run("Between five and zero days before concert", func(t *testing.T) {
		items := []*gildedrose.Item{
			{
				Name:    BackstagePasses,
				SellIn:  4,
				Quality: 5,
			},
		}

		gildedrose.UpdateQuality(items)
		require.Equal(t, 8, items[0].Quality)
		require.Equal(t, 3, items[0].SellIn)
	})

	t.Run("Quality drop to zero after concert.", func(t *testing.T) {
		items := []*gildedrose.Item{
			{
				Name:    BackstagePasses,
				SellIn:  0,
				Quality: 50,
			},
		}

		gildedrose.UpdateQuality(items)
		require.Equal(t, 0, items[0].Quality)
		require.Equal(t, -1, items[0].SellIn)
	})

}

func Test_AgedBrieIncreasesInQualityAsExpiryApproaches(t *testing.T) {
	items := []*gildedrose.Item{
		{
			Name:    AgedBrie,
			SellIn:  -1,
			Quality: 20,
		},
	}

	gildedrose.UpdateQuality(items)
	require.Equal(t, -2, items[0].SellIn)
	require.Equal(t, 22, items[0].Quality)
}
