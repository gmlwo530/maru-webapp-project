import React from 'react';
import styled from 'styled-components';

interface Props {
  text: string;
  onClick?: () => void;
}

class MaterialDefaultButton extends React.Component<Props> {
  render() {
    return (
      <Button onClick={this.props.onClick}>
        <span>{this.props.text}</span>
      </Button>
    );
  }
}

const Button = styled.button`
  color: #fff;
  background-color: #6200ee;
  padding: 8px 16px;
  min-width: 64px;
  font-size: 0.875rem;
  letter-spacing: calc(0.875em / 10);
  font-weight: 500;
  border-radius: 4px;
  border: none;
  outline: none;
  display: inline-flex;
  justify-content: center;
  align-items: center;
  vertical-align: middle;
  line-height: inherit;
  box-shadow: 0px 3px 1px -2px rgba(0, 0, 0, 0.2),
    0px 2px 2px 0px rgba(0, 0, 0, 0.14), 0px 1px 5px 0px rgba(0, 0, 0, 0.12);
`;

export default MaterialDefaultButton;
