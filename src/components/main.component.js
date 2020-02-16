import React, { Component } from "react";
import Dropdown from "react-bootstrap/Dropdown";

export default class SignUp extends Component {
  render() {
    return (
      <form>
        <h3>My pages</h3>
        <div className="form-group">
          <Dropdown className="dropdown">
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
