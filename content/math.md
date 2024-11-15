---
title: The Math
draft: false
tags:
  - internals
---

# The Math Behind SkyDriver 🧮

*This document is just a a summary of [notes.md](https://github.com/kociumba/SkyDriver/blob/main/notes.md)*

> [!warning]
> Math ahead! If you just want to use SkyDriver, you can skip this. But if you're curious how the predictions work, read on!

> [!note]
> Found a problem with the math? Please [report it](https://github.com/kociumba/SkyDriver/issues/new/choose)! I'm not a mathematician, so there might be mistakes.

## Overview

SkyDriver uses several indicators to predict if an item will be profitable to flip. Each indicator looks at different aspects of the market, and they work together to give us a prediction.

## Smoothing Functions

We smooth out the results of the other calculations to make them more manageable. We have a few ways to do this:

### No Smoothing
Just use the raw numbers (boring but sometimes useful)

### Sigmoid Smoothing
$$f(x) = \frac{200}{1 + e^{-kx}} - 100$$

### Tanh Smoothing
$$f(x) = 100 \tanh(kx)$$

### Saturating Smoothing
$$f(x) = \frac{100x}{\sqrt{1 + kx^2}}$$

### Piecewise Smoothing
For x > 0:
$$f(x) = \frac{x}{(1 + (\frac{x}{100})^n)^{\frac{1}{n}}}$$

For x ≤ 0:
$$f(x) = -\frac{-x}{(1 + (\frac{-x}{100})^n)^{\frac{1}{n}}}$$

> [!note]
> $k$ and $n$ control how aggressive the smoothing is

---

## Market Indicators

### 1. Price Spread (PS)
How different the buy and sell prices are:

$$PS = \frac{buyPrice - sellPrice}{sellPrice} \times 100$$

### 2. Volume Imbalance (VI)
Are more buys or sells going through?

$$VI = \frac{buyVolume - sellVolume}{buyVolume + sellVolume} \times 100$$

### 3. Order Imbalance (OI)
Are there more buy orders than sell orders?

$$OI = \frac{buyOrders - sellOrders}{buyOrders + sellOrders} \times 100$$

### 4. Moving Week Trend (MWT)
What's been happening over the last week:

$$MWT = \frac{buyMovingWeek - sellMovingWeek}{buyMovingWeek + sellMovingWeek} \times 100$$

### 5. Top Order Book Pressure (TOBP)
Looking at the top 30 orders:

$$TOBP = \frac{\sum_{i=1}^{30} (buyAmount_i \times buyPrice_i) - \sum_{i=1}^{30} (sellAmount_i \times sellPrice_i)}{\sum_{i=1}^{30} (buyAmount_i \times buyPrice_i) + \sum_{i=1}^{30} (sellAmount_i \times sellPrice_i)} \times 100$$

### 6. Volume Factor (VF)
How much trading is happening:

$V_{total} = buyMovingWeek + sellMovingWeek$

$$VF = -100 + \left( \frac{V_{total} - V_{low}}{V_{high} - V_{low}} \right) \times 200$$

### 7. Profit Margin Factor (PMF)
How much profit you might make:

$PM = sellPrice - buyPrice$

$PM_{percentage} = \frac{PM}{sellPrice}$

$$PMF = -100 + \left( \frac{PM_{percentage} - PM_{low}}{PM_{high} - PM_{low}} \right) \times 200$$

---

## Putting It All Together

### The Final Prediction
We combine all these indicators with different weights:

$$P_{pred} = w_1 \times PS + w_2 \times VI + w_3 \times OI + w_4 \times MWT + w_5 \times TOBP + w_6 \times VF + w_7 \times PMF$$

Where all the weights ($w_1$ through $w_7$) add up to 1.

> [!note]
> The weights are configurable in the [[config|Config]] if you want to change how the predictions are weighted

### What The Numbers Mean
- Positive number = Probably profitable
- Negative number = Probably not profitable
- Higher % = Stronger prediction

### Confidence Score
We also calculate how confident we are in the prediction:

$$Confidence = \frac{\text{Number of indicators agreeing with prediction}}{7} \times 100$$

> [!tip]
> Want to see how this works in code? Check out [internal/priceFluctuation](https://github.com/kociumba/SkyDriver/blob/main/internal/priceFluctuation.go)
