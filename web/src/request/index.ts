import axios from 'axios'
import { message } from 'ant-design-vue';
import router from "@/router"; // 引入axios
axios.defaults.baseURL = '/api/'

const request = {
    login: function (data:any, loginOK:any) {
        const d ={
            RequestData : data
        }
        const onRequestOK =(r:any) => {
            console.log('Request OK:', r);
            if (r.data["ResponseCode"]==0) {
                loginOK(r.data)
            }else{
                message.error(r.data["ErrMsg"]);
            }
        }
        const onRequestErr =(error:any) => {
            console.log("Request Err:",error);
            message.error("登录失败，请检查网络连接情况！");
        }

        return axios.post('login', d).then(onRequestOK).catch(onRequestErr);
    },
    do: function (token:string, apiType:string, data:any, doOK:any) {
        const d ={
            Token : token,
            ApiType : apiType,
            RequestData : data
        }
        const onRequestOK =(r:any) => {
            console.log('Request OK:', r);
            if (r.data["ResponseCode"]==0) {
                doOK(r.data)
            }else if (r.data["ResponseCode"]==1){
                message.error(r.data["ResponseCode"]);
            }else if (r.data["ResponseCode"]==2){
                message.error("登录超时!请重新登录");
                router.push({name:"login"});
            }else {
                message.error(r.data["ResponseCode"]);
            }
        }
        const onRequestErr =(error:any) => {
            console.log("Request Err:",error);
            message.error("操作失败，请检查网络连接情况！");
        }
        return axios.post('do', d).then(onRequestOK).catch(onRequestErr);
    },

}

export default request
