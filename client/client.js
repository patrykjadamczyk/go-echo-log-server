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
        let xhr = new XMLHttpRequest();
        xhr.open('POST', `${base_url}/logger`);
        xhr.setRequestHeader('Content-type', 'application/json; charset=utf-8');
        xhr.send(JSON.stringify(data));
        xhr.onload = resolve;
        xhr.onerror = reject;
    });
}
