import React from 'react';
import { action } from '@storybook/addon-actions';
import { MaterialBasicCard } from '../components/cards';

export default {
  title: 'Card',
  component: MaterialBasicCard,
};

export const MaterialDefault = () => (
  <MaterialBasicCard
    onClick={action('clicked')}
    imageUrl="https://dapp.dblog.org/img/default.jpg"
  />
);
