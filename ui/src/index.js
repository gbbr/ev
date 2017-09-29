import React from 'react';
import {render} from 'react-dom';
import App from './app';
import reducer from './ducks';
import {Provider} from 'react-redux';
import {createStore} from 'redux';
import {Diff2Html as diff} from 'diff2html';

document.addEventListener('DOMContentLoaded', () => {
    const store = createStore(reducer, {index: 0, history: window.GIT_HISTORY});

    render(
        <Provider store={store}><App /></Provider>,
        document.getElementById('app-root')
    );
});
