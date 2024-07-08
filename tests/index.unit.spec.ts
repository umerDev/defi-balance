import { alephiumBody, getBalance } from "..";

beforeEach(() => {
  const mockResponse: alephiumBody = {
    balance: "1",
    balanceHint: "1 ALPH",
    lockedBalance: "1",
    lockedBalanceHint: "1 ALPH",
    utxoNum: 0,
  };

  global.fetch = jest.fn(() =>
    Promise.resolve({
      json: () => Promise.resolve(mockResponse),
    })
  ) as jest.Mock;
});

describe("Defi Balance()", () => {
  it("should throw an error for an empty address", async () => {
    //arrange
    const address = "";

    // act
    const balance = await getBalance(address);

    // assert
    expect(balance).toEqual(new Error("Provide a valid address"));
  });

  it("should return 0 defi balance", async () => {
    //arrange
    const address = "1DrDyTr9RpRsQnDnXo2YRiPzPW4ooHX5LLoqXrqfMrpQH";

    // act
    const balance = await getBalance(address);

    // assert
    expect(balance).toEqual("1");
  });
});
