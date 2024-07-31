# Technical notes on math used in SkyDriver

These notes are public because I frankly suck at math and it's much easier for non coders to spot math problems
in this format than in code.

In other words if I fucked something up here please report it in [issues](https://github.com/kociumba/SkyDriver/issues/new/choose).

## Price Prediction Math

This is how SkyDriver calculates the price rise or fall predictions without using historical bazaar data.

- ### Average Buying Price

    The average buying price $( P_{avg} )$ is calculated as:

    
    $$P_{avg} = \frac{1}{n} \sum_{i=1}^{n} P_{buy,i}$$


    where $( P_{buy,i} )$ is the price per unit for each purchase and $( n )$ is the total number of purchases.

- ### Price Change

    The price change $( PC )$ in percentage is given by:

    
    $$PC = \frac{P_{sell} - P_{avg}}{P_{avg}} \times 100$$
    

    where $( P_{sell} )$ is the current selling price.

- ### Market Trend

    The market trend $( MT )$ in percentage is calculated as:

    
    $$MT = \frac{V_{sell} - V_{buy}}{V_{buy}} \times 100$$

    where $( V_{sell} )$ is the selling volume and $( V_{buy} )$ is the buying volume.

- ### Price Prediction

    The predicted price change $( P_{pred} )$ is computed by averaging the price change and the market trend:

    $$P_{pred} = \frac{PC + MT}{2}$$


    This is a pretty rough prediction that doesn not account for any historical data or current events in the market.
    
    So please never rely on it, for example while programming this the prediction for enchanted diamond blocks was 
    <span style="color: limegreen;">
    $▲ 20442.77%$ 
    </span>
    which is not very possible with a sell price of $200174.92$ and a buy price of $204504.24$, the discrapency came from there being almoast 2x more sells than buys with these numbers respectively $Sell=2093369, Buy=1097683$
