#!/usr/bin/env python

import os

class OrderBook:
    def __init__(self):
        self.bids = {}
        self.asks = {}
        
    def place_sell_order(self, price, shares):
        if price not in self.asks:
            self.asks[price] = shares
        else:
            self.asks[price] += shares

    def place_buy_order(self, price, shares):
        if price not in self.bids:
            self.bids[price] = shares
        else:
            self.bids[price] += shares

    def exe_orders(self):
        # bid from high to low
        # ask from low to high
        bid_prices = sorted(self.bids.keys(), reverse=False) # [10.2, 9.6, 8.5, 7.5, 7.2, 6.9, 6.2, 5.5, 5.1]
        ask_prices = sorted(self.asks.keys())                # [9.5, 9.6, 9.8, 10.1, 10.2, 10.3, 11.2, 12.5]

        # flatten bid shares
        bids = []
        for price in bid_prices:
            shares = self.bids[price]
            for i in range(shares):
                bids.append(bid)

        # flatten ask shares
        asks = []
        for price in ask_prices:
            shares = self.asks[price]
            for i in range(shares):
                asks.append(bid)

        # to maximize profit
        # match orders based on both spread and number of deals we want to make
        ask_value, bid_value = 0, 0
        shares = 0
        index = 0
        while index < len(asks) and index < len(bids):
            # ask and bid prices
            ask = asks[index]
            bid = bids[index]

            # stop or not
            if ask_value + ask > bid_value + bid:
                break
            else:
                ask_value += ask
                bid_value += bid

            # remove this ask share
            self.asks[ask] -= 1
            if self.asks[ask] == 0:
                del self.asks[ask]

            # remove this bid share
            self.bids[bid] -= 1
            if self.bids[bid] == 0:
                del self.bids[bid]
            
            shares += 1
            index += 1

if __name__ == "__main__":
    book = OrderBook()
    bids = {}
    bids[3.14] = 25
    bids[4.5] = 10
    print(bids)
