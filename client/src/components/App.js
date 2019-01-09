import React, { Component } from 'react';
import { BrowserRouter, Route } from 'react-router-dom';
import SignIn from './SignIn';

// Routes
const routes = [
  {
    path: "/",
    component: SignIn
  }
];

class App extends Component {
  render() {
    return (
      <BrowserRouter>
      <div className="container">
        {routes.map((route, i) => (
          <Route key={i} exact path={route.path} component={route.component} />
        ))}
      </div>
      </BrowserRouter>
    );
  }
}

export default App;
