import { combineReducers } from 'redux';
import { reducer as authReducer } from './auth';
import { reducer as featureReducer } from './feature';
import { reducer as formReducer } from 'redux-form';
import { reducer as overviewReducer} from './overview.reducer';
const rootReducer = combineReducers({
    form: formReducer,
    auth: authReducer,
    features: featureReducer,
    overview: overviewReducer
});

export default rootReducer;