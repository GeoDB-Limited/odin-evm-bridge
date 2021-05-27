// SPDX-License-Identifier: GPL-3.0
pragma solidity ^0.7.0;

import {Obi} from "./Obi.sol";

library ResultDecoder {
    using Obi for Obi.Data;

    struct Result {
        string counts;
    }

    function decodeResult(bytes memory _data) internal pure returns (Result memory) {
        Obi.Data memory data = Obi.from(_data);
        Result memory result;
        result.counts = data.decodeString();
        return result;
    }
}
