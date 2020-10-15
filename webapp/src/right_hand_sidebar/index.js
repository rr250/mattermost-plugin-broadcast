import {connect} from 'react-redux';
import {getUsers} from 'mattermost-redux/selectors/entities/common';
import {getAllChannels} from 'mattermost-redux/selectors/entities/channels';

import RHSView from './rhs_view';

const mapStateToProps = (state) => ({
    team: getUsers(state),
    channels: getAllChannels(state),
});

export default connect(mapStateToProps)(RHSView);

