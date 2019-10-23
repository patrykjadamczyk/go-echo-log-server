function log(base_url, trace = false, variable = {}) {
    return new Promise((resolve, reject) => {
        let data = variable;
        if (trace) {
            const err = new Error();
            data = {
                _var: variable,
                _trace: err.stack,
            };
        }
        data = JSON.stringify(data);
        const url = `${base_url}/logger`;
        let xhr;
        try {
            xhr = new XMLHttpRequest();
            xhr.open('POST', url);
            xhr.setRequestHeader('Content-type', 'application/json; charset=utf-8');
            xhr.send(data);
            xhr.onload = resolve;
            xhr.onerror = reject;
        } catch (e) {
            const http = require('http');
            const req = http.request(url, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json; charset=utf-8',
                },
            }, (res) => {
                res.on('end', resolve);
            });
            req.on('error', reject);
            req.write(data);
            req.end();
        }
    });
}
