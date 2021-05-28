// SPDX-License-Identifier: GPL-3.0
pragma solidity ^0.7.0;
pragma experimental ABIEncoderV2;

import {Obi} from "./Obi.sol";

library ResultDecoder {
    using Obi for Obi.Data;

    struct Result {
        uint64[] count;
    }

    function decodeResult(bytes memory _data) internal pure returns (Result memory result) {
        Obi.Data memory data = Obi.from(_data);
        uint32 length = data.decodeU32();
        uint64[] memory count = new uint64[](length);
        for (uint256 i = 0; i < length; i++) {
            count[i] = data.decodeU64();
        }
        return Result(count);
    }
}
