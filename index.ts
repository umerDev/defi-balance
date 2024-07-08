const url = (address: string) =>
  `https://wallet.mainnet.alephium.org/addresses/${address}/balance`;

export interface alephiumBody {
  balance: string;
  balanceHint: string;
  lockedBalance: string;
  lockedBalanceHint: string;
  utxoNum: number;
}

export async function getBalance(address: string): Promise<string | Error> {
  if (!address) return new Error("Provide a valid address");

  const response = await fetch(url(address));
  const json: alephiumBody = await response.json();
  return json.balance;
}

(async () => {
  try {
    const args = process.argv.slice(2);
    const addressArg = args.find((arg) => arg.startsWith("--arg="));
    const address = addressArg ? addressArg.split("=")[1] : "";
    const balance = await getBalance(address);
    console.log("balance for address: ", balance);
  } catch (e) {
    throw e;
  }
})();
