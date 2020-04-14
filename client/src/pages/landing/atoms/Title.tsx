import React from 'react';
import styled from 'styled-components';

interface Props {}

class Title extends React.Component<Props> {
  render() {
    return <TitlePTag>Maru Web App Project</TitlePTag>;
  }
}

const TitlePTag = styled.p`
  font-size: 4rem;
  font-weight: 800;
  margin: 0;
`;

export default Title;
