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

  rowClicked = (e) => {
    console.log(e.target.className)
    document.querySelectorAll(`span`).forEach(el => {
      el.style.color = 'black';
    });
    document.querySelectorAll(`.${e.target.className}`).forEach(el => {
      el.style.color = 'blue';
    });
    document.getElementById('view_trip').disabled = false;
  }


  render() {
    if ("_id" in this.props.user) {
      return (
        <div className="two_grid_container sm_top">
          <div className="title_container">
            <i className="back_button" onClick={this.backClick} />
            <h1 className="title">My Trips</h1>
          </div>

          <div className="trips_grid">
            <span>Name</span>
            <span>To</span>
            <span>From</span>
            {this.props.user.hasOwnProperty("trips") && this.props.user.trips.map((trip) => (
              <>
              <span className={`trip-${trip._id}`}>{trip.name}</span>
              <span className={`trip-${trip._id}`} onClick={this.rowClicked}>{trip.from}</span>
              <span className={`trip-${trip._id}`}>{trip.to}</span>
              </>
            ))}
          </div>
          <div className="button_grid">
          <button id="view_trip" className="button trip_btn" onClick={this.addTripClick} disabled>
              View
          </button>
          <button className="button trip_btn" onClick={this.addTripClick}>
            Add A Trip
          </button>
          </div>
          
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
