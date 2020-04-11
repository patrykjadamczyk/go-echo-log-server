import {Component, h} from 'preact';
import {parse} from 'query-string';
import 'semantic-ui-table/table.css';
import RelativeDateTime from './RelativeDateTime';

export default class Log extends Component {
    checkIsXML(data) {
        const parser = new DOMParser();
        const dom = parser.parseFromString(data, "text/xml");
        return dom.documentElement.nodeName !== "parsererror";
    }

    checkIsHTML(data) {
        const parser = new DOMParser();
        const dom = parser.parseFromString(data, "text/html");
        return dom.documentElement.nodeName !== "parsererror";
    }

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
        if (!this.checkIsXML(dataObj) && !this.checkIsHTML(dataObj)) {
            try {
                formattedDataObj = JSON.stringify(JSON.parse(dataObj), undefined, 4);
            } catch (e) {
                try {
                    formattedDataObj = JSON.stringify(parse(dataObj), undefined, 4);
                } catch (f) {
                }
            }
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
