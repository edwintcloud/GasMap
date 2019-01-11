import React, { Component } from "react";
import { connect } from "react-redux";
import { Redirect } from "react-router-dom";
import { getUser, logoutUser } from "../../actions/users";
import { GoogleLogout } from "react-google-login";

class Dashboard extends Component {
  logoutUserClick = e => {
    this.props.logoutUser();
    window.location = "/";
  };

  vehiclesClick = () => {
    this.props.history.push("/dashboard/vehicles");
  };

  render() {
    if ("_id" in this.props.user) {
      return (
        <div className="two_grid_container">
          <div className="two_grid_a">
            <h1 className="header greeting">
              Hello, {this.props.user.firstName}
            </h1>
          </div>

          <div className="two_grid_b">
            <button className="button lg" onClick={this.vehiclesClick}>
              My Vehicles
            </button>
            <button className="button lg">My Trips</button>
            <GoogleLogout buttonText="Logout" onLogoutSuccess={this.logoutUserClick} className="button" />
            <a href="/">Settings</a>
          </div>
        </div>
      );
    }
    return <Redirect to="/" />;
  }
}

const mapStateToProps = state => {
  return {
    user: state.user,
    hasErrored: state.usersHasErrored,
    isLoading: state.usersIsLoading
  };
};

const mapDispatchToProps = dispatch => {
  return {
    getUser: () => dispatch(getUser()),
    logoutUser: () => dispatch(logoutUser())
  };
};

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(Dashboard);
