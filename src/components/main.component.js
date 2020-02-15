import React, { Component } from "react";
import Dropdown from "react-bootstrap/Dropdown";

export default class SignUp extends Component {
  // constructor(props) {
  //   super(props);
  //   this.state = { IS: "" };
  //   this.handleChange = this.handleChange.bind(this);
  //   this.handleSubmit = this.handleSubmit.bind(this);
  // }

  // handleChange(eventKey) {
  //   console.log(eventKey);
  //   return;
  // }

  // handleSubmit() {
  //   return;
  // }

  render() {
    return (
      <form>
        <h3>My pages</h3>
        <div className="form-group">
          <Dropdown>
            <Dropdown.Toggle variant="success" id="dropdown-basic">
              Choose Insurance
            </Dropdown.Toggle>

            <Dropdown.Menu>
              <Dropdown.Item href="/option1">Health Insurance</Dropdown.Item>
              <Dropdown.Item href="/option2">Car Insurance</Dropdown.Item>
              <Dropdown.Item href="/option3">Home Insurance</Dropdown.Item>
            </Dropdown.Menu>
          </Dropdown>
        </div>

        <button type="submit" className="btn btn-primary btn-block">
          Submit
        </button>
      </form>
    );
  }
}
