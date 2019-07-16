import React, { Component } from 'react';
import { Container } from './styles';

import Sidebar from '../../components/Sidebar';
import Navbar from '../../components/Navbar';
import { } from '../../services/documentService';
import { MessageComponent } from '../../components/message';
import ReactTable from 'react-table';
import "react-table/react-table.css";

class Blacklist extends Component {
  state = {
    message: '',
    messageClass: 'hidden',
    documentList: []
  };

  componentDidMount() {
  
  }

  render() {

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
            <div>
                <h4>Blacklist</h4>
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
export default Blacklist;
