import React, { Component } from 'react';
import { Menu, Container, Dropdown } from 'semantic-ui-react';
import { listAccounts, dispatchAll } from '../actions/accountActions';
import { connect } from 'react-redux';


class ProphetMenuComponent extends Component {

  componentDidMount() {
    this.props.getAccounts();
  }

  render() {
    const fetchAll = this.props.fetchAll.bind(this);
    const currentAccount = this.props.account;
    const currentProvider = this.props.provider;
    return <Menu fixed={'top'} inverted>
      <Container>
        <Menu.Item as={'a'} header>
          Prophet
        </Menu.Item>

        <Dropdown item simple text={'账号和运营商'}>

          <Dropdown.Menu>{
            Object.entries(this.props.accountData).map(function (v, i) {
              return <Dropdown.Item key={i}>
                <i className={'dropdown icon'} />
                <span className={'text'}>{v[0]}</span>
                <Dropdown.Menu>{
                  v[1].map(function (name, j) {
                    return <Dropdown.Item
                      key={j}
                      active={currentAccount == name && currentProvider == v[0]}
                      onClick={() => fetchAll(name, v[0])}>
                      <span className={'text'}>{name}</span>
                    </Dropdown.Item>;
                  })

                }
                </Dropdown.Menu>
              </Dropdown.Item>;
            })
          }

          </Dropdown.Menu>
        </Dropdown>
        {
          function () {
            if (currentAccount != null && currentProvider != null) {
              return <Menu.Menu position={'right'}>
                <Menu.Item name={
                  `当前账号：${currentAccount}  运营商：${currentProvider}`} />
              </Menu.Menu>;
            }
            return <div></div>;
          }()
        }

      </Container>
    </Menu>;
  }
}

function mapStateToProps(state) {
  return {
    account: state.accountReducer.account,
    provider: state.accountReducer.provider,
    accountData: state.accountReducer.accountData,
  };
}

function mapDispatchToProps(dispatch) {
  return {
    getAccounts: () => dispatch(listAccounts()),
    fetchAll: (account, provider) => dispatch(dispatchAll(account, provider)),
  };
}

const ProphetMenu = connect(
  mapStateToProps, mapDispatchToProps)(ProphetMenuComponent);


export default ProphetMenu;
