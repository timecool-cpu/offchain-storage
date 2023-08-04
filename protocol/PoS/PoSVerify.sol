// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

interface IDataMarketplace {
    function registerStorageProvider(address addr, uint256 availableSpace, uint256 pricePerGBPerMonth) external;
}

contract Verifier {
    bytes public pk;   // public key to verify the proof
    uint256 public beta; // number of challenges needed
    bytes32 public root; // root hash
    IDataMarketplace public market;
    uint256 public availableSpace;
    uint256 public pricePerGBPerMonth;

    struct Graph {
        bytes pk;
        uint256 index;
        uint256 log2;
        uint256 pow2;
        uint256 size;
    }

    Graph public graph;
    uint256 public index; // index of the graphy in the family
    uint256 public size;
    uint256 public pow2;
    uint256 public log2;

    event Log_uint256(string message, uint256 value);
    event Log_uint256s(string message, uint256[] value);
    event Log_bytes32(string message, bytes32 value);
    event Log_bytes(string message, bytes value);
    event Log_bytess(string message, bytes[] value);

    constructor(bytes memory _pk, uint256 _index, uint256 _beta, bytes32 _root, uint256 _availableSpace, uint256 _pricePerGBPerMonth) public {
        availableSpace = _availableSpace;
        pricePerGBPerMonth = _pricePerGBPerMonth;
        // emit Log_uint256("Function started", 0);
        pk = _pk;
        beta = _beta;
        root = _root;

        size = numXi(_index);
        log2 = getlog2(size) + 1;
        pow2 = 2 ** log2;
        if (2 ** (log2 - 1) == size) {
            log2--;
            pow2 = 2 ** log2;
        }

        // initialize the graph
        Graph memory newGraph = Graph({
            pk: _pk,
            index: _index,
            log2: log2,
            pow2: pow2,
            size: size
        });

        graph = newGraph;

        index = _index;
    }


    function setAddress(address add) external {
        market = IDataMarketplace(add);
    }

    function verifySpace(
        uint256[] memory challenges,
        bytes[] memory hashes,
        bytes[][] memory parents,
        bytes[][] memory proofs,
        bytes[][][] memory pProofs
    ) public returns(bool) {
        for (uint256 i = 0; i < challenges.length; i++) {
            bool flag = true;
            // if(i == 0) emit Log_uint256("print pow2", pow2); //
            bytes32 buf = putVarint(challenges[i] + pow2);
            // if(i == 0) emit Log_uint256("print challenges[i]", challenges[i]); //
            // if(i == 0) emit Log_bytes32("print buf", buf); //
            bytes memory val = abi.encodePacked(pk, buf);
            for (uint256 j = 0; j < parents[i].length; j++) {
                val = abi.encodePacked(val, parents[i][j]);
            }
            // if(i == 0) emit Log_bytes("print val", val); //
            // bytes32 exp = keccak256(val);
            // if(i == 0) emit Log_bytes32("print exp", exp); //打印exp
            // for (uint256 j = 0; j < 32; j++) {
            //     if (exp[j] != hashes[i][j]) {
            //         return false;
            //     }
            // }

            if (!verify(challenges[i], hashes[i], proofs[i],flag)) {
                if(flag) flag = false;
                return false;
            }

            uint256[] memory ps = getParents(challenges[i], index);
            // emit Log_uint256s("print ps",ps);
            for (uint256 j = 0; j < ps.length; j++) {
                if (!verify(ps[j], parents[i][j], pProofs[i][j],flag)) {
                    return false;
                }
            }
        }
        market.registerStorageProvider(msg.sender, availableSpace,pricePerGBPerMonth);
        return true;
    }

    function verify(
        uint256 node,
        bytes memory hash,
        bytes[] memory proof,bool flag
    ) public view returns(bool) {
        // if(flag) {
        //     emit Log_uint256("node:",node);
        //     emit Log_bytes("hash", hash);
        //     emit Log_bytess("proof", proof);
        // }
        bytes memory curHash = hash;
        uint256 counter = 0;
        for (uint256 i = node + pow2; i > 1; i /= 2) {
            bytes memory val;
            if (i % 2 == 0) {
                val = abi.encodePacked(curHash, proof[counter]);
            } else {
                val = abi.encodePacked(proof[counter], curHash);
            }
            // emit Log_bytes("val", val);
            bytes32 hash = sha256(val);
            curHash = abi.encodePacked(hash);
            // emit Log_bytes("curHash",curHash);
            counter++;
        }
        for (uint256 i = 0; i < 32; i++) {
            // emit Log_bytes32("root",root);
            // emit Log_bytes("curHash",curHash);
            if (root[i] != curHash[i]) {
                return false;
            }
        }
        // emit Log_uint256("Yes!!!",1);
        return root.length == curHash.length;
    }


    // Determine parent nodes
    function getParents(uint256 node, uint256 index) public view returns(uint256[] memory){
        // emit Log_uint256("print node",node);
        // emit Log_uint256("print index",index);
        if(node < 2**index) {
            return new uint256[](0);
        }

        (uint256 offset0, uint256 offset1) = getGraph(node, index);
        // emit Log_uint256("print offset0",offset0);
        // emit Log_uint256("print offset1",offset1);

        uint256[] memory tempRes = new uint256[](2); // 创建一个临时数组用于存储结果
        uint256 count = 0;
        if(offset0 != 0) {
            tempRes[count] = node-offset0;
            count++;
        }
        if(offset1 != 0) {
            tempRes[count] = node-offset1;
            count++;
        }

        // 创建一个新的数组用于返回结果
        uint256[] memory res = new uint256[](count);
        for(uint256 i = 0; i < count; i++) {
            res[i] = tempRes[i];
        }
        // emit Log_uint256s("print res",res);
        return res;
    }

    // Separate the computeConstants function into multiple helper functions
    function computeBasic(uint256 index) public view returns(uint256, uint256, uint256) {
        uint256 pow2index = 2 ** index;
        uint256 pow2index_1 = 2 ** (index-1);
        uint256 sources = pow2index;

        return (pow2index, pow2index_1, sources);
    }

    function computeButterfly(uint256 index, uint256 sources) public view returns(uint256, uint256) {
        uint256 firstButter = sources + numButterfly(index-1);
        uint256 firstXi = firstButter + numXi(index-1);

        return (firstButter, firstXi);
    }

    function computeSink(uint256 index, uint256 firstXi) public view returns(uint256, uint256, uint256) {
        uint256 secondXi = firstXi + numXi(index-1);
        uint256 secondButter = secondXi + numButterfly(index-1);
        uint256 sinks = secondButter + 2 ** index;

        return (secondXi, secondButter, sinks);
    }

    // Assuming that numButterfly and numXi are other functions in the contract
    function getGraph(uint256 node, uint256 index) public view returns(uint256, uint256) {
        if (index == 1) {
            if (node < 2) {
                return (2, 0);
            } else if (node == 2) {
                return (1, 2);
            } else if (node == 3) {
                return (3, 2);
            }
        }

        (uint256 pow2index, uint256 pow2index_1, uint256 sources) = computeBasic(index);
        (uint256 firstButter, uint256 firstXi) = computeButterfly(index, sources);
        (uint256 secondXi, uint256 secondButter, uint256 sinks) = computeSink(index, firstXi);

        if (node < sources) {
            return (pow2index, 0);
        } else if (node >= sources && node < firstButter) {
            if (node < sources + pow2index_1) {
                return (pow2index, pow2index_1);
            } else {
                (uint256 parent0, uint256 parent1) = butterflyParents(sources, node, index);
                return (node - parent0, node - parent1);
            }
        } else if (node >= firstButter && node < firstXi) {
            node = node - firstButter;
            return getGraph(node, index-1);
        } else if (node >= firstXi && node < secondXi) {
            node = node - firstXi;
            return getGraph(node, index-1);
        } else if (node >= secondXi && node < secondButter) {
            if (node < secondXi + pow2index_1) {
                return (pow2index_1, 0);
            } else {
                (uint256 parent0, uint256 parent1) = butterflyParents(secondXi, node, index);
                return (node - parent0, node - parent1);
            }
        } else if (node >= secondButter && node < sinks) {
            uint256 offset = (node - secondButter) % pow2index_1;
            uint256 parent1 = sinks - numXi(index) + offset;
            if (offset + secondButter == node) {
                return (pow2index_1, node - parent1);
            } else {
                uint256 nodeMinusParent1 = node - parent1;  // Calculate this first
                return (pow2index, nodeMinusParent1 - pow2index_1);
            }
        } else {
            return (0, 0);
        }
    }


    function butterflyParents(uint256 begin, uint256 node, uint256 index) public pure returns(uint256, uint256) {
        uint256 pow2index_1 = 2 ** (index - 1);
        uint256 level = (node - begin) / pow2index_1;
        uint256 prev;
        uint256 shift = (index - 1) > level ? (index - 1) - level : level - (index - 1);
        uint256 i = (node - begin) % pow2index_1;
        if ((i >> shift) & 1 == 0) {
            prev = i + (1 << shift);
        } else {
            prev = i - (1 << shift);
        }
        uint256 parent0 = begin + (level - 1) * pow2index_1 + prev;
        uint256 parent1 = node - pow2index_1;
        return (parent0, parent1);
    }

    function numButterfly(uint256 index) public pure returns (uint256) {
        return 2 * (2 ** index) * index;
    }


    function numXi(uint256 index) private pure returns (uint256) {
        return (2 ** index) * (index + 1) * index;
    }

    function getlog2(uint256 x) private pure returns (uint256) {
        uint256 r = 0;
        while (x > 1) {
            x >>= 1;
            r++;
        }
        return r;
    }


    function selectChallenges(bytes32 seed) public view returns (uint256[] memory) {
        uint256[] memory challenges = new uint256[](beta * log2);
        for (uint256 i = 0; i < challenges.length; i++) {
            bytes32 hash = keccak256(abi.encodePacked(seed, i));
            challenges[i] = uint256(hash) % size;  // Assuming `size` is the range of the random index
        }
        return challenges;
    }

    function putVarint(uint256 x) public pure returns (bytes32) {
        uint256 ux = uint256(x) << 1;
        if (x < 0) {
            ux = ~ux;
        }
        return putUvarint(ux);
    }

    function putUvarint(uint256 x) public pure returns (bytes32) {
        bytes memory buf = new bytes(32);
        uint256 i = 0;
        while (x >= 0x80) {
            buf[i] = bytes1(uint8(x) | 0x80);
            x >>= 7;
            i++;
        }
        buf[i] = bytes1(uint8(x));

        bytes32 result;
        for (i = 0; i < 32; i++) {
            result |= bytes32(buf[i] & 0xFF) >> (i * 8);
        }
        return result;
    }

}

