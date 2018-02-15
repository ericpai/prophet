import React, { Component } from 'react';
import Plot from 'react-plotly.js';
import { Segment, Divider, Grid, Statistic, Container, Header } from 'semantic-ui-react';
import { connect } from 'react-redux';


class InstanceOverviewComponent extends Component {

  render() {
    let values = [];
    let labels = [];
    let total = 0;

    this.props.vmData.map(function (v, i) {
      values.push(v['count']);
      labels.push(v['type']);
      total += v['count'];
    });
    const pieData = [{
      values: values,
      labels: labels,
      type: 'pie',
    }];
    const barData = [{
      y: values,
      x: labels,
      text: values,
      textposition: 'auto',
      type: 'bar',
    }];
    const pieLayout = {
      width: 350,
      height: 300,
      autosize: true,
      title: '<b>类型占比</b>',
      titlefont: {
        family: `'Lato', 'Helvetica Neue', Arial, Helvetica, sans-serif`,
        size: 14,
      },
    };
    const barLayout = {
      width: 350,
      height: 300,
      autosize: true,
      title: '<b>类型总量</b>',
      titlefont: {
        family: `'Lato', 'Helvetica Neue', Arial, Helvetica, sans-serif`,
        size: 14,
      },
    };

    return <Container style={{ marginTop: '7em' }} textAlign={'center'}>
      <Header as='h1'>服务器</Header>
      <Divider fitted />
      <Grid columns={3} divided verticalAlign={'middle'} textAlign={'center'}>
        <Grid.Row>
          <Grid.Column verticalAlign={'middle'} textAlign={'center'}
            className={'ds_cell'}>
            <Header as='h5'>使用总量</Header>
            <Statistic size={'huge'}>
              <Statistic.Value>{total}</Statistic.Value>
              <Statistic.Label>instances</Statistic.Label>
            </Statistic>
          </Grid.Column>
          <Grid.Column verticalAlign={'middle'} textAlign={'center'}
            className={'ds_cell'}>
            <Plot
              style={{ width: '100%', height: '100%' }}
              useResizeHandler={true}
              data={barData}
              layout={barLayout}
              config={{
                displayModeBar: false, staticPlot: true,
                editable: false
              }}
            />
          </Grid.Column>
          <Grid.Column verticalAlign={'middle'} textAlign={'center'}
            className={'ds_cell'}>
            <Plot
              style={{ width: '100%', height: '100%' }}
              useResizeHandler={true}
              data={pieData}
              layout={pieLayout}
              config={{
                displayModeBar: false, staticPlot: true,
                editable: false
              }}
            />
          </Grid.Column>
        </Grid.Row>
      </Grid>
    </Container >


  }
}

function mapStateToProps(state) {
  return {
    vmData: state.vmReducer.vmData,
  };
}

function mapDispatchToProps(dispatch) {
  return {};
}

const InstanceOverview = connect(
  mapStateToProps, mapDispatchToProps)(InstanceOverviewComponent);

export default InstanceOverview;
