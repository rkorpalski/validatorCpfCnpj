import styled from 'styled-components';

export const Container = styled.div`
  #sidebar-wrapper {
    min-height: 100%;
    margin-left: -15rem;
    -webkit-transition: margin 0.25s ease-out;
    -moz-transition: margin 0.25s ease-out;
    -o-transition: margin 0.25s ease-out;
    transition: margin 0.25s ease-out;
    background-color: #2b2d3d;
    height: 100vh;
    a {
      text-decoration: none;
    }
  }

  #sidebar-wrapper .sidebar-heading {
    font-size: 1.5rem;
    color: #FFFFFF;
    font-weight: bold;
    display:flex;
    justify-content:center;
    align-items:center;
  }

  #sidebar-wrapper .list-group {
    width: 15rem;
    li {
      padding: 18px !important;
      color: #fff;
      font-weight: 600;
      text-transform: uppercase;
      font-size: 15px !important;
      cursor: pointer;

      &:hover {
        background-color: #27293a !important;
      }
      i {
        margin-right: 10px;
      }
    }
  }
  #wrapper.toggled #sidebar-wrapper {
    margin-left: 0;
  }
  @media (min-width: 768px) {
    #sidebar-wrapper {
      margin-left: 0;
    }

    #wrapper.toggled #sidebar-wrapper {
      margin-left: -15rem;
    }
  }
`;
