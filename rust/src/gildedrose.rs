use std::fmt::{self, Display};

pub struct Item {
    pub name: String,
    pub sell_in: i32,
    pub quality: i32,
}

impl Item {
    pub fn new(name: impl Into<String>, sell_in: i32, quality: i32) -> Item {
        Item {
            name: name.into(),
            sell_in,
            quality,
        }
    }
}

impl Display for Item {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        write!(f, "{}, {}, {}", self.name, self.sell_in, self.quality)
    }
}

pub struct GildedRose {
    pub items: Vec<Item>,
}

const BACKSTAGE_PASSES: &'static str = "Backstage passes to a TAFKAL80ETC concert";
const AGED_BRIE: &'static str = "Aged Brie";
const SULFURAS: &'static str = "Sulfuras, Hand of Ragnaros";
const CONJURED_PREFIX: &'static str = "Conjured";

const MAX_QUALITY: i32 = 50;
const MIN_QUALITY: i32 = 0;

impl GildedRose {
    pub fn new(items: Vec<Item>) -> GildedRose {
        GildedRose { items }
    }

    pub fn update_quality(&mut self) {
        for item in &mut self.items {
            if item.name == SULFURAS {
                continue;
            }

            if item.name == AGED_BRIE {
                item.quality = clamp_quality(item.quality + 1);
            } else if item.name == BACKSTAGE_PASSES {
                item.quality = clamp_quality(item.quality + 1);

                if item.sell_in < 11 {
                    item.quality = clamp_quality(item.quality + 1);
                }

                if item.sell_in < 6 {
                    item.quality = clamp_quality(item.quality + 1);
                }
            } else if item.name.starts_with(CONJURED_PREFIX) {
                item.quality = clamp_quality(item.quality - 2);
            } else {
                item.quality = clamp_quality(item.quality - 1);
            }

            item.sell_in = item.sell_in - 1;

            if item.sell_in < 0 {
                if item.name == AGED_BRIE {
                    item.quality = clamp_quality(item.quality + 1);
                } else if item.name == BACKSTAGE_PASSES {
                    item.quality = clamp_quality(item.quality - item.quality);
                } else if item.name.starts_with(CONJURED_PREFIX) {
                    item.quality = clamp_quality(item.quality - 2);
                } else {
                    item.quality = clamp_quality(item.quality - 1);
                }
            }
        }
    }
}

/*
    clamp_quality keeps quality from exceeding min/max
 */
fn clamp_quality(i: i32) -> i32 {
    match i {
        i if i > MAX_QUALITY => MAX_QUALITY,
        i if i < MIN_QUALITY => MIN_QUALITY,
        _ => i
    }
}

#[cfg(test)]
mod tests {
    use super::{GildedRose, Item};

    macro_rules! update_quality_test {
        ($name:ident, $item_name:expr, $input_quality:expr, $input_sell_in:expr, $expected_quality:expr, $expected_sell_in:expr) => {
            #[test]
            pub fn $name() {
                let items = vec![Item::new($item_name, $input_sell_in, $input_quality)];
                let mut rose = GildedRose::new(items);
                rose.update_quality();

                assert_eq!($expected_quality, rose.items[0].quality, "expected quality {} received {}", $expected_quality, rose.items[0].quality);
                assert_eq!($expected_sell_in, rose.items[0].sell_in, "expected sell_in {} received {}", $expected_sell_in, rose.items[0].sell_in);
            }
        };
    }

    update_quality_test!(normal_item_should_have_quality_and_sellin_decremented, "Shield of Sadness", 10, 10, 9, 9);
    update_quality_test!(normal_item_should_have_quality_decrease_twice_after_sell_in, "Shield of Sadness", 10, 0, 8, -1);
    update_quality_test!(normal_item_should_not_have_quality_below_0, "Shield of Sadness", 0, 10, 0, 9);
    update_quality_test!(normal_item_should_not_have_quality_below_0_after_sell_in, "Shield of Sadness", 1, 0, 0, -1);

    update_quality_test!(aged_brie_should_increase_in_quality, "Aged Brie", 10, 10, 11, 9);
    update_quality_test!(aged_brie_should_increase_quality_twice_after_sell_in, "Aged Brie", 10, 0, 12, -1);
    update_quality_test!(aged_brie_should_not_exceed_quality_of_50, "Aged Brie", 50, 10, 50, 9);
    update_quality_test!(aged_brie_should_not_exceed_quality_of_50_after_sell_in, "Aged Brie", 49, 0, 50, -1);

    update_quality_test!(sulfuras_should_not_change_sellin_or_quality, "Sulfuras, Hand of Ragnaros", 30, 20, 30, 20);
    update_quality_test!(sulfuras_should_not_change_sellin_or_quality_even_exceeding_max_quality, "Sulfuras, Hand of Ragnaros", 80, 30, 80, 30);
    update_quality_test!(sulfuras_should_not_change_sellin_or_quality_even_exceeding_min_quality, "Sulfuras, Hand of Ragnaros", -20, -10, -20, -10);

    update_quality_test!(concert_tickets_should_increase_in_quality_when_greater_than_10_days_out_15, "Backstage passes to a TAFKAL80ETC concert", 10, 15, 11, 14);
    update_quality_test!(concert_tickets_should_increase_in_quality_when_greater_than_10_days_out_10, "Backstage passes to a TAFKAL80ETC concert", 10, 11, 11, 10);
    update_quality_test!(concert_tickets_quality_should_not_exceed_50_when_greater_than_10_days_out, "Backstage passes to a TAFKAL80ETC concert", 50, 11, 50, 10);
    update_quality_test!(concert_tickets_should_increase_in_quality_twice_when_less_than_10_days_out_10, "Backstage passes to a TAFKAL80ETC concert", 10, 10, 12, 9);
    update_quality_test!(concert_tickets_should_increase_in_quality_twice_when_less_than_10_days_out_6, "Backstage passes to a TAFKAL80ETC concert", 10, 6, 12, 5);
    update_quality_test!(concert_tickets_quality_should_not_exceed_50_when_less_than_10_days_out, "Backstage passes to a TAFKAL80ETC concert", 49, 6, 50, 5);
    update_quality_test!(concert_tickets_should_increase_in_quality_thrice_when_less_than_5_days_out_5, "Backstage passes to a TAFKAL80ETC concert", 10, 5, 13, 4);
    update_quality_test!(concert_tickets_should_increase_in_quality_thrice_when_less_than_5_days_out_1, "Backstage passes to a TAFKAL80ETC concert", 10, 1, 13, 0);
    update_quality_test!(concert_tickets_quality_should_not_exceed_50_when_less_than_5_days_out, "Backstage passes to a TAFKAL80ETC concert", 48, 1, 50, 0);
    update_quality_test!(concert_tickets_should_have_quality_set_to_0_after_sell_in, "Backstage passes to a TAFKAL80ETC concert", 10, 0, 0, -1);
    update_quality_test!(concert_tickets_should_have_quality_set_to_0_long_after_sell_in, "Backstage passes to a TAFKAL80ETC concert", 10, -5, 0, -6);

    update_quality_test!(conjured_items_should_degrade_twice_as_fast, "Conjured Mana Cake", 10, 10, 8, 9);
    update_quality_test!(conjured_items_should_degrade_4x_after_sell_in, "Conjured Dressing Gown", 10, 0, 6, -1);
    update_quality_test!(conjured_items_quality_should_not_go_below_0, "Conjured Saddle Bags", 1, 10, 0, 9);
    update_quality_test!(conjured_items_quality_should_not_go_below_0_after_sell_in, "Conjured Saddle Bags", 3, 0, 0, -1);
}
