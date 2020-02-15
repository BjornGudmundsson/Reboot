import React, { Component } from "react";

export default class Login extends Component {
  constructor(props) {
    super(props);
    this.state = { PW: "" };
    this.handleChangeNum = this.handleChangeNum.bind(this);
    this.handleSubmit = this.handleSubmit.bind(this);
  }

  handleChangeNum(event) {
    this.setState({ Number: event.target.value });
  }

  handleSubmit(e) {
    e.preventDefault();
    var xhr = new XMLHttpRequest();

    console.log(this.state);
    xhr.open("POST", "http://localhost:8084");
    xhr.send();
    xhr.addEventListener("load", () => {
      if (xhr.status === 200) {
        this.props.history.push("/main");
      }
    });
    this.setState = { Number: "" };
  }
  render() {
    return (
      <form onSubmit={this.handleSubmit}>
        <h3>Sign In</h3>

        <div className="form-group">
          <label>Phone number</label>
          <input
            type="text"
            value={this.state.Number}
            name="Phone number"
            className="form-control"
            placeholder="Enter phone number"
            onChange={this.handleChangeNum}
          />
        </div>

        <button type="submit" className="btn btn-primary btn-block">
          Submit
        </button>
      </form>
    );
  }
}
