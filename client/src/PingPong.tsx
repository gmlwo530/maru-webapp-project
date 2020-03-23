import React from 'react';
import axios from 'axios';

interface State {
  pong?: string;
}

class PingPong extends React.Component<State> {
  state: State = {
    pong: 'pending',
  };

  componentWillMount() {
    axios
      .get('api/ping')
      .then(response => {
        this.setState(() => {
          return { pong: response.data.message };
        });
      })
      .catch(function(error) {
        console.log(error);
      });
  }

  render() {
    return <h1>Ping {this.state.pong}</h1>;
  }
}

export default PingPong;
