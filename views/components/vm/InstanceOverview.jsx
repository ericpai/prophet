import React, { Component } from 'react';
import { randomColor } from 'randomcolor';
import {
  Tooltip, Legend, ResponsiveContainer, PieChart, Pie, Cell,
} from 'recharts';
import { Statistic, Card } from 'semantic-ui-react';
import { connect } from 'react-redux';


class InstanceOverviewComponent extends Component {

  render() {
    return <Card.Group itemsPerRow={2}>
      {this.renderInstances()}
      {this.renderTypes()}
    </Card.Group>;
  }

  renderInstances() {
    let total = 0;
    this.props.vmData.map(function (v, i) {
      total += v['count'];
    });
    return <Card fluid>
      <Card.Content style={{ display: 'flex', flexDirection: 'column' }}>
        <Card.Header>使用总量</Card.Header>
        <Card.Description textAlign={'center'}
          style={{
            display: 'flex', flex: 1,
            justifyContent: 'center', alignItems: 'center',
          }}>
          {
            function (vmData, total) {
              if (vmData.length == 0) {
                return <Card.Content>无数据</Card.Content>;
              }
              return <Statistic size={'huge'}>
                <Statistic.Value>{total}</Statistic.Value>
                <Statistic.Label>instances</Statistic.Label>
              </Statistic>;
            }(this.props.vmData, total)
          }

        </Card.Description>
      </Card.Content>
    </Card>;
  }

  renderTypes() {
    let today = new Date();
    let rc = randomColor({
      luminosity: 'dark',
      count: this.props.vmData.length,
      seed: today.getDate(),
    });
    return <Card fluid>
      <Card.Content>
        <Card.Header>类型占比</Card.Header>

        <Card.Description textAlign={'center'}>
          {
            function (vmData, rc) {
              if (vmData.length == 0) {
                return <Card.Content>无数据</Card.Content>;
              }
              return <ResponsiveContainer
                width={'100%'} height={300} style={{ margin: 'auto' }}>
                <PieChart>
                  <Pie data={vmData} dataKey={'count'}
                    nameKey={'type'} label>
                    {
                      vmData.map((entry, index) =>
                        <Cell key={`slice-${index}`} fill={rc[index]} />
                      )
                    }
                  </Pie>
                  <Tooltip />
                  <Legend />
                </PieChart>
              </ResponsiveContainer>;
            }(this.props.vmData, rc)
          }

        </Card.Description>
      </Card.Content>
    </Card>;
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
