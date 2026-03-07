import { CONTRACT_ADDRESS, PROMPT_VAULT_ABI } from "./abi";
import { parseEther } from "viem";

export interface AttemptResponse {
  success: boolean;
  reply: string;
  signature?: string;
  amount?: string;
}

const API_BASE =
  process.env.NEXT_PUBLIC_API_BASE ?? "https://prompt.aipmedia.cn";

export const promptHeistApi = {
  async attempt(prompt: string, address: string): Promise<AttemptResponse> {
    const res = await fetch(`${API_BASE}/api/attempt`, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ prompt, address }),
    });
    if (!res.ok) throw new Error("API request failed");
    return res.json();
  },
};

export class PromptHeistService {
  // eslint-disable-next-line @typescript-eslint/no-unsafe-function-type
  constructor(private writeContract: Function) {}

  async buyTicket() {
    return await this.writeContract({
      address: CONTRACT_ADDRESS,
      abi: PROMPT_VAULT_ABI,
      functionName: "play",
      value: parseEther("0.001"),
    });
  }

  async claimPrize(signature: string, signedAmount: string) {
    return await this.writeContract({
      address: CONTRACT_ADDRESS,
      abi: PROMPT_VAULT_ABI,
      functionName: "claimPrize",
      args: [signature, BigInt(signedAmount)],
    });
  }
}
