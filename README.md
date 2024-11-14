# SkyDriver

<!-- TOC -->
- [SkyDriver](#skydriver)
  - [Features](#features)
  - [Usage](#usage)
  - [Examples](#examples)
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
> Always do your own research before flipping lots of coins.

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

To use SkyDriver, download the latest release from [releases](https://github.com/kociumba/SkyDriver/releases).

> [!TIP]
> If you have the latest version of [go](https://golang.org/doc/install) installed, you can run `go install -ldflags="-s -w" github.com/kociumba/SkyDriver` to install SkyDriver and automatically add it to path.
>
> On windows the built binary will lack an icon and any other typical data when istalled this way.

You can run SkyDriver to retrieve a list of the 10 products with the highest price differences, or use the `-limit` option to filter items within your flipping budget.

Use the `-sell` option followed by the number of sales per week to filter out items with fewer sales than specified.

**New in v0.0.4**

The `-search` option allows you to search for items by part of their name. For example, searching for flawless will include all items with "flawless" in their name. To display a specific number of items, use the `-max` flag to set the maximum number of items to display.

With the release of v0.0.4, you can now set these flags in the cli interface.

If a value of the arguments is not passed, SkyDriver will prompt you to enter it in.
This can be skipped by just leaving the prompts blank or passing the `-skip` flag.

**New in v0.0.6**

You can now use the `-json` flag to export results in JSON in this [schema](api/schema.json). This is all of the data shown in the normal table output.
The output is in a human readable format and ready for processing in tools like [jq](https://stedolan.github.io/jq/).

> [!TIP]
> The `-json` flag disables the prompts so you will have to set the parameters using other flags.

## Examples

```console
SkyDriver -search flawless -max 100
```

This command will display up to 100 items that contain the word "flawless" in their name.

```console
SkyDriver -json -search flawless -max 100 | jq '.results | length'
```

This command will pipe the output from SkyDriver into jq and return the number of search results found on the bazaar.

```console
SkyDriver -limit 10000000 -sell 100000 -max 100
```

This command will display the top 100 items that have more than 100,000 sales per week and a sell price below 10,000,000.

> [!NOTE]
> SkyDriver filters items that have less than 100 buy and sell price, as well as items that are sold less than 10 times a week.
> 
> This behaviour is turned off when searching and only items with 0 sell price and 0 buy price are filtered.

<details>
  <summary>Pro tip 😎</summary>

> [!TIP]
> Run `SkyDriver -max 10000000` to show a very long table of all the items on the bazaar.

</details>

## FAQ (not actually because this has 0 users 😭)

**Q:** Why? <br>
**A:** I enjoy command line interfaces, and retrieving data this way feels more intuitive to me. 🙂

**Q:** Why not create a mod? <br>
**A:** I am most familiar with Go, and the Java ecosystem is a shithole (my condolences to all java devs out there 💀).

**Q:** How do I [blank]? <br>
**A:** If you are having trouble with using SkyDriver, read any 
resource out there on cli apps and if you get past opening the app and your issues persists, 
please [open an issue](https://github.com/kociumba/SkyDriver/issues/new/choose)

**Q:** What is the prediction? <br>
**A:** The prediction comes from data collected through the bazaar API and helps you figure out if flipping an item will be profitable. If the prediction is positive, it means you're likely to make a profit; if it's negative, it's probably not worth it. The bigger the number, the higher the expected profit. Confidence shows how many of the seven internal indicators agree on whether the item will be profitable or not. For more info, check out the [model](notes.md)
> [!CAUTION]
> This model was made by me (a dumbass) and you should not soly rely on it.
> If you loose all your coins because you bet on something SkyDriver said was going to be profitable, that's on you.