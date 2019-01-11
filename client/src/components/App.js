import React, { Component } from 'react';
import { BrowserRouter, Route } from 'react-router-dom';
import SignIn from './SignIn';
import Dashboard from './Dashboard/Dashboard';
import Vehicle from './Dashboard/Vehicle/Vehicle';

// Routes
const routes = [
  {
    path: "/",
    component: SignIn
  },
  {
    path: "/dashboard",
    component: Dashboard
  },
  {
    path: "/dashboard/vehicles",
    component: Vehicle
  },
  {
    path: "/dashboard/vehicles/add",
    component: Vehicle
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
