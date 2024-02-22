import axios from 'axios'
import qs from 'qs'
import { message } from "antd";
/**
 *
 */
axios.interceptors.request.use(function (config) {
    const { method, data } = config;
    if (method.toLocaleLowerCase() === 'post' && typeof
        data === 'object') {
        config.data = qs.stringify(data);
    }
    return config;
}, function (error) {
    // Do something with request error
    return Promise.reject(error);
});
/*
响应拦截器，响应数据之前做⼯作
*/
axios.interceptors.response.use(function (response) {
    return response.data;
}, function (error) {
    message.error('Response Failed:' + error.message);
    // return Promise.reject(error);
    // 让错误处于pending状态，不再往下执⾏，
//返回⼀个pending的promise, 中断promise链
    return new Promise(()=>{})
});
export default axios;
