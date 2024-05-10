import $ from 'jquery'
import {backendUrl} from "@/config/backendUrl";

class request {
    method = ''
    host = ''
    uri = ''
    headers = {}
    data = {}
    callback = () => {}
    errorCallback = () => {}

    constructor(method, url, headers, data, callback, errorCallback) {

        this.method = method
        this.uri = url
        this.headers = headers
        this.data = data
        this.callback = callback
        this.errorCallback = errorCallback
    }

    withMethod(method) {
        this.method = method
        return this
    }

    withHost(host) {
        this.host = host
        return this
    }

    withUri(uri) {
        this.uri = uri
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
        if (this.host === '') this.host = backendUrl
        let url = this.host+this.uri
        $.ajax({
            method: this.method,
            url: url,
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