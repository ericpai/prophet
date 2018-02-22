import {
  ACTION_ACCOUNT_LIST, ACTION_SHOW_ACCOUNT,
} from '../common/prophetConstants';

function accountReducer(
  state = { accountData: {}, account: null, provider: null }, action) {
  switch (action.type) {
    case ACTION_ACCOUNT_LIST:
      return {
        accountData: action.accountData,
        account: state.account,
        provider: state.provider,
      };
    case ACTION_SHOW_ACCOUNT:
      return {
        accountData: state.accountData,
        account: action.account,
        provider: action.provider,
      };
    default:
      return state;
  }
}

export default accountReducer;
