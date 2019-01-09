import React, { Component } from "react";
import GoogleLogin from "react-google-login";
import axios from "axios";
import Cookies from "js-cookie";

class SignIn extends Component {
  constructor(props) {
    super(props);
    this.state = {
      session: ''
    };
  }

  signIn = response => {
    axios
      .post("/api/v1/users", {
        firstName: response.profileObj.givenName,
        lastName: response.profileObj.familyName,
        email: response.profileObj.email
      })
      .then(res => {
        Cookies.set('_session', res.data.token, { expires: 365, path: '/' });
        this.setState({ session: res.data });
      })
      .catch(err => {
        console.log(err);
      });
  };

  signInFailure = error => {
    console.log(error.error);
  };

  render() {
    return (
      <div className="signin_container">
        <h1>Welcome to Gas Map</h1>
        <div className="signin_buttons">
          <GoogleLogin
            clientId={process.env.REACT_APP_GOOGLE_OAUTH_CLIENT_ID}
            uxMode="popup"
            buttonText="Sign in With Google"
            onSuccess={this.signIn}
            onFailure={this.signInFailure}
          />
        </div>
      </div>
    );
  }
}

export default SignIn;
