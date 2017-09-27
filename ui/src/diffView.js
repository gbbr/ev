import React, {Component} from 'react';
import 'diff2html/dist/diff2html.css';
import {Diff2Html as diff} from 'diff2html';
import {connect} from 'react-redux';

class DiffView extends Component {
    render() {
        const __html = diff.getPrettyHtml(this.props.diff);
        return <div dangerouslySetInnerHTML={{__html}} />;
    }
}

export default connect(({history, index}) => ({diff: history[index].Diff}))(DiffView);
