import React, { Component } from "react";

export default class Option2 extends Component {
  constructor(props) {
    super(props);
    this.state = {
      name: "Car insurance",
      payment: "",
      description: ""
    };
    this.componentDidMount = this.componentDidMount.bind(this);
    this.request = this.request.bind(this);
    this.handleSubmit = this.handleSubmit.bind(this);
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
        method: "POST"
      };
      console.log(options);
    } else {
      options = {
        headers: {},
        method: "POST"
      };
    }

    (async () => {
      const fetchedResource = await fetch(url, options);
      const reader = await fetchedResource.body.getReader();

      let charsReceived = 0;
      let result = "";

      const test = await reader.read(); // this works? what do with

      const d = new TextDecoder("utf-8").decode(test.value);
      const d2 = d.split("\n").join("");
      const virkar = JSON.parse(d2); //!!!!1
      console.log(virkar.Name); // !!!!!
      console.log(JSON.parse(d2));
      this.setState({ name: virkar.Name });
      this.setState({ description: virkar.Desc });
      this.setState({ payment: virkar.Payment });
      console.log(this.state);

      //POGCHAMP

      await reader.read().then(function processText({ done, value }) {
        if (done) {
          console.log("Stream finished. Content received:");
          console.log(result);
          return;
        }

        console.log(`Received ${result.length} chars so far!`);

        result += value;

        return reader.read().then(processText);
      });
    })();
  }

  //vorum að fá readable stream frá server, svo þetta mix er bara ehv til að extracta data from that, prob not needed??
  async componentDidMount() {
    const a = await this.request({ Name: this.state.name });
  }

  async req(data) {
    const url = `http://localhost:8084/addInsurance`;
    let options = {};
    if (data) {
      options = {
        body: data.toString(),
        headers: {
          "content-type": "text",
          origin: "http://localhost:3000"
        },
        method: "POST"
      };
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
    console.log("test");
    console.log(this);
    const a = await this.req(2);
    // the fuck is this?? ahh
    // FATTAÐI
    // All work now
    //
    if (a === 200) {
      this.props.history.push("/insurances");
    }
  }
  render() {
    return (
      <form onSubmit={this.handleSubmit}>
        <h3>{this.state.name}</h3>
        <div className="form-group">
          <p align="center">{this.state.description}</p>
          <p align="center">{this.state.payment}kr.</p>
        </div>
        <button type="submit" className="btn btn-primary btn-block">
          Buy insurance
        </button>
      </form>
    );
  }
}
