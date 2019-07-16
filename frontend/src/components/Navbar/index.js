import React from 'react';

import { Container } from './styles';

const Navbar = () => (
  <Container>
    <nav className="navbar navbar-expand-lg  border-bottom">
      <div className="toggle-btn">
        <i className="fa fa-bars fa-lg" id="menu-toggle" />
      </div>

      <button
        className="navbar-toggler"
        type="button"
        data-toggle="collapse"
        data-target="#navbarSupportedContent"
        aria-controls="navbarSupportedContent"
        aria-expanded="false"
        aria-label="Toggle navigation"
      >
        <span className="navbar-toggler-icon" />
      </button>
    </nav>
  </Container>
);

export default Navbar;
