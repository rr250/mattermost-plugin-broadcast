/* eslint-disable react/jsx-filename-extension */
/* eslint-disable indent */
import React from 'react';
import {FormattedMessage} from 'react-intl';
import {Multiselect} from 'multiselect-react-dropdown';
import Button from 'react-bootstrap/Button';
import Form from 'react-bootstrap/Form';
import Modal from 'react-bootstrap/Modal';

import {add} from '../actions';

export default class RHSView extends React.PureComponent {
    constructor(props) {
        super(props);
        var options = [];
        // eslint-disable-next-line react/prop-types
        Object.values(this.props.team).forEach((user) => {
            options.push({name: user.username, id: user.id});
        });
        this.state = {
            options,
            selectedList: [],
            message: null,
        };
    }

    onSelect(selectedList, selectedItem) {
        this.setState({selectedList: [...this.state.selectedList, selectedItem]});
    }

    onRemove(selectedList, removedItem) {
        this.setState({selectedList: this.state.selectedList.filter((selectedItem) => selectedItem !== removedItem)});
    }
    onMessage(message) {
        this.setState({message});
    }
    submit() {
		add(this.state.message, this.state.selectedList);
		this.setState({
			selectedList: [],
            message: null,
            options: this.state.options,
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
                            options={this.state.options} // Options to display in the dropdown
                            selectedValues={this.state.selectedList} // Preselected value to persist in dropdown
                            onSelect={(selectedList, selectedItem) => this.onSelect(selectedList, selectedItem)} // Function will trigger on select event
                            onRemove={(selectedList, selectedItem) => this.onRemove(selectedList, selectedItem)} // Function will trigger on remove event
                            displayValue='name' // Property name to display in the dropdown options
                            placeholder='Select Users'
                            closeOnSelect={false}

                            // showCheckbox={true} //checkbox inactive
                            avoidHighlightFirstOption={true}
                        />
                    </div>
                    <br/>
                    <br/>
                    <br/>
                    <br/>
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
                        variant='success'
                        onClick={() => this.submit()}
                    // eslint-disable-next-line react/jsx-no-literals
                    >Send Message</Button>
                </Modal.Footer>
            </div>

        );
    }
}