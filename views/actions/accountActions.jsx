import $ from 'jquery';
import {
  ACTION_ACCOUNT_LIST, ACTION_SHOW_ACCOUNT
} from '../common/prophetConstants';
import { getVMOverview, getVMOfferings } from './vmActions';

export function listAccounts() {
  return function (dispatch) {
    $.ajax(
      `/api/accounts/`,
      {
        method: 'GET',
        dataType: 'json',
        success: function (data) {
          dispatch(renderAccountAction(data));
        },
        error: function (xhr, status, err) {
          dispatch(renderAccountAction({}));
        },
      },
    );
  };
}

export function dispatchAll(account, provider) {
  return function (dispatch) {
    dispatch(getVMOverview(account, provider));
    dispatch(showAccountAction(account, provider));
    dispatch(getVMOfferings(account, provider));
  }
}

function renderAccountAction(data) {
  return {
    type: ACTION_ACCOUNT_LIST,
    accountData: data,
  };
}

function showAccountAction(account, provider) {
  return {
    type: ACTION_SHOW_ACCOUNT,
    account: account,
    provider: provider,
  };
}
