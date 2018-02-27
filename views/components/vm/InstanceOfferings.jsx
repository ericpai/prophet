import React, { Component } from 'react';
import { connect } from 'react-redux';
import { Table, Card } from 'semantic-ui-react';

class InstanceOfferingsComponent extends Component {
  render() {
    return <Card.Group itemsPerRow={2}>
      <Card fluid>
        <Card.Content>
          <Card.Header>付费类型</Card.Header>
          <Card.Description textAlign={'center'}>
            {this.renderOfferingTableContent()}
          </Card.Description>
        </Card.Content>
      </Card>
    </Card.Group>;
  }

  renderOfferingTableContent() {
    const offeringData = this.props.offeringData;
    if (!('offerings' in offeringData && 'offering_types' in offeringData)) {
      return <Card.Content>无数据</Card.Content>
    }
    const offerings = Object.entries(offeringData['offerings']);
    return <Table definition>
      <Table.Header>
        <Table.Row>
          <Table.HeaderCell></Table.HeaderCell>
          {
            offeringData['offering_types'].map((name, i) =>
              <Table.HeaderCell key={i}>{name}</Table.HeaderCell>)
          }
        </Table.Row>
      </Table.Header>
      <Table.Body>
        {
          offerings.map((value, i) => <Table.Row key={i}>
            <Table.Cell>{value[1]['type']}</Table.Cell>
            {
              value[1]['counts'].map((count, j) => <Table.Cell key={j}>
                {count}</Table.Cell>)
            }
          </Table.Row>)
        }

      </Table.Body>
    </Table>;
  }
}

function mapStateToProps(state) {
  return {
    offeringData: state.vmOfferingReducer.offeringData,
  };
}

function mapDispatchToProps(dispatch) {
  return {};
}

const InstanceOfferings = connect(
  mapStateToProps, mapDispatchToProps)(InstanceOfferingsComponent);

export default InstanceOfferings;
