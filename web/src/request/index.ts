import axios from 'axios' // 引入axios
axios.defaults.baseURL = '/api/'

const request = {
    login: function (data:any) {
        const d ={
            RequestData : data
        }
        return axios.post('login', d);
    }
}

export default request
