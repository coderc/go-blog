import {request} from "@/common/http/request";
import {backendUrl} from "@/config/backendUrl";
import {MethodPOST} from "@/common/http/method";

const ping = (callback, errorCallback) => {
    let req;

    req = new request().withCallback(callback)
        .withErrorCallback(errorCallback)
        .withHost(backendUrl)
        .withUri('/api/ping')
        .withMethod(MethodPOST);

    req.send()
}

export {
    ping
}