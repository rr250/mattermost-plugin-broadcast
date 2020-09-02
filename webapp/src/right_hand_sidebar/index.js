import {connect} from 'react-redux';
import {getUsers} from 'mattermost-redux/selectors/entities/common';

import RHSView from './rhs_view';

const mapStateToProps = (state) => ({
    team: getUsers(state),
});

export default connect(mapStateToProps)(RHSView);

