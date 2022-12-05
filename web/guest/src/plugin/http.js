import Vue from 'vue'
import axios from 'axios'

let Url = 'https://nu1l.online:8443/api/v1/'

axios.defaults.baseURL = Url

Vue.prototype.$http = axios

export { Url }
