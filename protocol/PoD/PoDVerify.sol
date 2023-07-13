// SPDX-License-Identifier: MIT
pragma solidity >=0.4.22 <0.9.0;

contract Verify {
    function verifyProof(
        bytes32 hash, // SHA-256 hash of the file
        uint8 v,
        bytes32 r,
        bytes32 s,
        address signer // expected signer
    )
    public pure returns(bool)
    {
        bytes32 message = prefixed(hash);
        return ecrecover(message, v, r, s) == signer;
    }

    function prefixed(bytes32 hash) internal pure returns (bytes32) {
        return keccak256(abi.encodePacked("\x19Ethereum Signed Message:\n32", hash));
    }
}
