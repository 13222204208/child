/*! 
 Build based on gin-vue-admin 
 Time : 1680141068000 */
import e from"./menuItem.9d7dcf2f.js";import o from"./asyncSubmenu.2a4cd848.js";import{z as t,aA as r,o as n,a0 as s,w as i,c as a,U as f,aC as l,f as p,a1 as u,J as c}from"./index.9fcd1cf4.js";/* empty css                *//* empty css               *//* empty css             */import"./_plugin-vue_export-helper.cdc0426e.js";import"./index.9028de34.js";import"./index.f783073c.js";import"./index.bb4181cc.js";const m={name:"AsideComponent"},d=Object.assign(m,{props:{routerInfo:{type:Object,default:()=>null},isCollapse:{default:function(){return!1},type:Boolean},theme:{default:function(){return{}},type:Object}},setup(m){const d=m,h=t((()=>d.routerInfo.children&&d.routerInfo.children.filter((e=>!e.hidden)).length?o:e));return(e,o)=>{const t=r("AsideComponent");return m.routerInfo.hidden?p("",!0):(n(),s(u(c(h)),{key:0,"is-collapse":m.isCollapse,theme:m.theme,"router-info":m.routerInfo},{default:i((()=>[m.routerInfo.children&&m.routerInfo.children.length?(n(!0),a(f,{key:0},l(m.routerInfo.children,(e=>(n(),s(t,{key:e.name,"is-collapse":!1,"router-info":e,theme:m.theme},null,8,["router-info","theme"])))),128)):p("",!0)])),_:1},8,["is-collapse","theme","router-info"]))}}});export{d as default};