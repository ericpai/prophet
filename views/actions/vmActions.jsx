import $ from 'jquery';
import { ACTION_VM_OVERVIEW } from '../common/prophetConstants';

export function getVMOverview(account, provider) {
  return function (dispatch) {
    $.ajax(
      `/api/vm/overview?account=${account}&provider=${provider}`,
      {
        method: 'GET',
        dataType: 'json',
        success: function (data) {
          dispatch(renderVMOverviewAction(data));
        },
        error: function (xhr, status, err) {
          dispatch(renderVMOverviewAction([]));
        },
      },
    );
  };
}

function renderVMOverviewAction(data) {
  return {
    type: ACTION_VM_OVERVIEW,
    vmData: data,
  };
}
