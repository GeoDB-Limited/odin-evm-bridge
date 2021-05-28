// SPDX-License-Identifier: GPL-3.0
pragma solidity ^0.7.0;
pragma experimental ABIEncoderV2;

import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";
import {SafeMath} from "@openzeppelin/contracts/math/SafeMath.sol";
import {ICacheBridge} from "./interfaces/ICacheBridge.sol";
import {IOdinDataset} from "./interfaces/IOdinDataset.sol";
import {ResultDecoder} from "./libraries/ResultDecoder.sol";


contract OdinDataset is IOdinDataset, Ownable {
    using SafeMath for uint256;
    using ResultDecoder for bytes;

    struct ModelData {
        uint64 oracleScriptId;
        uint8 calldataId;
        uint8 modelId;
    }

    ICacheBridge bridge;

    mapping(string => ModelData) dataFromModel;
    bytes[] calldataArray = new bytes[](1);

    constructor(address _bridge) {
        bridge = ICacheBridge(_bridge);

        // ["iPhone 8", "iPhone X", "iPhone 11", "iPhone 12"]
        calldataArray[0] = hex"00000004000000086950686f6e652038000000086950686f6e652058000000096950686f6e65203131000000096950686f6e65203132";

        // Models
        dataFromModel["iPhone 8"] = ModelData({
            oracleScriptId: 2,
            calldataId: 0,
            modelId: 0
        });
        dataFromModel["iPhone X"] = ModelData({
            oracleScriptId: 2,
            calldataId: 0,
            modelId: 1
        });
        dataFromModel["iPhone 11"] = ModelData({
            oracleScriptId: 2,
            calldataId: 0,
            modelId: 2
        });
        dataFromModel["iPhone 12"] = ModelData({
            oracleScriptId: 2,
            calldataId: 0,
            modelId: 3
        });
    }

    function setBridge(ICacheBridge _bridge) external onlyOwner {
        bridge = _bridge;
    }

    function getReferenceData(string[] memory _models) external override view returns (ReferenceData[] memory) {
        ReferenceData[] memory results = new ReferenceData[](_models.length);
        for (uint256 i = 0; i < _models.length; i++) {
            (uint64 count, uint64 lastUpdated) = getModelCount(_models[i]);

            results[i].count = uint256(count);
            results[i].lastUpdated = uint256(lastUpdated);
        }
        return results;
    }

    function getModelCount(string memory _model) internal view returns (uint64, uint64) {
        ICacheBridge.RequestPacket memory req;
        req.clientId = "1";
        req.askCount = 1;
        req.minCount = 1;

        ModelData storage data = dataFromModel[_model];
        req.oracleScriptId = data.oracleScriptId;
        req.params = calldataArray[data.calldataId];

        ICacheBridge.ResponsePacket memory packet = bridge.getLatestResponse(req);
        ResultDecoder.Result memory result = packet.result.decodeResult();
        return (result.count[data.modelId], packet.resolveTime);
    }
}
