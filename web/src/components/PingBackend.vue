<script setup>

import {ping} from "@/common/api/ping";

let callback = (resp) => {
    console.info("ping success, resp: ", resp)
}

let errorCallback = (error) => {
  console.error("ping error, error: ", error)

  // clearInterval(intervalId)
}

let goPing = () => {
  ping(callback, errorCallback)
}

setInterval(goPing, 3000);

</script>