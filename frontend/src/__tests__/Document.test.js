import React from 'react';
import { shallow } from 'enzyme';
import Document from '../pages/Document';

// Testa o component puro, sem redenrizar os filhos
it('renders without crashing', () => {
  const component = shallow(<Document />);

  expect(component).toMatchSnapshot();
});
