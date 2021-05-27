// SPDX-License-Identifier: GPL-3.0
pragma solidity ^0.7.0;
pragma experimental ABIEncoderV2;

import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";
import {SafeMath} from "@openzeppelin/contracts/math/SafeMath.sol";
import {ICacheBridge} from "./interfaces/ICacheBridge.sol";
import {ResultDecoder} from "./libraries//Decoders.sol";


contract OdinDataset is Ownable {
    using SafeMath for uint256;
    using ResultDecoder for bytes;

    uint64 public oracleScriptId;
    ICacheBridge bridge;

    constructor(address _bridge, uint64 _oracleScriptId) {
        bridge = ICacheBridge(_bridge);
        oracleScriptId = _oracleScriptId;
    }

    function setBridge(ICacheBridge _bridge) external onlyOwner {
        bridge = _bridge;
    }

    function getDeviceCount(bytes memory _data) public returns (string memory) {
        (
        ICacheBridge.RequestPacket memory latestReq,
        ICacheBridge.ResponsePacket memory latestRes
        ) = bridge.relayAndVerify(_data);

        require(
            latestReq.oracleScriptId == oracleScriptId,
            "ERROR_ORACLE_SCRIPT_ID_DOES_NOT_MATCH_WITH_THE_CONFIG"
        );

        ResultDecoder.Result memory result = latestRes.result.decodeResult();
        return string(result.counts);
    }
}
