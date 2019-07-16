import React from 'react';
import { Link } from 'react-router-dom';

import { Container } from './styles';

const Sidebar = () => (
  <Container>
    <div className="border-right" id="sidebar-wrapper">
      <div className="sidebar-heading">
        <p>CPF/CNPJ</p>
      </div>
      <div className="list-group list-group-flush">
        <ul>
          <Link to="/">
            <li className="list-group-item list-group-item-action bg-dark-blue">
              <i className="fa fa-id-card fa-lg" /> Consulta CPF/CNPJ
            </li>
          </Link>
          <Link to="/blacklist">
            <li className="list-group-item list-group-item-action bg-dark-blue">
              <i className="fa fa-list fa-lg" /> Blacklist
            </li>
          </Link>
        </ul>
      </div>
    </div>
  </Container>
);

export default Sidebar;
