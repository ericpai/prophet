import React, { Component } from 'react';
import { render } from 'react-dom';
import { Container } from 'semantic-ui-react';
import VMIndex from './components/vm/VMIndex';
import ProphetMenu from './components/ProphetMenu';
import { Provider } from 'react-redux';
import store from './common/store';


const MOUNT_NODE = document.getElementById('root');


class App extends Component {

  render() {
    return <Container>
      <Provider store={store}>
        <ProphetMenu />
      </Provider>

      <Provider store={store}>
        <VMIndex />
      </Provider>
    </Container>
  }
}
render(<App />, MOUNT_NODE);
