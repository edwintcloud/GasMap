import React, { Component } from "react";
import { connect } from "react-redux";
import { Redirect } from "react-router-dom";
import { getUser } from "../../../actions/users";
import Axios from "axios";

class ViewTrip extends Component {
  constructor(props) {
    super(props);

    this.state = {};
  }

  componentWillMount() {
    if (
      this.props.user.hasOwnProperty("trips") &&
      this.props.user.hasOwnProperty("vehicles")
    ) {
      let trip = this.props.user.trips.filter(
        el => el._id === this.props.match.params.id
      )[0];
      const vehicle = this.props.user.vehicles.filter(
        el => el._id === trip.vehicle
      )[0];
      trip.vehicle = vehicle;
      this.setState({ trip });
    } else {
      this.props.history.push("/dashboard/trips");
    }
  }

  backClick = () => {
    this.props.history.push("/dashboard/trips");
  };

  deleteClick = () => {
    if (window.confirm("Are you sure you want to delete this trip?")) {
      Axios.delete(`/api/v1/trips?id=${this.state.trip._id}`, {
        headers: { Authorization: `Bearer ${this.props.user.token}` }
      })
        .then(res => {
          this.props.getUser();
          this.props.history.push("/dashboard/trips");
        })
        .catch(err => {
          console.log(err);
        });
    }
  };

  navigateClick = () => {
    this.openMaps();
  };

  openMaps = () => {
    console.log(this.state.trip)
    // open using app if on iOS, otherwise open with google url
    if (
      navigator.platform.indexOf("iPhone") !== -1 ||
      navigator.platform.indexOf("iPad") !== -1 ||
      navigator.platform.indexOf("iPod") !== -1
    ) {
      window.open("maps://maps.google.com/maps/dir/");
    } else {
      window.open("https://maps.google.com/maps/dir/");
    }
  };

  render() {
    if ("_id" in this.props.user) {
      return (
        <div className="two_grid_container sm_top">
          <div className="title_container">
            <i className="back_button" onClick={this.backClick} />
            <h1 className="title">{this.state.trip.name}</h1>
          </div>
          <div className="view_trip_container">
            <div className="trip_info">
              <span>From</span>
              <span>{this.state.trip.from}</span>
            </div>
            <div className="trip_info">
              <span>To</span>
              <span>{this.state.trip.to}</span>
            </div>
            <div className="trip_info">
              <span>Distance</span>
              <span>
                {(this.state.trip.hasOwnProperty("distance") &&
                  this.state.trip.distance) ||
                  `-`}
              </span>
            </div>
            <div className="trip_info">
              <span>Gas Station Stops</span>
              <span>
                {(this.state.trip.hasOwnProperty("stops") &&
                  this.state.trip.stops) ||
                  `-`}
              </span>
            </div>
            <div className="trip_info">
              <span>Estimated Gallons Used</span>
              <span>
                {(this.state.trip.hasOwnProperty("gallons") &&
                  this.state.trip.gallons + " gal") ||
                  `-`}
              </span>
            </div>
            <div className="trip_info">
              <span>Estimated Gas Cost</span>
              <span>
                {(this.state.trip.hasOwnProperty("price") &&
                  "$" + this.state.trip.price) ||
                  `-`}
              </span>
            </div>
          </div>
          <div className="button_grid">
            <button
              id="view_trip"
              className="button trip_btn delete_btn"
              onClick={this.deleteClick}
            >
              Delete
            </button>
            <button className="button trip_btn" onClick={this.navigateClick}>
              Navigate
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

const mapDispatchToProps = dispatch => {
  return {
    getUser: () => dispatch(getUser())
  };
};

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(ViewTrip);
