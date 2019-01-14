import React, { Component } from "react";
import { connect } from "react-redux";
import { Redirect } from "react-router-dom";
import { getUser } from "../../../actions/users";
import Axios from "axios";

class AddTrip extends Component {
  constructor(props) {
    super(props);

    this.state = {};
  }

  backClick = () => {
    this.props.history.push("/dashboard/trips");
  };

  formSubmit = e => {
    e.preventDefault();
    console.log(this.state)
    console.log(this.props.user.token)
    Axios
      .post("/api/v1/trip", this.state, {
        headers: { 'Authorization': `Bearer ${this.props.user.token}` }
      })
      .then(res => {
        this.props.getUser();
        this.props.history.push("/dashboard/trip");
      })
      .catch(err => {
        console.log(err);
      });
  };

  inputChanged = e => {
    // update state
    this.setState({
      [e.target.name]: e.target.value
    });
  };

  render() {
    if ("_id" in this.props.user) {
      return (
        <div className="two_grid_container sm_top">
          <div className="title_container">
            <i className="back_button" onClick={this.backClick} />
            <h1 className="title">Add a Trip</h1>
          </div>

          <form onSubmit={this.formSubmit} className="add_vehicle_form">
            <div className="form-group">
              <label htmlFor="from">From</label>
              <input
                id="from"
                name="from"
                type="text"
                onChange={this.inputChanged}
              />
            </div>
            <div className="form-group">
              <label htmlFor="to">To</label>
              <input
                id="to"
                name="to"
                type="text"
                onChange={this.inputChanged}
              />
            </div>
            <div className="form-group">
              <label htmlFor="vehicle">Vehicle</label>
              <input
                id="vehicle"
                name="vehicle"
                type="text"
                onChange={this.inputChanged}
              />
            </div>
            <div className="form-group">
              <label htmlFor="mpg">Current MTE</label>
              <input
                id="currentMte"
                name="currentMte"
                type="text"
                onChange={this.inputChanged}
              />
            </div>
            <button type="submit" className="button form-submit-btn" disabled>
              Add Trip
            </button>
          </form>
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

export default connect(mapStateToProps, mapDispatchToProps)(AddTrip);
