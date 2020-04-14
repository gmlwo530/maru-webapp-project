import React from 'react';
import Landing from './pages/landing';
import styled from 'styled-components';

function App() {
  return (
    <Container>
      <Landing />
    </Container>
  );
}

const Container = styled.div`
  width: 100%;
  height: 100%;
`;

export default App;
