import React, { Component } from 'react';
import InputMask from 'react-input-mask';
import { Container } from './styles';

import Sidebar from '../../components/Sidebar';
import Navbar from '../../components/Navbar';
import { saveDocument, getAllDocuments, moveToBlacklist, deleteDocument, findDocument } from '../../services/documentService';
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
    this.findAllDocuments();
  }

  handleDocumentChange = (e) => {
    this.setState({document: e.target.value});
  }

  handleSubmit = (e) => {
    this.setState({ message: '', messageClass: ''});
    e.preventDefault();
    saveDocument(this.state.document, this.state.selectedOption).then(() => {
      this.findAllDocuments();
    }).catch((error) => {
      this.setState({ message: error.response.data, messageClass: 'error-message'});
    });
  }

  handleTipoChange = (e) => {
    this.setState({selectedOption: e.target.value, document: ''});
  }

  handleBlacklist = (documentId) => {
    moveToBlacklist(documentId).then(() => {
      this.findAllDocuments();
    }).catch((error) => {
      this.setState({ message: 'Ocorreu um erro ao mover o documento para a blacklist', messageClass: 'error-message'});
    });
  }

  handleDelete = (documentId) => {
    deleteDocument(documentId).then(() => {
      this.findAllDocuments();
    }).catch((error) => {
      this.setState({ message: "Ocorreu um erro ao remover o documento", messageClass: 'error-message'});
    });
  }

  handleFind = () => {
    const {document, selectedOption} = this.state
    if(document === ''){
      this.findAllDocuments();
    } else {
      let regexp
      if (selectedOption === 'CPF') {
        regexp = new RegExp(/^\d{3}\.?\d{3}\.?\d{3}-?\d{2}$/);
      } else {
        regexp = new RegExp(/^\d{2}\.?\d{3}\.?\d{3}\/?(:?\d{3}[1-9]|\d{2}[1-9]\d|\d[1-9]\d{2}|[1-9]\d{3})-?\d{2}$/);
      }
      const isValid = regexp.test(document)
      if(isValid){
        findDocument(document).then((response) => {
          this.setState({documentList: response.data});
        });
      } else {
        this.setState({ message: 'O documento está incompleto. Por favor forneça um documento válido', messageClass: 'error-message'});
      }
    }
  }

  findAllDocuments = () => {
    this.setState({ message: '', messageClass: 'hidden'});
    getAllDocuments().then((response) => {
      this.setState({documentList: response.data});
    });
  }

  handleClean = () => {
    this.setState({document: ''});
    this.findAllDocuments();
  }

  render() {
    const {selectedOption} = this.state

    const columns = [{
        Header:'CPF/CNPJ',
        accessor: 'number',
        sortMethod: (a, b) => {
         return a > b ? 1 : -1;
        }
      },
      {
        Header:'Tipo',
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
              <option value='all'>Mostrar tudo</option>
              <option value='CPF'>CPF</option>
              <option value='CNPJ'>CNPJ</option>
          </select>
      },
      {
        Header: 'Data de criação',
        accessor: 'createdate',
        sortMethod: (a, b) => {
          return a > b ? 1 : -1;
        }
      },
      {
        Header: 'Ações',
        accessor: 'actions',
        filterable: false,
        width: 150,
        style: {'textAlign': 'center'},
        Cell: row => (
          <div>
            <button title='Blacklist' className="btn btn-dark fa fa-list action-button" onClick={() => this.handleBlacklist(row.original.id)}/>
            <button title='Remover' className="btn btn-danger fa fa-trash action-button" onClick={() => this.handleDelete(row.original.id)} />
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
            <div className= "form_main">
                <h4 className="heading">Consultar CPF/CNPJ</h4>
                <form onSubmit={this.handleSubmit}>
                  <div >
                  <label>Tipo:  </label>
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
                        <input type="submit" value="Cadastrar" className="btn btn-dark button"/>
                        <input type="button" value="Consultar" className="btn btn-dark button" onClick={() => this.handleFind()}/>
                        <input type="button" value="Limpar" className="btn btn-dark button" onClick={() => this.handleClean()}/>
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
