import { combineReducers } from 'redux';
import vmReducer from './vm';
import accountReducer from './account';

const RootReducer = combineReducers({
  vmReducer,
  accountReducer,
});

export default RootReducer;
