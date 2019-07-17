import styled from 'styled-components';

export const Container = styled.div`
  #page-content-wrapper {
    min-width: 100vw;
  }

  #wrapper.toggled #sidebar-wrapper {
    margin-left: 0;
  }

  #page-content-wrapper .input-tipo {
    padding-right: 1rem;
    padding-left: 1.5rem;
  }
  #page-content-wrapper .action-button {
    margin-right: 0.5rem;
  }

  .form_main h4 {
      font-family: roboto;
      font-size: 20px;
      font-weight: 300;
      margin-bottom: 15px;
      margin-top: 20px;
      text-transform: uppercase;
  }
  .heading {
      border-bottom: 1px solid #fcab0e;
      padding-bottom: 9px;
      position: relative;
  }
  .heading span {
      background: #9e6600 none repeat scroll 0 0;
      bottom: -2px;
      height: 3px;
      left: 0;
      position: absolute;
      width: 75px;
  }   

  .button {
    margin-right: 1rem;
  }

  @media (min-width: 768px) {
    #sidebar-wrapper {
      margin-left: 0;
    }

    #page-content-wrapper {
      min-width: 0;
      width: 100%;
    }

    #wrapper.toggled #sidebar-wrapper {
      margin-left: -15rem;
    }
  }
`;
