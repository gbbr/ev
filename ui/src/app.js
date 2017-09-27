import './styles.css';
import React, {Component} from 'react';
import DiffView from './diffView';
import SidePanel from './sidePanel';

export default class App extends Component {
    render() {
        return (
            <div className="ev-root">
                <DiffView />
                <SidePanel />
            </div>
        );
    }
}
