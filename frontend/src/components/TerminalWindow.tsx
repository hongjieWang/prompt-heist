import React, { useEffect, useRef, useState } from "react";
import { motion, AnimatePresence } from "framer-motion";

interface TerminalWindowProps {
  children: React.ReactNode;
  className?: string;
  title?: string;
  status?: "online" | "offline" | "warning";
}

export const TerminalWindow: React.FC<TerminalWindowProps> = ({
  children,
  className = "",
  title = "TERMINAL_V1.0",
  status = "online",
}) => {
  const scrollRef = useRef<HTMLDivElement>(null);
  const [glitch, setGlitch] = useState(false);
  const [scanLine, setScanLine] = useState(0);

  // Auto-scroll
  useEffect(() => {
    if (scrollRef.current) {
      scrollRef.current.scrollTop = scrollRef.current.scrollHeight;
    }
  }, [children]);

  // Random glitch effect
  useEffect(() => {
    const interval = setInterval(() => {
      if (Math.random() > 0.85) {
        setGlitch(true);
        setTimeout(() => setGlitch(false), 120);
      }
    }, 2500);
    return () => clearInterval(interval);
  }, []);

  // Scan line animation
  useEffect(() => {
    let frame: number;
    let pos = 0;
    const animate = () => {
      pos = (pos + 0.4) % 100;
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
    warning: "SYS_WARNING2",
  };

  return (
    <div
      className={`relative overflow-hidden rounded-lg font-mono ${className}`}
      style={{
        background:
          "linear-gradient(135deg, #020c02 0%, #000a00 50%, #010801 100%)",
        border: `1px solid ${statusColors[status]}22`,
        boxShadow: `
          0 0 0 1px ${statusColors[status]}15,
          0 0 20px ${statusColors[status]}10,
          0 0 60px ${statusColors[status]}08,
          inset 0 0 60px rgba(0,0,0,0.5)
        `,
      }}
    >
      {/* CRT scanline overlay */}
      <div
        className="absolute inset-0 pointer-events-none z-20"
        style={{
          background: `linear-gradient(
            to bottom,
            transparent ${scanLine - 2}%,
            ${statusColors[status]}06 ${scanLine}%,
            transparent ${scanLine + 2}%
          )`,
          transition: "none",
        }}
      />

      {/* Static noise texture */}
      <div
        className="absolute inset-0 pointer-events-none z-10 opacity-[0.03]"
        style={{
          backgroundImage: `url("data:image/svg+xml,%3Csvg viewBox='0 0 256 256' xmlns='http://www.w3.org/2000/svg'%3E%3Cfilter id='noise'%3E%3CfeTurbulence type='fractalNoise' baseFrequency='0.9' numOctaves='4' stitchTiles='stitch'/%3E%3C/filter%3E%3Crect width='100%25' height='100%25' filter='url(%23noise)'/%3E%3C/svg%3E")`,
          backgroundSize: "128px 128px",
        }}
      />

      {/* Horizontal scan lines (CRT effect) */}
      <div
        className="absolute inset-0 pointer-events-none z-10"
        style={{
          backgroundImage:
            "repeating-linear-gradient(0deg, transparent, transparent 2px, rgba(0,0,0,0.08) 2px, rgba(0,0,0,0.08) 4px)",
        }}
      />

      {/* Top bar */}
      <div
        className="relative flex items-center px-4 py-2.5 z-30"
        style={{
          borderBottom: `1px solid ${statusColors[status]}20`,
          background: `linear-gradient(90deg, ${statusColors[status]}08 0%, transparent 60%)`,
        }}
      >
        {/* Traffic lights */}
        <div className="flex gap-1.5 mr-4">
          {["#ff5f57", "#ffbd2e", "#28c840"].map((color, i) => (
            <div
              key={i}
              className="w-3 h-3 rounded-full cursor-pointer transition-all duration-150 hover:brightness-125"
              style={{
                background: color,
                boxShadow: `0 0 6px ${color}80`,
              }}
            />
          ))}
        </div>

        {/* Title */}
        <motion.span
          animate={glitch ? { x: [0, -2, 2, 0], opacity: [1, 0.5, 1] } : {}}
          transition={{ duration: 0.1 }}
          className="text-xs font-bold tracking-widest flex-1 text-center"
          style={{
            color: statusColors[status],
            textShadow: `0 0 8px ${statusColors[status]}80`,
            letterSpacing: "0.2em",
          }}
        >
          {title}
        </motion.span>

        {/* Status indicator */}
        <div className="flex items-center gap-2">
          <motion.div
            animate={{ opacity: [1, 0.2, 1] }}
            transition={{ duration: 1.4, repeat: Infinity, ease: "easeInOut" }}
            className="w-1.5 h-1.5 rounded-full"
            style={{
              background: statusColors[status],
              boxShadow: `0 0 6px ${statusColors[status]}`,
            }}
          />
          <span
            className="text-[10px] tracking-widest"
            style={{ color: `${statusColors[status]}90` }}
          >
            {statusLabel[status]}
          </span>
        </div>
      </div>

      {/* Corner decorations */}
      {[
        "top-0 left-0 border-t border-l",
        "top-0 right-0 border-t border-r",
        "bottom-0 left-0 border-b border-l",
        "bottom-0 right-0 border-b border-r",
      ].map((pos, i) => (
        <div
          key={i}
          className={`absolute w-4 h-4 z-30 ${pos}`}
          style={{ borderColor: `${statusColors[status]}60` }}
        />
      ))}

      {/* Content area */}
      <div
        ref={scrollRef}
        className="relative z-30 overflow-y-auto p-4"
        style={{
          maxHeight: "60vh",
          scrollbarWidth: "thin",
          scrollbarColor: `${statusColors[status]}40 transparent`,
        }}
      >
        <AnimatePresence mode="sync">{children}</AnimatePresence>

        {/* Blinking cursor */}
        <motion.span
          animate={{ opacity: [1, 0] }}
          transition={{
            duration: 0.6,
            repeat: Infinity,
            repeatType: "reverse",
          }}
          className="inline-block w-2 h-[1em] ml-0.5 align-middle translate-y-[1px]"
          style={{
            background: statusColors[status],
            boxShadow: `0 0 6px ${statusColors[status]}`,
          }}
        />
      </div>

      {/* Bottom status bar */}
      <div
        className="relative z-30 flex items-center justify-between px-4 py-1.5"
        style={{
          borderTop: `1px solid ${statusColors[status]}15`,
          background: `rgba(0,0,0,0.4)`,
        }}
      >
        <span
          className="text-[10px] tracking-widest"
          style={{ color: `${statusColors[status]}40` }}
        >
          SECURE_CHANNEL_ESTABLISHED
        </span>
        <motion.span
          animate={{ opacity: [0.3, 0.7, 0.3] }}
          transition={{ duration: 3, repeat: Infinity }}
          className="text-[10px] tracking-widest"
          style={{ color: `${statusColors[status]}40` }}
        >
          ENC:AES-256
        </motion.span>
      </div>
    </div>
  );
};
