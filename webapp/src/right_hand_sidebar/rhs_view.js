/* eslint-disable react/prop-types */
/* eslint-disable no-console */
/* eslint-disable react/jsx-filename-extension */
/* eslint-disable indent */
import React from 'react';
import {FormattedMessage} from 'react-intl';
import {Multiselect} from 'multiselect-react-dropdown';
import Button from 'react-bootstrap/Button';
import Form from 'react-bootstrap/Form';
import Modal from 'react-bootstrap/Modal';

import {broadcast, getAllUsersInTeam} from '../actions';

export default class RHSView extends React.PureComponent {
    constructor(props) {
        super(props);
        var options1 = [];

        // Object.values(this.props.team).forEach((user) => {
        //     options1.push({name: user.username, id: user.id});
        // });
        var options2 = [];
        Object.values(this.props.channels).forEach((channel) => {
            if (channel.type !== 'D') {
                options2.push({name: channel.display_name + ' (url:/' + channel.name + ')', id: channel.id});
            }
        });
        this.state = {
            options1,
            options2,
            selectedList1: [],
            selectedList2: [],
            message: null,
        };
    }

    componentDidMount() {
        console.log(this.props.theme);
        getAllUsersInTeam(this.props.currentTeamId).then((users) => {
            var options1 = [];
            Object.values(users).forEach((user) => {
                options1.push({name: user.username, id: user.id});
            });
            this.setState({options1});
        });
    }

    onSelect1(selectedList1, selectedItem) {
        this.setState({selectedList1: [...this.state.selectedList1, selectedItem]});
    }

    onRemove1(selectedList1, removedItem) {
        this.setState({selectedList1: this.state.selectedList1.filter((selectedItem) => selectedItem !== removedItem)});
    }
    onSelect2(selectedList2, selectedItem) {
        this.setState({selectedList2: [...this.state.selectedList2, selectedItem]});
    }

    onRemove2(selectedList2, removedItem) {
        this.setState({selectedList2: this.state.selectedList2.filter((selectedItem) => selectedItem !== removedItem)});
    }
    onMessage(message) {
        this.setState({message});
    }
    submit(e) {
        e.preventDefault();
        broadcast(this.state.message, this.state.selectedList1, this.state.selectedList2);
		this.setState({
            message: '',
		});
    }
    render() {
        return (
            <div>
                <Modal.Body>
                    <div>
                        {/* <Form.Label>Select Users</Form.Label> */}
                        <FormattedMessage
                            id='plugin.name'
                            defaultMessage='Select Users'
                        />
                        <Multiselect
                            options={this.state.options1} // Options to display in the dropdown
                            selectedValues={this.state.selectedList1} // Preselected value to persist in dropdown
                            onSelect={(selectedList1, selectedItem) => this.onSelect1(selectedList1, selectedItem)} // Function will trigger on select event
                            onRemove={(selectedList1, selectedItem) => this.onRemove1(selectedList1, selectedItem)} // Function will trigger on remove event
                            displayValue='name' // Property name to display in the dropdown options1
                            placeholder='Select Users'
                            closeOnSelect={false}

                            // showCheckbox={true} //checkbox inactive
                            avoidHighlightFirstOption={true}
                            style={{option: { // To change css for dropdown options
                                color: this.props.theme.sidebarHeaderTextColor, background: this.props.theme.sidebarHeaderBg,
                            }}}
                        />
                        <br/>
                        <br/>
                        <Multiselect
                            options={this.state.options2} // Options to display in the dropdown
                            selectedValues={this.state.selectedList2} // Preselected value to persist in dropdown
                            onSelect={(selectedList2, selectedItem) => this.onSelect2(selectedList2, selectedItem)} // Function will trigger on select event
                            onRemove={(selectedList2, selectedItem) => this.onRemove2(selectedList2, selectedItem)} // Function will trigger on remove event
                            displayValue='name' // Property name to display in the dropdown options1
                            placeholder='Select Channels'
                            closeOnSelect={false}

                            // showCheckbox={true} //checkbox inactive
                            avoidHighlightFirstOption={true}
                            style={{option: { // To change css for dropdown options
                                color: this.props.theme.sidebarHeaderTextColor, background: this.props.theme.sidebarHeaderBg,
                            }}}
                        />
                    </div>
                    <br/>
                    <br/>
                    <div>
                        <Form.Group controlId='exampleForm.ControlTextarea1'>
                            <FormattedMessage
                                id='plugin.name'
                                defaultMessage='Message'
                            />
                            <Form.Control
                                value={this.state.message}
                                onChange={(event) => {
                                    this.onMessage(event.target.value);
                                }}
                                as='textarea'
                                placeholder='Message'
                                rows='3'
                            />
                        </Form.Group>
                    </div>
                </Modal.Body>
                <br/>
                <br/>
                <br/>
                <Modal.Footer>
                    <Button
                        style={{color: this.props.theme.buttonColor, background: this.props.theme.buttonBg, border: 0}}
                        variant='success'
                        onClick={(e) => this.submit(e)}
                    // eslint-disable-next-line react/jsx-no-literals
                    >Send Message</Button>
                </Modal.Footer>
            </div>

        );
    }
}