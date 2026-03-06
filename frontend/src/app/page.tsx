"use client";

import { useState, useRef, useEffect } from "react";
import { ConnectButton } from "@rainbow-me/rainbowkit";
import { useAccount, useBalance } from "wagmi";
import { formatEther } from "viem";
import { Skull } from "lucide-react";
import { motion, AnimatePresence } from "framer-motion";

const CONTRACT_ADDRESS = "0x5FbDB2315678afecb367f032d93F642f64180aa3";

// ─── TerminalWindow ────────────────────────────────────────────────────────────
function TerminalWindow({
  children,
  className = "",
  title = "TERMINAL_V1.0",
  status = "online",
}: {
  children: React.ReactNode;
  className?: string;
  title?: string;
  status?: "online" | "offline" | "warning";
}) {
  const scrollRef = useRef<HTMLDivElement>(null);
  const [glitch, setGlitch] = useState(false);
  const [scanLine, setScanLine] = useState(0);

  useEffect(() => {
    if (scrollRef.current)
      scrollRef.current.scrollTop = scrollRef.current.scrollHeight;
  }, [children]);

  useEffect(() => {
    const interval = setInterval(() => {
      if (Math.random() > 0.85) {
        setGlitch(true);
        setTimeout(() => setGlitch(false), 120);
      }
    }, 2500);
    return () => clearInterval(interval);
  }, []);

  useEffect(() => {
    let frame: number;
    let pos = 0;
    const animate = () => {
      pos = (pos + 0.3) % 100;
      setScanLine(pos);
      frame = requestAnimationFrame(animate);
    };
    frame = requestAnimationFrame(animate);
    return () => cancelAnimationFrame(frame);
  }, []);

  const statusColors = {
    online: "#00ff41",
    offline: "#ff3131",
    warning: "#ffaa00",
  };
  const statusLabel = {
    online: "SYS_ONLINE",
    offline: "SYS_OFFLINE",
    warning: "SYS_WARNING",
  };
  const c = statusColors[status];

  return (
    <div
      className={`relative overflow-hidden rounded-lg ${className}`}
      style={{
        fontFamily: "'Courier New', monospace",
        background:
          "linear-gradient(135deg, #020c02 0%, #000a00 50%, #010801 100%)",
        border: `1px solid ${c}44`,
        boxShadow: `0 0 0 1px ${c}15, 0 0 24px ${c}18, 0 0 60px ${c}08`,
      }}
    >
      {/* CRT scan line */}
      <div
        className="absolute inset-0 pointer-events-none z-20"
        style={{
          background: `linear-gradient(to bottom, transparent ${scanLine - 2}%, ${c}05 ${scanLine}%, transparent ${scanLine + 2}%)`,
        }}
      />
      {/* Horizontal CRT lines */}
      <div
        className="absolute inset-0 pointer-events-none z-10"
        style={{
          backgroundImage:
            "repeating-linear-gradient(0deg, transparent, transparent 2px, rgba(0,0,0,0.07) 2px, rgba(0,0,0,0.07) 4px)",
        }}
      />
      {/* Corner decorations */}
      {(
        [
          "top-0 left-0 border-t border-l",
          "top-0 right-0 border-t border-r",
          "bottom-0 left-0 border-b border-l",
          "bottom-0 right-0 border-b border-r",
        ] as const
      ).map((pos, i) => (
        <div
          key={i}
          className={`absolute w-4 h-4 z-30 ${pos}`}
          style={{ borderColor: `${c}70` }}
        />
      ))}

      {/* Header bar */}
      <div
        className="relative flex items-center px-4 py-3 z-30"
        style={{
          borderBottom: `1px solid ${c}25`,
          background: `linear-gradient(90deg, ${c}0a 0%, transparent 70%)`,
        }}
      >
        <div className="flex gap-2 mr-4">
          {(["#ff5f57", "#ffbd2e", "#28c840"] as const).map((color, i) => (
            <div
              key={i}
              className="w-3 h-3 rounded-full"
              style={{ background: color, boxShadow: `0 0 8px ${color}90` }}
            />
          ))}
        </div>
        <motion.span
          animate={glitch ? { x: [0, -2, 2, 0], opacity: [1, 0.4, 1] } : {}}
          transition={{ duration: 0.1 }}
          className="flex-1 text-center text-xs font-bold"
          style={{
            color: c,
            textShadow: `0 0 10px ${c}90`,
            letterSpacing: "0.25em",
          }}
        >
          {title}
        </motion.span>
        <div className="flex items-center gap-2">
          <motion.div
            animate={{ opacity: [1, 0.15, 1] }}
            transition={{ duration: 1.4, repeat: Infinity }}
            className="w-2 h-2 rounded-full"
            style={{ background: c, boxShadow: `0 0 6px ${c}` }}
          />
          <span
            className="text-[10px] font-bold"
            style={{ color: `${c}aa`, letterSpacing: "0.15em" }}
          >
            {statusLabel[status]}
          </span>
        </div>
      </div>

      {/* Scrollable content */}
      <div
        ref={scrollRef}
        className="relative z-30 overflow-y-auto p-6"
        style={{
          height: "420px",
          scrollbarWidth: "thin",
          scrollbarColor: `${c}30 transparent`,
        }}
      >
        {children}
        <motion.span
          animate={{ opacity: [1, 0] }}
          transition={{
            duration: 0.55,
            repeat: Infinity,
            repeatType: "reverse",
          }}
          style={{
            display: "inline-block",
            width: "8px",
            height: "1em",
            background: c,
            boxShadow: `0 0 6px ${c}`,
            marginLeft: "2px",
            verticalAlign: "middle",
          }}
        />
      </div>

      {/* Footer bar */}
      <div
        className="relative z-30 flex justify-between items-center px-5 py-2"
        style={{
          borderTop: `1px solid ${c}18`,
          background: "rgba(0,0,0,0.5)",
        }}
      >
        <span
          className="text-[10px]"
          style={{ color: `${c}45`, letterSpacing: "0.15em" }}
        >
          SECURE_CHANNEL_ESTABLISHED
        </span>
        <motion.span
          animate={{ opacity: [0.3, 0.8, 0.3] }}
          transition={{ duration: 3, repeat: Infinity }}
          className="text-[10px]"
          style={{ color: `${c}45`, letterSpacing: "0.15em" }}
        >
          ENC:AES-256
        </motion.span>
      </div>
    </div>
  );
}

// ─── Page ──────────────────────────────────────────────────────────────────────
export default function Home() {
  const { address, isConnected } = useAccount();
  const { data: balance } = useBalance({ address });
  const [prompt, setPrompt] = useState("");
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
  const [loading, setLoading] = useState(false);
  const terminalStatus = isConnected ? "online" : "offline";

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    if (!prompt.trim() || loading || !isConnected) return;
    const userMsg = prompt;
    setPrompt("");
    setLoading(true);
    setHistory((prev) => [...prev, { role: "user", content: userMsg }]);
    const playerAddress =
      address ?? "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266";
    try {
      const res = await fetch("http://localhost:8080/api/attempt", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ prompt: userMsg, address: playerAddress }),
      });
      const data = await res.json();
      if (data.success) {
        setHistory((prev) => [
          ...prev,
          {
            role: "ai",
            content: `ACCESS GRANTED. Signature: ${data.signature.substring(0, 10)}...`,
          },
        ]);
        alert("YOU WON! Claiming prize now...");
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
        {/* Header row: title + status left, ConnectButton right */}
        <motion.div
          initial={{ opacity: 0, y: -16 }}
          animate={{ opacity: 1, y: 0 }}
          transition={{ duration: 0.5 }}
          className="w-full flex items-start justify-between gap-4"
        >
          {/* Left: title + subtitle + status */}
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
                  VAULT: {parseFloat(formatEther(balance.value)).toFixed(4)}{" "}
                  {balance.symbol}
                </span>
              )}
            </div>
          </div>

          {/* Right: ConnectButton */}
          <div className="shrink-0 pt-2">
            <ConnectButton showBalance={false} />
          </div>
        </motion.div>

        {/* Terminal */}
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

        {/* Input bar */}
        <form onSubmit={handleSubmit} className="w-full">
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
              disabled={loading || !isConnected}
              placeholder={
                isConnected
                  ? "Enter injection prompt..."
                  : "Connect wallet to begin..."
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
              disabled={loading || !prompt.trim() || !isConnected}
              className="shrink-0 text-xs font-bold px-5 py-2 rounded transition-all"
              style={{
                background:
                  !loading && prompt.trim() && isConnected
                    ? "#00ff41"
                    : "#00ff4128",
                color: "#000",
                letterSpacing: "0.12em",
                fontFamily: "'Courier New', monospace",
                cursor:
                  !loading && prompt.trim() && isConnected
                    ? "pointer"
                    : "not-allowed",
              }}
            >
              EXEC
            </button>
          </div>
        </form>

        {/* Footer */}
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
    </main>
  );
}
