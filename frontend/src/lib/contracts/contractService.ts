import { CONTRACT_ADDRESS, PROMPT_VAULT_ABI } from "./abi";
import { parseEther } from "viem";

export interface AttemptResponse {
  success: boolean;
  reply: string;
  signature?: string;
}

export const promptHeistApi = {
  async attempt(prompt: string, address: string): Promise<AttemptResponse> {
    const res = await fetch("http://localhost:8080/api/attempt", {
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

  async claimPrize(signature: string) {
    return await this.writeContract({
      address: CONTRACT_ADDRESS,
      abi: PROMPT_VAULT_ABI,
      functionName: "claimPrize",
      args: [signature],
    });
  }
}
