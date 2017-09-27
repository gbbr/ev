import React from 'react';
import {render} from 'react-dom';
import App from './app';
import reducer from './ducks';
import {Provider} from 'react-redux';
import {createStore} from 'redux';
import {Diff2Html as diff} from 'diff2html';

function makeStore() {
    let history = window.GIT_HISTORY;

    history.forEach((entry) => {
        entry.DiffJSON = diff.getJsonFromDiff(entry.Diff);
    });

    return createStore(reducer, {index: 0, history});
}

document.addEventListener('DOMContentLoaded', () => {
    const store = makeStore();

    render(
        <Provider store={store}><App /></Provider>,
        document.getElementById('app-root')
    );
});
