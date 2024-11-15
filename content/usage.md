---
title: Basic usage
draft: false
tags:
  - user-guide
---

# SkyDriver User Guide

SkyDriver is a powerful command-line companion application for Hypixel Skyblock, focusing on bazaar data analysis and flipping opportunities. This guide will walk you through everything you need to know to use SkyDriver effectively.

## Quick Start

1. Download the latest release from the [releases page](https://github.com/kociumba/SkyDriver/releases)
2. Run SkyDriver in your terminal of choice to get started!

For more detail, keep reading below.

## Installation Options

### Direct Download
Download the latest release from our [releases page](https://github.com/kociumba/SkyDriver/releases).

### Go Install
If you have Go installed, you can use:
```bash
go install -ldflags="-s -w" github.com/kociumba/SkyDriver
```

> [!note]
> On Windows, installing via `go install` will result in an executable without an icon and typical Windows metadata.

## Basic Usage

SkyDriver provides a clean, table-based output showing bazaar opportunities:

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

> [!note]
> The output is normally colored but it's impossible to showcase this here without a screenshot.

## Command-Line Options

### Core Parameters

| Flag | Description | Example |
|------|-------------|---------|
| `-limit` | Maximum buy price filter | `-limit 1000000` |
| `-sell` | Minimum weekly sales filter | `-sell 100000` |
| `-search` | Search items by name | `-search flawless` |
| `-max` | Maximum items to display | `-max 100` |
| `-json` | Output in JSON format | `-json` |
| `-skip` | Skip interactive prompts | `-skip` |

### Example Commands

1. Search for top 100 items containing "flawless" in their name:
```bash
SkyDriver -search flawless -max 100
```

2. Find the top 100 items with a price under 10,000,000 and more than 100,000 sales per week:
```bash
SkyDriver -limit 10000000 -sell 100000 -max 100
```

3. Get JSON output for processing with tools like `jq`:
```bash
SkyDriver -json -search flawless -max 100 | jq '.results | length'
```

## Understanding the Output

### Table Columns

- **Product**: Item name(This is the internal hypixel api name, as I didn't implement a way to get the real names)
- **SellPrice**: Current instant sell price
- **BuyPrice**: Current instant buy price
- **Difference**: Simple difference between buy and sell prices
- **Weekly Traffic**: Amounts of buys and sells from the last 7 days
- **Predicted/Confidence**: Prediction of profitability and confidence level of the prediction

### Prediction System

The prediction system helps evaluate potential profitability:

- **Positive numbers**: Likely profitable - bigger numbers means higher predicted prifitability
- **Negative numbers**: Potentially risky - bigger negative numbers means higher predicted unprofitability
- **Confidence percentage**: Agreement level among 7 internal indicators, see [math](math.md) for more

> [!warning]
> **Important:** The prediction model was made by me (a dumbass) and you should not rely on it. It will make some mistakes, so always use common sense before making any decisions.

## Filtering Behavior

By default, SkyDriver filters out items that have:
- Less than 100 buy and sell price
- Less than 10 sales per week

> [!note]
> This filtering is disabled when using the `-search` flag, where only items with 0 sell price and 0 buy price are filtered.

## Advanced Usage

### JSON Output
Use the `-json` flag for programmatic access to data. The output follows a [defined schema](https://github.com/kociumba/SkyDriver/blob/main/api/schema.json) and is compatible with tools like `jq`.

This allows for more complex filtering like for example:

```bash
SkyDriver -json -search flawless -max 2 | jq -r '.results[] | "\(.product_id), \((.prediction * 100 | round) / 100), \((.confidence * 100 | round) / 100)"'
```

Which outputs something like:
```bash
FLAWLESS_JASPER_GEM, 37.58, 71.43
FLAWLESS_ONYX_GEM, 9.35, 57.14
```

> [!tip]
> The `-json` flag automatically disables interactive prompts.

### Miscellaneous Stuff
To see a very long list of all the items on the bazaar run this command:
```bash
SkyDriver -max 10000000
```

## Getting Help

If you encounter issues:
1. Read up on CLI tools and how they work
2. Read these [docs](index.md) under the #user-guide tag, they should contain anwsers to most of your questions
3. [Open an issue](https://github.com/kociumba/SkyDriver/issues/new/choose) if your problems persist

## Mathematical Model

The prediction system uses seven internal indicators to evaluate the bazaar info. For detailed information about the mathematical model and its components, see our [math](math.md) page.
> [!important]
> Once again I am not a mathematician and the prediction model will make some mistakes. Always use common sense before flipping.
