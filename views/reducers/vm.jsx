import {
  ACTION_VM_OVERVIEW, ACTION_VM_OFFERINGS,
} from '../common/prophetConstants';

function vmReducer(state = { vmData: [] }, action) {
  switch (action.type) {
    case ACTION_VM_OVERVIEW:
      return {
        vmData: action.vmData,
      };
    default:
      return state;
  }
}

function vmOfferingReducer(state = { offeringData: {} }, action) {
  switch (action.type) {
    case ACTION_VM_OFFERINGS:
      return {
        offeringData: action.offeringData,
      };
    default:
      return state;
  }
}

export { vmReducer, vmOfferingReducer };
