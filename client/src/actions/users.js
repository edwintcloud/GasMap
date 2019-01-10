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
    const curUser = JSON.parse(localStorage.getItem("user"));
    if (curUser != null) {
      dispatch(usersFetchDataSuccess(curUser));
    }
  };
}

// delete user from local storage
export function logoutUser() {
  return dispatch => {
    localStorage.removeItem("user");
  }
}
