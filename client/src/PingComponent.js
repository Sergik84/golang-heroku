import React, { Component } from 'react'
import axios from 'axios';
class PingComponent extends Component {

    constructor() {
        super();
        this.state = {
            health: 'pending'
        }
    }

    componentWillMount() {
        axios.get('api/health')
            .then((response) => {
                this.setState(() => {
                    return { health: response.data.status }
                })
            })
            .catch(function (error) {
                console.log(error);
            });

    }

    render() {
        return <h1>Application Status: {this.state.health}</h1>;
    }
}

export default PingComponent;