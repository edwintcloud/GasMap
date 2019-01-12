import Axios from "axios";

export function usersHasErrored(bool) {
  return {
    type: "USERS_HAS_ERRORED",
    hasErrored: bool
  };
}

export function usersIsLoading(bool) {
  return {
    type: "USERS_IS_LOADING",
    isLoading: bool
  };
}

export function usersFetchDataSuccess(users) {
  return {
    type: "USERS_FETCH_DATA_SUCCESS",
    users
  };
}

export function usersSignin(url, data) {
  return dispatch => {
    dispatch(usersIsLoading(true));

    fetch(url, {
      method: "POST",
      headers: {
        "Content-Type": "application/json"
      },
      body: JSON.stringify(data)
    })
      .then(response => {
        if (!response.ok) {
          throw Error(response.statusText);
        }

        dispatch(usersIsLoading(false));

        return response;
      })
      .then(response => response.json())
      .then(user => {
        // create local storage instance of result
        localStorage.setItem("user", JSON.stringify(user));

        dispatch(usersFetchDataSuccess(user));
      })
      .catch(() => dispatch(usersHasErrored(true)));
  };
}

// get user from local storage
export function getUser() {
  return dispatch => {
    // get user from local storage
    const curUser = JSON.parse(localStorage.getItem("user"));
    if (curUser == null) {
      return;
    }

    // get user with vehicles
    Axios.get("/api/v1/users", {
      headers: { Authorization: `Bearer ${curUser.token}` }
    })
      .then(res => {
        localStorage.setItem("user", JSON.stringify(res.data));
        dispatch(usersFetchDataSuccess(res.data));
      })
      .catch(err => {
        console.log(err);
      });
  };
}

// delete user from local storage
export function logoutUser() {
  return dispatch => {
    localStorage.removeItem("user");
  };
}
