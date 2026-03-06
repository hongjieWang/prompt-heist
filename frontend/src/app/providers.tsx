"use client";

import * as React from "react";
import {
  RainbowKitProvider,
  darkTheme,
  connectorsForWallets,
} from "@rainbow-me/rainbowkit";
import {
  metaMaskWallet,
  walletConnectWallet,
  coinbaseWallet,
  okxWallet,
} from "@rainbow-me/rainbowkit/wallets";
import { WagmiProvider, createConfig, http } from "wagmi";
import { bscTestnet } from "wagmi/chains";
import { QueryClientProvider, QueryClient } from "@tanstack/react-query";

const projectId = "3a8170812b534d0ff9d794f19a901d64";

// ✅ 移到独立文件，避免模块重复执行
function makeConfig() {
  const connectors = connectorsForWallets(
    [
      {
        groupName: "Popular",
        wallets: [
          metaMaskWallet,
          walletConnectWallet,
          coinbaseWallet,
          okxWallet,
        ],
      },
    ],
    { appName: "Prompt Heist", projectId },
  );

  return createConfig({
    connectors,
    chains: [bscTestnet],
    multiInjectedProviderDiscovery: true,
    transports: {
      [bscTestnet.id]: http("https://data-seed-prebsc-1-s1.binance.org:8545"),
    },
    ssr: false,
  });
}

// ✅ 用 useState 确保只初始化一次
const queryClient = new QueryClient();

export function Providers({ children }: { children: React.ReactNode }) {
  const [mounted, setMounted] = React.useState(false);
  // ✅ config 用 useState 保证只创建一次，不随重渲染重复初始化
  const [config] = React.useState(() => makeConfig());

  React.useEffect(() => {
    setMounted(true);
  }, []);

  return (
    <WagmiProvider config={config}>
      <QueryClientProvider client={queryClient}>
        <RainbowKitProvider
          initialChain={bscTestnet}
          theme={darkTheme({
            accentColor: "#0f0",
            accentColorForeground: "black",
            borderRadius: "small",
            fontStack: "system",
            overlayBlur: "small",
          })}
          modalSize="compact"
        >
          {mounted ? children : null}
        </RainbowKitProvider>
      </QueryClientProvider>
    </WagmiProvider>
  );
}
