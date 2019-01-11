import React, { Component } from "react";
import { connect } from "react-redux";
import { Redirect } from "react-router-dom";

class Vehicle extends Component {
  backClick = () => {
    this.props.history.push("/dashboard");
  };

  addVehicleClick = () => {
    this.props.history.push("/dashboard/vehicles/add");
  };

  render() {
    if ("_id" in this.props.user) {
      return (
        <div className="two_grid_container sm_top">
          <div className="title_container">
            <i className="back_button" onClick={this.backClick} />
            <h1 className="title">My Vehicles</h1>
          </div>

          <div className="vehicles_grid">
              <span></span>
              <span>Name</span>
              <span>Miles Per Gallon</span>
              <span>Miles Till Empty</span>
              <object className="vehicle_picture" aria-label="Vehicle Photo" />
              <span className="vehicle_name">2013 GMC Sierra</span>
              <span className="vehicle_mpg">17</span>
              <span className="vehicle_mte">467</span>
              <object className="vehicle_picture" aria-label="Vehicle Photo" />
              <span className="vehicle_name">2013 GMC Sierra</span>
              <span className="vehicle_mpg">17</span>
              <span className="vehicle_mte">467</span>
          </div>
          <button className="button" onClick={this.addVehicleClick}>Add A Vehicle</button>
        </div>
      );
    }
    return <Redirect to="/" />;
  }
}

const mapStateToProps = state => {
  return {
    user: state.user
  };
};

export default connect(mapStateToProps)(Vehicle);
