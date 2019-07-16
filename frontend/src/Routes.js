import React from 'react';
import { BrowserRouter, Switch, Route } from 'react-router-dom';

import Document from './pages/Document';

const Routes = () => (
  <BrowserRouter>
    <Switch>
      <Route path="/" exact component={Document} />
    </Switch>
  </BrowserRouter>
);

export default Routes;
