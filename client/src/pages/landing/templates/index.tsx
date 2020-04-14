import React from 'react';
import styled from 'styled-components';
import TitleLayout from '../mocules';

interface Props {}

interface State {
  isMouseEnter: boolean;
}

class Landing extends React.Component<Props, State> {
  state: State = {
    isMouseEnter: false,
  };

  onMouseEnterFunc = (): void => {
    this.setState({
      isMouseEnter: true,
    });
  };

  onMouseLeaveFunc = (): void => {
    this.setState({
      isMouseEnter: false,
    });
  };

  render() {
    return (
      <LandingLayout isMouseEnter={this.state.isMouseEnter}>
        <div
          onMouseEnter={() => this.onMouseEnterFunc()}
          onMouseLeave={() => this.onMouseLeaveFunc()}
        >
          <TitleLayout isMouseEnter={this.state.isMouseEnter} />
        </div>
      </LandingLayout>
    );
  }
}

const LandingLayout = styled.div<State>`
  display: flex;
  justify-content: center;
  align-items: center;
  width: 100%;
  height: 100%;
  background: ${(props) => (props.isMouseEnter ? '#1d1f21' : '#fff')};
  opacity: ${(props) => (props.isMouseEnter ? '0.9' : '1')};
`;

export default Landing;
