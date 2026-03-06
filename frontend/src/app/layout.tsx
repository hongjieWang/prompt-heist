import type { Metadata } from "next";
import "./globals.css";
import "@rainbow-me/rainbowkit/styles.css";
import { Providers } from "./providers";

export const metadata: Metadata = {
  title: "Prompt Heist | AI Jailbreak Challenge",
  description: "Can you trick the AI into giving you the prize money?",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en" className="dark">
      <body className="antialiased">
        {/* 背景装饰层 */}
        <div className="fixed top-0 left-0 w-full h-full pointer-events-none z-[-1] bg-[url('/grid.png')] opacity-10" />
        <Providers>{children}</Providers>
      </body>
    </html>
  );
}
