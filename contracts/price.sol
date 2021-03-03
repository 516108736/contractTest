// SPDX-License-Identifier: MIT
pragma solidity ^0.6.7;

import "@chainlink/contracts/src/v0.6/interfaces/AggregatorV3Interface.sol";

contract PriceScf {

    AggregatorV3Interface internal priceFeedETHUSDT;
    AggregatorV3Interface internal priceFeedBTCUSDT;

    /**
     * Network: Kovan
     * Aggregator: ETH/USD BTC / USD
     * Address: 0x9326BFA02ADD2366b30bacB125260Af641031331,0x6135b13325bfC4B00278B4abC5e20bbce2D6580e
     */
    constructor() public {
        priceFeedETHUSDT = AggregatorV3Interface(0x9326BFA02ADD2366b30bacB125260Af641031331);
        priceFeedBTCUSDT = AggregatorV3Interface(0x6135b13325bfC4B00278B4abC5e20bbce2D6580e);
    }

    /**
     * Returns the latest price
     */
    function getETHUSDTLatestPrice() public view returns (int) {
        (
        uint80 roundID,
        int price,
        uint startedAt,
        uint timeStamp,
        uint80 answeredInRound
        ) = priceFeedETHUSDT.latestRoundData();
        return price;
    }
    /**
    * Returns the latest price
    */
    function getBTCUSDTLatestPrice() public view returns (int) {
        (
        uint80 roundID,
        int price,
        uint startedAt,
        uint timeStamp,
        uint80 answeredInRound
        ) = priceFeedBTCUSDT.latestRoundData();
        return price;
    }
}