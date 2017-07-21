import React from "react";
import PropTypes from "prop-types";

import "./hello.css";

const Hello = ({ hello }) => {
  const { message, helloError } = hello;

  return (
    <div>
      <div className="hello">
          {message}
      </div>
      <div className="helloError">
          {helloError}
      </div>
    </div>
  );
};

Hello.propTypes = {
  hello: PropTypes.shape({
      message: PropTypes.string.isRequired,
      helloError: PropTypes.string.isRequired,
  }).isRequired
};

export default Hello;
