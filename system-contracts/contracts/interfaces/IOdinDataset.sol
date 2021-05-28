// SPDX-License-Identifier: GPL-3.0
pragma solidity ^0.7.0;
pragma experimental ABIEncoderV2;

interface IOdinDataset {
    struct ReferenceData {
        uint256 count;
        uint256 lastUpdated;
    }

    function getReferenceData(string[] memory _models) external view returns (ReferenceData[] memory);
}
