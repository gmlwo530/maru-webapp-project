import React from 'react';
import { action } from '@storybook/addon-actions';
import { MaterialDefaultButton } from '../components/buttons';

export default {
  title: 'Button',
  component: MaterialDefaultButton,
};

export const MaterialDefault = () => (
  <MaterialDefaultButton onClick={action('clicked')} text="Buttonnnnnnnn" />
);
