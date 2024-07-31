# SkyDriver

<!-- TOC -->
- [SkyDriver](#skydriver)
  - [Features](#features)
  - [Usage](#usage)
  - [FAQ (not actually because this has 0 users 😭)](#faq-not-actually-because-this-has-0-users-)
<!-- TOC -->

SkyDriver is a companion application for Hypixel Skyblock, currently in early development.

## Features

The primary focus of SkyDriver is on bazaar data and flipping, with plans for future expansion. If you have ideas or suggestions for improvements, please [open an issue](https://github.com/kociumba/SkyDriver/issues/new/choose).

SkyDriver generates a detailed breakdown of bazaar data sourced from the Hypixel API. This data is accessible via a public endpoint, meaning you do not need a Hypixel API key to utilize SkyDriver.

The current output format is as follows:

- Command:
```console
SkyDriver -limit 10000
```

- Output:
```
╔═══════════════════════════════╦═════════╦════════╦══════════╦═══════════════════════════╗
║Product/price limit: 10000.00  ║SellPrice║BuyPrice║Difference║       Weekly Trafic       ║
╠═══════════════════════════════╬═════════╬════════╬══════════╬═══════════════════════════╣
║1.  ENCHANTMENT_TURBO_CARROT_1 ║1931.36  ║9901.22 ║7969.86   ║Sell:14741   |  Buy:4278   ║
║2.  MIDAS_JEWEL                ║501.90   ║8348.29 ║7846.39   ║Sell:1091    |  Buy:342    ║
║3.  ENDER_MONOCLE              ║2953.02  ║9745.29 ║6792.27   ║Sell:27715   |  Buy:27676  ║
║4.  COMPACTOR                  ║2003.08  ║8544.46 ║6541.37   ║Sell:15314   |  Buy:20142  ║
║5.  GOBLIN_EGG_RED             ║3574.81  ║9590.39 ║6015.58   ║Sell:75278   |  Buy:57206  ║
║6.  ENCHANTMENT_TURBO_WHEAT_1  ║963.94   ║6688.90 ║5724.96   ║Sell:20099   |  Buy:5970   ║
║7.  HYPERGOLIC_IONIZED_CERAMICS║2994.68  ║8362.60 ║5367.92   ║Sell:7792    |  Buy:2811   ║
║8.  HAZMAT_ENDERMAN            ║100.40   ║5128.86 ║5028.46   ║Sell:10748   |  Buy:10580  ║
║9.  HEALING_TISSUE             ║499.15   ║5474.30 ║4975.15   ║Sell:74783   |  Buy:35194  ║
║10. EXPIRED_PUMPKIN            ║5005.42  ║9768.20 ║4762.78   ║Sell:268191  |  Buy:227983 ║
╚═══════════════════════════════╩═════════╩════════╩══════════╩═══════════════════════════╝
```

As you may have noticed, there is a price limit feature that can be enforced. This is done by passing `-limit (price)` as an argument to SkyDriver, which restricts the results to products with an instant buy price below this value.
 
## Usage

To use SkyDriver, download the latest release from the [releases](https://github.com/kociumba/SkyDriver/releases).
You can run SkyDriver to retrieve a list of the 10 products with the highest price differences, or use the `-limit` option to filter items within your flipping budget.

There is also `-sell (amount of sells per week)` this essentially filters all items that have less sales per week than the specified value.

```console
SkyDriver -limit 100000 -sell 10000
```

This command will show the list of 10 products with the highest price differences up to the sell price of 100000 and have more than 10000 sales per week.

> [!NOTE]
> SkyDriver filters items that have less than 100 buy and sell price, as well as items that are sold less than 10 times a week.


## FAQ (not actually because this has 0 users 😭)

**Q:** Why? <br>
**A:** I enjoy command line interfaces, and retrieving data this way feels more intuitive to me. 🙂

**Q:** Why not create a mod? <br>
**A:** I am most familiar with Go, and the Java ecosystem is a shithole (my condolences to all java devs out there 💀).

**Q:** How do I [blank]? <br>
**A:** If you are having trouble with using SkyDriver, read any 
resource out there on cli apps and if you get past opening the app and your issues persists, 
please [open an issue](https://github.com/kociumba/SkyDriver/issu.es/new/choose)