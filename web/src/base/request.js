import axios from 'axios' // 引入axios
axios.defaults.baseURL = '/api/'

const request = {}


request.login= (data) => {
    let d ={
        RequestData : data
    }
    return axios.post('login', d);
}

export default request
