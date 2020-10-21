import {connect} from 'react-redux';
import {getAllChannels} from 'mattermost-redux/selectors/entities/channels';
import {getUsers} from 'mattermost-redux/selectors/entities/common';
import {getTheme} from 'mattermost-redux/selectors/entities/preferences';

import RHSView from './rhs_view';

const mapStateToProps = (state) => ({
    state,
    channels: getAllChannels(state),
    theme: getTheme(state),
    currentTeamId: state.entities.teams.currentTeamId,
    team: getUsers(state),
});

export default connect(mapStateToProps)(RHSView);

