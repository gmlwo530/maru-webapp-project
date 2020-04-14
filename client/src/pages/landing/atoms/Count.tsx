import React from 'react';
import styled from 'styled-components';

interface Props {}

class Count extends React.Component<Props> {
  render() {
    return <CountPTag>10</CountPTag>;
  }
}

const CountPTag = styled.p`
  font-size: 2rem;
  font-weight: 600;
  letter-spacing: calc(2em / 10);
  margin: 29px 0;
`;

export default Count;
