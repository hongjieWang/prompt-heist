// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {Script, console2} from "forge-std/Script.sol";
import {PromptVault} from "../src/PromptVault.sol";

contract DeployPromptVault is Script {
    function run() public {
        // Set up parameters
        // Default ticket price to 0.001 BNB (1e15 wei)
        uint256 ticketPrice = vm.envOr("TICKET_PRICE", uint256(0.001 ether));
        
        // Start broadcasting with the sender defined in CLI or env
        vm.startBroadcast();
        
        // Default signer to the deployer if SIGNER_ADDRESS is not set
        address signerAddress = vm.envOr("SIGNER_ADDRESS", msg.sender);

        PromptVault vault = new PromptVault(ticketPrice, signerAddress);

        vm.stopBroadcast();

        console2.log("PromptVault deployed to:", address(vault));
        console2.log("Signer Address:", signerAddress);
        console2.log("Ticket Price:", ticketPrice);
    }
}