// Copyright (c) RoochNetwork
// SPDX-License-Identifier: Apache-2.0
// Author: Jason Jo

import "@fontsource-variable/plus-jakarta-sans";
import "@fontsource-variable/raleway";
import { RoochProvider, WalletProvider } from "@roochnetwork/rooch-sdk-kit";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import ReactDOM from "react-dom/client";
import App from "./App.tsx";
import "./index.css";
import { networkConfig } from "./networks.ts";

const queryClient = new QueryClient();

ReactDOM.createRoot(document.getElementById("root")!).render(
    <QueryClientProvider client={queryClient}>
      <RoochProvider networks={networkConfig} defaultNetwork="testnet">
        <WalletProvider chain={"bitcoin"} autoConnect>
          <App />
        </WalletProvider>
      </RoochProvider>
    </QueryClientProvider>
);
