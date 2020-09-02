
import {Client4} from 'mattermost-redux/client';

import {id as pluginId} from '../manifest';

export const add = async (message, usersList) => {
    var usersid = [];
    usersList.forEach((element) => {
        usersid.push(element.id);
    });
    await fetch(window.location.origin + '/plugins/' + pluginId + '/broadcast', Client4.getOptions({
        method: 'post',
        body: JSON.stringify({Message: message, Usersid: usersid}),
    }));
};

