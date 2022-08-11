import axios from 'axios' // 引入axios
axios.defaults.baseURL = '/api/'

const request = {
    login: function (data:any, requestOK:any ) {
        const d ={
            RequestData : data
        }
        return axios.post('login', d).then(requestOK);
    },
    do: function (token:string, apiType:string, data:any) {
        const d ={
            Token : token,
            ApiType : apiType,
            RequestData : data
        }
        return axios.post('do', d);
    }
}

export default request
