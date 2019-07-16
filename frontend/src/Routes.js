import React from 'react';
import { BrowserRouter, Switch, Route } from 'react-router-dom';

import Document from './pages/Document';
import Blacklist from './pages/Blacklist';

const Routes = () => (
  <BrowserRouter>
    <Switch>
      <Route path="/" exact component={Document} />
      <Route path="/blacklist" exact component={Blacklist} />
    </Switch>
  </BrowserRouter>
);

export default Routes;
