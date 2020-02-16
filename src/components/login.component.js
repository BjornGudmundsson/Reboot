import React, { Component } from "react";

export default class Login extends Component {
  constructor(props) {
    super(props);
    this.state = { Number: "" };
    this.handleChangeNum = this.handleChangeNum.bind(this);
    this.handleSubmit = this.handleSubmit.bind(this);
  }

  handleChangeNum(event) {
    this.setState({ Number: event.target.value });
  }

  async request(data) {
    const url = `http://localhost:8084/loginForm`;
    let options = {};
    console.log(data);
    if (data) {
      options = {
        body: JSON.stringify(data),
        headers: {
          "content-type": "application/json",
          origin: "http://localhost:3000",
          accept: "application/json"
        },
        method: "POST"
      };
      console.log(options);
    } else {
      options = {
        headers: {},
        method: "POST"
      };
    }

    const response = await fetch(url, options);

    return response.status;
  }

  async handleSubmit(e) {
    e.preventDefault();
    console.log(this.state);

    const a = await this.request(this.state);
    console.log(a);

    if (a === 200) {
      this.props.history.push("/main");
    }
  }
  render() {
    return (
      <form onSubmit={this.handleSubmit}>
        <h3>Sign In</h3>

        <div className="form-group">
          <label>Fone number</label>
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
