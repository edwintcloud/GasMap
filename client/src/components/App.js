import React, { Component } from 'react';
import { BrowserRouter, Route } from 'react-router-dom';
import SignIn from './SignIn';
import Dashboard from './Dashboard/Dashboard';
import Vehicle from './Dashboard/Vehicle/Vehicle';
import AddVehicle from './Dashboard/Vehicle/AddVehicle';
import Trip from './Dashboard/Trip/Trip';
import AddTrip from './Dashboard/Trip/AddTrip';
import ViewTrip from './Dashboard/Trip/ViewTrip';

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
    component: AddVehicle
  },
  {
    path: "/dashboard/trips",
    component: Trip
  },
  {
    path: "/dashboard/trips/add",
    component: AddTrip
  },
  {
    path: "/dashboard/trips/view/:id",
    component: ViewTrip
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
