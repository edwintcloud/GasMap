import React, { Component } from "react";
import { connect } from "react-redux";
import { Redirect } from "react-router-dom";
import { getUser, logoutUser } from "../../actions/users";

class Dashboard extends Component {

  logoutUserClick = (e) => {
    e.preventDefault();
    this.props.logoutUser();
    window.location = "/"
  }

  render() {
    if ("_id" in this.props.user) {
      return (
        <div className="signin_container">
        <h1>Hello, {this.props.user.firstName}</h1>
        <div className="signin_buttons">
          <button className="button bigger_button">My Vehicles</button>
          <button className="button bigger_button">My Trips</button>
          <a href="/">Settings</a>
          <a href="/" onClick={this.logoutUserClick}>Logout</a>
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
