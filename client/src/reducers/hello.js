import {
  HELLO_REQUEST,
  HELLO_SUCCESS,
  HELLO_FAILURE
} from "../actions/hello";

const initialState = {
  message: null,
  helloError: null
};

function initializeState() {
  return Object.assign({}, initialState);
}

export default function hello(state = initializeState(), action = {}) {
  switch (action.type) {
    case HELLO_REQUEST:
      return Object.assign({}, state);
    case HELLO_SUCCESS:
      return Object.assign({}, state, {
          ...state,
          message: action.message,
          helloError: "no errors"
      });
    case HELLO_FAILURE:
      return {
          ...state,
          message: "empty message",
          helloError: "hello error"
      };
    default:
      return state;
  }
}
