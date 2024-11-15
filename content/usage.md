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

## Installation Options

### Direct Download
Download the latest release from our [releases page](https://github.com/kociumba/SkyDriver/releases).

### Go Install
If you have Go installed, you can use:
```bash
go install -ldflags="-s -w" github.com/kociumba/SkyDriver
```

> [!note]
> On Windows, installing via `go install` will result in a binary without an icon and typical Windows metadata.

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

1. Search for specific items:
```bash
SkyDriver -search flawless -max 100
```

2. Find high-volume trading opportunities:
```bash
SkyDriver -limit 10000000 -sell 100000 -max 100
```

3. Get JSON output for automation:
```bash
SkyDriver -json -search flawless -max 100 | jq '.results | length'
```

## Understanding the Output

### Table Columns

- **Product**: Item name
- **SellPrice**: Current instant sell price
- **BuyPrice**: Current instant buy price
- **Difference**: Potential profit margin
- **Weekly Traffic**: Trading volume data
- **Predicted/Confidence**: Market prediction and confidence level

### Prediction System

The prediction system helps evaluate potential profitability:

- **Positive numbers**: Likely profitable
- **Negative numbers**: Potentially risky
- **Confidence percentage**: Agreement level among internal indicators

> [!warning]
> **Important:** The prediction system should be used as one of many tools in your decision-making process, not as the sole indicator for trading decisions.

## Filtering Behavior

By default, SkyDriver filters out items that have:
- Less than 100 buy and sell price
- Less than 10 sales per week

> [!note]
> This filtering is disabled when using the `-search` flag, where only items with 0 sell price and 0 buy price are filtered.

## Advanced Usage

### JSON Output
Use the `-json` flag for programmatic access to data. The output follows a [defined schema](api/schema.json) and is compatible with tools like `jq`.

> [!tip]
> The `-json` flag automatically disables interactive prompts.

### View All Items
To see a complete list of bazaar items:
```bash
SkyDriver -max 10000000
```

## Getting Help

If you encounter issues:
1. Review this documentation
2. Check common CLI usage patterns
3. [Open an issue](https://github.com/kociumba/SkyDriver/issues/new/choose) if problems persist

## Mathematical Model

The prediction system uses seven internal indicators to evaluate market conditions. For detailed information about the mathematical model and its components, see our [[math|Mathematical Model]] page.

> [!tip]
> Remember that market conditions can change rapidly. Always verify predictions against current market conditions and your own research.
