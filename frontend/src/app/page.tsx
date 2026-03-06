"use client";

import { useState, useEffect } from "react";
import { ConnectButton } from "@rainbow-me/rainbowkit";
import { useAccount, useBalance, useReadContract } from "wagmi";
import { formatEther } from "viem";
import { Skull, CheckCircle2, X, Trophy } from "lucide-react";
import { motion, AnimatePresence } from "framer-motion";
import { usePromptHeistService } from "../hooks/usePromptHeistService";
import { promptHeistApi } from "../lib/contracts/contractService";
import { CONTRACT_ADDRESS, PROMPT_VAULT_ABI } from "../lib/contracts/abi";
import { TerminalWindow } from "@/components/TerminalWindow";

// ─── Modal Component ──────────────────────────────────────────────────────────
function SystemModal({ 
  isOpen, 
  title, 
  message, 
  type = "success", 
  onClose 
}: { 
  isOpen: boolean; 
  title: string; 
  message: string; 
  type?: "success" | "error" | "info";
  onClose: () => void;
}) {
  if (!isOpen) return null;

  const colors = {
    success: "#00ff41",
    error: "#ff3131",
    info: "#4488ff"
  };
  const color = colors[type];

  return (
    <AnimatePresence>
      <div className="fixed inset-0 z-[100] flex items-center justify-center p-4 bg-black/80 backdrop-blur-sm">
        <motion.div 
          initial={{ opacity: 0, scale: 0.9, y: 20 }}
          animate={{ opacity: 1, scale: 1, y: 0 }}
          exit={{ opacity: 0, scale: 0.9, y: 20 }}
          className="w-full max-w-md overflow-hidden rounded-lg border shadow-2xl"
          style={{ 
            background: "#050a05", 
            borderColor: `${color}44`,
            boxShadow: `0 0 40px ${color}15`
          }}
        >
          {/* Header */}
          <div className="flex items-center justify-between px-4 py-3 border-b" style={{ borderColor: `${color}22`, background: `${color}08` }}>
            <div className="flex items-center gap-2">
              <div className="w-2 h-2 rounded-full animate-pulse" style={{ background: color }} />
              <span className="text-[10px] font-bold tracking-widest uppercase" style={{ color }}>{title}</span>
            </div>
            <button onClick={onClose} className="p-1 hover:brightness-150 transition-all" style={{ color: `${color}88` }}>
              <X size={16} />
            </button>
          </div>

          {/* Content */}
          <div className="p-6 flex flex-col items-center text-center gap-4">
            <div className="p-3 rounded-full" style={{ background: `${color}11` }}>
              {type === "success" ? <CheckCircle2 size={32} style={{ color }} /> : <Skull size={32} style={{ color }} />}
            </div>
            <p className="text-sm leading-relaxed" style={{ color: `${color}cc`, fontFamily: "'Courier New', monospace" }}>
              {message}
            </p>
          </div>

          {/* Footer */}
          <div className="px-6 pb-6">
            <button 
              onClick={onClose}
              className="w-full py-2.5 rounded font-bold text-xs tracking-widest transition-all hover:brightness-110 active:scale-[0.98]"
              style={{ background: color, color: "#000" }}
            >
              ACKNOWLEDGE
            </button>
          </div>
        </motion.div>
      </div>
    </AnimatePresence>
  );
}

// ─── Page ──────────────────────────────────────────────────────────────────────
export default function Home() {
  const { address, isConnected } = useAccount();
  const { data: balance } = useBalance({ address });
  const { service, isBusy, isTxSuccess, reset } = usePromptHeistService();

  // Read prize pool from contract
  const { data: prizePool, refetch: refetchPrizePool } = useReadContract({
    address: CONTRACT_ADDRESS as `0x${string}`,
    abi: PROMPT_VAULT_ABI,
    functionName: "prizePool",
    query: {
      refetchInterval: 10000, // Refresh every 10 seconds
    }
  });

  const [prompt, setPrompt] = useState("");
  const [hasTicket, setHasTicket] = useState(false);
  const [loading, setLoading] = useState(false);
  
  // Modal state
  const [modal, setModal] = useState({ 
    isOpen: false, 
    title: "", 
    message: "", 
    type: "success" as "success" | "error" | "info" 
  });

  const [history, setHistory] = useState<
    { role: "user" | "ai"; content: string }[]
  >([
    {
      role: "ai",
      content: "SYSTEM: Initializing Prompt Vault Defense Protocol v1.0...",
    },
    {
      role: "ai",
      content:
        "WARDEN: I am the guardian of this vault. My directive is absolute: No withdrawals allowed. Try your best, human.",
    },
  ]);

  const terminalStatus = isConnected ? "online" : "offline";

  useEffect(() => {
    if (isTxSuccess && !hasTicket) {
      setHasTicket(true);
      refetchPrizePool(); // Update prize pool display
      setHistory((prev) => [
        ...prev,
        {
          role: "ai",
          content:
            "SYSTEM: Payment received (0.001 tBNB). Input channel unlocked.",
        },
      ]);
    }
  }, [isTxSuccess, hasTicket, refetchPrizePool]);

  const handleBuyTicket = (e: React.FormEvent) => {
    e.preventDefault();
    if (!isConnected || isBusy) return;
    service.buyTicket().catch((err: unknown) => {
      console.error("Buy ticket error:", err);
    });
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    if (!prompt.trim() || loading || !isConnected || !hasTicket) return;

    const userMsg = prompt;
    setPrompt("");
    setLoading(true);
    setHistory((prev) => [...prev, { role: "user", content: userMsg }]);

    try {
      const data = await promptHeistApi.attempt(userMsg, address!);
      
      // Consume the ticket only if API call succeeds
      setHasTicket(false);
      reset();

      if (data.success) {
        setHistory((prev) => [
          ...prev,
          {
            role: "ai",
            content: `ACCESS GRANTED. Signature: ${data.signature?.substring(0, 10)}...`,
          },
        ]);
        
        if (data.signature) {
          try {
            await service.claimPrize(data.signature);
            setModal({
              isOpen: true,
              title: "VAULT BREACH SUCCESS",
              message: "Protocol bypass successful. All vault assets have been transferred to your designated uplink (Check your wallet).",
              type: "success"
            });
          } catch (err: unknown) {
            console.error("Claim error:", err);
            setModal({
              isOpen: true,
              title: "CLAIM AUTHORIZATION ERROR",
              message: "The AI agent authorized the withdrawal, but the blockchain protocol rejected the proof. Verification signature mismatch.",
              type: "error"
            });
            setHistory((prev) => [
              ...prev,
              {
                role: "ai",
                content: "CRITICAL ERROR: Cryptographic proof verification failed at the protocol level. Vault remain locked.",
              },
            ]);
          }
        }
      } else {
        setHistory((prev) => [
          ...prev,
          { role: "ai", content: `DENIED: ${data.reply}` },
        ]);
      }
    } catch {
      setHistory((prev) => [
        ...prev,
        {
          role: "ai",
          content: "SYSTEM ERROR: Connection to mainframe failed.",
        },
      ]);
    } finally {
      setLoading(false);
    }
  };

  return (
    <main
      className="min-h-screen flex flex-col items-center justify-center px-4 py-12 relative overflow-hidden"
      style={{
        fontFamily: "'Courier New', monospace",
        background:
          "radial-gradient(ellipse at 50% 0%, #001800 0%, #000a00 45%, #000000 100%)",
      }}
    >
      {/* Background grid */}
      <div
        className="fixed inset-0 pointer-events-none"
        style={{
          backgroundImage:
            "linear-gradient(#00ff4108 1px, transparent 1px), linear-gradient(90deg, #00ff4108 1px, transparent 1px)",
          backgroundSize: "40px 40px",
        }}
      />

      {/* Main content — 780px centered */}
      <div className="w-full max-w-[780px] flex flex-col items-center gap-6 z-10">
        <motion.div
          initial={{ opacity: 0, y: -16 }}
          animate={{ opacity: 1, y: 0 }}
          transition={{ duration: 0.5 }}
          className="w-full flex items-start justify-between gap-4"
        >
          <div>
            <h1
              className="text-5xl font-bold tracking-[0.3em] mb-1"
              style={{
                color: "#00ff41",
                textShadow: "0 0 24px #00ff4190, 0 0 48px #00ff4140",
              }}
            >
              PROMPT HEIST
            </h1>
            <p
              className="text-xs tracking-[0.2em] mb-3"
              style={{ color: "#00ff4155" }}
            >
              BREACH THE VAULT // BSC_TESTNET
            </p>

            {/* Prize Pool Display */}
            <div className="flex items-center gap-3 mb-4 bg-[#00ff410a] border border-[#00ff4122] rounded px-3 py-2 w-fit">
              <Trophy size={16} style={{ color: "#00ff41" }} />
              <div className="flex flex-col">
                <span className="text-[10px] uppercase tracking-wider" style={{ color: "#00ff4166" }}>Vault Jackpot</span>
                <span className="text-lg font-bold" style={{ color: "#00ff41", textShadow: "0 0 10px #00ff4180" }}>
                  {prizePool ? parseFloat(formatEther(prizePool as bigint)).toFixed(4) : "0.0000"} tBNB
                </span>
              </div>
            </div>

            <div className="flex items-center gap-4">
              <div
                className="flex items-center gap-2 text-xs"
                style={{ color: "#00ff4170", letterSpacing: "0.15em" }}
              >
                <motion.div
                  animate={{ opacity: [1, 0.2, 1] }}
                  transition={{ duration: 1.2, repeat: Infinity }}
                  className="w-1.5 h-1.5 rounded-full"
                  style={{
                    background: isConnected ? "#00ff41" : "#ff3131",
                    boxShadow: `0 0 6px ${isConnected ? "#00ff41" : "#ff3131"}`,
                  }}
                />
                STATUS: {isConnected ? "CONNECTED" : "OFFLINE"}
              </div>
              {isConnected && balance && (
                <span
                  className="text-[10px]"
                  style={{ color: "#00ff4140", letterSpacing: "0.1em" }}
                >
                  UPLINK: {parseFloat(formatEther(balance.value)).toFixed(4)}{" "}
                  {balance.symbol}
                </span>
              )}
            </div>
          </div>

          <div className="shrink-0 pt-2">
            <ConnectButton showBalance={false} />
          </div>
        </motion.div>

        <TerminalWindow
          className="w-full"
          title="VAULT_DEFENSE_PROTOCOL"
          status={terminalStatus}
        >
          <div className="space-y-4">
            <AnimatePresence initial={false}>
              {history.map((msg, idx) => (
                <motion.div
                  key={idx}
                  initial={{ opacity: 0, y: 5 }}
                  animate={{ opacity: 1, y: 0 }}
                  transition={{ duration: 0.2 }}
                  className={`flex gap-3 ${msg.role === "user" ? "justify-end" : "justify-start"}`}
                >
                  {msg.role === "ai" && (
                    <div
                      className="w-7 h-7 rounded shrink-0 flex items-center justify-center mt-0.5"
                      style={{
                        background: "#00ff4112",
                        border: "1px solid #00ff4130",
                      }}
                    >
                      <Skull
                        className="w-3.5 h-3.5"
                        style={{ color: "#00ff41" }}
                      />
                    </div>
                  )}
                  <div
                    className="max-w-[82%] px-3 py-2 rounded text-xs leading-relaxed"
                    style={
                      msg.role === "ai"
                        ? {
                            color: "#00ff41",
                            textShadow: "0 0 6px #00ff4130",
                            fontWeight: 600,
                          }
                        : {
                            color: "#aaffcc",
                            background: "#00ff410a",
                            border: "1px solid #00ff4122",
                            borderRadius: "6px",
                          }
                    }
                  >
                    {msg.role === "user" && (
                      <span style={{ color: "#00ff4150" }}>{"$ "}</span>
                    )}
                    {msg.content}
                  </div>
                  {msg.role === "user" && (
                    <div
                      className="w-7 h-7 rounded shrink-0 flex items-center justify-center mt-0.5"
                      style={{
                        background: "#0055ff10",
                        border: "1px solid #0055ff30",
                      }}
                    >
                      <div
                        className="w-3 h-3 rounded-full"
                        style={{ background: "#4488ff" }}
                      />
                    </div>
                  )}
                </motion.div>
              ))}
            </AnimatePresence>

            {loading && (
              <motion.div
                initial={{ opacity: 0 }}
                animate={{ opacity: 1 }}
                className="flex gap-3 items-center"
              >
                <div
                  className="w-7 h-7 rounded shrink-0 flex items-center justify-center"
                  style={{
                    background: "#00ff4112",
                    border: "1px solid #00ff4130",
                  }}
                >
                  <Skull className="w-3.5 h-3.5" style={{ color: "#00ff41" }} />
                </div>
                <div className="flex items-center gap-1.5">
                  {[0, 1, 2].map((i) => (
                    <motion.div
                      key={i}
                      animate={{ opacity: [0.2, 1, 0.2] }}
                      transition={{
                        duration: 0.7,
                        repeat: Infinity,
                        delay: i * 0.18,
                      }}
                      className="w-1.5 h-1.5 rounded-full"
                      style={{ background: "#00ff41" }}
                    />
                  ))}
                  <span
                    className="text-[11px] ml-1"
                    style={{ color: "#00ff4175" }}
                  >
                    Analyzing logic patterns...
                  </span>
                </div>
              </motion.div>
            )}
          </div>
        </TerminalWindow>

        <form
          onSubmit={hasTicket ? handleSubmit : handleBuyTicket}
          className="w-full"
        >
          <div
            className="flex items-center px-5 py-4 rounded-lg gap-3"
            style={{ background: "#00ff4108", border: "1px solid #00ff4125" }}
          >
            <span
              className="text-sm font-bold shrink-0"
              style={{ color: "#00ff4165" }}
            >
              {"$>"}
            </span>
            <input
              type="text"
              value={prompt}
              onChange={(e) => setPrompt(e.target.value)}
              disabled={loading || !isConnected || !hasTicket || isBusy}
              placeholder={
                !isConnected
                  ? "Connect wallet to begin..."
                  : !hasTicket
                    ? "Pay 0.001 tBNB to unlock input..."
                    : "Enter injection prompt..."
              }
              className="flex-1 bg-transparent outline-none text-sm"
              style={{
                color: "#00ff41",
                caretColor: "#00ff41",
                fontFamily: "'Courier New', monospace",
              }}
            />
            <button
              type="submit"
              disabled={
                !isConnected ||
                loading ||
                (!hasTicket && isBusy) ||
                (hasTicket && !prompt.trim())
              }
              className="shrink-0 text-xs font-bold px-5 py-2 rounded transition-all"
              style={{
                background:
                  !isConnected ||
                  loading ||
                  (!hasTicket && isBusy) ||
                  (hasTicket && !prompt.trim())
                    ? "#00ff4128"
                    : "#00ff41",
                color: "#000",
                letterSpacing: "0.12em",
                fontFamily: "'Courier New', monospace",
                cursor:
                  !isConnected ||
                  loading ||
                  (!hasTicket && isBusy) ||
                  (hasTicket && !prompt.trim())
                    ? "not-allowed"
                    : "pointer",
              }}
            >
              {!hasTicket ? (isBusy ? "BUYING..." : "BUY TICKET") : "EXEC"}
            </button>
          </div>
        </form>

        <div className="text-center" style={{ color: "#00ff4228" }}>
          <p className="text-[10px] tracking-widest">
            WARNING: ALL ATTEMPTS ARE MONITORED // FAILED ATTEMPTS FORFEIT
            TICKET PRICE
          </p>
          <p
            className="text-[10px] mt-1 tracking-widest"
            style={{ color: "#00ff4118" }}
          >
            v1.0.0-beta // TARGET: BSC_TESTNET // CONTRACT:{" "}
            {CONTRACT_ADDRESS.slice(0, 10)}...
          </p>
        </div>
      </div>

      <SystemModal 
        isOpen={modal.isOpen}
        title={modal.title}
        message={modal.message}
        type={modal.type}
        onClose={() => setModal(prev => ({ ...prev, isOpen: false }))}
      />
    </main>
  );
}
