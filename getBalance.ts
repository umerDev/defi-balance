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

  if (response.status != 200)
    return new Error(`unable to decode balance data for ${address}`);

  const json: alephiumBody = await response.json();
  return json.balance;
}
