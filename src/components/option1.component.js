import React, { Component } from "react";

export default class Option1 extends Component {
  constructor(props) {
    super(props);
    this.state = {
      name: "health insurance",
      payment: "",
      description: ""
    };
    this.componentDidMount = this.componentDidMount.bind(this);
  }
  async request(data) {
    const url = `http://localhost:8084/search`;
    let options = {};
    if (data) {
      options = {
        body: JSON.stringify(data),
        headers: {
          "content-type": "application/json",
          origin: "http://localhost:3000",
          accept: "application/json"
        },
        method: "GET"
      };
      console.log(options);
    } else {
      options = {
        headers: {},
        method: "GET"
      };
    }
    const response = await fetch(url, options);

    return response.status;
  }
  async componentDidMount() {
    const a = await this.request(this.state.name);
    console.log(a);
    const son = JSON.parse(a);
    if (a === 200) {
      this.setState({ name: son.Name });
      this.setState({ payment: son.Payment });
      this.setState({ description: son.Desc });
    }
  }
  render() {
    return (
      <form>
        <h3>{this.state.name}</h3>
        <div className="form-group">
          <p>{this.state.description}</p>
          <p>{this.state.payment}kr.</p>
        </div>
      </form>
    );
  }
}
