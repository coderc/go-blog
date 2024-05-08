import $ from 'jquery'

class request {
    method = ''
    url = ''
    headers = {}
    data = {}
    callback = () => {}
    errorCallback = () => {}

    constructor(method, url, headers, data, callback, errorCallback) {

        this.method = method
        this.url = url
        this.headers = headers
        this.data = data
        this.callback = callback
        this.errorCallback = errorCallback
    }

    withMethod(method) {
        this.method = method
        return this
    }

    withUrl(url) {
        this.url = url
        return this
    }

    withHeaders(headers) {
        this.headers = headers
        return this
    }

    withData(data) {
        this.data = data
        return this
    }

    withCallback(callback) {
        this.callback = callback
        return this
    }

    withErrorCallback(errorCallback) {
        this.errorCallback = errorCallback
        return this
    }

    send() {
        $.ajax({
            method: this.method,
            url: this.url,
            headers: this.headers,
            data: this.data,
            success: this.callback,
            error: this.errorCallback
        })
    }
}

export {
    request
}