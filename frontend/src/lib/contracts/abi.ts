export const CONTRACT_ADDRESS = "0xD9930219566eED251e6757FEAfA35D768d20c9c5";

export const PROMPT_VAULT_ABI = [
  {
    inputs: [],
    name: "play",
    outputs: [],
    stateMutability: "payable",
    type: "function",
  },
  {
    inputs: [
      {
        internalType: "bytes",
        name: "signature",
        type: "bytes",
      },
    ],
    name: "claimPrize",
    outputs: [],
    stateMutability: "nonReentrant",
    type: "function",
  },
  {
    inputs: [],
    name: "prizePool",
    outputs: [
      {
        internalType: "uint256",
        name: "",
        type: "uint256",
      },
    ],
    stateMutability: "view",
    type: "function",
  },
] as const;
