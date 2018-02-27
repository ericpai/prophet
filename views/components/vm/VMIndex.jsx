import React, { Component } from 'react';
import InstanceOverview from './InstanceOverview';
import InstanceOfferings from './InstanceOfferings';
import { Header, Divider, Container } from 'semantic-ui-react';

export default class VMIndex extends Component {
  render() {
    return <Container style={{ marginTop: '7em' }} textAlign={'center'}>
      <Header as={'h1'} textAlign={'left'}>服务器</Header>
      <Divider fitted style={{ marginBottom: '2em' }} />
      <InstanceOverview />
      <InstanceOfferings />
    </Container >;
  }
}
