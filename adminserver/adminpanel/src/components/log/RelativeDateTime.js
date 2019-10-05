import {Component, h} from 'preact';
import 'semantic-ui-table/table.css';
import moment from 'moment';

export const GOLANG_TIMEFORMAT = 'YYYY-MM-DD HH:mm:ss.SSSSSSS zz';

export default class Log extends Component {
    refresh = () => {
        const { id } = this.state;
        this.setState({
            id: (id || 0) + 1,
        });
    };

    componentDidMount() {
        const { interval } = this.state;
        if (interval === undefined) {
            this.setState({
                interval: setInterval(() => {
                    this.refresh();
                }, 100),
            });
        }
    }

    render({ time }) {
        return moment(time, GOLANG_TIMEFORMAT, false)
            .fromNow();
    }
}
