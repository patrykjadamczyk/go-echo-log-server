import {Component} from 'preact';
import style from './style';
import axios from 'axios';
import Log from '../../components/log';
import moment from 'moment';
import 'semantic-ui-button/button.css';
import {GOLANG_TIMEFORMAT} from '../../components/log/RelativeDateTime';

export default class Home extends Component {
	componentDidMount() {
		const { interval } = this.state;
		if (interval === undefined) {
			this.setState({
				interval: setInterval(() => {
					this.initData();
				}, 1000),
			});
		}
		this.initData();
	}

	initData = () => {
		let url = '/_data/';
		if (process.env.NODE_ENV === 'development') {
			const BASE_URL = 'http://127.0.0.1:7778';
			url = `${BASE_URL}/_data/`;
		}
		axios.get(url)
			.then(response => response.data)
			.then(this.sortItems)
			.then((data) => {
				this.setState({
					err: false,
					data,
				});
			})
			.catch((err) => {
				this.setState({
					err: true,
				});
				console.error(err);
			});
	};

	sortItems(requests) {
		const getDate = date => moment(date, GOLANG_TIMEFORMAT, false);
		const sortedKeys = Object
			.keys(requests)
			.sort((a,b) => {
				const realA = getDate(requests[a].Start).unix();
				const realB = getDate(requests[b].Start).unix();
				return realA - realB;
			});
		const sortedRequests = {};
		sortedKeys.forEach((key) => {
			sortedRequests[key] = requests[key];
		});
		return sortedRequests;
	}

	render(props, { data }) {
		return (
			<div className={style.home}>
				<br/>
				<button
					class="ui fluid primary button"
					onClick={() => this.initData()}
				>
					Refresh Now
				</button>
				{
					data !== null && data !== undefined
						? Object.keys(data).reverse().map((reqId) => (
							<Log identifier={reqId} data={data[reqId]} />
						))
						: null
				}
			</div>
		);
	}
}
