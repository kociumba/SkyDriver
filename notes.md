# Technical notes on math used in SkyDriver

These notes are public because I frankly suck at math and it's much easier for non coders to spot math problems
in this format than in code.

In other words if I fucked something up here please report it in [issues](https://github.com/kociumba/SkyDriver/issues/new/choose).

If you do want to read the code for this it's in [internal/priceFluctuation.go](internal/priceFluctuation.go)

## Price Prediction Math

This is how SkyDriver calculates the price rise or fall predictions without using historical bazaar data.

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

    This measures the pressure from the visible order book.

- ### 6. Volume Factor (VF) [New]
  
    Let $V_{total} = buyMovingWeek + sellMovingWeek$

    
    1. **If the total volume is less than or equal to the low volume threshold  $V_{low}$**:

        $$VF = -100$$

    2. **If the total volume is greater than or equal to the high volume threshold $V_{high}$**:

        $$VF = 100$$

    3. **If the total volume is between the low volume threshold $V_{low}$ and the high volume threshold $V_{high}$**:
    
        $$VF = -100 + \left( \frac{V_{total} - V_{low}}{V_{high} - V_{low}} \right) \times 200$$
  
    Where:

    $V_{low}$ is the low volume threshold
    $V_{high}$ is the high volume threshold

- ### Price Prediction Formula
    Combine these factors with appropriate weights:

    $$P_{pred} = w_1 \times PS + w_2 \times VI + w_3 \times OI + w_4 \times MWT + w_5 \times TOBP + w_6 \times VF$$

    Where $w_1$, $w_2$, $w_3$, $w_4$, $w_5$, and $w_6$ are weights that sum to 1.

- ### Interpretation

    A positive $P_{pred}$ suggests a potential price increase.

    A negative $P_{pred}$ suggests a potential price decrease.

    The magnitude of $P_{pred}$ indicates the strength of the prediction.

- ### Confidence Measure
    We can create a simple confidence measure based on the consistency of our indicators:

    $$Confidence = \frac{\text{Number of indicators with the same sign as } P_{pred}}{6} \times 100$$

    This gives us a percentage confidence in our prediction.
