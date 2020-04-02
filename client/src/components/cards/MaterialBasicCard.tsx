import React from 'react';
import styled from 'styled-components';

interface Props {
  imageUrl: string;
  onClick?: () => void;
}

class MaterialBasicCard extends React.Component<Props> {
  render() {
    return (
      <CardLayout>
        <ImageLayout imageUrl={this.props.imageUrl} />
        <TextLayout>
          <HeaderLayout>
            <Title>Title</Title>
            <SubTitle>SubTitle</SubTitle>
          </HeaderLayout>
          <BodyLayout>
            <Body>
              body is bodybody is bodybody is bodybody is bodybody is bodybody
              is bodybody is bodybody is bodybody is bodybody is bodybody is
              body
            </Body>
          </BodyLayout>
        </TextLayout>
      </CardLayout>
    );
  }
}

const CardLayout = styled.div`
  width: 400px;
  border-radius: 4px;
  box-shadow: 0px 3px 1px -2px rgba(0, 0, 0, 0.2),
    0px 2px 2px 0px rgba(0, 0, 0, 0.14), 0px 1px 5px 0px rgba(0, 0, 0, 0.12);
`;

const ImageLayout = styled.div<Props>`
  width: 100%;
  padding-top: 56%;
  position: relative;
  background-image: url(${props => props.imageUrl});
  background-position: center center;
  background-repeat: no-repeat;
  border-top-left-radius: inherit;
  border-top-right-radius: inherit;
`;

const TextLayout = styled.div`
  width: 100%;
  display: flex;
  flex-direction: column;
  background-color: white;
`;

const HeaderLayout = styled.div`
  width: 100%;
  padding-top: 16px;
  padding-left: 16px;
  padding-right: 16px;
`;

const BodyLayout = styled.div`
  width: 100%;
  padding-top: 16px;
  padding-left: 16px;
  padding-right: 16px;
  padding-bottom: 16px;
`;

const Title = styled.div`
  width: 100%;
  font-size: 1.25rem;
  line-height: 2rem;
  font-weight: 500;
  letter-spacing: 0.0125em;
`;

const SubTitle = styled.div`
  width: 100%;
  font-size: 0.875rem;
  line-height: 1.25rem;
  font-weight: 400;
  letter-spacing: 0.0125em;
  opacity: 0.6;
`;

const Body = styled.div`
  width: 100%;
  font-size: 0.875rem;
  line-height: 1.25rem;
  font-weight: 400;
  letter-spacing: 0.0125em;
  opacity: 0.6;
`;

export default MaterialBasicCard;
