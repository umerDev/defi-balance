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

  it("should return an error for incorrect address", async () => {
    //arrange
    const address = "fakeaddress";

    // act
    const balance = await getBalance(address);

    // assert
    expect(balance).toEqual(
      new Error("unable to decode balance data for fakeaddress")
    );
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
