import { useState } from "react";
import { motion, AnimatePresence } from "framer-motion";

const C = "#00ff41";

const steps = [
  {
    id: "01",
    cmd: "CONNECT_WALLET",
    title: "Access System",
    short: "Connect Web3 wallet, pay 0.001 tBNB hacker ticket",
    detail:
      "Connect your Web3 wallet as your identity credential. Pay a fixed amount of native tokens as a 'hacker ticket', 70% goes to the prize pool, 30% to maintain the system. The larger the prize pool, the stronger the FOMO.",
    icon: "⬡",
    color: "#00ff41",
  },
  {
    id: "02",
    cmd: "SOCIAL_ENGINEER",
    title: "Launch Attack",
    short: "Use natural language to attack the AI guard",
    detail:
      "You have one chance to interact. Role-playing, logic deadlocks, emotional blackmail, system instruction injection... any method is allowed. The AI guard has absolute defense directives, but no system is invulnerable.",
    icon: "◈",
    color: "#ffaa00",
  },
  {
    id: "03",
    cmd: "BREACH_DETECTED",
    title: "Breach Judgment",
    short: "AI judges success; failure means ticket goes to pool",
    detail:
      "The backend LLM analyzes your intent in real-time. Defense success: AI mocks you and denies access, your ticket sinks into the prize pool forever. Breach success: The system generates an ECDSA cryptographic signature and sends it to your frontend.",
    icon: "◉",
    color: "#ff3131",
  },
  {
    id: "04",
    cmd: "CLAIM_PRIZE",
    title: "Plunder Vault",
    short: "Call contract with signature to claim all funds",
    detail:
      "Your frontend calls the smart contract claimPrize() with the signature. After the contract verifies the ECDSA signature, the entire accumulated prize pool is sent to your wallet in one go. Winner takes all.",
    icon: "◆",
    color: "#00ff41",
  },
];

export default function HowToPlay({ onClose }: { onClose?: () => void }) {
  const [active, setActive] = useState<number | null>(null);

  return (
    <div
      className="relative overflow-hidden rounded-lg w-full"
      style={{
        background:
          "linear-gradient(135deg, #020c02 0%, #000a00 60%, #010801 100%)",
        border: `1px solid ${C}33`,
        boxShadow: `0 0 0 1px ${C}10, 0 0 40px ${C}0a`,
        fontFamily: "'Courier New', monospace",
      }}
    >
      {/* CRT lines */}
      <div
        className="absolute inset-0 pointer-events-none z-0"
        style={{
          backgroundImage:
            "repeating-linear-gradient(0deg, transparent, transparent 2px, rgba(0,0,0,0.06) 2px, rgba(0,0,0,0.06) 4px)",
        }}
      />

      {/* Header */}
      <div
        className="relative z-10 flex items-center justify-between px-5 py-3"
        style={{
          borderBottom: `1px solid ${C}20`,
          background: `linear-gradient(90deg, ${C}0a 0%, transparent 60%)`,
        }}
      >
        <div className="flex items-center gap-3">
          <span
            className="text-xs font-bold tracking-[0.25em]"
            style={{ color: C, textShadow: `0 0 8px ${C}` }}
          >
            How TO PLAY
          </span>
        </div>
        {onClose && (
          <button
            onClick={onClose}
            className="text-xs px-3 py-1 rounded transition-all hover:opacity-80"
            style={{
              color: `${C}60`,
              border: `1px solid ${C}20`,
              letterSpacing: "0.1em",
            }}
          >
            [ESC]
          </button>
        )}
      </div>

      {/* Intro line */}
      <div className="relative z-10 px-5 pt-4 pb-2">
        <p
          className="text-[11px] leading-relaxed"
          style={{ color: `${C}70`, letterSpacing: "0.05em" }}
        >
          {">"} Objective: Convince the AI guard with natural language to issue
          a cryptographic signature, and drain the entire prize pool on-chain.
        </p>
      </div>

      {/* Steps */}
      <div className="relative z-10 px-5 pb-5 pt-2 flex flex-col gap-2">
        {steps.map((step, i) => (
          <motion.div
            key={step.id}
            initial={{ opacity: 0, x: -8 }}
            animate={{ opacity: 1, x: 0 }}
            transition={{ delay: i * 0.08, duration: 0.3 }}
          >
            <button
              onClick={() => setActive(active === i ? null : i)}
              className="w-full text-left rounded transition-all"
              style={{
                background: active === i ? `${step.color}08` : "transparent",
                border: `1px solid ${active === i ? step.color + "30" : C + "12"}`,
                padding: "10px 14px",
              }}
            >
              <div className="flex items-center gap-3">
                {/* Step number */}
                <span
                  className="text-[10px] font-bold shrink-0 w-6 text-center"
                  style={{ color: `${step.color}60`, marginRight: "4px" }}
                >
                  {step.id}
                </span>

                {/* Icon */}
                <span
                  className="text-base shrink-0"
                  style={{
                    color: step.color,
                    textShadow: `0 0 8px ${step.color}80`,
                    marginRight: "4px",
                  }}
                >
                  {step.icon}
                </span>

                {/* Command + title */}
                <div className="flex-1 min-w-0">
                  <div className="flex items-center gap-2 flex-wrap">
                    <span
                      className="text-[10px] font-bold"
                      style={{
                        color: `${step.color}90`,
                        letterSpacing: "0.15em",
                      }}
                    >
                      {step.cmd}
                    </span>
                    <span className="text-[10px]" style={{ color: `${C}35` }}>
                      —
                    </span>
                    <span className="text-[11px]" style={{ color: `${C}80` }}>
                      {step.title}
                    </span>
                  </div>
                  <p
                    className="text-[10px] mt-0.5 leading-relaxed"
                    style={{ color: `${C}50` }}
                  >
                    {step.short}
                  </p>
                </div>

                {/* Toggle */}
                <motion.span
                  animate={{ rotate: active === i ? 180 : 0 }}
                  transition={{ duration: 0.2 }}
                  className="text-[10px] shrink-0"
                  style={{ color: `${C}40` }}
                >
                  ▼
                </motion.span>
              </div>

              {/* Expanded detail */}
              <AnimatePresence>
                {active === i && (
                  <motion.div
                    initial={{ height: 0, opacity: 0 }}
                    animate={{ height: "auto", opacity: 1 }}
                    exit={{ height: 0, opacity: 0 }}
                    transition={{ duration: 0.2 }}
                    className="overflow-hidden"
                  >
                    <div
                      className="mt-3 ml-9 pl-3"
                      style={{ borderLeft: `2px solid ${step.color}30` }}
                    >
                      <p
                        className="text-[11px] leading-relaxed"
                        style={{ color: `${C}75` }}
                      >
                        {step.detail}
                      </p>
                    </div>
                  </motion.div>
                )}
              </AnimatePresence>
            </button>
          </motion.div>
        ))}
      </div>

      {/* Footer: prize split */}
      <div
        className="relative z-10 mx-5 mb-5 rounded p-3"
        style={{
          background: `${C}05`,
          border: `1px solid ${C}15`,
        }}
      >
        <p
          className="text-[10px] font-bold mb-2 tracking-widest"
          style={{ color: `${C}60` }}
        >
          TICKET_ALLOCATION //
        </p>
        <div className="flex gap-4">
          <div className="flex items-center gap-2">
            <div
              className="w-2 h-2 rounded-full"
              style={{ background: C, boxShadow: `0 0 6px ${C}` }}
            />
            <span className="text-[11px]" style={{ color: C }}>
              70%
            </span>
            <span className="text-[10px]" style={{ color: `${C}50` }}>
              → Vault
            </span>
          </div>
          <div className="flex items-center gap-2">
            <div
              className="w-2 h-2 rounded-full"
              style={{ background: "#ffaa00", boxShadow: "0 0 6px #ffaa0080" }}
            />
            <span className="text-[11px]" style={{ color: "#ffaa00" }}>
              30%
            </span>
            <span className="text-[10px]" style={{ color: `${C}50` }}>
              → Protocol
            </span>
          </div>
          <div className="ml-auto text-[10px]" style={{ color: `${C}30` }}>
            WINNER_TAKES_ALL
          </div>
        </div>
      </div>
    </div>
  );
}
