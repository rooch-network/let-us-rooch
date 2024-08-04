// Copyright (c) RoochNetwork
// SPDX-License-Identifier: Apache-2.0
// Author: Jason Jo

import { LoadingButton } from "@mui/lab";
import { Button, Chip, Divider, Stack, Typography } from "@mui/material";
import { Args, Transaction } from "@roochnetwork/rooch-sdk";
import {
  UseSignAndExecuteTransaction,
  useConnectWallet,
  useCreateSessionKey,
  useCurrentAddress,
  useCurrentSession,
  useRemoveSession,
  useRoochClientQuery,
  useWalletStore,
  useWallets,
} from "@roochnetwork/rooch-sdk-kit";
import { useState } from "react";
import "./App.css";
import { shortAddress } from "./utils";

// Publish address of the counter contract
const counterAddress =
  "0x0cc94c5429368b2dcd7ebfca18b65e891d8ae0fad6371514d42f4c7d6f50d9cf";

function App() {
  const wallets = useWallets();
  const currentAddress = useCurrentAddress();
  const sessionKey = useCurrentSession();
  const connectionStatus = useWalletStore((state) => state.connectionStatus);
  const setWalletDisconnected = useWalletStore(
    (state) => state.setWalletDisconnected
  );
  const { mutateAsync: connectWallet } = useConnectWallet();

  const { mutateAsync: createSessionKey } = useCreateSessionKey();
  const { mutateAsync: removeSessionKey } = useRemoveSession();
  const { mutateAsync: signAndExecuteTransaction } =
    UseSignAndExecuteTransaction();
  const { data, refetch } = useRoochClientQuery(
    "executeViewFunction",
    {
      target: `${counterAddress}::quick_start_counter::value`,
      args: [
        currentAddress != null
          ? Args.address(currentAddress)
          : Args.address(""),
      ],
    },
    {
      enabled: currentAddress != null,
    }
  );

  const [sessionLoading, setSessionLoading] = useState(false);
  const [txnLoading, setTxnLoading] = useState(false);
  const handlerCreateSessionKey = async () => {
    if (sessionLoading) {
      return;
    }
    setSessionLoading(true);

    const defaultScopes = [`${counterAddress}::*::*`];
    createSessionKey(
      {
        appName: "my_first_rooch_dapp",
        appUrl: "http://localhost:5173",
        maxInactiveInterval: 1000,
        scopes: defaultScopes,
      },
      {
        onSuccess: (result) => {
          console.log("session key", result);
        },
        onError: (why) => {
          console.log(why);
        },
      }
    ).finally(() => setSessionLoading(false));
  };

  return (
    <Stack
      className="font-sans min-w-[1024px]"
      direction="column"
      sx={{
        minHeight: "calc(100vh - 4rem)",
      }}
    >
      <Stack justifyContent="space-between" className="w-full">
        <img src="./rooch_black_combine.svg" width="120px" alt="" />
        <Stack spacing={1} justifyItems="flex-end">
          <Chip
            label="Rooch Testnet"
            variant="filled"
            className="font-semibold !bg-slate-950 !text-slate-50 min-h-10"
          />
          <Button
            variant="outlined"
            onClick={async () => {
              if (connectionStatus === "connected") {
                setWalletDisconnected();
                return;
              }
              await connectWallet({ wallet: wallets[0] });
            }}
          >
            {connectionStatus === "connected"
              ? shortAddress(currentAddress?.genRoochAddress().toStr(), 8, 6)
              : "Connect Wallet"}
          </Button>
        </Stack>
      </Stack>
      <Typography className="text-4xl font-semibold mt-6 text-left w-full mb-4">
        My First Rooch dApp | <span className="text-2xl">Counter</span>
      </Typography>
      <Divider className="w-full" />
      <Stack
        direction="column"
        className="mt-4 font-medium font-serif w-full text-left"
        spacing={2}
        alignItems="flex-start"
      >
        <Typography className="text-xl">
          Rooch Address:{" "}
          <span className="underline tracking-wide underline-offset-8 ml-2">
            {currentAddress?.genRoochAddress().toStr()}
          </span>
        </Typography>
        <Typography className="text-xl">
          Hex Address:
          <span className="underline tracking-wide underline-offset-8 ml-2">
            {currentAddress?.genRoochAddress().toHexAddress()}
          </span>
        </Typography>
        <Typography className="text-xl">
          Bitcoin Address:
          <span className="underline tracking-wide underline-offset-8 ml-2">
            {currentAddress?.toStr()}
          </span>
        </Typography>
      </Stack>
      <Divider className="w-full !mt-12" />
      <Stack
        className="mt-4 w-full font-medium "
        direction="column"
        alignItems="flex-start"
      >
        <Typography className="text-3xl font-bold">Session Key</Typography>
        {/* <Typography className="mt-4">
          Status: Session Key not created
        </Typography> */}
        <Stack
          className="mt-4 text-left"
          spacing={2}
          direction="column"
          alignItems="flex-start"
        >
          <Typography className="text-xl">
            Session Rooch address:{" "}
            <span className="underline tracking-wide underline-offset-8 ml-2">
              {sessionKey?.getRoochAddress().toStr()}
            </span>
          </Typography>
          <Typography className="text-xl">
            Key scheme:{" "}
            <span className="underline tracking-wide underline-offset-8 ml-2">
              {sessionKey?.getKeyScheme()}
            </span>
          </Typography>
          <Typography className="text-xl">
            Create time:{" "}
            <span className="underline tracking-wide underline-offset-8 ml-2">
              {sessionKey?.getCreateTime()}
            </span>
          </Typography>
        </Stack>
        {!sessionKey ? (
          <LoadingButton
            loading={sessionLoading}
            variant="contained"
            className="!mt-4"
            disabled={connectionStatus !== "connected"}
            onClick={() => {
              handlerCreateSessionKey();
            }}
          >
            {connectionStatus !== "connected"
              ? "Please connect wallet first"
              : "Create"}
          </LoadingButton>
        ) : (
          <Button
            variant="contained"
            className="!mt-4"
            onClick={() => {
              removeSessionKey({ authKey: sessionKey.getAuthKey() });
            }}
          >
            Clear Session
          </Button>
        )}
      </Stack>
      <Divider className="w-full !mt-12" />
      <Stack
        className="mt-4 w-full font-medium "
        direction="column"
        alignItems="flex-start"
      >
        <Typography className="text-3xl font-bold">
          dApp integration
          <span className="text-base font-normal ml-4">({counterAddress})</span>
        </Typography>
        <Stack
          className="mt-4"
          spacing={2}
          direction="column"
          alignItems="flex-start"
        >
          <Typography className="text-xl">
            Counter Value:{" "}
            <span className="underline tracking-wide underline-offset-8 ml-2">
              {data?.return_values?.[0]?.decoded_value.toString()}
            </span>
          </Typography>
          {data?.return_values == null ? (
            <LoadingButton
              loading={txnLoading}
              variant="contained"
              fullWidth
              disabled={!sessionKey}
              onClick={async () => {
                console.log("refetching");
                try {
                  setTxnLoading(true);
                  const txn = new Transaction();
                  txn.callFunction({
                    address: counterAddress,
                    module: "quick_start_counter",
                    function: "mint",
                    args: [],
                  });
                  await signAndExecuteTransaction({ transaction: txn });
                  await refetch();
                } catch (error) {
                  console.error(String(error));
                } finally {
                  setTxnLoading(false);
                }
              }}
            >
              Mint Counter
            </LoadingButton>
          ) : (
            <LoadingButton
              loading={txnLoading}
              variant="contained"
              fullWidth
              // disabled={!sessionKey}
              onClick={async () => {
                console.log("refetching");
                try {
                  setTxnLoading(true);
                  const txn = new Transaction();
                  txn.callFunction({
                    address: counterAddress,
                    module: "quick_start_counter",
                    function: "increase",
                    args: [],
                  });
                  await signAndExecuteTransaction({ transaction: txn });
                  await refetch();
                } catch (error) {
                  console.error(String(error));
                } finally {
                  setTxnLoading(false);
                }
              }}
            >
              {sessionKey
                ? "Increase Counter Value"
                : "Please create Session Key first"}
            </LoadingButton>
          )}
        </Stack>
      </Stack>
    </Stack>
  );
}

export default App;
