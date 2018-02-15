import { ACTION_VM_OVERVIEW } from '../common/prophetConstants';

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

export default vmReducer;
