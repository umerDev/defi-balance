import { getBalance } from "..";

describe("Defi Balance()", () => {
  it("should return defi balance", () => {
    //arrange
    const address = "1";

    // act
    const balance = getBalance(address);

    // assert
    expect(balance).toEqual("1");
  });
});
