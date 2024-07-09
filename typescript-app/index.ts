import { getBalance } from "./getBalance";

(async () => {
  try {
    const args = process.argv.slice(2);
    const addressArg = args.find((arg) => arg.startsWith("--address="));
    const address = addressArg ? addressArg.split("=")[1] : "";
    const balance = await getBalance(address);
    console.log("balance for address: ", balance);
  } catch (e) {
    throw e;
  }
})();
