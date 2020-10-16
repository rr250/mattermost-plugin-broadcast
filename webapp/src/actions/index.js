
import {Client4} from 'mattermost-redux/client';

import {id as pluginId} from '../manifest';

export const broadcast = async (message, userList, channelList) => {
    var channelIdList = [];
    channelList.forEach((element) => {
        channelIdList.push(element.id);
    });
    var userIdList = [];
    userList.forEach((element) => {
        userIdList.push(element.id);
    });
    await fetch(window.location.origin + '/plugins/' + pluginId + '/broadcast', Client4.getOptions({
        method: 'post',
        body: JSON.stringify({message, userIdList, channelIdList}),
    }));
};

export const getAllUsersInTeam = (teamId) => {
    getAllUsersInCurrentTeam(teamId).then((users) => {
        return users;
    });
};

export const getAllUsersInCurrentTeam = async (teamId) => {
    const userListPromise = await fetch(window.location.origin + '/plugins/' + pluginId + '/getallusersinteam', Client4.getOptions({
        method: 'post',
        body: JSON.stringify({teamId}),
    }));
    const userList = await userListPromise.json();
    return userList;
};