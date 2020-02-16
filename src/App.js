import React from "react";
import "../node_modules/bootstrap/dist/css/bootstrap.min.css";
import "./App.css";
import { BrowserRouter as Router, Switch, Route, Link } from "react-router-dom";

import Login from "./components/login.component";
import Main from "./components/main.component";
import Option1 from "./components/option1.component";
import Option2 from "./components/option2.component";
import Option3 from "./components/option3.component";

function App() {
  return (
    <Router>
      <div className="App">
        <nav className="navbar navbar-expand-lg navbar-light fixed-top">
          <div className="container">
            <div className="collapse navbar-collapse" id="navbarTogglerDemo02">
              <ul className="navbar-nav ml-auto">
                <li className="nav-item">
                  {/* <Link className="nav-link" to={"/sign-in"}>
                    Login
                  </Link> */}
                </li>
                <li className="nav-item">
                  <Link className="nav-link" to={"/main"}>
                    Main
                  </Link>
                </li>
              </ul>
            </div>
          </div>
        </nav>

        <div className="auth-wrapper">
          <div className="auth-inner">
            <Switch>
              <Route exact path="/" component={Login} />
              <Route path="/sign-in" component={Login} />
              <Route path="/main" component={Main} />
              <Route path="/option1" component={Option1} />
              <Route path="/option2" component={Option2} />
              <Route path="/option3" component={Option3} />
            </Switch>
          </div>
        </div>
      </div>
    </Router>
  );
}

export default App;
