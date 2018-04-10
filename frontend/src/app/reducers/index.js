import { combineReducers } from 'redux';
import { reducer as authReducer } from './auth.reducer';
import { reducer as featureReducer } from './feature';
import { reducer as formReducer } from 'redux-form';
import { reducer as overviewReducer } from './overview.reducer';
import { reducer as sendReducer } from './send.reducer';
const rootReducer = combineReducers({
    form: formReducer,
    auth: authReducer,
    features: featureReducer,
    overview: overviewReducer,
    send: sendReducer
});

export default rootReducer;