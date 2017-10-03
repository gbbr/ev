import './styles.css';
import 'diff2html/dist/diff2html.css';

import React, {Component} from 'react';
import {render} from 'react-dom';
import reducer from './ducks';
import {Provider} from 'react-redux';
import {createStore} from 'redux';
import {Diff2Html as diff} from 'diff2html';
import SidePanel from './sidePanel';
import {connect} from 'react-redux';

class DiffViewComponent extends Component {
    render() {
        const __html = diff.getPrettyHtml(this.props.diff);
        return <div dangerouslySetInnerHTML={{__html}} />;
    }
}

const DiffView = connect(({history, index}) => ({diff: history[index].Diff}))(DiffViewComponent);

document.addEventListener('DOMContentLoaded', () => {
    const store = createStore(reducer, {index: 0, history: window.GIT_HISTORY});

    render(
        <Provider store={store}>
            <div className="ev-root">
                <DiffView />
                <SidePanel />
            </div>
        </Provider>,
        document.getElementById('app-root')
    );
});
