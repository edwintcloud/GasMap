import React, { Component } from "react";
import { connect } from "react-redux";
import { Redirect } from "react-router-dom";
import { getUser } from "../../actions/users";

class Dashboard extends Component {

  render() {
    if ("_id" in this.props.user) {
      return (
        <div className="signin_container">
          hello {this.props.user.firstName}
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
    getUser: () => dispatch(getUser())
  };
};

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(Dashboard);
