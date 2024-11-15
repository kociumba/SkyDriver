---
title: Config
draft: false
tags:
  - config
  - user-guide
---

# The Config File 💾

SkyDriver uses a TOML file to store its configuration. The file is located at `~/.skydriver/config.toml`.

> [!note]
> For windows users, this is `C:/Users/<username>/.skydriver/config.toml` or `$HOME/.skydriver/config.toml`. But in recent times `~` also works on windows.

## What Can You Configure?

Right now, you can tweak how SkyDriver makes its predictions by adjusting the weights of different bazaar indicators.

### Prediction Weights

Here's what your config file looks like:
```toml
[prediction.weights]
# How important is the difference between buy and sell prices
price_spread = 0.1428571429

# How much do we care about whether more people are buying or selling
volume_imbalance = 0.1428571429

# How much do we care about the number of orders
order_imbalance = 0.1428571429

# How important is last week's trend
moving_week_trend = 0.1428571429

# How much do we look at the order book
order_book_pressure = 0.1428571429

# How much do we care about total trading volume
volume_factor = 0.1428571429

# How important is the potential profit margin
profit_margin_factor = 0.1428571429
```

> [!warning]
> All the weights must add up to 1.0! If they don't, you will end up with weird and unpredictable results.

### Default Values

By default, SkyDriver gives equal weight (≈0.143) to each factor. This means no single factor dominates the predictions.

## Customizing Your Config

<details>
<summary>Want to make SkyDriver focus more on certain factors? Here are some example configurations:</summary>

### Volume-Focused
If you want you recommendations to fill in buy orders and sell the sell orders fast:
```toml
[prediction.weights]
price_spread = 0.1
volume_imbalance = 0.2
order_imbalance = 0.1
moving_week_trend = 0.1
order_book_pressure = 0.1
volume_factor = 0.3        # More weight on volume
profit_margin_factor = 0.1
```

### Profit-Focused
If you want to focus on pure potential profit and don't care about speed:
```toml
[prediction.weights]
price_spread = 0.2
volume_imbalance = 0.1
order_imbalance = 0.1
moving_week_trend = 0.1
order_book_pressure = 0.1
volume_factor = 0.1
profit_margin_factor = 0.3  # More weight on profit
```
</details>

## How It Works

1. When you first run SkyDriver, it creates a default config file
    - if you want to reset your config to the defaults, just delete `config.toml` and SkyDriver will create a new one nex time you run it
2. You can edit this file with any text editor (if you have a code editor like VSCode, you can use intelisense with a provided schema)
3. SkyDriver reads this file every time it is run
4. If there's a problem with your config, it falls back to defaults and gives you a warning

> [!tip]
> Want to understand what each factor means? Check out the [[math|Math page]] for detailed explanations!

## Future Plans

More config options are coming! Some ideas:
- Custom operations on the bazaar data
- Output themes

> [!note]
> Got ideas for what else should be configurable? Open an issue on [GitHub](https://github.com/kociumba/SkyDriver/issues/new/choose)!
