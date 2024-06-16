<script setup>

import {ping} from "@/common/api/ping";
import {cInfo, cError} from "@/common/console/console"

let callback = (resp) => {
  cInfo("ping success, resp: ", resp)
}

let errorCallback = (error) => {
  cError("ping error, error: ", error)

  // clearInterval(intervalId)
}

let goPing = () => {
  ping(callback, errorCallback)
}

setInterval(goPing, 200);

</script>