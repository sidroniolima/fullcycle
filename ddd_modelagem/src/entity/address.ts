export default class Address {
  _street: string = "";
  _number: number = 0;
  _city: string = "";
  _state: string = "";
  _zip: string = "";

  constructor(
    street: string,
    number: number,
    city: string,
    state: string,
    zip: string
  ) {
    this._street = street;
    this._number = number;
    this._city = city;
    this._state = state;
    this._zip = zip;

    this.validate();
  }

  validate() {
    if (this._street.length === 0) {
      throw new Error("Street is required");
    }
  }
}
