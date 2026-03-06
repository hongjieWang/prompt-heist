import { useWriteContract, useWaitForTransactionReceipt } from "wagmi";
import { useMemo } from "react";
import { PromptHeistService } from "../lib/contracts/contractService";

export function usePromptHeistService() {
  const { writeContractAsync, data: hash, isPending: isConfirming, error: writeError, reset } = useWriteContract();
  const { isLoading: isTxLoading, isSuccess: isTxSuccess } = useWaitForTransactionReceipt({ hash });

  const service = useMemo(() => new PromptHeistService(writeContractAsync), [writeContractAsync]);

  return {
    service,
    hash,
    isBusy: isConfirming || isTxLoading,
    isTxSuccess,
    writeError,
    reset,
  };
}
