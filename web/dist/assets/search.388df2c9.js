/*! 
 Build based on gin-vue-admin 
 Time : 1679909256000 */
import{u as s,c2 as e,r as a,o,c as r,e as t,w as c,U as l,aC as i,J as n,a0 as u,d,n as p,V as v,ck as m}from"./index.61dd0a2d.js";/* empty css              */import"./tag.3ecd1fed.js";import{E as f,a as b}from"./select.31449efc.js";import"./scrollbar.ef9dc6d6.js";/* empty css               */import j from"./index.3c42b07d.js";import{_ as x}from"./_plugin-vue_export-helper.cdc0426e.js";import"./index.cbcb4c8f.js";import"./index.ba9b69da.js";import"./strings.ed5ba909.js";import"./isEqual.3dfeee36.js";import"./_Uint8Array.0caf8bf8.js";const h={class:"search-component"},g={key:0,class:"transition-box"},k={class:"user-box"},y={class:"user-box"},_={name:"BtnBox"},I=x(Object.assign(_,{setup(x){const _=s(),I=e(),w=s=>{s.indexOf("http:")>-1||s.indexOf("https:")>-1?window.open(s):_.push({name:s})},C=a(!0),q=async()=>{setTimeout((()=>{C.value=!1}),100)},B=a(null),O=async()=>{C.value=!0,await v(),B.value.focus()},E=a(!1),T=()=>{E.value=!0,m.emit("reload"),setTimeout((()=>{E.value=!1}),500)},U=()=>{window.open("https://support.qq.com/product/371961")};return(s,e)=>{const a=f,v=b;return o(),r("div",h,[C.value?(o(),r("div",g,[t(v,{ref_key:"searchInput",ref:B,filterable:"",placeholder:"请选择",onBlur:q,onChange:w},{default:c((()=>[(o(!0),r(l,null,i(n(I).routerList,(s=>(o(),u(a,{key:s.value,label:s.label,value:s.value},null,8,["label","value"])))),128))])),_:1},512)])):(o(),r(l,{key:1},[d("div",k,[d("div",{class:p(["gvaIcon gvaIcon-refresh",[E.value?"reloading":""]]),onClick:T},null,2)]),d("div",{class:"user-box"},[d("div",{class:"gvaIcon gvaIcon-search",onClick:O})]),d("div",y,[t(j,{class:"search-icon",style:{cursor:"pointer"}})]),d("div",{class:"user-box"},[d("div",{class:"service gvaIcon-customer-service",onClick:U})])],64))])}}}),[["__scopeId","data-v-fa45d217"]]);export{I as default};