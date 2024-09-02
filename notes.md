# Technical notes on math used in SkyDriver

These notes are public because I frankly suck at math and it's much easier for non coders to spot math problems
in this format than in code.

**In other words if I fucked something up here please report it in [issues](https://github.com/kociumba/SkyDriver/issues/new/choose).**

If you do want to read the code for this it's in [internal/priceFluctuation](internal/priceFluctuation.go)

## Price Prediction Math

This is how SkyDriver profit predictions without using historical bazaar data.

Smoothing Functions
Before applying the calculations, we now use smoothing functions to normalize the inputs. The current implementation allows for different smoothing functions:

**No Smoothing:** Raw values are used without modification.

**Sigmoid Smoothing:**
$$f(x) = \frac{200}{1 + e^{-kx}} - 100$$

**Tanh Smoothing:**
$$f(x) = 100 \tanh(kx)$$

**Saturating Smoothing:**
$$f(x) = \frac{100x}{\sqrt{1 + kx^2}}$$

**Piecewise Smoothing:**
$$f(x) = \begin{cases}
\frac{x}{(1 + (\frac{x}{100})^n)^{\frac{1}{n}}} & \text{if } x > 0 \
-\frac{-x}{(1 + (\frac{-x}{100})^n)^{\frac{1}{n}}} & \text{if } x \leq 0
\end{cases}$$

Where $k$ and $n$ are adjustable parameters controlling the steepness of the function.

- ### Price Spread $(PS)$
    $$PS = \frac{buyPrice - sellPrice}{sellPrice} \times 100$$

    This gives us the percentage spread between the buy and sell prices.
    This calculation assumes that the user will be flipping sell and buy orders.

- ### Volume Imbalance $(VI)$
    $$VI = \frac{buyVolume - sellVolume}{buyVolume + sellVolume} \times 100$$

    This measures the imbalance between buy and sell volumes.

- ### Order Imbalance $(OI)$
    $$OI = \frac{buyOrders - sellOrders}{buyOrders + sellOrders} \times 100$$

    This measures the imbalance between the number of buy and sell orders.

- ### Moving Week Trend $(MWT)$
    $$MWT = \frac{buyMovingWeek - sellMovingWeek}{buyMovingWeek + sellMovingWeek} \times 100$$

    This gives us a sense of the longer-term trend based on the past week's activity.

- ### Top Order Book Pressure $(TOBP)$
    Using the top 30 orders from buy_summary and sell_summary:

    $$TOBP = \frac{\sum_{i=1}^{30} (buyAmount_i \times buyPrice_i) - \sum_{i=1}^{30} (sellAmount_i \times sellPrice_i)}{\sum_{i=1}^{30} (buyAmount_i \times buyPrice_i) + \sum_{i=1}^{30} (sellAmount_i \times sellPrice_i)} \times 100$$

    This measures the pressure from the visible orders.

- ### Volume Factor (VF)
  
    $V_{total} = buyMovingWeek + sellMovingWeek$

    Total volume:
    
    $$VF = -100 + \left( \frac{V_{total} - V_{low}}{V_{high} - V_{low}} \right) \times 200$$

- ### Profit Margin Factor Calculation $(PMF)$ *[New]*

    The profit margin factor $PMF$ is calculated based on the profit margin as a percentage of the sell price.

    $PM = sellPrice - buyPrice$

    $PM_{percentage} = \frac{PM}{sellPrice}$

    Margin percentage:
   
    $$PMF = -100 + \left( \frac{PM_{percentage} - PM_{low}}{PM_{high} - PM_{low}} \right) \times 200$$

- ### Price Prediction Formula
    Combine these factors with appropriate weights:

    $$P_{pred} = w_1 \times PS + w_2 \times VI + w_3 \times OI + w_4 \times MWT + w_5 \times TOBP + w_6 \times VF + w_7 \times PMF$$

    Where $w_1$, $w_2$, $w_3$, $w_4$, $w_5$, $w_6$, and $w_7$ are weights that sum to 1.

- ### Interpretation

    A positive $P_{pred}$ suggests a potential price increase.

    A negative $P_{pred}$ suggests a potential price decrease.

    The magnitude of $P_{pred}$ indicates the strength of the prediction.

- ### Confidence Measure
    We can create a simple confidence measure based on the consistency of our indicators:

    $$Confidence = \frac{\text{Number of indicators with the same sign as } P_{pred}}{7} \times 100$$

    This gives us a percentage confidence in our prediction.
