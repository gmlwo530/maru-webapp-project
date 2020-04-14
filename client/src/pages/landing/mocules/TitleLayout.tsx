import React from 'react';
import styled from 'styled-components';
import { Title, Count } from '../atoms';

interface Props {
  isMouseEnter: boolean;
}

class TitleLayout extends React.Component<Props> {
  render() {
    return (
      <TitleLayoutDivTag isMouseEnter={this.props.isMouseEnter}>
        <Title />
        <Count />
      </TitleLayoutDivTag>
    );
  }
}

const TitleLayoutDivTag = styled.div<Props>`
  color: ${(props) => (props.isMouseEnter ? '#fff' : '#1d1f21')};
  opacity: ${(props) => (props.isMouseEnter ? '1' : '0.9')};
  text-align: center;
  cursor: pointer;
  -webkit-touch-callout: none;
  -webkit-user-select: none;
  -khtml-user-select: none;
  -moz-user-select: none;
  -ms-user-select: none;
  user-select: none;
`;

export default TitleLayout;
