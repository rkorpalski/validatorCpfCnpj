import React from 'react';
import { shallow } from 'enzyme';
import Blacklist from '../pages/Blacklist';

// Testa o component puro, sem redenrizar os filhos
it('renders without crashing', () => {
  const component = shallow(<Blacklist />);

  expect(component).toMatchSnapshot();
});
