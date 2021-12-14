using Xunit;
using System.Collections.Generic;
using System.Runtime.InteropServices;
using GildedRoseKata;

namespace GildedRoseTests
{
    public class GildedRoseTest
    {
        [Theory]
        [InlineData("Shield of Sorrow", 10, 10, 9, 9)] // Normal Item decrements SellIn and Quality
        [InlineData("Shield of Sorrow", 10, 0, 9, 0)] // Normal Item quality can not be negative
        [InlineData("Shield of Sorrow", 0, 10, -1, 8)] // Normal Item quality degrades twice once past sellin
        
        [InlineData("Aged Brie", 10, 10, 9, 11)] // Aged brie increases over time
        [InlineData("Aged Brie", 0, 10, -1, 12)] // Aged brie increases twice as fast past sellin
        [InlineData("Aged Brie", 10, 50, 9, 50)] // Aged brie quality cannot exceed 50
        [InlineData("Aged Brie", 0, 49, -1, 50)] // Aged brie increases twice but is clipped to 50
        
        [InlineData("Sulfuras, Hand of Ragnaros", 0, 80, 0, 80)] // Sulfuras is constant
        [InlineData("Sulfuras, Hand of Ragnaros", -1, 50, -1, 50)] // Sulfuras is constant
        
        [InlineData("Backstage passes to a TAFKAL80ETC concert", 15, 10, 14, 11)] // Back stage passes > 10 days
        [InlineData("Backstage passes to a TAFKAL80ETC concert", 11, 10, 10, 11)] // Back stage passes > 10 days
        [InlineData("Backstage passes to a TAFKAL80ETC concert", 10, 10, 9, 12)] // Back stage passes < 10 days
        [InlineData("Backstage passes to a TAFKAL80ETC concert", 6, 10, 5, 12)] // Back stage passes < 10 days
        [InlineData("Backstage passes to a TAFKAL80ETC concert", 5, 10, 4, 13)] // Back stage passes < 5 days
        [InlineData("Backstage passes to a TAFKAL80ETC concert", 1, 10, 0, 13)] // Back stage passes < 5 days
        [InlineData("Backstage passes to a TAFKAL80ETC concert", 0, 10, -1, 0)] // Back stage passes 0 days
        
        [InlineData("Conjured Mana Cake", 10, 10, 9, 8)] // Conjured items degrade twice as fast
        [InlineData("Conjured Ram's Head", 0, 10, -1, 6)] // Conjured item past sell in date degrades 4x
        public void About_UpdateQuality(string name, int sellIn, int quality, int expectedSellIn, int expectedQuality)
        {
            IList<Item> Items = new List<Item> { new Item { Name = name, SellIn = sellIn, Quality = quality } };
            GildedRose app = new GildedRose(Items);
            app.UpdateQuality();
            
            Assert.Equal(expectedQuality, Items[0].Quality);
            Assert.Equal(expectedSellIn, Items[0].SellIn);
        }
    }
}