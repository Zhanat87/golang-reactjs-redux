import React, { Component } from "react";
import PropTypes from "prop-types";
import { connect } from "react-redux";

import Hello from "../../components/hello/Hello";

import {
    getHelloMessage
} from "../../actions/hello";

class HelloPage extends Component {
    constructor(props) {
        super(props);
    }

    componentDidMount() {
        const { dispatch } = this.props;
        dispatch(getHelloMessage());
    }

    render() {
        const { message, helloError } = this.props;

        const hello = {message, helloError};

        return (
            <div>
                <Hello hello={hello} />
            </div>
        );
    }
}

HelloPage.propTypes = {
    message: PropTypes.string.isRequired,
    dispatch: PropTypes.func.isRequired,
    helloError: PropTypes.string.isRequired
};

function mapStateToProps(state) {
    const { message, helloError } = state.hello;
    return {
        message: message,
        helloError: helloError
    };
}

export default connect(mapStateToProps)(HelloPage);
