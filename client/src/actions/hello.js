import { callApi, loadIdToken } from "../utils/apiUtils";

export const HELLO_REQUEST = "HELLO_REQUEST";
export const HELLO_SUCCESS = "HELLO_SUCCESS";
export const HELLO_FAILURE = "HELLO_FAILURE";

function helloRequest() {
  return {
    type: HELLO_REQUEST
  };
}

function helloSuccess(data) {
  return {
    type: HELLO_SUCCESS,
    message: data.message
  };
}

function helloFailure(error) {
  return {
    type: HELLO_FAILURE,
      helloError: error
  };
}

export function getHelloMessage() {
  const idToken = loadIdToken();
  const config = {
    method: "get",
    headers: {
      Accept: "application/json",
      "Content-Type": "application/json",
      Authorization: `Bearer ${idToken}`
    }
  };

  return callApi(
    "/auth/hello",
    config,
    helloRequest(),
    helloSuccess,
    helloFailure
  );
}
