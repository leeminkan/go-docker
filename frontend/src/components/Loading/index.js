import React, { Component } from "react";
import LoadingOverlay from "react-loading-overlay";
import * as command from "../../constants/ui";
import { observer } from "../../observer";

export default class Loading extends Component {
  constructor(props) {
    super(props);
    this.state = {
      apiCallStack: 0,
    };
  }

  componentDidMount() {
    observer.subscribe(command.SHOW_LOADING, {
      id: "loading",
      update: this.update,
    });
  }

  componentWillUnmount() {
    observer.unsubscribe(command.SHOW_LOADING, {
      id: "loading",
      update: this.update,
    });
  }

  update = (cmd, data) => {
    if (cmd === command.SHOW_LOADING) {
      if (data) {
        this.setState({ apiCallStack: this.state.apiCallStack + 1 });
      } else {
        this.setState({ apiCallStack: this.state.apiCallStack - 1 });
      }
    }
  };

  render() {
    return (
      <LoadingOverlay
        active={this.state.apiCallStack > 0}
        spinner
        styles={{
          overlay: (base) => ({
            ...base,
            height: "100vh",
          }),
        }}
      >
        {this.props.children}
      </LoadingOverlay>
    );
  }
}
