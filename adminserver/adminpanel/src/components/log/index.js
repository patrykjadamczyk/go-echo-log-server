import {Component, h} from 'preact';
import 'semantic-ui-table/table.css';
import RelativeDateTime from './RelativeDateTime';

export default class Log extends Component {
    render({ identifier, data }) {
        const {
            Start: start,
            RequestInfo: info,
            RequestDataForm: dataForm,
            RequestData: dataObj,
        } = data;
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
                                {dataForm}
                            </td>
                        </tr>
                        <tr>
                            <td class="ui center aligned">
                                Request Data
                            </td>
                            <td>
                                {dataObj}
                            </td>
                        </tr>
                    </tbody>
                </table>
            </div>
        );
    }
}
