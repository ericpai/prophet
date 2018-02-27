import $ from 'jquery';
import {
  ACTION_VM_OVERVIEW, ACTION_VM_OFFERINGS,
} from '../common/prophetConstants';

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

export function getVMOfferings(account, provider) {
  return function (dispatch) {
    $.ajax(
      `/api/vm/offerings?account=${account}&provider=${provider}`,
      {
        method: 'GET',
        dataType: 'json',
        success: function (data) {
          dispatch(renderVMOfferingsAction(data));
        },
        error: function (xhr, status, err) {
          dispatch(renderVMOfferingsAction([]));
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

function renderVMOfferingsAction(data) {
  return {
    type: ACTION_VM_OFFERINGS,
    offeringData: data,
  };
}
