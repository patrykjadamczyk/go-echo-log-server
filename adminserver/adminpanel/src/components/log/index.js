import {Component, h} from 'preact';
import 'semantic-ui-table/table.css';
import RelativeDateTime from './RelativeDateTime';

function parseQuery(queryString) {
    const queryRegex = new RegExp(/\??[^&"'=]+=[^&"'=]+&*/);
    if (!queryRegex.test(queryString)) {
        throw new Error("Not a query string");
    }
    const query = {};
    const pairs = (queryString[0] === '?' ? queryString.substr(1) : queryString).split('&');
    for (let i = 0; i < pairs.length; i++) {
        const pair = pairs[i].split('=');
        query[decodeURIComponent(pair[0])] = decodeURIComponent(pair[1] || '');
    }
    return query;
}

export default class Log extends Component {
    render({ identifier, data }) {
        const {
            Start: start,
            RequestInfo: info,
            RequestDataForm: dataForm,
            RequestData: dataObj,
        } = data;
        let formattedDataForm = dataForm;
        try {
            formattedDataForm = JSON.stringify(JSON.parse(dataForm), undefined, 4);
        } catch (e) {}

        let formattedDataObj = dataObj;
        try {
            formattedDataObj = JSON.stringify(JSON.parse(dataObj), undefined, 4);
        } catch (e) {
            try {
                formattedDataObj = JSON.stringify(parseQuery(dataObj), undefined, 4);
            } catch (f) {}
        }

        return (
            <div>
                <h2>{identifier}</h2>
                <table class="ui celled table">
                    <thead>
                        <tr>
                            <th>Information</th>
                            <th>Data</th>
                        </tr>
                    </thead>
                    <tbody>
                        <tr>
                            <td className="ui center aligned">
                                Time Ago
                            </td>
                            <td>
                                <RelativeDateTime time={start}/>
                            </td>
                        </tr>
                        <tr>
                            <td class="ui center aligned">
                                Date
                            </td>
                            <td>
                                {start}
                            </td>
                        </tr>
                        <tr>
                            <td class="ui center aligned">
                                Request Info
                            </td>
                            <td>
                                <pre>
                                    {info}
                                </pre>
                            </td>
                        </tr>
                        <tr>
                            <td class="ui center aligned">
                                Request Form Data
                            </td>
                            <td>
                                <pre>
                                    {formattedDataForm}
                                </pre>
                            </td>
                        </tr>
                        <tr>
                            <td class="ui center aligned">
                                Request Data
                            </td>
                            <td>
                                <pre>
                                    {formattedDataObj}
                                </pre>
                            </td>
                        </tr>
                    </tbody>
                </table>
            </div>
        );
    }
}
