// SPDX-License-Identifier: MIT
const PSCF = artifacts.require("PriceScf");

module.exports = function (deployer) {
deployer.deploy(PSCF);
};