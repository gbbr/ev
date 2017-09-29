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
        this.chart = null;
    }

    toggleCollapsed() {
        this.setState({collapsed: !this.state.collapsed});
    }

    componentDidMount() {
        let {history, goto, total} = this.props;
        let changes = ['Lines Changed'];
        let dates = ['x'];

        history.forEach(({Changes, CommitterDate}) => {
            changes.push(Changes);
            dates.push(moment(CommitterDate).toDate());
        });

        this.chart = c3.generate({
            bindto: '#chart',
            data: {
                x: 'x',
                columns: [dates, changes],
                onclick: ({index}) => goto(total - index - 1),
                selection: {
                    enabled: true,
                    multiple: false
                }
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

        this.selectIndex(0);
    }

    selectIndex(i) {
        this.chart.select(null, [this.props.total - i - 1], true);
    }

    componentDidUpdate(prevProps) {
        const {index} = this.props;

        if (prevProps.index !== index && this.chart !== null) {
            this.selectIndex(index)
        }
    }

    render() {
        const {next, prev, total, entry, index} = this.props;
        const {collapsed} = this.state;

        return (
            <div className={"panel" + (collapsed ? " collapsed" : "")}>
                <div className="nav">
                    <button disabled={index === total - 1} onClick={next}>&lt; Older</button>
                    <button disabled={index === 0} onClick={prev}>Newer &gt;</button>
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
