import React from 'react';
import { shallow } from 'enzyme';
import Message from '../components/message';

describe('Message', () => {
  // testa o component puro, sem renderizar os filhos
  it('should render correctly with no props', () => {
    const component = shallow(<Message />);
    expect(component).toMatchSnapshot();
  });
});
