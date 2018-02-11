import React, { Component } from 'react';
import { render } from 'react-dom';
import Plot from 'react-plotly.js';
import { Container, Header } from 'semantic-ui-react';

const MOUNT_NODE = document.getElementById('root');


class App extends Component {

  render() {
    const data = [{
      values: [19, 26, 55],
      labels: ['Residential', 'Non-Residential', 'Utility'],
      type: 'pie',
    }];

    const layout = {
      height: 400,
      width: 500,
    };

    return <Container>
      <Header as="h1">Hello world!</Header>
      <Plot layout={layout} data={data} config={{ displayModeBar: false }} />
    </Container>;
  }
}
render(<App />, MOUNT_NODE);
