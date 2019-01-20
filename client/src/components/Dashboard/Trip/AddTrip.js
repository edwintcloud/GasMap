import React, { Component } from "react";
import { connect } from "react-redux";
import { Redirect } from "react-router-dom";
import { getUser } from "../../../actions/users";
import Axios from "axios";
import GooglePlacesAutocomplete from "react-google-places-autocomplete";
import ReactTooltip from "react-tooltip";

class AddTrip extends Component {
  constructor(props) {
    super(props);

    this.state = {
      stations : []
    };
  }

  componentDidMount() {
    // Below is a fix for multiple autocompletes on the page displaying suggstions correctly
    let index = 1;
    document
      .querySelectorAll(".google-places-autocomplete")
      .forEach(element => {
        element.id = `places${index}`;
        element.classList.remove("google-places-autocomplete");
        index++;
      });
    const from = document.getElementById("places1");
    const to = document.getElementById("places2");
    from.addEventListener("click", e => {
      to.classList.remove("google-places-autocomplete");
      from.classList.add("google-places-autocomplete");
    });
    to.addEventListener("click", e => {
      from.classList.remove("google-places-autocomplete");
      to.classList.add("google-places-autocomplete");
    });

    // set current mte to selected vehicle
    const vehicles = document.querySelectorAll("#vehicle > option");
    if (vehicles.length > 0) {
      const vehicle = this.props.user.vehicles.filter(
        vehicle => vehicle._id === vehicles[0].value
      )[0];
      const mte = Number(vehicle.mpg) * Number(vehicle.tankSize);
      this.setState({ vehicle: vehicle._id });
      if (!isNaN(mte)) {
        document.getElementById("currentMte").value = mte;
        this.setState({ currentMte: String(mte) });
      }
    }

    // disable autocomplete on all input elements
    document.querySelectorAll("input").forEach(el => {
      el.setAttribute("autocomplete", "off");
    });
  }

  backClick = () => {
    this.props.history.push("/dashboard/trips");
  };

  formSubmit = e => {
    e.preventDefault();
    this.calculateDistance();
  };

  calculateDistance = () => {
    // use distance matrix api to calculate distance of trip and update the state with result
    const matrix = new window.google.maps.DistanceMatrixService();
    let distance = "";
    const that = this;
    matrix.getDistanceMatrix(
      {
        origins: [this.state.from],
        destinations: [this.state.to],
        travelMode: window.google.maps.TravelMode.DRIVING,
        unitSystem: window.google.maps.UnitSystem.IMPERIAL
      },
      function(res, status) {
        if (status === "OK") {
          distance = res.rows[0].elements[0].distance.text;
          that.setState({ distance: distance }, () => {
            that.getStations();
          });
        }
      }
    );
  };

  createTrip = () => {
    Axios.post("/api/v1/trips", this.state, {
      headers: { Authorization: `Bearer ${this.props.user.token}` }
    })
      .then(res => {
        this.props.getUser();
        this.props.history.push("/dashboard/trips");
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

    if (
      e.target.name === "vehicle" &&
      this.props.user.hasOwnProperty("vehicles")
    ) {
      const vehicle = this.props.user.vehicles.filter(
        vehicle => vehicle._id === e.target.value
      )[0];
      const mte = Number(vehicle.mpg) * Number(vehicle.tankSize);
      if (!isNaN(mte)) {
        document.querySelector("#currentMte").value = mte;
      }
    }
  };

  fromSelected = e => {
    this.setState({ from: e.description });
  };

  toSelected = e => {
    this.setState({ to: e.description });
  };

  getStops = () => {
    const mte = this.state.currentMte - 50;  
    const distance = Number(this.state.distance.substring(0, this.state.distance.length-3));
    return Math.ceil(distance / mte);
  };

  getStations = () => {
    const that = this;
    const stops = this.getStops();
    const routeBoxer = new window.RouteBoxer();
    const directionService = new window.google.maps.DirectionsService();
    const placesService = new window.google.maps.places.PlacesService(
      window.document.createElement("div")
    );
    const request = {
      origin: this.state.from,
      destination: this.state.to,
      travelMode: window.google.maps.DirectionsTravelMode.DRIVING
    };

    // make directions request to direction api
    directionService.route(request, function(result, status) {
      if (status == window.google.maps.DirectionsStatus.OK) {
        console.log(result);

        const path = result.routes[0].overview_path;

        // box around overview path of the first route, 1 mile (in km) distance
        const boxes = routeBoxer.box(path, 1.609344);
        console.log(boxes);

        // find places
        const interval = Math.ceil(boxes.length / stops);
        let counter = 0;
        for (var i = 0; i < boxes.length; i += interval) {
          placesService.nearbySearch(
            {
              bounds: boxes[i],
              type: "gas_station"
            },
            function(results, status) {
              if (status == window.google.maps.places.PlacesServiceStatus.OK) {
                let station = results[0];
                results.forEach(el => {
                  if (el.rating > station.rating) {
                    station = el;
                  }
                });
                that.setState(
                  { stations: [station.vicinity, ...that.state.stations] },
                  () => {
                    console.log("added station")
                    console.log(that.state.stations)
                  }
                );
              }
            }
          );
        }
        console.log(counter)
        console.log(stops)
        // once all the stations are added, then create the trip

        // while (counter < stops * 2) {
        //   console.log(counter)
        //   if (counter === stops) {
        //     that.createUser();
        //     break;
        //   }
        // }
      }
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

          <form
            onSubmit={this.formSubmit}
            className="add_vehicle_form"
            autoComplete="off"
          >
            <div className="form-group">
              <label htmlFor="name">Name</label>
              <input
                id="name"
                name="name"
                type="text"
                onChange={this.inputChanged}
              />
            </div>
            <div className="form-group">
              <label htmlFor="from">From</label>
              <GooglePlacesAutocomplete id="one" onSelect={this.fromSelected} />
            </div>
            <div className="form-group">
              <label htmlFor="to">To</label>
              <GooglePlacesAutocomplete id="two" onSelect={this.toSelected} />
            </div>
            <div className="form-group">
              <label htmlFor="vehicle">Vehicle</label>
              <select name="vehicle" id="vehicle" onChange={this.inputChanged}>
                {this.props.user.hasOwnProperty("vehicles") &&
                  this.props.user.vehicles.map((vehicle, index) => (
                    <option key={index} value={vehicle._id}>
                      {vehicle.year} {vehicle.make} {vehicle.model}
                    </option>
                  ))}
              </select>
            </div>
            <div className="form-group">
              <ReactTooltip />
              <label htmlFor="mpg">
                Current <span data-tip="Miles Till Empty">MTE</span>
              </label>
              <input
                id="currentMte"
                name="currentMte"
                type="text"
                onChange={this.inputChanged}
              />
            </div>
            <button type="submit" className="button form-submit-btn">
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

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(AddTrip);
