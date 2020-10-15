/* eslint-disable react/jsx-filename-extension */
/* eslint-disable lines-around-comment */
/* eslint-disable no-console */
import {Component} from 'react';
import {FormattedMessage} from 'react-intl';

import RHSView from './right_hand_sidebar';

import {id as pluginId} from './manifest';

const Icon = () => <i className='icon fa fa-bullhorn'/>;

// eslint-disable-next-line react/require-optimization
class BroadcastPlugin extends Component {
    initialize(registry, store) {
        // console.log(store);

        const {toggleRHSPlugin} = registry.registerRightHandSidebarComponent(
            RHSView,
            <FormattedMessage
                id='plugin.name'
                defaultMessage='Broadcast'
            />);
        registry.registerChannelHeaderButtonAction(
            // icon - JSX element to use as the button's icon
            <Icon/>,
            () => store.dispatch(toggleRHSPlugin),
            // dropdown_text - string or JSX element shown for the dropdown button description
            <FormattedMessage
                id='plugin.name'
                defaultMessage='Broadcast'
            />,
        );
    }

    uninitialize() {
        console.log(pluginId + '::uninitialize()');
    }
}

export default BroadcastPlugin;
