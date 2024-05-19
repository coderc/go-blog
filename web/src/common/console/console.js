import {getTimer} from "@/common/time/time";

export const cInfo = (...optionParams) => {
    let timer = getTimer()
    console.info(timer, ...optionParams)
}

export const cWarn = (...optionParams) => {
    let timer = getTimer()
    console.warn(timer, ...optionParams)
}

export const cError = (...optionParams) => {
    let timer = getTimer()
    console.error(timer, ...optionParams)
}