import React, { Component } from 'react';
import InputMask from 'react-input-mask';
import { Container } from './styles';

import Sidebar from '../../components/Sidebar';
import Navbar from '../../components/Navbar';
import { saveDocument, getAllDocuments } from '../../services/documentService';
import { MessageComponent } from '../../components/message';
import ReactTable from 'react-table';
import "react-table/react-table.css";


class Document extends Component {
  state = {
    document: '',
    message: '',
    messageClass: 'hidden',
    selectedOption: 'CPF',
    documentList: []
  };

  componentDidMount() {
    getAllDocuments().then((response) => {
      this.setState({documentList: response.data})
    });
  }

  handleDocumentChange = (e) => {
    this.setState({document: e.target.value})
  }

  handleSubmit = (e) => {
    this.setState({ message: '', messageClass: ''});
    e.preventDefault();
    saveDocument(this.state.document, this.state.selectedOption).then(() => {
      this.setState({ message: 'Cadastro realizado com sucesso', messageClass: 'success-message'});
      getAllDocuments().then((response) => {
        this.setState({documentList: response.data})
      });
    }).catch((error) => {
      this.setState({ message: error.response.data, messageClass: 'error-message'});
    });
  }

  handleTipoChange = (e) => {
    this.setState({selectedOption: e.target.value, document: ''})
  }

  moveToBlackList = (row) => {
    console.log('teste', row)
  }

  render() {
    const {selectedOption} = this.state

    const columns = [{
        Header:'Number',
        accessor: 'number',
        sortMethod: (a, b) => {
         return a > b ? 1 : -1;
        }
      },
      {
        Header:'Type',
        accessor: 'type',
        sortMethod: (a, b) => {
          return a > b ? 1 : -1;
        },
        filterMethod: (filter, row) => {
          if (filter.value === "all") {
            return true;
          }
          if (filter.value === 'CPF') {
            return row.type === 'CPF';
          }
          return row.type === 'CNPJ';
        },
        Filter: ({ filter, onChange }) =>
          <select
            onChange={event => onChange(event.target.value)}
            style={{ width: "100%" }}
            value={filter ? filter.value : "all"}>
              <option value='all'>Show All</option>
              <option value='CPF'>CPF</option>
              <option value='CNPJ'>CNPJ</option>
          </select>
      },
      {
        Header: 'Create Date',
        accessor: 'createdate',
        sortMethod: (a, b) => {
          return a > b ? 1 : -1;
        }
      },
      {
        Header: 'Actions',
        accessor: 'actions',
        filterable: false,
        width: 150,
        style: {'textAlign': 'center'},
        Cell: row => (
          <div>
            <button title='Blacklist' className="btn btn-dark fa fa-list action-button" onClick={this.moveToBlackList(row.original.id)}/>
            <button title='Remover' className="btn btn-danger fa fa-trash action-button" />
          </div>  
         
        )
      }
    ]

    return (
      <Container>
        <div className="d-flex" id="wrapper">
          <Sidebar />
          <div id="page-content-wrapper">
            <Navbar />
            <MessageComponent text = {this.state.message} classe = {this.state.messageClass}/>
            <br/>
            <div className="container-fluid">
            <div>
                <h4>Cadastrar CPF/CNPJ</h4>
                <form onSubmit={this.handleSubmit}>
                  <div className="form-check">
                    <label className="input-tipo">
                      <input 
                        type="radio"
                        name="tipo"
                        value='CPF'
                        checked={this.state.selectedOption === 'CPF'}
                        onChange={this.handleTipoChange}
                        className="form-check-input"
                      />
                      CPF
                    </label>

                    <label className="input-tipo">
                      <input 
                        type="radio"
                        name="tipo"
                        value='CNPJ'
                        checked={this.state.selectedOption === 'CNPJ'}
                        onChange={this.handleTipoChange}
                        className="form-check-input"
                      />
                      CNPJ
                    </label>  
                  </div>
                  {selectedOption === 'CPF' ? (
                    <div className="form-group">
                      <label>CPF:  </label>
                      <InputMask mask="999.999.999-99" value={this.state.document} onChange={this.handleDocumentChange} required/>
                    </div>
                  ): (
                    <div className="form-group">
                      <label>CNPJ:  </label>
                      <InputMask mask="99.999.999/9999-99" value={this.state.document} onChange={this.handleDocumentChange} required/>
                   </div>
                  )}
                    <div className="form-group">
                        <input type="submit" value="Cadastrar" className="btn btn-primary"/>
                    </div>
                </form>
            </div>
              <div>
                <ReactTable data={this.state.documentList} columns={columns}  defaultPageSize={10}
                  className="-striped -highlight" filterable />
              </div>
            </div>
          </div>
        </div>
      </Container>
    );
  }
}
export default Document;
