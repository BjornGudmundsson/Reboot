import React, { Component } from "react";

export default class Insurances extends Component {
  constructor(props) {
    super(props);
    this.state = {
      name: "",
      payment: "",
      description: ""
    };
    this.componentDidMount = this.componentDidMount.bind(this);
    this.request = this.request.bind(this);
  }
  async request() {
    const url = `http://localhost:8084/myInsurances`;
    let options = {};
    options = {
      headers: {
        "content-type": "application/json",
        origin: "http://localhost:3000",
        accept: "application/json"
      },
      method: "GET"
    };

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
    const a = await this.request();
  }

  render() {
    return (
      <div>
        <form onSubmit={this.handleSubmit}>
          <h2 className="hecks" align="center">
            Your insurances
          </h2>
          <h3>{this.state.name}</h3>
          <div className="form-group">
            <h4 align="center">{this.state.payment}kr.</h4>
          </div>
        </form>
      </div>
    );
  }
}
