import { combineReducers } from 'redux';
import { vmReducer, vmOfferingReducer } from './vm';
import accountReducer from './account';

const RootReducer = combineReducers({
  vmReducer,
  vmOfferingReducer,
  accountReducer,
});

export default RootReducer;
