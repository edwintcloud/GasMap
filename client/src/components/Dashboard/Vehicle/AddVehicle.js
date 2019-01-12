import React, { Component } from "react";
import { connect } from "react-redux";
import { Redirect } from "react-router-dom";
import { getUser } from "../../../actions/users";
import Axios from "axios";

class AddVehicle extends Component {
  constructor(props) {
    super(props);

    this.state = {};
  }

  backClick = () => {
    this.props.history.push("/dashboard/vehicles");
  };

  formSubmit = e => {
    e.preventDefault();
    console.log(this.state)
    console.log(this.props.user.token)
    Axios
      .post("/api/v1/vehicles", this.state, {
        headers: { 'Authorization': `Bearer ${this.props.user.token}` }
      })
      .then(res => {
        this.props.getUser();
        this.props.history.push("/dashboard/vehicles");
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
            <h1 className="title">Add a Vehicle</h1>
          </div>

          <form onSubmit={this.formSubmit} className="add_vehicle_form">
            <div className="form-group">
              <label htmlFor="year">Year</label>
              <input
                id="year"
                name="year"
                type="text"
                onChange={this.inputChanged}
              />
            </div>
            <div className="form-group">
              <label htmlFor="make">Make</label>
              <input
                id="make"
                name="make"
                type="text"
                onChange={this.inputChanged}
              />
            </div>
            <div className="form-group">
              <label htmlFor="model">Model</label>
              <input
                id="model"
                name="model"
                type="text"
                onChange={this.inputChanged}
              />
            </div>
            <div className="form-group">
              <label htmlFor="mpg">MPG</label>
              <input
                id="mpg"
                name="mpg"
                type="text"
                onChange={this.inputChanged}
              />
            </div>
            <div className="form-group">
              <label htmlFor="tankSize">Tank Size</label>
              <input
                id="tankSize"
                name="tankSize"
                type="text"
                onChange={this.inputChanged}
              />
            </div>
            <div className="form-group">
              <label htmlFor="fuelQuality">Fuel Quality</label>
              <input
                id="fuelQuality"
                name="fuelQuality"
                type="text"
                onChange={this.inputChanged}
              />
            </div>
            <button type="submit" className="button form-submit-btn">
              Add Vehicle
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

export default connect(mapStateToProps, mapDispatchToProps)(AddVehicle);
