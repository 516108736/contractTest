// SPDX-License-Identifier: MIT
const PV3 = artifacts.require("PriceConsumerV3");

module.exports = function (deployer) {
deployer.deploy(PV3);
};