import React, {Component} from 'react';
import {next, prev, goto} from './ducks';
import {connect} from 'react-redux';
import moment from 'moment';
import 'c3/c3.min.css';
import c3 from 'c3';

class SidePanel extends Component {
    constructor(props) {
        super(props);

        this.state = {collapsed: false};
        this.toggleCollapsed = this.toggleCollapsed.bind(this);
    }

    toggleCollapsed() {
        this.setState({collapsed: !this.state.collapsed});
    }

    componentDidMount() {
        let {history, goto} = this.props;
        let changes = ['Lines Changed'];
        let dates = ['x'];

        history.forEach(({DiffJSON, CommitterDate}) => {
            //TODO(gbbr): Try to get change weight only for excerpt (not overall commit)
            changes.push(DiffJSON[0].addedLines + DiffJSON[0].deletedLines);
            dates.push(moment(CommitterDate).toDate());
        });

        c3.generate({
            bindto: '#chart',
            data: {
                x: 'x',
                columns: [dates, changes],
                onclick: ({index}) => goto(index)
            },
            axis: {
                x: {
                    type: 'timeseries',
                    tick: {
                        format: '%d-%m-%Y %H:%M'
                    }
                }
            }
        });
    }

    render() {
        const {next, prev, total, entry, index} = this.props;
        const {collapsed} = this.state;
        const prevClass = index === 0 ? 'disabled' : '';
        const nextClass = index === total - 1 ? 'disabled' : '';

        return (
            <div className={"panel" + (collapsed ? " collapsed" : "")}>
                <div className="nav">
                    <a className={prevClass} onClick={prev}>Prev</a>
                    <a className={nextClass} onClick={next}>Next</a>
                    <a className="pull-right" onClick={this.toggleCollapsed}>
                        {collapsed ? 'Show' : 'Hide'}
                    </a>
                </div>
                <div className="details">
                    <div>SHA: {entry.SHA}</div>
                    <div>
                        Author: {entry.AuthorName} &lt;{entry.AuthorEmail}&gt;
                        &nbsp;{moment(entry.AuthorDate).fromNow()}
                    </div>
                    <div>
                        Committer: {entry.CommitterName} &lt;{entry.CommitterEmail}&gt;
                        &nbsp;{moment(entry.CommitterDate).fromNow()}
                    </div>
                </div>
                <div id="chart" />
                <div className="txt"><pre>{entry.Msg}</pre></div>
            </div>
        );
    }
}

export default connect(({history, index}) => ({
    entry: history[index],
    total: history.length,
    history,
    index
}), {next, prev, goto})(SidePanel);
