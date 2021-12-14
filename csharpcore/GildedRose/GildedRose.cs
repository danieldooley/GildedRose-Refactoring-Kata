using System;
using System.Collections.Generic;

namespace GildedRoseKata
{
    public class GildedRose
    {
        private const string AgedBrie = "Aged Brie";
        private const string BackstagePasses = "Backstage passes to a TAFKAL80ETC concert";
        private const string Sulfuras = "Sulfuras, Hand of Ragnaros";

        IList<Item> Items;

        public GildedRose(IList<Item> Items)
        {
            this.Items = Items;
        }

        public void AdjustQuality(Item item, int delta)
        {
            item.Quality += delta;
            item.Quality = Math.Clamp(item.Quality, 0, 50);
        }

        public void UpdateQuality()
        {
            foreach (var item in Items)
            {
                if (item.Name == Sulfuras)
                {
                    continue;
                }

                item.SellIn -= 1;

                var delta = item.Name switch
                {
                    AgedBrie => item.SellIn < 0 ? 2 : 1,
                    BackstagePasses => item.SellIn switch
                    {
                        var i when i < 0 => -item.Quality,
                        var i when i < 5 => 3,
                        var i when i < 10 => 2,
                        _ => 1
                    },
                    var name when name.StartsWith("Conjured") => item.SellIn < 0 ? -4 : -2,
                    _ => item.SellIn < 0 ? -2 : -1
                };

                AdjustQuality(item, delta);
            }
        }
    }
}