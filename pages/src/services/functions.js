import React,{ Component } from 'react';

function add0(m){
    return m<10?'0'+m:m;
}

export function formatDate(shijianchuo){
    var time = new Date(shijianchuo);
    var y = time.getFullYear();
    var m = time.getMonth()+1;
    var d = time.getDate();
    var h = time.getHours();
    var mm = time.getMinutes();
    var s = time.getSeconds();
    return y+'-'+add0(m)+'-'+add0(d)+' '+add0(h)+':'+add0(mm)+':'+add0(s);
}

export function getDay(shijianchuo){
    var time = new Date(shijianchuo);
    var d = time.getDate();
    return d;
}

export function getYMD(shijianchuo){
    var time = new Date(shijianchuo);
    var y = time.getFullYear();
    var m = time.getMonth()+1;
    var d = time.getDate();
    var h = time.getHours();
    var mm = time.getMinutes();
    var s = time.getSeconds();
    return y+'-'+add0(m)+'-'+add0(d);
}

export function formatDateZero(shijianchuo){
    var time = new Date(shijianchuo);
    var y = time.getFullYear();
    var m = time.getMonth()+1;
    var d = time.getDate();
    var h = time.getHours();
    var mm = time.getMinutes();
    var s = time.getSeconds();
    return y+'-'+add0(m)+'-'+add0(d)+' 00:00:00';
}

export function formatDateEnd(shijianchuo){
    var time = new Date(shijianchuo);
    var y = time.getFullYear();
    var m = time.getMonth()+1;
    var d = time.getDate();
    var h = time.getHours();
    var mm = time.getMinutes();
    var s = time.getSeconds();
    return y+'-'+add0(m)+'-'+add0(d)+' 23:59:59';
}

export function fenToYuan(fen){
    if(isNaN(fen)){
        return '-';
    }
    return fen / 100;
}

export function fenToWan(fen){
    if(isNaN(fen)){
        return '-';
    }
    return fen / 100 / 10000;
}

