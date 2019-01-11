import React, { Component } from "react";
import GoogleLogin from "react-google-login";
import { Redirect } from "react-router-dom";
import { connect } from "react-redux";
import { usersSignin, getUser } from "../actions/users";

class SignIn extends Component {
  componentDidMount() {
    this.props.getUser();
  }

  signIn = response => {
    this.props.userSignin("/api/v1/users", {
      firstName: response.profileObj.givenName,
      lastName: response.profileObj.familyName,
      email: response.profileObj.email
    });
    this.props.history.push("/dashboard");
  };

  signInFailure = error => {
    console.log(error);
  };

  render() {
    if ("_id" in this.props.user) {
      return <Redirect to="/dashboard" />;
    }

    return (
      <div className="two_grid_container even">
        <div className="two_grid_a">
          <h1 className="header">
            Welcome
            <br /> to <br />
            Gas Map
          </h1>
        </div>

        <div className="signin_buttons">
          <GoogleLogin
            clientId={process.env.REACT_APP_GOOGLE_OAUTH_CLIENT_ID}
            uxMode="redirect"
            buttonText="Sign in With Google"
            onSuccess={this.signIn}
            onFailure={this.signInFailure}
            className="button"
            isSignedIn={true}
          />
        </div>
      </div>
    );
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
    userSignin: (url, data) => dispatch(usersSignin(url, data)),
    getUser: () => dispatch(getUser())
  };
};

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(SignIn);
