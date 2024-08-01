# SkyDriver

<!-- TOC -->
- [SkyDriver](#skydriver)
  - [Features](#features)
  - [Usage](#usage)
  - [FAQ (not actually because this has 0 users 😭)](#faq-not-actually-because-this-has-0-users-)
<!-- TOC -->

SkyDriver is a companion application for Hypixel Skyblock, currently in early development.

## Features

> [!IMPORTANT]
> There is a big update in v0.0.4 with a new prediction algorithm.
> To read more about it go to [notes](notes.md)

> [!WARNING]
> **Do not depend on the predictions.**
> They are purely for providing broad insights into the market. <br>
> Always do you own research before flipping lots of coins.

The primary focus of SkyDriver is on bazaar data and flipping, with plans for future expansion. If you have ideas or suggestions for improvements, please [open an issue](https://github.com/kociumba/SkyDriver/issues/new/choose).

SkyDriver generates a detailed breakdown of bazaar data sourced from the Hypixel API. This data is accessible via a public endpoint, meaning you do not need a Hypixel API key to utilize SkyDriver.

The current output format is as follows:

- Command:
```console
SkyDriver -limit 100000 -sell 100000
```

- Output:
```
╔═══════════════════════════════╦═════════╦════════╦══════════╦═══════════════════════════════╦════════════════════╗
║Product/price limit: 100000.00 ║SellPrice║BuyPrice║Difference║         Weekly Trafic         ║Predicted/Confidence║
╠═══════════════════════════════╬═════════╬════════╬══════════╬═══════════════════════════════╬════════════════════╣
║1.  FINE_AQUAMARINE_GEM        ║22217.20 ║39409.20║17192.00  ║Sell:317284    |  Buy:190413   ║▼ -43.47/100.00%    ║
║2.  ENCHANTED_CAKE             ║18028.53 ║34975.14║16946.61  ║Sell:102213    |  Buy:149330   ║▲ 0.32/60.00%       ║
║3.  MAGMA_CREAM_DISTILLATE     ║4611.41  ║16657.69║12046.28  ║Sell:110370    |  Buy:31200    ║▼ -63.92/100.00%    ║
║4.  HAMSTER_WHEEL              ║59745.58 ║68212.58║8466.99   ║Sell:432596    |  Buy:429592   ║▼ -57.53/100.00%    ║
║5.  GLACITE_JEWEL              ║44873.70 ║53025.92║8152.22   ║Sell:199886    |  Buy:122002   ║▼ -12.57/80.00%     ║
║6.  GOBLIN_EGG_YELLOW          ║2503.59  ║9990.90 ║7487.31   ║Sell:112769    |  Buy:74976    ║▼ -70.28/100.00%    ║
║7.  FOUL_FLESH                 ║27898.44 ║34999.35║7100.91   ║Sell:337081    |  Buy:301615   ║▼ -58.41/100.00%    ║
║8.  ENCHANTMENT_REJUVENATE_1   ║19854.46 ║26949.95║7095.49   ║Sell:317132    |  Buy:35152    ║▼ -59.01/100.00%    ║
║9.  ENCHANTED_GLOWSTONE        ║51016.56 ║56943.24║5926.68   ║Sell:167883    |  Buy:147181   ║▲ 17.95/40.00%      ║
║10. FINE_PERIDOT_GEM           ║42509.63 ║48399.20║5889.57   ║Sell:644392    |  Buy:749316   ║▼ -37.09/80.00%     ║
╚═══════════════════════════════╩═════════╩════════╩══════════╩═══════════════════════════════╩════════════════════╝
```

<!--As you may have noticed, there is a price limit feature that can be enforced. This is done by passing `-limit (price)` as an argument to SkyDriver, which restricts the results to products with an instant buy price below this value.-->
 
## Usage

To use SkyDriver, download the latest release from the [releases](https://github.com/kociumba/SkyDriver/releases).
You can run SkyDriver to retrieve a list of the 10 products with the highest price differences, or use the `-limit` option to filter items within your flipping budget.

There is also `-sell (amount of sells per week)` this essentially filters all items that have less sales per week than the specified value.

**New for v0.0.4**
The `-search (part of an item name)` option can be used to search for items matching the name you provide.

```console
SkyDriver -limit 100000 -sell 10000 -search gold
```

This command will show the list of all products up to the sell price of 100000 and 10000 sales per week that contain the word "gold" in their name.

> [!NOTE]
> SkyDriver filters items that have less than 100 buy and sell price, as well as items that are sold less than 10 times a week.
> 
> This behaviour is turned off when searching and only items with 0 sell price and 0 buy price are filtered.

With the release of v0.0.4, you can now set these in the cli interface.

If a value of the arguments is not passed, SkyDriver will prompt you to enter it in.
This can be skipped by just leaving the prompts blank or passing the `-skip` flag.

When searching for example for `flawless` all of the flawless gemstones will be considered, to display all 12
use the `-max (number of items to display)` flag which let's you set the maximum number of items to display.

<details>
  <summary>Pro tip</summary>

  run `SkyDriver -max 10000000` to show a very long table of all the items on the bazaar.

</details>

## FAQ (not actually because this has 0 users 😭)

**Q:** Why? <br>
**A:** I enjoy command line interfaces, and retrieving data this way feels more intuitive to me. 🙂

**Q:** Why not create a mod? <br>
**A:** I am most familiar with Go, and the Java ecosystem is a shithole (my condolences to all java devs out there 💀).

**Q:** How do I [blank]? <br>
**A:** If you are having trouble with using SkyDriver, read any 
resource out there on cli apps and if you get past opening the app and your issues persists, 
please [open an issue](https://github.com/kociumba/SkyDriver/issu.es/new/choose)