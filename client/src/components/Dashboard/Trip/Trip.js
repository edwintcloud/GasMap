import React, { Component } from "react";
import { connect } from "react-redux";
import { Redirect } from "react-router-dom";

class Trip extends Component {
  backClick = () => {
    this.props.history.push("/dashboard");
  };

  addTripClick = () => {
    this.props.history.push("/dashboard/trips/add");
  };

  render() {
    if ("_id" in this.props.user) {
      return (
        <div className="two_grid_container sm_top">
          <div className="title_container">
            <i className="back_button" onClick={this.backClick} />
            <h1 className="title">My Trips</h1>
          </div>

          <div className="trips_grid">
            <span>From</span>
            <span>To</span>
            <span></span>
            <span>1 Main St Brinson, GA 39845</span>
            <span>244 McAllister St San Francisco, CA 94102</span>
            <span className="trips_button_container"><button className="button trips_button">View</button><button className="button trips_button">Go</button></span>
            {/* {this.props.user.hasOwnProperty("vehicles") && this.props.user.vehicles.map((vehicle) => (
              <>
              <object className="vehicle_picture" aria-label="Vehicle Photo" />
              <span className="vehicle_name">{`${vehicle.year} ${vehicle.make} ${vehicle.model}`}</span>
              <span className="vehicle_mpg">{this.getMpg(vehicle.mpg)}</span>
              <span className="vehicle_mte">{this.getMte(vehicle.mpg, vehicle.tankSize)}</span>
              </>
            ))} */}
          </div>
          <button className="button" onClick={this.addTripClick}>
            Add A Trip
          </button>
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

export default connect(mapStateToProps)(Trip);
