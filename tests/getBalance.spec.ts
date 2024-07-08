import { getBalance } from "../getBalance";

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
    expect(balance).toEqual("0");
  });
});
