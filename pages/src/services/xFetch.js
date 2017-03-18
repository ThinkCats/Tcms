// 自定义访问服务器api函数

import {notification} from 'antd';

const errorMessages = (res) => `${res.status} ${res.statusText}`;

function check401(res) {
  if (res.status === 401) {
    location.href = '/401';
  }
  return res;
}

function check404(res) {
  if (res.status === 404) {
    return Promise.reject(errorMessages(res));
  }
  return res;
}

function jsonParse(res) {
  return res.json();
}

function errorMessageParse(res) {
  const success = res.success;
  const code = res.code;
  if (!success || code != '200') {
    notification.warn({
      message: '获取服务器数据异常',
      description: res.msg
    });
    return Promise.reject(res.msg);
  }
  return res;
}

function catchError(error){
    if(error){
        console.error('ERROR:',error);
    }
}

function xFetch(url) {
    let seperator = '';
    let token = '';

    let options = {
        method: 'get'
    };

    //开发时候指定SERVER_URL,需带token,正式发布不需要携带该token
    //开发环境下，不携带cookie
    if(SERVER_URL){
        token = '_security_token_inc=91481170693197946';
    }else{
        options.credentials = 'include';
    }

    if(url.endsWith('.json')){
        seperator = '?';
    }else{
        seperator = '&';
    }
    url = url + seperator + token;


    return fetch(url, options)
        .then(check401)
        .then(check404)
        .then(jsonParse)
        .then(errorMessageParse)
        .catch(catchError);
}

export default xFetch;
